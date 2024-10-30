package watch

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/jakecoffman/cron"
)

func RunCrontab() {
	c := cron.New()
	if config.ApplicationConfig.Mode != "test" {
		c.AddFunc("@every 2m", WatchHostBandwidth, "重算总带宽")
		c.AddFunc("@every 3m", WatchAssetBindHost, "巡检资产列表和CMDB关联关系")
		c.AddFunc("@every 4m", WatchCombinationHost, "巡检资产组合空主机数据")
		c.AddFunc("@every 5m", WatchCombinationCustom, "巡检组合的客户")
	}
	c.Start()
	fmt.Println("增加cron成功")
}
