package main

import (
	"github.com/candy44777/k8s-dashboard/conf"
	"github.com/candy44777/k8s-dashboard/pkg"
	"github.com/candy44777/k8s-dashboard/pkg/namespace"
	"github.com/candy44777/k8s-dashboard/protocol"
	"log"
)

func main() {
	if err := conf.LoadConfigFromToml("./etc/config.toml"); err != nil {
		log.Fatal(err)
	}

	conf.InitClientSet()
	pkg.Informer()

	protocol.NewHttpServer().Mount("v1", namespace.NewCtl()).Run(":9999")
}
