package configs

/********************************************************************************
* Temancode Load Config Package			                                        *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

import (
	"github.com/spf13/viper"
)

var Configs Config

// func init() {
// 	LoadConfig(".")
// }

type Config struct {
	APP_NAME         string `mapstructure:"APP_NAME"`
	PORT             string `mapstructure:"PORT"`
	GIN_MODE         string `mapstructure:"GIN_MODE"`
	APP_ENV          string `mapstructure:"APP_ENV"`
	TIMEOUT_REQ      int64  `mapstructure:"TIMEOUT_REQUEST"`
	LOCAL            string `mapstructure:"LOCAL"`
	BASE_URL_EDM     string `mapstructure:"BASE_URL_EDM"`
	USERNAME_BRIGATE string `mapstructure:"USERNAME_BRIGATE"`
	PASSWORD_BRIGATE string `mapstructure:"PASSWORD_BRIGATE"`
	DEBUG            bool   `mapstructure:"DEBUG"`
	DB_USERNAME      string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD      string `mapstructure:"DB_PASSWORD"`
	DB_HOST          string `mapstructure:"DB_HOST"`
	DB_PORT          string `mapstructure:"DB_PORT"`
	DB_DATABASE      string `mapstructure:"DB_DATABASE"`
	TIME_LOCATION    string `mapstructure:"TIME_LOCATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	Configs = config
	return
}
