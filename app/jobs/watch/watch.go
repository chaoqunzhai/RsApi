package watch

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/jakecoffman/cron"
)

func RunCrontab() {
	c := cron.New()
	if config.ApplicationConfig.Mode != "test" {
		c.AddFunc("@every 3m", WatchAssetBindHost, "巡检资产列表和CMDB关联关系")
		//c.AddFunc("@every 20m", WatchHostAndAssetStatus, "在线/离线 机器状态更新")
		c.AddFunc("@every 6h", WatchOnlineUsage, "在线机器利用率检测")
	}
	c.Start()
	fmt.Println("增加cron成功")
}
