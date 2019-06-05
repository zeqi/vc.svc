package services

import (
	"fmt"
	"time"

	"github.com/micro/go-grpc"
	// "github.com/micro/go-grpc"

	// "github.com/micro/go-plugins/registry/etcd"

	// "github.com/micro/go-plugins/registry/kubernetes"
	// k8s "github.com/micro/kubernetes/go/micro"
	// k8s "github.com/micro/kubernetes/go/micro"
	// static selector offloads load balancing to k8s services
	// note: requires user to create k8s services

	"vc.svc/models"
	sequenceService "vc.svc/services/sequence"

	"github.com/micro/go-micro"
)

func Init(config models.MicroConfig) {
	// adds := strings.Split(config.Etcd.Addrs, ",")
	// ros := registry.Addrs(adds...)
	// r := etcdv3.NewRegistry(func(op *registry.Options) {
	// 	op.Addrs = strings.Split(config.Etcd.Addrs, ",")
	// })
	service := grpc.NewService(
		micro.Name(config.Name),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	// micro.Registry(r),
	// micro.Selector(static.NewSelector()),
	)
	service.Init()
	sequenceService.Register(service)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
