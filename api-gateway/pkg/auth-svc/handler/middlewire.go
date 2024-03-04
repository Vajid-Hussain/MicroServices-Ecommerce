package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/pb"
)

type Middlewire struct {
	Clind pb.AuthServiceClient
}

func NewAuthMiddlewire(clind pb.AuthServiceClient) *Middlewire {
	return &Middlewire{Clind: clind}
}

func (m *Middlewire) UserAuthMiddlewire(c *gin.Context) {
	accessToken := c.Request.Header.Get("authorization")
	refreshToken := c.Request.Header.Get("refreshtoken")

	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"err": "there is no access token"})
		c.Abort()
		return
	}

	result, err := m.Clind.ValidateUserToken(context.Background(), &pb.ValidateTokenRequest{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token is not valid"})
		c.Abort()
		return
	}
	c.Set("userID", result.UserID)
	c.Next()
}

func (m *Middlewire) AdminAuthMiddlewire(c *gin.Context) {
	token := c.Request.Header.Get("authorization")

	_, err := m.Clind.ValidateAdminToken(context.Background(), &pb.ValidateAdminTokenRequest{
		Token: token,
	})
	fmt.Println("----", err, token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token is not valid"})
		return
	}
	c.Next()
}
