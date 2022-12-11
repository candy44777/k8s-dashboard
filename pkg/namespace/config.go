package namespace

import (
	"github.com/candy44777/k8s-dashboard/conf"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
)

type Config struct {
	Handler *Handler
}

func NewConfig() *Config {
	return &Config{
		Handler: NewHandler(),
	}
}

// InitInformer 初始化 informer
func (c *Config) InitInformer() informers.SharedInformerFactory {
	factory := informers.NewSharedInformerFactory(conf.K8sCLientset, 0)
	factory.Core().V1().Namespaces().Informer().AddEventHandler(c.Handler)
	factory.Start(wait.NeverStop)
	return factory
}
