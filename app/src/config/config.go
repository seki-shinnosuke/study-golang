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
		AppApiPort  string `mapstructure:"APP_API_PORT" validate:"required"`
		GinMode     string `mapstructure:"GIN_MODE"`
		CorsOrigins string `mapstructure:"CORS_ORIGINS"`
	}
)

func NewConfig(confPath string) *Config {
	c := new(Config)

	viper.SetConfigFile(confPath)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("Failed to read Configuration file. err: %v", err)
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		logger.Fatal("Failed to read environment. err: %v", err)
	}
	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		logger.Fatal("validation error. err: %v", err)
	}

	return c
}
