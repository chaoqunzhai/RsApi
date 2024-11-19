package watch

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk"
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/common/utils"
	"go-admin/global"
	"regexp"
	"strconv"
	"strings"
)

func WatchCombinationCustom() {
	dbList := sdk.Runtime.GetDb()
	fmt.Println("巡检资产组合的客户匹配")
	for _, d := range dbList {
		var CombinationList []models2.Combination
		//只自动化检索 在线的
		d.Model(&models2.Combination{}).Select("id,idc_id,host_id,custom_id").Where("status = 3").Find(&CombinationList)
		count := 0
		hostIds := make([]int, 0)
		idcList := make([]int, 0)
		CombinationBindHost := make(map[int]int)
		CombinationBindIdc := make(map[int][]int)
		CombinationHostBindIdc := make(map[int][]int)
		for _, row := range CombinationList {

			if row.HostId > 0 {
				hostIds = append(hostIds, row.Id)
				CombinationBindHost[row.HostId] = row.Id
			}
			if row.IdcId > 0 {
				idcList = append(idcList, row.IdcId)

				bindIdc, bindOk := CombinationBindIdc[row.IdcId]
				if !bindOk {
					bindIdc = make([]int, 0)
				}
				bindIdc = append(bindIdc, row.Id)
				CombinationBindIdc[row.IdcId] = bindIdc
			}
		}

		hostIds = utils.RemoveRepeatInt(hostIds)

		var hostListModel []models2.Host
		//只查询在线的机器
		d.Model(&models2.Host{}).Select("id,idc,sn").Where("id in ? and status = 1", hostIds).Find(&hostListModel)
		//先进行资产组合关联的主机 查询到对应的IDC

		for _, hostRow := range hostListModel {
			CombinationId, ok := CombinationBindHost[hostRow.Id]
			if !ok {
				continue
			}
			fmt.Printf("更新组合 Combination:%v 新IDC-ID:%v\n", CombinationId, hostRow.Idc)
			d.Model(&models2.Combination{}).Where("id = ?", CombinationId).Updates(map[string]interface{}{
				"idc_id": hostRow.Idc,
			})
			if hostRow.Idc == 0 {
				continue
			}
			bindIdc, bindOk := CombinationBindIdc[hostRow.Idc]
			if !bindOk {
				bindIdc = make([]int, 0)
			}
			bindIdc = append(bindIdc, CombinationId)
			CombinationBindIdc[hostRow.Idc] = bindIdc
			idcList = append(idcList, hostRow.Idc)
		}

		idcList = utils.RemoveRepeatInt(idcList)
		fmt.Println("查询IDC列表", idcList)
		var IdcListHost []models2.Idc
		d.Model(&models2.Idc{}).Select("id,custom_id").Where("id in ?", idcList).Find(&IdcListHost)
		fmt.Printf("CombinationBindIdc 关联:%+v\n", CombinationBindIdc)
		for _, v := range IdcListHost {
			CombinationIds, ok := CombinationBindIdc[v.Id]
			if !ok {
				continue
			}
			if v.CustomId == 0 {
				continue
			}
			count += 1
			//更新资产组合关联的客户
			fmt.Printf("更新组合 Combination:%v 新客户ID:%v\n", CombinationIds, v.CustomId)
			d.Model(&models2.Combination{}).Where("id in ?", CombinationIds).Updates(map[string]interface{}{
				"custom_id": v.CustomId,
			})

			//更新拨号关联的客户
			bindHost, bindOk := CombinationHostBindIdc[v.Id]
			if !bindOk {
				continue
			}
			if len(bindHost) == 0 {
				continue
			}
			d.Model(&models2.Dial{}).Where("host_id in ?", bindHost).Updates(map[string]interface{}{
				"custom_id": v.CustomId,
			})
		}
		fmt.Printf("巡检资产组合的客户匹配,总更新了%v条数据", count)

	}

}

func ReRemark(remark string) float64 {
	compile, _ := regexp.Compile("\\((\\d{1,3})\\*(\\d{1,5})M\\)")
	ttt := compile.FindStringSubmatch(remark)
	if len(ttt) < 2 {
		return 0
	}
	line, err1 := strconv.ParseInt(ttt[1], 10, 64)
	width, err2 := strconv.ParseInt(ttt[2], 10, 64)
	if err1 == nil && err2 == nil {
		Value := float64(line) * float64(width)
		return Value
	}
	return 0
}
func WatchHostBandwidth() {
	fmt.Println("重算总带宽任务开始")
	dbList := sdk.Runtime.GetDb()
	for _, d := range dbList {
		var HostList []models2.Host
		count := 0
		d.Model(&models2.Host{}).Select("remark,id,balance").Find(&HostList)

		for _, host := range HostList {
			if host.Remark == "" {
				continue
			}
			algBalance := ReRemark(host.Remark)

			if algBalance == 0 {
				continue
			}
			if algBalance == host.Balance {
				continue
			}
			d.Model(&models2.Host{}).Where("id = ?", host.Id).Updates(map[string]interface{}{
				"balance": algBalance,
			})
			count += 1
		}
		fmt.Printf("重算总带宽任务结束,总更新了%v条数据", count)

	}

}
func WatchCombinationHost() {
	fmt.Println("巡检资产组合空主机数据 任务开始")
	dbList := sdk.Runtime.GetDb()
	for _, d := range dbList {
		var CombinationList []models2.Combination
		d.Model(&models2.Combination{}).Select("code,id").Where("host_id = 0 OR host_id IS  NULL ").Find(&CombinationList)

		for _, v := range CombinationList {
			count := 0
			//先查hostname,因为有的机器SN一样,那就用hostName来做sn
			var host models2.Host
			d.Model(&host).Where("host_name = ?", v.Code).Limit(1).Find(&host)
			updateHostId := 0
			if host.Id > 0 {
				d.Model(&models2.Combination{}).Where("id = ?", v.Id).Updates(map[string]interface{}{
					"host_id": host.Id,
				})
				updateHostId = host.Id

			} else {
				//sn查一下
				var snHost models2.Host
				d.Model(&snHost).Where("sn = ?", v.Code).Limit(1).Find(&snHost)
				if snHost.Id > 0 {
					d.Model(&models2.Combination{}).Where("id = ?", v.Id).Updates(map[string]interface{}{
						"host_id": snHost.Id,
					})
					updateHostId = snHost.Id
				}

			}
			if updateHostId == 0 {
				continue
			}
			d.Model(&models2.Combination{}).Where("id = ?", v.Id).Updates(map[string]interface{}{
				"host_id": updateHostId,
			})

			d.Model(&models2.AdditionsWarehousing{}).Where("combination_id = ?", v.Id).Updates(map[string]interface{}{
				"host_id": updateHostId,
			})
			count += 1

			//如果资产是一个待出库 并且资产 状态不是一个待上架  那就更改为待上架
			if v.Status == 2 && host.Status != 3 {

				d.Model(&models2.Host{}).Where("id = ?", updateHostId).Updates(map[string]interface{}{
					"status": 3,
				})

			}
			fmt.Printf("巡检资产组合空主机数据,总更新了%v条数据", count)

		}
	}
}
func WatchAssetBindHost() {
	dbList := sdk.Runtime.GetDb()

	fmt.Println("巡检资产列表和CMDB关联关系")
	for _, d := range dbList {
		var hostIds []string
		d.Raw("SELECT ac.host_id FROM asset_combination ac JOIN rs_host rh ON ac.host_id = rh.id WHERE ac.status = 2   AND rh.status != 3;").Scan(&hostIds)

		if len(hostIds) > 0 {
			d.Exec(fmt.Sprintf("update rs_host set status =3 where id in  (%s)", strings.Join(hostIds, ",")))
		}
		//只查询 资产分类为主机 + host_id = 0 的数据
		var Warehousing []models2.AdditionsWarehousing
		d.Model(&models2.AdditionsWarehousing{}).Select("sn,id,combination_id").Where("sn IS NOT NULL AND (host_id = 0 OR host_id IS  NULL)  AND category_id = 1").Find(&Warehousing)
		snList := make([]string, 0)
		snBindMap := make(map[string]int, 0)
		assetBindCombinedMap := make(map[int]int, 0)
		for _, v := range Warehousing {

			if isDirty := global.BlackMap[v.Sn]; isDirty { //对于有问题的SN直接跳出
				continue
			}

			snList = append(snList, v.Sn)
			snBindMap[v.Sn] = v.Id
			if v.CombinationId > 0 {
				assetBindCombinedMap[v.Id] = v.CombinationId
			}
		}
		if len(snList) == 0 {
			continue
		}
		var HostList []models2.Host
		d.Model(&models2.Host{}).Select("id,sn,idc").Where("sn in ?", snList).Find(&HostList)

		updateAssetBindHost := make(map[int]int, 0)
		if len(HostList) == 0 {
			continue
		}

		hostBindIdc := make(map[int]int, 0)
		idcList := make([]int, 0)
		for _, v := range HostList {
			assetId, ok := snBindMap[v.Sn]
			if !ok {
				continue
			}
			updateAssetBindHost[assetId] = v.Id
			hostBindIdc[v.Id] = v.Idc
			idcList = append(idcList, v.Idc)
		}
		idcList = utils.RemoveRepeatInt(idcList)

		var IdcListHost []models2.Idc
		d.Model(&models2.Idc{}).Where("id in ?", idcList).Find(&IdcListHost)
		idcBindCustom := make(map[int]int, 0)
		for _, v := range IdcListHost {
			idcBindCustom[v.Id] = v.CustomId
		}
		//需要完善一点就是。如果自动绑定关系,那就是如果这个主机关联了IDC，IDC关联了客户,那这个资产也就是这个客户了。
		//因为是自动化配比组合，需要自动化关联客户。 因为手动出库的时候 就需要选客户
		for assetId, hostId := range updateAssetBindHost {
			d.Model(&models2.AdditionsWarehousing{}).Where("id = ?", assetId).Updates(map[string]interface{}{
				"host_id": hostId,
			})

			if CombinedId, ok := assetBindCombinedMap[assetId]; ok {

				hostIdcId := hostBindIdc[hostId] //拿主机Id获取关联的IDC

				CustomId := idcBindCustom[hostIdcId] //idcID 获取关联的客户
				d.Model(&models2.Combination{}).Where("id = ?", CombinedId).Updates(map[string]interface{}{
					"host_id":   hostId,
					"idc_id":    hostIdcId,
					"custom_id": CustomId,
				})
				d.Model(&models2.AdditionsWarehousing{}).Where("combination_id = ?", CombinedId).Updates(map[string]interface{}{
					"host_id": hostId,
				})

			}
		}

	}

}
