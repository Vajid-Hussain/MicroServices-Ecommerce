package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/config"
)

func InitGinServer(c config.Config) (*gin.Engine, error) {
	engin := gin.Default()

	err := engin.Run(c.Port)
	if err != nil {
		return nil, err
	}

	return engin, nil
}
