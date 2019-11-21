package client

import (
	"micro_learn/micro/go-micro/client"
	"micro_learn/micro/go-micro/client/grpc"
	user "micro_learn/tutorials/learn1/user_server/proto/user"

)

var userService user.UserService

// 本列子使用grpc通信
func Init(opts ...client.Option) {
	rpcClient := grpc.NewClient(opts...)
	userService = user.NewUserService("smtl.micro.learn.srv.user", rpcClient)
}

func GetUserService() user.UserService {
	return userService
}


