package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server/grpc"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"
	"micro_learn/tutorials/learn1/user_server/basic"
	"micro_learn/tutorials/learn1/user_server/handler"
	user "micro_learn/tutorials/learn1/user_server/proto/user"
)

func main() {
	// init mysql
	basic.Init("")

	// register
	etcdEndpoints := basic.GetEtdcConfig()
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = etcdEndpoints
	})


	// New Service
	service := micro.NewService(
		micro.Server(grpc.NewServer()),
		micro.Name("smtl.micro.learn.srv.user"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}


}
