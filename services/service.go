package services

import (
	"fmt"
	"time"

	// "github.com/micro/go-grpc"
	// "github.com/micro/go-plugins/registry/etcd"
	// "github.com/micro/go-plugins/registry/kubernetes"
	// k8s "github.com/micro/kubernetes/go/micro"
	// static selector offloads load balancing to k8s services
	// note: requires user to create k8s services
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"vc.svc/models"
	sequenceService "vc.svc/services/sequence"
)

func Init(config models.MicroConfig) {
	service := grpc.NewService(
		micro.Name(config.Name),
		// micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
	service.Init()
	sequenceService.Register(service)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
