package config

import (
	"fmt"
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
	//_, err := toml.DecodeFile("config.tml", &AppCon)
	_, err := toml.DecodeFile("/Users/tanan/Programming/go/src/github.com/google-oauth/config.tml", &AppCon)
	fmt.Println(AppCon.Api.GoogleClientID)
	if err != nil {
		panic(err)
	}
}
