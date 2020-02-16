package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPPort int `envconfig:"http:port" default:"8000"`

	DBHost     string `envconfig:"db_host" default:"localhost"`
	DBPort     uint   `envconfig:"db_port" default:"5432"`
	DBName     string `envconfig:"db_name"`
	DBUser     string `envconfig:"db_user"`
	DBPassword string `envconfig:"db_password"`
}

func GetConfig() Config {
	var config Config
	err := envconfig.Process("", &config)

	if err != nil {
		panic(err)
	}

	return config
}
