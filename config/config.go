package config

import (
	foundationConfig "bisale/foundation/config"
)

var Config = struct {
	Version    string
	Host       string
	Port       string `env:"port"`
	LogLevel   string `yaml:"log_level" default:"debug"`
	JWTToken   string `yaml:"jwt_token" default:"123456"`
	KYCBucket  string `yaml:"kyc_bucket" default:"bisale-test-huadong"`
	OldKYCHost string `yaml:"old_kyc_bucket" default:"http://bi-sale.oss-cn-hongkong.aliyuncs.com/"`

	AccountService struct {
		Host        string
		Port        string
		MaxConn     uint32 `yaml:"max_conn" default:"10"`
		ConnTimeout uint32 `yaml:"conn_timeout" default:"3"`
		IdleTimeout uint32 `yaml:"idle_timeout" default:"2"`
	} `yaml:"account_service"`

	MessageService struct {
		Host        string
		Port        string
		MaxConn     uint32 `yaml:"max_conn" default:"10"`
		ConnTimeout uint32 `yaml:"conn_timeout" default:"3"`
		IdleTimeout uint32 `yaml:"idle_timeout" default:"2"`
	} `yaml:"message_service"`

	CaptchaService struct {
		Host        string
		Port        string
		MaxConn     uint32 `yaml:"max_conn" default:"10"`
		ConnTimeout uint32 `yaml:"conn_timeout" default:"1"`
		IdleTimeout uint32 `yaml:"idle_timeout" default:"1"`
	} `yaml:"captcha_service"`

	StorageService struct {
		Host        string
		Port        string
		MaxConn     uint32 `yaml:"max_conn" default:"10"`
		ConnTimeout uint32 `yaml:"conn_timeout" default:"1"`
		IdleTimeout uint32 `yaml:"idle_timeout" default:"1"`
	} `yaml:"storage_service"`

	BisaleUserService struct {
		Host        string
		Port        string
		MaxConn     uint32 `yaml:"max_conn" default:"10"`
		ConnTimeout uint32 `yaml:"conn_timeout" default:"1"`
		IdleTimeout uint32 `yaml:"idle_timeout" default:"1"`
	} `yaml:"bisale_user_service"`

	BisaleBusinessService struct {
		Host        string
		Port        string
		MaxConn     uint32 `yaml:"max_conn" default:"10"`
		ConnTimeout uint32 `yaml:"conn_timeout" default:"1"`
		IdleTimeout uint32 `yaml:"idle_timeout" default:"1"`
	} `yaml:"bisale_business_service"`
}{}

func init() {

	foundationConfig.Load(&Config)
}

func GetListenNetAddress() string {
	return Config.Host + ":" + Config.Port
}
