package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google-oauth/controllers"
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
		controller := controllers.NewConnectController()
		c.Redirect(302, controller.GetRedirectUrl())
	})
	r.GET("/google/callback", func(c *gin.Context) {
		code := c.Query("code")
		session := c.Query("session")
		controller := controllers.NewCallbackController(code, session)
		tokenInfo := controller.GetUserInfo()

		c.JSON(200, gin.H{
			"userid": tokenInfo.UserId,
			"email":  tokenInfo.Email,
		})
	})
	Router = r
}
