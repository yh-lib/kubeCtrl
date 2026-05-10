package metrics

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1beta1"
)

// 定义kubeutils的接口
type MetricsUtilser interface {
	// namespace
	Metrics(namespace string) (any, error)
}

// 定义结构体
type Node struct {
	InstanceInterface v1beta1.NodeMetricsInterface
	Item              *corev1.Node
}

type ResourceInstance struct {
	Kubeconfig string
	Clientset  *versioned.Clientset
}

func (c *Node) Metrics(namespace string) (items any, err error) {
	list, err := c.InstanceInterface.List(context.TODO(), metav1.ListOptions{})
	items = list.Items
	return items, err
}

// 构造 Node 结构体
func NewNode(kubeconfig string, item *corev1.Node) *Node {
	// 构造ResourceInstance实例
	instance := ResourceInstance{}
	instance.Init(kubeconfig)
	// 定义一个Node实例
	resource := Node{}
	resource.InstanceInterface = instance.Clientset.MetricsV1beta1().NodeMetricses()
	resource.Item = item
	return &resource
}

func (c *ResourceInstance) Init(kubeconfig string) {
	// kubeconfig 赋值
	c.Kubeconfig = kubeconfig
	// 实例化 metrics clientSet
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(c.Kubeconfig))
	if err != nil {
		msg := "解析kubeconfig错误:" + err.Error()
		panic(msg)
	}
	// 设置超时时间
	restConfig.Timeout = 3 * time.Second
	clientset, err := versioned.NewForConfig(restConfig)
	if err != nil {
		msg := "创建metrics clientset失败:" + err.Error()
		panic(msg)
	}
	c.Clientset = clientset
}
