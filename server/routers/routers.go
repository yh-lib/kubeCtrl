// 路由层 管理程序的路由信息
package routers

import (
	"server/routers/auth"
	"server/routers/cluster"
	"server/routers/configMap"
	"server/routers/cronJob"
	"server/routers/daemonSet"
	"server/routers/deployment"
	"server/routers/event"
	"server/routers/namespace"
	"server/routers/node"
	"server/routers/pod"
	"server/routers/pvc"
	"server/routers/secret"
	"server/routers/service"
	"server/routers/statefulSet"

	"github.com/gin-gonic/gin"
)

// 注册路由的方法
func RegistrerRouters(r *gin.Engine) {
	// 根路由
	apiGroup := r.Group("/api")
	// 认证
	auth.RegisterSubRouter(apiGroup)
	// 集群资源
	cluster.RegisterSubRouter(apiGroup)
	node.RegisterSubRouter(apiGroup)
	namespace.RegisterSubRouter(apiGroup)
	// pod
	pod.RegisterSubRouter(apiGroup)
	// 控制器资源
	deployment.RegisterSubRouter(apiGroup)
	statefulSet.RegisterSubRouter(apiGroup)
	daemonSet.RegisterSubRouter(apiGroup)
	// 任务
	cronJob.RegisterSubRouter(apiGroup)
	// 存储
	secret.RegisterSubRouter(apiGroup)
	configMap.RegisterSubRouter(apiGroup)
	pvc.RegisterSubRouter(apiGroup)
	// 网络
	service.RegisterSubRouter(apiGroup)
	// 事件
	event.RegisterSubRouter(apiGroup)
}
