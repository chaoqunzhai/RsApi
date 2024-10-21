package watch

import (
	"github.com/go-admin-team/go-admin-core/sdk"
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/common/utils"
	"go-admin/global"
)

func WatchAssetBindHost() {
	dbList := sdk.Runtime.GetDb()

	for _, d := range dbList {
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
				"status":  3,
			})

			if CombinedId, ok := assetBindCombinedMap[assetId]; ok {

				hostIdcId := hostBindIdc[hostId] //拿主机Id获取关联的IDC

				CustomId := idcBindCustom[hostIdcId] //idcID 获取关联的客户
				d.Model(&models2.Combination{}).Where("id = ?", CombinedId).Updates(map[string]interface{}{
					"host_id": hostId,
					"status":  3,
					//"idc_id":    hostIdcId, 因为资产的数据 只是一周上报一次, 但是IDC是随着CMDB资产数据一直在变化 所以是由着下面的
					"custom_id": CustomId,
				})
				d.Model(&models2.AdditionsWarehousing{}).Where("combination_id = ?", CombinedId).Updates(map[string]interface{}{
					"host_id": hostId,
				})

			}
		}

	}

	for _, d := range dbList {
		var CombinationList []models2.Combination
		d.Model(&models2.Combination{}).Find(&CombinationList)

		hostIds := make([]int, 0)
		CombinationBindHost := make(map[int]int, 0)
		CombinationBindIdc := make(map[int]int, 0)
		CombinationHostBindIdc := make(map[int][]int, 0)
		for _, row := range CombinationList {
			hostIds = append(hostIds, row.HostId)
			CombinationBindHost[row.HostId] = row.Id
		}
		hostIds = utils.RemoveRepeatInt(hostIds)

		var hostListModel []models2.Host
		d.Model(&models2.Host{}).Select("id,idc").Where("id in ?", hostIds).Find(&hostListModel)
		//先进行资产组合关联的主机 查询到对应的IDC
		idcList := make([]int, 0)

		for _, hostRow := range hostListModel {
			CombinationId, ok := CombinationBindHost[hostRow.Id]
			if !ok {
				continue
			}
			d.Model(&models2.Combination{}).Where("id = ?", CombinationId).Updates(map[string]interface{}{
				"idc_id": hostRow.Idc,
			})
			idcList = append(idcList, hostRow.Idc)
			//资产和组合对应起来
			CombinationBindIdc[hostRow.Idc] = CombinationId

			bindHostList, toOk := CombinationHostBindIdc[hostRow.Idc]
			if !toOk {
				bindHostList = make([]int, 0)
			}
			bindHostList = append(bindHostList, hostRow.Id)
		}

		idcList = utils.RemoveRepeatInt(idcList)
		var IdcListHost []models2.Idc
		d.Model(&models2.Idc{}).Select("id,custom_id").Where("id in ?", idcList).Find(&IdcListHost)

		for _, v := range IdcListHost {
			CombinationId, ok := CombinationBindIdc[v.Id]
			if !ok {
				continue
			}
			if v.CustomId == 0 {
				continue
			}
			//更新资产组合关联的客户
			d.Model(&models2.Combination{}).Where("id = ?", CombinationId).Updates(map[string]interface{}{
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

	}

	for _, d := range dbList {
		var CombinationList []models2.Combination
		d.Model(&models2.Combination{}).Select("code,id").Where("host_id = 0 OR host_id IS  NULL ").Find(&CombinationList)

		for _, v := range CombinationList {

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

		}
	}

}
