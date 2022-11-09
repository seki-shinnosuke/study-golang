package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/seki-shinnosuke/study-golang/logger"
	"github.com/spf13/viper"
)

type (
	Config struct {
		APIServer `mapstructure:",squash" validate:"required"`
	}

	APIServer struct {
		AppApiPort string `mapstructure:"APP_API_PORT" validate:"required"`
	}
)

func NewConfig(confPath string) *Config {
	c := new(Config)

	viper.SetConfigFile(confPath)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("設定読み込み失敗. err: %v", err)
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		logger.Fatal("環境変数読み込み失敗. err: %v", err)
	}
	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		logger.Fatal("バリデーションエラー. err: %v", err)
	}

	return c
}
