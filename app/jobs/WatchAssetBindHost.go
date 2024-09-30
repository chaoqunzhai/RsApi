package jobs

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk"
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/common/utils"
	"time"
)

// InitJob
// 需要将定义的struct 添加到字典中；
// 字典 key 可以配置到 自动任务 调用目标 中；
func InitJob() {
	jobList = map[string]JobExec{
		"ExamplesOne":        ExamplesOne{},
		"WatchAssetBindHost": WatchAssetBindHost{},
		// ...
	}
}

// ExamplesOne
// 新添加的job 必须按照以下格式定义，并实现Exec函数
type ExamplesOne struct {
}

func (t ExamplesOne) Exec(arg interface{}) error {
	str := time.Now().Format(timeFormat) + " [INFO] JobCore ExamplesOne exec success"
	// TODO: 这里需要注意 Examples 传入参数是 string 所以 arg.(string)；请根据对应的类型进行转化；
	switch arg.(type) {

	case string:
		if arg.(string) != "" {
			fmt.Println("string", arg.(string))
			fmt.Println(str, arg.(string))
		} else {
			fmt.Println("arg is nil")
			fmt.Println(str, "arg is nil")
		}
		break
	}

	return nil
}

type WatchAssetBindHost struct {
}

func (t WatchAssetBindHost) Exec(arg interface{}) error {
	dbList := sdk.Runtime.GetDb()

	for _, d := range dbList {
		//只查询 资产分类为主机 + host_id = 0 的数据
		var Warehousing []models2.AdditionsWarehousing
		d.Model(&models2.AdditionsWarehousing{}).Select("sn,id,combination_id").Where("sn IS NOT NULL AND (host_id = 0 OR host_id IS  NULL)  AND category_id = 1").Find(&Warehousing)
		snList := make([]string, 0)
		snBindMap := make(map[string]int, 0)
		assetBindCombinedMap := make(map[int]int, 0)
		for _, v := range Warehousing {
			snList = append(snList, v.Sn)
			snBindMap[v.Sn] = v.Id
			if v.CombinationId > 0 {
				assetBindCombinedMap[v.Id] = v.CombinationId
			}
		}
		if len(snList) == 0 {
			return nil
		}
		var HostList []models2.Host
		d.Model(&models2.Host{}).Select("id,sn,idc").Where("sn in ?", snList).Find(&HostList)

		updateAssetBindHost := make(map[int]int, 0)
		if len(HostList) == 0 {
			return nil
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
					"host_id":   hostId,
					"status":    3,
					"idc_id":    hostIdcId,
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
		d.Model(&models2.Combination{}).Where("idc_id > 0 and custom_id  = 0 ").Find(&CombinationList)

		idcList := make([]int, 0)
		CombinationBindIdc := make(map[int]int, 0)
		for _, row := range CombinationList {
			idcList = append(idcList, row.IdcId)
			CombinationBindIdc[row.IdcId] = row.Id
		}
		idcList = utils.RemoveRepeatInt(idcList)

		var IdcListHost []models2.Idc
		d.Model(&models2.Idc{}).Where("id in ?", idcList).Find(&IdcListHost)

		for _, v := range IdcListHost {
			CombinationId, ok := CombinationBindIdc[v.Id]
			if !ok {
				continue
			}
			d.Model(&models2.Combination{}).Where("id = ?", CombinationId).Updates(map[string]interface{}{
				"custom_id": v.CustomId,
			})
		}

	}
	return nil
}
