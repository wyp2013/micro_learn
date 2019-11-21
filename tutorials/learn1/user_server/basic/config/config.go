package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"micro_learn/micro/go-micro/config"
	"micro_learn/micro/go-micro/config/source/file"
	"os"
	"path/filepath"
	"strings"
)

type UserConfig struct {
	MySql MysqlConfig  `json:"mysql"`
	Etcds  []EtcdConfig `json:"etcds"`
}

type MysqlConfig struct {
	URL               string `json:"url"`
	Enable            bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
}

type EtcdConfig struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

// 这个有bug
func Init0() {
	dirPath, _ := filepath.Abs(filepath.Dir("./"))
	last := strings.LastIndex(dirPath, "basic/config")
	appPath := dirPath[:last]
	fmt.Println(appPath)

	pt := filepath.Join(appPath, "config")
	os.Chdir(appPath)
	if err := config.Load(file.NewSource(file.WithPath(pt + "/config.yaml"))); err != nil {
		panic(err)
	}

	//
	var etcd []MysqlConfig

	if err := config.Get("app", "etcd").Scan(&etcd); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(etcd)
}

var gConf UserConfig

func Init() {
	appPath, _ := filepath.Abs(filepath.Dir("./"))
	confPath := filepath.Join(appPath, "config/config.yaml")

	if err := configor.Load(&gConf, confPath); err != nil {
		panic(err)
	}

	fmt.Println(gConf)
}

func GetUserConf() *UserConfig {
	return  &gConf
}
