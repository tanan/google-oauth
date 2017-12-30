package controllers

import "github.com/google-oauth/google"

type ConnectController struct {
}

func NewConnectController() *ConnectController {
	return &ConnectController{}
}

func (controller *ConnectController) GetRedirectUrl() string {
	config := google.GetConnect()
	return config.AuthCodeURL("")
}
