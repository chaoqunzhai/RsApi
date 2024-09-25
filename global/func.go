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
