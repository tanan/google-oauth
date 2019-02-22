package controllers

import (
	"context"
	"errors"
	"github.com/tanan/google-oauth/google"
	v2 "google.golang.org/api/oauth2/v2"
)

type CallbackController struct {
	*CallbackRequest
}

type CallbackRequest struct {
	code    string
	session string
}

func NewCallbackController(code string, session string) *CallbackController {
	return &CallbackController{
		&CallbackRequest{
			code:    code,
			session: session,
		},
	}
}

func (controller *CallbackController) GetUserInfo() *v2.Tokeninfo {
	config := google.GetConnect()
	ctx := context.Background()

	tok, err := config.Exchange(ctx, controller.code)
	if err != nil {
		panic(err)
	}
	if tok.Valid() == false {
		panic(errors.New("valid token"))
	}

	service, _ := v2.New(config.Client(ctx, tok))
	tokenInfo, _ := service.Tokeninfo().AccessToken(tok.AccessToken).Context(ctx).Do()
	return tokenInfo
}
