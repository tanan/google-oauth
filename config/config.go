package config

import (
	"github.com/BurntSushi/toml"
)

var AppCon AppConfig

type AppConfig struct {
	Api ApiConfig
}

type ApiConfig struct {
	GoogleClientID     string
	GoogleClientSecret string
	ApiURL             string
	GoogleCallBack     string
}

func init() {
	_, err := toml.DecodeFile("config.tml", &AppCon)
	if err != nil {
		panic(err)
	}
}
