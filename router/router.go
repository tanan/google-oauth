package router

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google-oauth/google"
	v2 "google.golang.org/api/oauth2/v2"
)

var Router *gin.Engine

func init() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/login/google", func(c *gin.Context) {
		config := google.GetConnect()
		url := config.AuthCodeURL("")
		c.Redirect(302, url)
	})
	r.GET("/google/callback", func(c *gin.Context) {
		code := c.Query("code")
		config := google.GetConnect()
		ctx := context.Background()

		tok, err := config.Exchange(ctx, code)
		if err != nil {
			panic(err)
		}
		if tok.Valid() == false {
			panic(errors.New("valid token"))
		}

		service, _ := v2.New(config.Client(ctx, tok))
		tokenInfo, _ := service.Tokeninfo().AccessToken(tok.AccessToken).Context(ctx).Do()

		c.JSON(200, gin.H{
			"userid": tokenInfo.UserId,
			"email":  tokenInfo.Email,
		})
	})
	Router = r
}
