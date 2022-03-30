package main

import (
	"jhr.com/datarelay/global"
	"jhr.com/datarelay/initialize"
	"jhr.com/datarelay/web"
)

func main() {
	initialize.InitLogger()
	// 初始化配置文件
	initialize.InitConfig()
	// 初始化数据源配置
	initialize.InitDataSourceConnectionInfo()
	// 初始化服务发现客户端
	initialize.InitNamingClient()
	// 根据配置动态注册路由
	r := web.RegisterRouter()
	// 测试服务
	if global.ServerConfig.Test {
		web.TestServer(r)
		return
	}
	// 启动服务端
	web.RunServer(r)
}
