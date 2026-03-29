package namespace

import (
	"server/controllers/namespace"

	"github.com/gin-gonic/gin"
)

func create(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.POST("/create", namespace.Create)
}
func delete(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.GET("/delete", namespace.Delete)
}
func deleteList(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.POST("/deleteList", namespace.DeleteList)
}
func update(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.POST("/update", namespace.Update)
}
func get(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.GET("/get", namespace.Get)
}
func list(namespaceGroup *gin.RouterGroup) {
	namespaceGroup.GET("/list", namespace.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置登录功能的路由策略
	namespaceGroup := g.Group("namespace")
	create(namespaceGroup)
	delete(namespaceGroup)
	deleteList(namespaceGroup)
	update(namespaceGroup)
	get(namespaceGroup)
	list(namespaceGroup)
}
