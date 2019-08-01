package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"
	"micro_learn/tutorials/learn2/auth_server/handler"
	"micro_learn/tutorials/learn2/auth_server/models"
	"micro_learn/tutorials/learn2/basic"

	auth "micro_learn/tutorials/learn2/auth_server/proto/auth"
)

func main() {
	// init
	basic.Init("./config/config.yaml", "./../log/authxorm")

	// register
	etcdEndpoints := basic.GetEtdcConfig()
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = etcdEndpoints
	})

	// New Service
	service := micro.NewService(
		micro.Name("smtl.micro.learn.srv.auth"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化handler
			models.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// Register Handler
	auth.RegisterAuthHandler(service.Server(), new(handler.Auth))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
