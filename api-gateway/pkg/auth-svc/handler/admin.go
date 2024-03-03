package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/pb"
	requestmodel_auth_svc "github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/requestModel"
)

type AdminHandler struct {
	clind  pb.AuthServiceClient
	server *gin.Engine
}

func NewAdminHandler(clind pb.AuthServiceClient, engin *gin.Engine) *AdminHandler {
	return &AdminHandler{clind: clind,
		server: engin}
}

func (u *AdminHandler) AdminLogin(c *gin.Context) {
	var loginCredential requestmodel_auth_svc.AdminLoginData

	err := c.BindJSON(&loginCredential)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	result, err := u.clind.AdminLogin(context.Background(), &pb.AdminLoginRequest{
		Email:    loginCredential.Email,
		Password: loginCredential.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

// func (u *AdminHandler) Google
