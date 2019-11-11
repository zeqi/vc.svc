package services

import (
	"context"
	"fmt"
	"sync"
	"time"

	// "github.com/micro/go-grpc"
	// "github.com/micro/go-plugins/registry/etcd"
	// "github.com/micro/go-plugins/registry/kubernetes"
	// k8s "github.com/micro/kubernetes/go/micro"
	// static selector offloads load balancing to k8s services
	// note: requires user to create k8s services
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/broker/kafka"
	"vc.svc/models"
	sequenceService "vc.svc/services/sequence"
)

var once sync.Once

type registrar struct {
}

func Init(config models.MicroConfig) {
	once.Do(func() {
		r := registrar{}
		r.init(config)
	})
}

func (r *registrar) init(config models.MicroConfig) {

	service := grpc.NewService(
		micro.Name(config.Name),
		micro.Version(config.Version),
		micro.Broker(kafka.NewBroker(func(o *broker.Options) {
			o.Addrs = []string{0: "127.0.0.1:9092"}
			// o.Addrs = []string{0: "redis://localhost:6379/0"}
		})),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.BeforeStart(r.beforeStart),
		micro.BeforeStop(r.beforeStop),
		micro.AfterStart(r.afterStart),
		micro.AfterStop(r.afterStop),
		micro.WrapHandler(r.logWrapper),
	)

	broker.Init()
	broker.Connect()

	broker.Subscribe("endpoint", func(p broker.Publication) error {
		// brokerHeader := p.Message().Header
		// aaa := brokerHeader["AAA"]
		fmt.Printf("[%v] topic endpoint value: %s\n", time.Now(), string(p.Message().Body))
		return nil
	})

	service.Init()
	sequenceService.Register(service)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func (r *registrar) beforeStart() error {
	fmt.Println("BeforeStart")
	return nil
}
func (r *registrar) beforeStop() error {
	fmt.Println("BeforeStop")
	return nil
}
func (r *registrar) afterStart() error {
	fmt.Println("AfterStart")
	return nil
}
func (r *registrar) afterStop() error {
	fmt.Println("AfterStop")
	return nil
}

func (r *registrar) logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		// fmt.Printf("[%v] server request endpoint: %s header: %s body: %s \n", time.Now(), req.Endpoint(), req.Header(), req.Body())
		broker.Publish("endpoint", &broker.Message{
			Header: map[string]string{},
			Body:   []byte(req.Endpoint()),
		})
		return fn(ctx, req, rsp)
	}
}
