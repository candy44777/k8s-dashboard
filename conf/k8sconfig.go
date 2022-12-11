package conf

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var K8sCLientset *kubernetes.Clientset

func K8sConfig() *rest.Config {
	cfg, err := clientcmd.BuildConfigFromFlags("", C().K8s.ConfigPath)
	if err != nil {
		panic("k8s config failed")
	}

	return cfg
}

func K8sClientSet() *kubernetes.Clientset {
	clientSet, err := kubernetes.NewForConfig(K8sConfig())
	if err != nil {
		panic("k8s init client failed")
	}

	return clientSet
}

func InitClientSet() {
	if K8sCLientset == nil {
		K8sCLientset = K8sClientSet()
	}
}
