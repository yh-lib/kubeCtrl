package namespace

import (
	"context"
	"server/config"
	"server/controllers"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Create(c *gin.Context) {
	logs.Info(nil, "创建逻辑")
	returnData := config.ReturnData{}
	clientSet, basicInfo, err := controllers.BasicInit(c)
	var newNameSpace corev1.Namespace
	newNameSpace.Name = basicInfo.NameSpace
	if err != nil {
		logs.Error(map[string]any{"Error": err}, "clientSet 初始化失败")
	}
	_, err = clientSet.CoreV1().Namespaces().Create(context.TODO(), &newNameSpace, metav1.CreateOptions{})
	if err != nil {
		returnData.Status = 400
		returnData.Message = "创建 namespace: " + newNameSpace.Name + " 失败:" + err.Error()
	} else {
		returnData.Status = 200
		returnData.Message = "创建 namespace: " + newNameSpace.Name + " 成功:"
	}
	logs.Error(map[string]any{"Error": err}, "创建 namespace: "+newNameSpace.Name+" 失败:")
	c.JSON(200, returnData)
}
