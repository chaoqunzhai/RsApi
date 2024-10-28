package global

import (
	"go-admin/app/admin/models"
	"gorm.io/gorm"
	"sync"
)

type MapEvent struct {
	sync.RWMutex
	M map[int]*models.SysUser
}

var (
	UserDatMap = &MapEvent{M: make(map[int]*models.SysUser)}
)

func (e *MapEvent) MakeAllCacheHost(orm *gorm.DB) {
	e.RLock()
	defer e.RUnlock()
	var sysUserList []*models.SysUser
	orm.Model(&models.SysUser{}).Find(&sysUserList)
	for _, sysUser := range sysUserList {

		e.M[sysUser.UserId] = sysUser
	}
}

func (e *MapEvent) Get(userId int) (*models.SysUser, bool) {
	e.RLock()
	defer e.RUnlock()
	event, exists := e.M[userId]
	return event, exists
}
func (e *MapEvent) Set(userId int, event *models.SysUser) {
	e.Lock()
	defer e.Unlock()
	e.M[userId] = event
}

func (e *MapEvent) Delete(userId int) {
	e.Lock()
	defer e.Unlock()

	delete(e.M, userId)
}

// 资产和主机状态相互置换
func AssetToHostStatus(v int) int {
	var updateHost int
	switch v {
	case 1: //在库
		updateHost = 0
	case 2: //出库 - 主机待上线
		updateHost = 3
	case 3: //在线 - 主机在线
		updateHost = 1
	case 4: //下架 - 主机下架
		updateHost = 2
	case 5: //闲置 - 主机连接中
		updateHost = 0
	case 6: //离线 - 主机离线
		updateHost = -1
	}
	return updateHost
}

// 主机状态和资产 相互置换
func HostToAssetStatus(v int) int {
	var updateHost int
	switch v {
	case 0: //链接中
		updateHost = 3
	case 1: //在线 - 资产在线
		updateHost = 3
	case -1: //离线 - 资产离线
		updateHost = 6
	case 2: //下架 - 资产下架
		updateHost = 4
	case 3:
		updateHost = 2

	}
	return updateHost
}
