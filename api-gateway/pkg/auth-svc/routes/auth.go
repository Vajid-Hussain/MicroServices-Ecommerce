package routes_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/handler"
)

func RouteAuthUser(engin *gin.Engine, user *handler.UserHandler, authMiddlewire *handler.Middlewire) {

	engin.GET("/google_callback", user.GoogleCallback)

	userManagement := engin.Group("user")
	{
		userManagement.GET("/google_login", user.Authv2)
		userManagement.POST("/signup", user.UserSignup)
		userManagement.POST("/otpverification", user.OtpVerification)
		userManagement.POST("/userlogin", user.UserLogin)
	}

	// engin.Use(authMiddlewire.AdminAuthMiddlewire)
	// {
	// 	engin.POST("/check", func(ctx *gin.Context) {
	// 		fmt.Println("work well")
	// 	})
	// }
}
