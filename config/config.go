package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	Log LogConfig `json:"log"`
	Mysql MysqlConfig `json:"mysql"`
}

type LogConfig struct {
	Folder       string `json:"folder"`
	FileName     string `json:"filename"`
	MaxAge       int    `json:"maxAge"`
	RotationTime int    `json:"rotationTime"`
}

type MysqlConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Addr     string `json:"addr"`
	DbName   string `json:"dbname"`
	Option   string `json:"option"`
}


var config Config
var configOnce sync.Once

func GetConfig()*Config{
	configOnce.Do(func() {
		v:=viper.New()
		v.SetConfigName("config")
		v.AddConfigPath("./config")
		v.AddConfigPath("../config")
		v.AddConfigPath("../../config")
		v.AddConfigPath("../")
		v.SetConfigType("json")
		err := v.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %s \n", err))
		}

		if err := v.Unmarshal(&config); err != nil {
			panic("parse config file failed")
		}
	})
	return &config
}
