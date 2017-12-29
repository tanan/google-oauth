package google

import (
	"github.com/google-oauth/config"
	"golang.org/x/oauth2"
)

const (
	authorizedEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenEndpoint      = "https://www.googleapis.com/oauth2/v4/token"
)

func GetConnect() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.AppCon.Api.GoogleClientID,
		ClientSecret: config.AppCon.Api.GoogleClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizedEndpoint,
			TokenURL: tokenEndpoint,
		},
		Scopes:      []string{"openid", "email", "profile"},
		RedirectURL: config.AppCon.Api.ApiURL + config.AppCon.Api.GoogleCallBack,
	}
}
