package client

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/grpc"
	user "micro_learn/tutorials/learn1/user_server/proto/user"

)

var userService user.UserService

func Init(opts ...client.Option) {
	rpcClient := grpc.NewClient(opts...)
	userService = user.NewUserService("smtl.micro.learn.srv.user", rpcClient)
}

func GetUserService() user.UserService {
	return userService
}


