package main

import (
	"github.com/micro/go-micro"
	mc "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"
	"micro_learn/tutorials/learn2/user_web/client"
	"net/http"

	"github.com/micro/go-micro/web"
	"micro_learn/tutorials/learn2/user_web/handler"
)

func main() {

	// register
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{"127.0.0.1:2379"}
	})

	client.Init(mc.Registry(reg))

	// create new web service
	service := web.NewService(
		web.Name("smtl.micro.learn.web.user"),
		web.Version("latest"),
		web.Registry(reg),
		web.Address(":8091"),
		web.MicroService(micro.NewService()),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/user/login", handler.UserLogIn)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
