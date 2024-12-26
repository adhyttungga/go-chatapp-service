package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Database struct {
	URI string `env:"MONGO_DB_URI"`
}

type Origin struct {
	AllowOrigin string `env:"ALLOW_ORIGIN"`
}

type ServerConfig struct {
	GinMode     string `env:"GIN_MODE" default:"debug"`
	ServiceHost string `env:"SERVICE_HOST" default:"localhost"`
	ServicePort string `env:"SERVICE_PORT" default:"5000"`
	DB          Database
	Origin      Origin
	PrivateKey  string `env:"PRIVATE_KEY"`
	PublicKey   string `env:"PUBLIC_KEY"`
}

var Config ServerConfig

func init() {
	if err := loadConfig(); err != nil {
		panic(err)
	}
}

func loadConfig() error {
	if err := godotenv.Load(".env"); err != nil {
		log.Warn().Msg("Cannot find .env fils. OS Environtments will be used")
	}

	err := env.Parse(&Config)
	return err
}
