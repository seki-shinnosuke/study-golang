package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/seki-shinnosuke/study-golang/util/logger"
	"github.com/spf13/viper"
)

type (
	Config struct {
		APIServer `mapstructure:",squash" validate:"required"`
		RDB       `mapstructure:",squash" validate:"required"`
		TimeZone  string `mapstructure:"TIME_ZONE" validate:"required"`
	}

	APIServer struct {
		AppApiPort  string `mapstructure:"APP_API_PORT" validate:"required"`
		GinMode     string `mapstructure:"GIN_MODE"`
		CorsOrigins string `mapstructure:"CORS_ORIGINS"`
	}

	RDB struct {
		Host   string `mapstructure:"RDB_HOST" validate:"required"`
		DBName string `mapstructure:"RDB_NAME" validate:"required"`
		User   string `mapstructure:"RDB_USER" validate:"required"`
		Passwd string `mapstructure:"RDB_PASSWORD" validate:"required"`
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
