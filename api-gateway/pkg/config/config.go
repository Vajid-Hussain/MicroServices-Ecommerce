package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string `mapstructure:"PORT"`
	AuthPort string `mapstructure:"AUTH_PORT"`
}

func InitConfig() (*Config, error) {
	var c = Config{}

	viper.AddConfigPath("./")
	viper.SetConfigType("env")
	viper.SetConfigName("dev")

	viper.AutomaticEnv()
	err := viper.Unmarshal(&c)
	if err != nil {
		return nil, fmt.Errorf("can't resolve unmarshel env %v", err)
	}
	return &c, nil
}
