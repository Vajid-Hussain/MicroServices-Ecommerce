package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/pb"
	requestmodel_auth_svc "github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/requestModel"
)

type UserHandler struct {
	clind  pb.AuthServiceClient
	server *gin.Engine
}

func NewUserHandler(clind pb.AuthServiceClient, engin *gin.Engine) *UserHandler {
	return &UserHandler{clind: clind,
		server: engin}
}

func (h *UserHandler) UserSignup(ctx *gin.Context) {
	userDetails := requestmodel_auth_svc.UserDetails{}
	err := ctx.BindJSON(&userDetails)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.clind.Sighup(context.Background(), &pb.SignupRequest{
		Name:            userDetails.Name,
		Email:           userDetails.Email,
		Phone:           userDetails.Phone,
		Password:        userDetails.Password,
		ConfirmPassword: userDetails.ConfirmPassword,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (h *UserHandler) OtpVerification(c *gin.Context) {
	var otpData requestmodel_auth_svc.OtpVerification

	token := c.Request.Header.Get("Authorization")

	if err := c.BindJSON(&otpData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := h.clind.OtpVerifiction(context.Background(), &pb.OtpRequest{
		Otp:            otpData.Otp,
		TemperverToken: token,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, result)
}

func (h *UserHandler) UserLogin(c *gin.Context) {
	var loginCredential requestmodel_auth_svc.UserLogin
	if err := c.BindJSON(&loginCredential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := h.clind.UserLogin(context.Background(), &pb.LoginRequest{
		Phone:    loginCredential.Phone,
		Password: loginCredential.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, result)
}
