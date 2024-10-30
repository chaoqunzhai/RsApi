package global

import "time"

const (
	RoleSuper = 80 //超管
	RoleAdmin = 81 //子管理员
	RoleUser  = 82 //用户

	HostLoading = 0
	HostSuccess = 1  //在线
	HostOffline = -1 //离线
	HostRemove  = 2
	HostWait    = 3 //待上架
)

var (
	LOC, _ = time.LoadLocation("Asia/Shanghai")

	IspMap = map[int]string{
		1: "移动",
		2: "电信",
		3: "联通",
		4: "其他",
	}
	RsRemark = []string{
		"20000", "10000",
	}
)

// 黑名单的SN， 因为有些SN都是一样的，只能通过主机名来确定唯一性
var BlackMap = map[string]bool{
	"01234567890123456789AB": true,
	"Default string":         true,
}
