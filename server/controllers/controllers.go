// 控制器层  实现路由的处理逻辑
package controllers

import (
	"errors"
	"server/config"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type BasicInfo struct {
	ClusterId string `json:"clusterId" form:"clusterId"`
	NameSpace string `json:"nameSpace" form:"nameSpace"`
}

func BasicInit(c *gin.Context) (clientSet *kubernetes.Clientset, basicInfo BasicInfo, err error) {
	// 绑定前端传递的数据到后端变量中
	basicInfo = BasicInfo{}
	requestMethod := c.Request.Method
	switch requestMethod {
	case "GET":
		err = c.ShouldBindQuery(&basicInfo)
	case "POST":
		err = c.ShouldBindJSON(&basicInfo)
	default:
		err = errors.New("不支持的请求类型")
	}
	if basicInfo.NameSpace == "" {
		basicInfo.NameSpace = "default"
	}
	if err != nil {
		msg := "请求数据绑定后端失败: " + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	// 实例化clientSet
	kubeConfig := config.ClusterKubeconfig[basicInfo.ClusterId]
	kubeConfigByte := []byte(kubeConfig)
	// kubeConfig
	restConfig, err := clientcmd.RESTConfigFromKubeConfig(kubeConfigByte)
	if err != nil {
		msg := "实例化 kubeconfig 失败:" + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	// clientSet
	clientSet, err = kubernetes.NewForConfig(restConfig)
	if err != nil {
		msg := "实例化clientSet失败:" + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	// 校验 kubeconfig
	serverVersion, err := clientSet.Discovery().ServerVersion()
	if err != nil {
		msg := "kubeconfig 校验失败: " + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	logs.Debug(nil, "集群版本："+serverVersion.String())
	return clientSet, basicInfo, nil
}
