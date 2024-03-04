package routes_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/handler"
)

func RouteAuthAdmin(engin *gin.Engine, admin *handler.AdminHandler, authMiddlewire *handler.Middlewire) {
	userManagement := engin.Group("admin")
	{
		userManagement.POST("/adminlogin", admin.AdminLogin)
	}
}
