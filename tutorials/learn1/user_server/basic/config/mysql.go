package config

func GetMysqlConf() *MysqlConfig {
	conf := GetUserConf()

	if conf == nil {
		return nil
	}

	return &conf.MySql
}
