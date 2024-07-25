package main

import (
	"go-admin/cmd"
)

//go:generate swag init --parseDependency --parseDepth=6 --instanceName admin -o ./docs/admin

// @title go-admin API
// @version 2.0.0
// @description 基于Gin + Vue + Ant UI的前后端分离权限管理系统的接口文档
// @license.name MIT

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
