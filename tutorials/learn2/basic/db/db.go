package db

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"micro_learn/tutorials/learn2/basic/config"
	"micro_learn/tutorials/utils"
)

var engine *xorm.Engine

func Init(xormLogPath string) {
	sqlConf := config.GetMysqlConf()
	if (!sqlConf.Enabled) {
		return
	}

	if sqlConf == nil {
		panic("mysql 配置为空")
	}

	if len(xormLogPath) == 0 {
		xormLogPath = "./../../log/xrom"
	}

	var err error
	engine, err = utils.InitDb(sqlConf.URL, xormLogPath)
	if err != nil {
		panic(err.Error())
	}

	engine.SetMaxOpenConns(sqlConf.MaxOpenConnection)
	engine.SetMaxIdleConns(sqlConf.MaxIdleConnection)

	fmt.Println(xormLogPath)
}

func GetDbEngine() *xorm.Engine {
	if engine == nil {
		panic("mysql is not initial")
	}

	return engine
}


