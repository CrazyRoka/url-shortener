package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost string `mapstructure:"DB_HOST"`
	Port   uint   `mapstructure:"PORT"`
}

func LoadConfig() (*Config, error) {
	log.Info("Reading environment variables")

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Error("Error reading config from environment")
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.WithError(err).Error("Error unmarshalling config")
		return nil, err
	}

	return &config, nil
}
