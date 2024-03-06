package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string `mapstructure:"PORT"`
	AuthPort string `mapstructure:"AUTH_SVC"`
	ProductPort string `mapstructure:"PRODUCT_SVC"`
	OrderPort string `mapstructure:"ORDER_SVC"`
}

func InitConfig() (*Config, error) {
	var c = &Config{}

	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil,err
	}

	err = viper.Unmarshal(c)
	if err != nil {
		return nil, fmt.Errorf("can't resolve unmarshel env %v", err)
	}
	return c, nil
}
