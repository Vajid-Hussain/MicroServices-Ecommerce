package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/pb"
	requestmodel_auth_svc "github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/requestModel"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/config"
)

type UserHandler struct {
	clind  pb.AuthServiceClient
	server *gin.Engine
	config *config.Config
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

func (h *UserHandler) Authv2(c *gin.Context) {
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")

	c.Redirect(http.StatusSeeOther, url)
}

func (h *UserHandler) GoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != "randomstate" {
		c.JSON(http.StatusBadRequest, "state is not matching")
		return
	}

	code := c.Query("code")

	googlecon := config.GoogleConfig()

	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusBadRequest, "make code excahnge lead to error")
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, "get some access token follow to error")
		return
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, "fetch user data followt ot error")
		return
	}

	fmt.Println("userdata", string(userData))
	userDetails := requestmodel_auth_svc.UserDataOAuth{}
	json.Unmarshal(userData, &userDetails)

	c.JSON(http.StatusOK, userDetails)
}
