package main

import (
	"micro_learn/micro/go-micro"
	"micro_learn/micro/go-micro/registry"
	"micro_learn/micro/go-micro/server/grpc"
	"micro_learn/micro/go-micro/util/log"
	"micro_learn/micro/go-plugins/registry/etcdv3"
	"micro_learn/tutorials/learn1/user_server/basic"
	"micro_learn/tutorials/learn1/user_server/handler"
	user "micro_learn/tutorials/learn1/user_server/proto/user"
	"time"
)

func main() {
	// init mysql
	basic.Init("")

	// register
	etcdEndpoints := basic.GetEtdcConfig()
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = etcdEndpoints
	})


	// New Service, use grpc as server
	service := micro.NewService(
		micro.Server(grpc.NewServer()),
		micro.Name("smtl.micro.learn.srv.user"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.Address(":2990"),
		micro.RegisterInterval(10*time.Second),
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
