package config_auth_svc

import "github.com/spf13/viper"

type General struct {
	Port   string `mapstructure:"PORT"`
	DBName string `mapstructure:"DBNAME"`
	DbUrl  string `mapstructure:"DBCONNECTION"`
}

type Token struct {
	AdminSecurityKey  string `mapstructure:"ADMIN_TOKENKEY"`
	SellerSecurityKey string `mapstructure:"SELLER_TOKENKEY"`
	UserSecurityKey   string `mapstructure:"USER_TOKENKEY"`
	TemperveryKey     string `mapstructure:"TEMPERVERY_TOKENKEY"`
}

type OTP struct {
	AccountSid string `mapstructure:"Account_SID"`
	AuthToken  string `mapstructure:"Auth_Token"`
	ServiceSid string `mapstructure:"Service_SID"`
}

type Config struct {
	Otp     OTP
	Token   Token
	General General
}

func InitConfig() (*Config, error) {
	var c = &Config{}
	viper.AddConfigPath("./auth-svc/")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.Unmarshal(&c.Token)
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&c.General)
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&c.Otp)
	if err != nil {
		return nil, err
	}

	return c, nil
}
