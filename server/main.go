// 项目的总入口
package main

import (
	"server/config"
	_ "server/controllers/initController"
	"server/middlerwares"
	"server/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化程序配置	package initController
	// 注册gin引擎
	r := gin.Default()
	// 开启跨域中间件
	r.Use(middlerwares.CorsMiddleware())
	// // 使用中间件鉴权
	r.Use(middlerwares.JwtAuth)
	// // 注册路由
	routers.RegistrerRouters(r)
	// // 启动程序
	r.Run(config.Port)
}
