package utils

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
)

func InitDb(dsn, dblog string) (engine *xorm.Engine, err error) {
	// engine, err = xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:6379)/test_db_name?charset=utf8")
	engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		return
	}

	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LogLevel(core.LOG_INFO))

	f := NewXormLogger(dblog)
	if f == nil {
		panic("Init db log failed")
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))
	if err != nil {
		engine.Logger().Error(err.Error())
		return
	} else {
		engine.Logger().Info("New Enging Ok")
	}

	//校验连接
	err = engine.Ping()
	if err != nil {
		engine.Logger().Error(err.Error())
		return
	} else {
		engine.Logger().Info("Ping Ok")
	}

	return
}

