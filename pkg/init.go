package pkg

import (
	"github.com/candy44777/k8s-dashboard/pkg/namespace"
)

func Informer() {
	// 初始化 namespace Informer
	ns := namespace.NewConfig()
	ns.InitInformer()
}
