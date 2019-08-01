package config

import "fmt"

type EtcdConfig struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

func GetEtdcConfig() []string {
	conf := GetUserConf()

	if conf == nil {
		return nil
	}

	var etcds []string
	for _, etcd := range conf.Etcds {
		if etcd.Enabled {
			etcds = append(etcds, fmt.Sprintf("%s:%d", etcd.Host, etcd.Port))
		}
	}

	return etcds
}
