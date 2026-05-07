// 控制器层  实现路由的处理逻辑
package controllers

import (
	"errors"
	"server/config"
	"server/utils/logs"
	"server/utils/metrics"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type BasicInfo struct {
	ClusterId  string   `json:"clusterId" form:"clusterId"`
	NameSpace  string   `json:"nameSpace" form:"nameSpace"`
	Name       string   `json:"name" form:"name"`
	NameList   []string `json:"nameList" form:"nameList"`
	Item       any      `json:"item"`
	DeleteList []string `json:"deleteList"`
}

type Info struct {
	BasicInfo
	ReturnData    config.ReturnData
	LabelSelector string `json:"labelSelector" form:"labelSelector"`
	FieldSelector string `json:"fieldSelector" form:"fieldSelector"`
	// 判断是否是强制删除
	Force bool `json:"force" form:"force"`
}

func NewInfo(c *gin.Context, info *Info) (kubeconfig string) {
	var (
		err        error
		returnData config.ReturnData
		// 赋值变量部分
		requestMethod = c.Request.Method
	)
	// 前后端数据绑定
	switch requestMethod {
	case "GET":
		err = c.ShouldBindQuery(&info)
	case "POST":
		err = c.ShouldBindJSON(&info)
	default:
		err = errors.New("不支持的请求类型")
	}
	logs.Debug(map[string]any{"Info": info}, "绑定前端传递的数据到后端info变量中的结果")
	if err != nil {
		msg := "请求数据绑定后端失败: " + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(200, returnData)
		logs.Error(nil, msg)
		return
	}
	// 获取kubeconfig
	return config.ClusterKubeconfig[info.ClusterId]
}

func BasicInit(c *gin.Context, item any) (clientSet *kubernetes.Clientset, basicInfo BasicInfo, err error) {
	// 绑定前端传递的数据到后端变量中
	basicInfo = BasicInfo{}
	basicInfo.Item = item
	requestMethod := c.Request.Method
	switch requestMethod {
	case "GET":
		err = c.ShouldBindQuery(&basicInfo)
	case "POST":
		err = c.ShouldBindJSON(&basicInfo)
	default:
		err = errors.New("不支持的请求类型")
	}
	if err != nil {
		msg := "请求数据绑定后端失败: " + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	if basicInfo.NameSpace == "" {
		basicInfo.NameSpace = "default"
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

func errorReturnData(c *gin.Context, returndata config.ReturnData, err error, info Info, resourceType string, msg string) {
	logs.Error(map[string]any{"ERROR": err.Error(), "clusterId": info.ClusterId, "namespace": info.NameSpace, "resourceType": resourceType, "resourceName": info.Name}, msg)
	returndata.Status = 400
	returndata.Message = msg + ": " + err.Error()
	c.JSON(200, returndata)
}

func successReturnData(c *gin.Context, returndata config.ReturnData, info Info, resourceType string, msg string) {
	logs.Debug(map[string]any{"clusterId": info.ClusterId, "namespace": info.NameSpace, "resourceType": resourceType, "resourceName": info.Name}, msg)
	returndata.Status = 200
	returndata.Message = msg
	c.JSON(200, returndata)
}

// 入口函数
func KubectlFunc(c *gin.Context, resourceType string, opMethod string) {
	logs.Debug(nil, "操作资源类型: "+resourceType)
	// 定义变量
	var (
		kubeUtilser        kubeutils.KubeUtilser
		metricUtilser      metrics.MetricsUtilser
		returndata         config.ReturnData
		info               Info
		gracePeriodSeconds int64
		// 资源类型
		node        corev1.Node
		namespace   corev1.Namespace
		pod         corev1.Pod
		deployment  appsv1.Deployment
		statefulSet appsv1.StatefulSet
		daemonSet   appsv1.DaemonSet
		// 存储资源
		configMap corev1.ConfigMap
		secret    corev1.Secret
		pvc       corev1.PersistentVolumeClaim
		// 服务发布
		service corev1.Service
		// 任务
		cronJob batchv1.CronJob
	)
	returndata.Data = map[string]any{}
	// 初始化 Item 类型
	switch resourceType {
	case "node":
		info.Item = &node
	case "namespace":
		info.Item = &namespace
	case "pod":
		info.Item = &pod
	case "deployment":
		info.Item = &deployment
	case "statefulSet":
		info.Item = &statefulSet
	case "daemonSet":
		info.Item = &daemonSet
	// 存储资源
	case "configMap":
		info.Item = &configMap
	case "secret":
		info.Item = &secret
	case "pvc":
		info.Item = &pvc
	// 服务发布
	case "service":
		info.Item = &service
	// 任务
	case "cronJob":
		info.Item = &cronJob
	default:
		logs.Error(nil, "不支持该资源类型")
		return
	}
	// 绑定前端传递的数据到后端info变量中
	kubeconfig := NewInfo(c, &info)
	// 匹配资源类型 构建接口需要得结构体
	switch resourceType {
	case "node":
		kubeUtilser = kubeutils.NewNode(kubeconfig, &node)
		metricUtilser = metrics.NewNode(kubeconfig, &node)
	case "namespace":
		kubeUtilser = kubeutils.NewNamespace(kubeconfig, &namespace)
	case "pod":
		kubeUtilser = kubeutils.NewPod(kubeconfig, &pod)
	case "deployment":
		kubeUtilser = kubeutils.NewDeployment(kubeconfig, &deployment)
	case "statefulSet":
		kubeUtilser = kubeutils.NewStatefulSet(kubeconfig, &statefulSet)
	case "daemonSet":
		kubeUtilser = kubeutils.NewDaemonSet(kubeconfig, &daemonSet)
	// 存储资源
	case "configMap":
		kubeUtilser = kubeutils.NewConfigMap(kubeconfig, &configMap)
	case "secret":
		kubeUtilser = kubeutils.NewSecret(kubeconfig, &secret)
	case "pvc":
		kubeUtilser = kubeutils.NewPersistentVolumeClaim(kubeconfig, &pvc)
	// 服务发布
	case "service":
		kubeUtilser = kubeutils.NewService(kubeconfig, &service)
	// 任务
	case "cronJob":
		kubeUtilser = kubeutils.NewCronJob(kubeconfig, &cronJob)
	default:
		logs.Error(nil, "不支持该资源类型")
		return
	}
	// 匹配操作方法
	switch opMethod {
	case "create":
		err := kubeUtilser.Create(info.NameSpace)
		if err != nil {
			errorReturnData(c, returndata, err, info, resourceType, "创建失败")
			return
		}
	case "delete":
		if info.Force {
			gracePeriodSeconds = 1
		}
		err := kubeUtilser.Delete(info.NameSpace, info.Name, &gracePeriodSeconds)
		if err != nil {
			errorReturnData(c, returndata, err, info, resourceType, "删除失败")
			return
		}
	case "deleteList":
		if info.Force {
			gracePeriodSeconds = 1
		}
		err := kubeUtilser.DeleteList(info.NameSpace, info.NameList, &gracePeriodSeconds)
		if err != nil {
			errorReturnData(c, returndata, err, info, resourceType, "删除失败")
			return
		}
	case "list":
		items, err := kubeUtilser.List(info.NameSpace, info.LabelSelector, info.FieldSelector)
		if err != nil {
			errorReturnData(c, returndata, err, info, resourceType, "获取列表失败")
			return
		}
		returndata.Data["items"] = items
	case "get":
		item, err := kubeUtilser.Get(info.NameSpace, info.Name)
		if err != nil {
			errorReturnData(c, returndata, err, info, resourceType, "获取详情失败")
			return
		}
		returndata.Data["items"] = item
	case "update":
		err := kubeUtilser.Update(info.NameSpace)
		if err != nil {
			errorReturnData(c, returndata, err, info, resourceType, "更新失败")
			return
		}
	case "metrics":
		items, err := metricUtilser.Metrics(info.NameSpace)
		if err != nil {
			errorReturnData(c, returndata, err, info, resourceType, "获取资源使用指标失败")
			return
		}
		returndata.Data["items"] = items
	default:
		logs.Error(nil, "不支持该操作方法")
		return
	}
	successReturnData(c, returndata, info, resourceType, "操作成功")
}
