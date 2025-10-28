package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Services ServicesConfig
}

type ServerConfig struct {
	AllowOrigin string
	Port        string
}

type ServicesConfig struct {
	Account ServiceConfig
	Auth    ServiceConfig
	Herd    ServiceConfig
}

type ServiceConfig struct {
	Host string
	Port string
}

const configPath = "config/config.yaml"

func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigFile(configPath)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
