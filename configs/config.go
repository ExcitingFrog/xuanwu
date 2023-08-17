package configs

import (
	"github.com/spf13/viper"
)

var conf *Config

const (
	PortKey = "SERVER_PORT"
)

type Config struct {
	Port int
}

func NewConfig() *Config {
	v := viper.New()
	v.AutomaticEnv()

	v.SetDefault(PortKey, 3000)

	return &Config{
		Port: v.GetInt(PortKey),
	}
}

func GetConfig() *Config {
	if conf == nil {
		conf = NewConfig()
	}
	return conf
}
