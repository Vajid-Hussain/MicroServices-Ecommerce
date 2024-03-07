package config

import (
	"fmt"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	Port              string `mapstructure:"PORT"`
	AuthPort          string `mapstructure:"AUTH_SVC"`
	ProductPort       string `mapstructure:"PRODUCT_SVC"`
	OrderPort         string `mapstructure:"ORDER_SVC"`
	GoogleClindID     string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClindSecret string `mapstructure:"GOOGLE_CLIENT_SECRET"`
}

type ConfigAuth struct {
	GoogleLoginConfig oauth2.Config
}

var AppConfig ConfigAuth

var c = &Config{}

func InitConfig() (*Config, error) {

	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(c)
	if err != nil {
		return nil, fmt.Errorf("can't resolve unmarshel env %v", err)
	}
	GoogleConfig()
	return c, nil
}

func GoogleConfig() oauth2.Config {
	AppConfig.GoogleLoginConfig = oauth2.Config{
		ClientID:     c.GoogleClindID,
		ClientSecret: c.GoogleClindSecret,
		RedirectURL:  "http://localhost:6005/google_callback",
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}

	return AppConfig.GoogleLoginConfig
}

