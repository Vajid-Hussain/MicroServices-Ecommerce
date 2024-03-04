package routes_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/handler"
)

func RouteAuthUser(engin *gin.Engine, user *handler.UserHandler, authMiddlewire *handler.Middlewire) {

	userManagement := engin.Group("user")
	{
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
