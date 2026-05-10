package initcontroller

import (
	"context"
	"server/config"
	"server/utils/logs"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	newNamespace corev1.Namespace
)

func metaDataInit() {
	logs.Info(nil, "初始化程序配置...")
	// step1 初始化 kubeConfig 实例
	var (
		restConfig *rest.Config
		err        error
	)
	// 初始化kubeconfig
	if config.InCluster {
		// in-cluster
		restConfig, err = rest.InClusterConfig()
	} else {
		// out-cluster
		restConfig, err = clientcmd.BuildConfigFromFlags("", config.KubeConfigPath)
	}
	if err != nil {
		logs.Error(map[string]any{"Error": err.Error()}, "初始化 kubeconfig 实例失败")
		panic(err.Error())
	}
	// step2 创建客户端工具 clientSet
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		logs.Error(map[string]any{"Error": err.Error()}, "客户端工具 clientSet 创建失败")
		panic(err.Error())
	}
	// incluster集群管理客户端
	config.InClusterClientSet = clientset
	// step3 创建元数据 nameSpace
	newNamespace.Name = config.MetadataNamespace
	// 判断 nameSpace 是否已存在
	_, err = clientset.CoreV1().Namespaces().Get(context.TODO(), newNamespace.Name, metav1.GetOptions{})
	if err != nil {
		logs.Warning(nil, "namespace: "+newNamespace.Name+" 不存在")
		logs.Warning(nil, "开始创建 namespace: "+newNamespace.Name)
		_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &newNamespace, metav1.CreateOptions{})
		if err != nil {
			logs.Error(map[string]any{"Error": err.Error()}, "namespace: "+newNamespace.Name+" 创建失败")
			panic(err.Error())
		}
		logs.Warning(nil, "namespace: "+newNamespace.Name+" 创建成功")
	} else {
		logs.Warning(nil, "namespace: "+newNamespace.Name+" 已存在")
	}
	// 获取集群kubeconfig
	config.ClusterKubeconfig = make(map[string]string)
	clusterList, _ := config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).List(context.TODO(), metav1.ListOptions{})
	for _, secret := range clusterList.Items {
		clusterId := secret.Name
		kubeConfigStr := string(secret.Data["kubeconfig"])
		config.ClusterKubeconfig[clusterId] = kubeConfigStr
	}
	logs.Info(nil, "初始化程序配置已完成")
}
