package config

import (
	"fmt"
	"github.com/jinzhu/configor"
)

type UserConfig struct {
	MySql  MysqlConfig    `json:"mysql"`
	Etcds  []EtcdConfig  `json:"etcds"`
	Redis  RedisConfig `json:"redis"`
}


var gConf UserConfig

func Init(confPath string) {
	fmt.Println(confPath)

	if err := configor.Load(&gConf, confPath); err != nil {
		panic(err)
	}

	fmt.Println(gConf)
}

func GetUserConf() *UserConfig {
	return  &gConf
}
