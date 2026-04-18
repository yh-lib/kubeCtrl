// 路由层 管理程序的路由信息
package routers

import (
	"server/routers/auth"
	"server/routers/cluster"
	"server/routers/deployment"
	"server/routers/namespace"
	"server/routers/node"
	"server/routers/pod"
	"server/routers/pvc"
	"server/routers/secret"

	"github.com/gin-gonic/gin"
)

// 注册路由的方法
func RegistrerRouters(r *gin.Engine) {
	// 1. 登录: /api/auth/login
	// 2. 退出: /api/auth/logout
	apiGroup := r.Group("/api")
	auth.RegisterSubRouter(apiGroup)
	cluster.RegisterSubRouter(apiGroup)
	node.RegisterSubRouter(apiGroup)
	namespace.RegisterSubRouter(apiGroup)
	pod.RegisterSubRouter(apiGroup)
	deployment.RegisterSubRouter(apiGroup)
	secret.RegisterSubRouter(apiGroup)
	pvc.RegisterSubRouter(apiGroup)
}
