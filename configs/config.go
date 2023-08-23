package configs

import (
	"github.com/spf13/viper"
)

var conf *Config

const (
	PortKey  = "SERVER_PORT"
	XuyuHost = "XUYU_HOST"
)

type Config struct {
	Port     int
	XuyuHost string
}

func NewConfig() *Config {
	v := viper.New()

	v.SetDefault(PortKey, 3000)
	v.SetDefault(XuyuHost, "localhost:6060")

	v.AutomaticEnv()

	return &Config{
		Port:     v.GetInt(PortKey),
		XuyuHost: v.GetString(XuyuHost),
	}
}

func GetConfig() *Config {
	if conf == nil {
		conf = NewConfig()
	}
	return conf
}
