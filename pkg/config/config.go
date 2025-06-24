package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	viper *viper.Viper
}

var (
	cfg *Config
)

func New() error {
	v := viper.New()
	v.SetConfigFile("./etc/config/config.yaml")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	cfg = &Config{
		viper: v,
	}

	return nil
}

func GetConfig() *Config {
	return cfg
}
