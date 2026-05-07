package metrics

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// 定义kubeutils的接口
type MetricsUtilser interface {
	// namespace
	metrics(namespace string) (any, error)
}

// 定义结构体
type Node struct {
	InstanceInterface typedv1.CoreV1Interface
	Item              *corev1.Node
}

// var (
// 	node corev1.Node
// )

func (c *Node) metrics() {
	fmt.Println("执行metrics逻辑")
}

// // New函数可以用于配置一些默认值
// func NewNode(kubeconfig string, item *corev1.Node) *Node {
// 	// 首先调用instance的init函数，生成一个ResourceInstance的实例，并配置默认值和生成clientset
// 	instance := ResourceInstance{}
// 	instance.Init(kubeconfig)
// 	// 定义一个Node实例
// 	resource := Node{}
// 	resource.InstanceInterface = instance.Clientset.CoreV1()
// 	resource.Item = item
// 	return &resource
// }

// func (c *Node) Create(namespace string) error {
// 	logs.Info(map[string]interface{}{"名称": c.Item.Name}, "创建Node")
// 	c.Item.ResourceVersion = ""
// 	_, err := c.InstanceInterface.Nodes().Create(context.TODO(), c.Item, v1.CreateOptions{})
// 	return err
// }
