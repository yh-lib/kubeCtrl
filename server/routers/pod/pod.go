package pod

import (
	"server/controllers/pod"

	"github.com/gin-gonic/gin"
)

func create(podGroup *gin.RouterGroup) {
	podGroup.POST("/create", pod.Create)
}
func delete(podGroup *gin.RouterGroup) {
	podGroup.POST("/delete", pod.Delete)
}
func deleteList(podGroup *gin.RouterGroup) {
	podGroup.POST("/deleteList", pod.DeleteList)
}
func update(podGroup *gin.RouterGroup) {
	podGroup.POST("/update", pod.Update)
}
func get(podGroup *gin.RouterGroup) {
	podGroup.GET("/get", pod.Get)
}
func list(podGroup *gin.RouterGroup) {
	podGroup.GET("/list", pod.List)
}
func status(podGroup *gin.RouterGroup) {
	podGroup.GET("/status", pod.Status)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置登录功能的路由策略
	podGroup := g.Group("pod")
	create(podGroup)
	delete(podGroup)
	deleteList(podGroup)
	update(podGroup)
	get(podGroup)
	list(podGroup)
	status(podGroup)
}
