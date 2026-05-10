package kubeutils

import (
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 定义kubeutils的接口
type KubeUtilser interface {
	// namespace, item
	Create(string) error
	// namespace, name, gracePeriodSeconds
	Delete(string, string, *int64) error
	// namespace, nameList, gracePeriodSeconds
	DeleteList(string, []string, *int64) error
	// namespace, item
	Update(string) error
	// namespace, labelSelector, fieldSelector
	List(string, string, string) (interface{}, error)
	// namespace name
	Get(string, string) (interface{}, error)
}

type ResourceInstance struct {
	Kubeconfig string
	// Namespace  string
	// Name       string
	// NameList   []string
	Clientset *kubernetes.Clientset
}

func (c *ResourceInstance) Init(kubeconfig string) {
	c.Kubeconfig = kubeconfig
	// c.Namespace = namespace
	// c.Name = name
	// // 可以是空的
	// c.NameList = nameList
	// 生成Clientset
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(c.Kubeconfig))
	if err != nil {
		msg := "解析kubeconfig错误:" + err.Error()
		panic(msg)
	}
	// 设置超时时间
	restConfig.Timeout = 3 * time.Second
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		msg := "创建clientset失败:" + err.Error()
		panic(msg)
	}
	c.Clientset = clientset

}
