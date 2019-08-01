package handler

import (
	"context"
	"fmt"
	"micro_learn/tutorials/learn2/user_server/models"
	user "micro_learn/tutorials/learn2/user_server/proto/user"
)

type User struct {
}

func (h *User) QueryUserByName(ctx context.Context, req *user.Request, rsp *user.Response) error {
	fmt.Println(req.GetUserName())
	userModel := models.NewUserModle()
	qUser, err := userModel.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Error = &user.Error{Code: 1, Detail: err.Error()}
		return err
	}

	if qUser == nil {
		rsp.Error = &user.Error{Code: 2, Detail: "not found user"}
		return nil
	}

	rsp.User = &user.User{
		Id:          int64(qUser.UserId),
		Name:        qUser.UserName,
		Pwd:         qUser.Pwd,
		CreatedTime: uint64(qUser.CreatedTime.Unix()),
		UpdatedTime: uint64(qUser.UpdatedTime.Unix()),
	}

	return nil
}
