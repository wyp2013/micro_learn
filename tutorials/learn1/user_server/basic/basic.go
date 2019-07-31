package basic

import (
	"github.com/go-xorm/xorm"
	"micro_learn/tutorials/learn1/user_server/basic/config"
	"micro_learn/tutorials/learn1/user_server/basic/db"
)

func Init(logPath string) {
	config.Init()
	db.Init(logPath)
}


func GetDbEngine() *xorm.Engine {
	return db.GetDbEngine()
}

func GetEtdcConfig() []string {
	return config.GetEtdcConfig()
}

func GetMysqlConf() *config.MysqlConfig {
	return config.GetMysqlConf()
}