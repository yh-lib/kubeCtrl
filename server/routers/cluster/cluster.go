package cluster

import (
	"server/controllers/cluster"

	"github.com/gin-gonic/gin"
)

func add(clusterGroup *gin.RouterGroup) {
	clusterGroup.POST("/add", cluster.Add)
}
func delete(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/delete", cluster.Delete)
}
func update(clusterGroup *gin.RouterGroup) {
	clusterGroup.POST("/update", cluster.Update)
}
func get(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/get", cluster.Get)
}
func list(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/list", cluster.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置登录功能的路由策略
	clusterGroup := g.Group("cluster")
	add(clusterGroup)
	delete(clusterGroup)
	update(clusterGroup)
	get(clusterGroup)
	list(clusterGroup)
}
