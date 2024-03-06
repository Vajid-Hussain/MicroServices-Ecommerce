package config_order_svc

import "github.com/spf13/viper"

type Config struct {
	Port        string `mapstructure:"PORT"`
	DBName      string `mapstructure:"DBNAME"`
	DBUrl       string `mapstructure:"DBURL"`
	ProductPort string `mapstructure:"PRODUCTPORT"`
}

func InitConfig() (*Config, error) {
	var c = &Config{}
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
		return nil, err
	}

	return c, nil
}
