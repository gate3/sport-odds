package config

import (
	"github.com/spf13/viper"
)

type EnvVariables struct {
	DatabaseUri     	string 		`mapstructure:"DATABASE_URI"`
	OddsApiBaseUrl 		string 		`mapstructure:"ODDS_API_BASE_URL"`
	OddsApiKey     		string 		`mapstructure:"ODDS_API_KEY"`
	DatabaseName		string		`mapstructure:"MONGO_DATABASE_NAME"`
}

func LoadEnvironmentVariables (path string) (envVariables EnvVariables, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return EnvVariables{}, err
	}

	err = viper.Unmarshal(&envVariables)
	return
}
