package config

import (
	"fmt"
)

type (
	DatabaseConfig struct {
		Address  string `mapstructure:"address"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	}
)

func (dbConf *DatabaseConfig) Validate() bool {
	return dbConf != nil && dbConf.Address != "" && dbConf.Username != "" && dbConf.Password != ""
}

func (conf *DatabaseConfig) Complete() {
}

func (dbConf *DatabaseConfig) DSN(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", dbConf.Username, dbConf.Password, dbConf.Address, dbName)
}
