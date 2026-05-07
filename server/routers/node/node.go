package node

import (
	"server/controllers/node"

	"github.com/gin-gonic/gin"
)

func create(nodeGroup *gin.RouterGroup) {
	nodeGroup.POST("/create", node.Create)
}
func delete(nodeGroup *gin.RouterGroup) {
	nodeGroup.GET("/delete", node.Delete)
}
func update(nodeGroup *gin.RouterGroup) {
	nodeGroup.POST("/update", node.Update)
}
func get(nodeGroup *gin.RouterGroup) {
	nodeGroup.GET("/get", node.Get)
}
func list(nodeGroup *gin.RouterGroup) {
	nodeGroup.GET("/list", node.List)
}
func metrics(nodeGroup *gin.RouterGroup) {
	nodeGroup.GET("/metrics", node.Metrics)
}
func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置登录功能的路由策略
	nodeGroup := g.Group("node")
	create(nodeGroup)
	delete(nodeGroup)
	update(nodeGroup)
	get(nodeGroup)
	list(nodeGroup)
	metrics(nodeGroup)
}
