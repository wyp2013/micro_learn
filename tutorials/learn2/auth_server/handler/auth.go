package handler

import (
	"context"
	"micro_learn/tutorials/learn2/auth_server/models/access"
	"strconv"

	"github.com/micro/go-micro/util/log"

	auth "micro_learn/tutorials/learn2/auth_server/proto/auth"
)

var tokenService access.Service

func Init() {
	var err error

	tokenService, err = access.GetService()
	if err != nil {
		panic(err)
	}
}


type Auth struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Auth) MakeAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Log("Received Auth.Call request")

	token, err := tokenService.MakeAccessToken(&access.Subject{
		ID: strconv.FormatUint(req.UserId, 10),
		Name: req.UserName,
	})
	if err != nil {
		rsp.Success = false
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Logf("[MakeAccessToken] token生成失败，err：%s", err)
		return err
	}

	rsp.Token = token
	rsp.Success = true

	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Auth) DelUserAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Log("Received Auth.Call request")
	err := tokenService.DelUserAccessToken(req.Token)
	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Logf("[DelUserAccessToken] 清除用户token失败，err：%s", err)
		return err
	}

	return nil
}






