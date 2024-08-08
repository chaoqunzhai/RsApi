package global

import "time"

const (
	RoleSuper = 80 //超管
	RoleAdmin = 81 //子管理员
	RoleUser  = 82 //用户

	HostLoading = 0
	HostSuccess = 1
	HostOffline = -1
)

var (
	LOC, _ = time.LoadLocation("Asia/Shanghai")
)
