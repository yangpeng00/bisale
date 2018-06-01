package config

import (
	foundationConfig "bisale/foundation/config"
)

var Config = struct {
	Version  string
	Host     string
	Port     string `env:"port"`
	LogLevel string `yaml:"log_level" default:"debug" `
}{}

func init() {

	foundationConfig.Load(&Config)
}

func GetListenNetAddress() string {
	return Config.Host + ":" + Config.Port
}
