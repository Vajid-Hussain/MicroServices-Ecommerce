package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/pb"
	requestmodel "github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/requestModel"
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
	userDetails := requestmodel.UserDetails{}
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
	if err!=nil{
		fmt.Println("err: ",err)
	}

	fmt.Println("--", result)
}
