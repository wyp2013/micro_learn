package main

import (
	"flag"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"
	"micro_learn/tutorials/learn2/basic"
	"micro_learn/tutorials/learn2/user_server/handler"
	user "micro_learn/tutorials/learn2/user_server/proto/user"
	"os"
)

func main() {
	flagSet := flag.NewFlagSet("config-load", flag.ExitOnError)
	confPath := flagSet.String("conf", "", "conf")
	logPath := flagSet.String("log", "", "log")
	flagSet.Parse(os.Args[1:])
	fmt.Println(*confPath, *logPath)

	// init mysql
	basic.Init(*confPath, *logPath)

	// register
	etcdEndpoints := basic.GetEtdcConfig()
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = etcdEndpoints
	})

	// New Service
	service := micro.NewService(
		micro.Name("smtl.micro.learn.srv.user"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.Address(":58090"),
		micro.Flags(
			cli.StringFlag{
				Name:   "conf",
				EnvVar: "./config/config.yaml",
				Usage:  "This is a config paht flag",
			},
			cli.StringFlag{
				Name:   "log",
				EnvVar: "./../log/xorm",
				Usage:  "This is a log path flag",
			},
		),
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
