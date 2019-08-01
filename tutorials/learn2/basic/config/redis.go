package config

type RedisConfig struct {
	Enabled      bool        `json:"enabled"`
	Conn         string      `json:"conn"`
	DbNum        int         `json:"dbNum"`
	Password     string      `json:"password"`
	Timeout      int         `json:"timeout"`
}

func GetRedisConf() *RedisConfig {
	conf := GetUserConf()

	if conf == nil {
		return nil
	}

	return &conf.Redis
}

