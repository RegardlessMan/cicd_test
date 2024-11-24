package main

import (
	"cicd_test/config"
	"cicd_test/router"
)

func main() {
	//初始化配置
	config.InitConfig()
	//初始化路由
	engine := router.SetRouter()
	port := config.AppConfig.App.Port
	if port == "" {
		port = ":8080"
	}

	engine.Run(port)
}
