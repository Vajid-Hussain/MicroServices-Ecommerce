package routes_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/handler"
)

func RouteAuthUser(engin *gin.Engine, user *handler.UserHandler) {
	userManagement := engin.Group("user")
	{
		userManagement.POST("/signup", user.UserSignup)
		userManagement.POST("/otpverification", user.OtpVerification)
		userManagement.POST("/userlogin", user.UserLogin)
	}
}
