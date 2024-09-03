package main

import (
	"flag"
	"fmt"
	"go-admin/cmd"
	"os"
)

//go:generate swag init --parseDependency --parseDepth=6 --instanceName admin -o ./docs/admin

// @title go-admin API
// @version 2.0.0
// @description 基于Gin + Vue + Ant UI的前后端分离权限管理系统的接口文档
// @license.name MIT

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
var (
	Version string
)

func main() {
	versionFlag := flag.Bool("version", false, "print the version")
	flag.Parse()
	if *versionFlag {
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}
	cmd.Execute()
}
