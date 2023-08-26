package lib

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var configPath = "./config.yml"
var configDefault = Config{
	Name: "go-gin-setup",
	Http: &HttpConfig{
		Host: "192.0.0.1",
		Port: 8080,
	},
	Database: &DatabaseConfig{},
}

type Config struct {
	Name     string          `mapstructure:"Name"`
	Http     *HttpConfig     `mapstructure:"Http"`
	Database *DatabaseConfig `mapstructure:"Database"`
}

type HttpConfig struct {
	Host string `mapstructure:"Host" validate:"ipv4"`
	Port int    `mapstructure:"Port" validate:"gte=1,lte=65535"`
}

type DatabaseConfig struct {
	Name     string `mapstructure:"Name"`
	Host     string `mapstructure:"Host"`
	Port     int    `mapstructure:"Port"`
	Username string `mapstructure:"Username"`
	Password string `mapstructure:"Password"`
	SslMode  string `mapstructure:"SslMode"`
	TimeZone string `mapstructure:"TimeZone"`
}

func NewConfig() Config {
	config := configDefault
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(errors.Wrap(err, "failed to read config"))
	}
	if err := viper.Unmarshal(&config); err != nil {
		panic(errors.Wrap(err, "failed to marshal config"))
	}
	return config
}

func SetConfigPath(path string) {
	configPath = path
}
