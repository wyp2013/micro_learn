package client

import (
	"github.com/micro/go-micro/client"
	auth "micro_learn/tutorials/learn2/auth_server/proto/auth"
	user "micro_learn/tutorials/learn2/user_server/proto/user"
)

var userService user.UserService
var authService auth.AuthService

func Init(opts ...client.Option) {
	rpcClient := client.NewClient(opts...)
	userService = user.NewUserService("smtl.micro.learn.srv.user", rpcClient)
	authService = auth.NewAuthService("smtl.micro.learn.srv.auth", rpcClient)
}

func GetUserService() user.UserService {
	return userService
}

func GetAuthService() auth.AuthService {
	return authService
}
