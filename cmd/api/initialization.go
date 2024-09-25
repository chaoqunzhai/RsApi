package api

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/global"
)

func InitializationMap() {
	dbList := sdk.Runtime.GetDb()
	//
	fmt.Println("加载全部的用户列表")
	for _, d := range dbList {
		global.UserDatMap.MakeAllCacheHost(d)

	}
}
