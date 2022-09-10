package config

import "github.com/spf13/viper"

type Config struct {
	Port        string `mapstructure:"PORT"`
	AuthSrvcUrl string `mapstructure:"AUTH_SRVC_URL"`
}

func LoadConfig() (conf Config, err error) {
	viper.AddConfigPath("./config/envs")

	viper.SetConfigName("default")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&conf)

	return
}
