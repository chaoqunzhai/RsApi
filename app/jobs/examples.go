package jobs

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk"
	models2 "go-admin/cmd/migrate/migration/models"
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
		d.Model(&models2.Host{}).Select("id,sn").Where("sn in ?", snList).Find(&HostList)

		updateAssetBindHost := make(map[int]int, 0)
		if len(HostList) == 0 {
			return nil
		}
		for _, v := range HostList {
			assetId, ok := snBindMap[v.Sn]
			if !ok {
				continue
			}
			updateAssetBindHost[assetId] = v.Id
		}
		for assetId, hostId := range updateAssetBindHost {
			d.Model(&models2.AdditionsWarehousing{}).Where("id = ?", assetId).Updates(map[string]interface{}{
				"host_id": hostId,
			})

			if CombinedId, ok := assetBindCombinedMap[assetId]; ok {
				d.Model(&models2.Combination{}).Where("id = ?", CombinedId).Updates(map[string]interface{}{
					"host_id": hostId,
				})
			}
		}

	}
	return nil
}
