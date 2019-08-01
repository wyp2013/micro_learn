package config

type MysqlConfig struct {
	URL               string `json:"url"`
	Enabled           bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
}

func GetMysqlConf() *MysqlConfig {
	conf := GetUserConf()

	if conf == nil {
		return nil
	}

	return &conf.MySql
}
