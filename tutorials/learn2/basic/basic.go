package basic

import (
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	"micro_learn/tutorials/learn2/basic/config"
	"micro_learn/tutorials/learn2/basic/db"
	mr "micro_learn/tutorials/learn2/basic/redis"
	"sync"
)

var once sync.Once

func Init(configPath, logPath string) {
	once.Do(func() {
		config.Init(configPath)
		db.Init(logPath)
		mr.Init()
	})
}


func GetDbEngine() *xorm.Engine {
	return db.GetDbEngine()
}

func GetRedisClient() *redis.Client {
	return mr.GetRedisClient()
}

func GetEtdcConfig() []string {
	return config.GetEtdcConfig()
}

func GetMysqlConf() *config.MysqlConfig {
	return config.GetMysqlConf()
}