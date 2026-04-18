package pvc

import (
	"server/controllers/pvc"

	"github.com/gin-gonic/gin"
)

//	func create(pvcGroup *gin.RouterGroup) {
//		pvcGroup.POST("/create", pvc.Create)
//	}
//
//	func delete(pvcGroup *gin.RouterGroup) {
//		pvcGroup.POST("/delete", pvc.Delete)
//	}
//
//	func deleteList(pvcGroup *gin.RouterGroup) {
//		pvcGroup.POST("/deleteList", pvc.DeleteList)
//	}
//
//	func update(pvcGroup *gin.RouterGroup) {
//		pvcGroup.POST("/update", pvc.Update)
//	}
//
//	func get(pvcGroup *gin.RouterGroup) {
//		pvcGroup.GET("/get", pvc.Get)
//	}
func list(pvcGroup *gin.RouterGroup) {
	pvcGroup.GET("/list", pvc.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置登录功能的路由策略
	pvcGroup := g.Group("pvc")
	// create(pvcGroup)
	// delete(pvcGroup)
	// deleteList(pvcGroup)
	// update(pvcGroup)
	// get(pvcGroup)
	list(pvcGroup)
}
