package models

import (
	"github.com/go-xorm/xorm"
	"micro_learn/tutorials/learn1/user_server/basic"
)

type UserModel struct {
	engine *xorm.Engine
}

func NewUserModle() *UserModel {
	engine := basic.GetDbEngine()

	return &UserModel{
		engine: engine,
	}
}

func (m *UserModel) QueryUserByName(name string) (*User, error) {
	var user User
	bFind, err := m.engine.Where("user_name = ?", name).Get(&user)
	if bFind == false {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}
