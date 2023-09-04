package lib

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net/http"
)

var configPath = "./config.yml"
var configDefault = Config{
	Name: "go-gin-setup",
	Http: &HttpConfig{
		Host: "192.0.0.1",
		Port: 8080,
	},
	Auth: &AuthConfig{},
	Database: &DatabaseConfig{
		MigrationDir: "migrations",
	},
	Cors: &CorsConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	},
	Swagger: &SwaggerConfig{
		Title:       "Go Gin Setup Docs",
		Description: "Collection of Endpoints",
		Version:     "1.0",
		DocUrl:      "/swagger/*any",
	},
}

type Config struct {
	Name     string          `mapstructure:"Name"`
	Http     *HttpConfig     `mapstructure:"Http"`
	Database *DatabaseConfig `mapstructure:"Database"`
	Auth     *AuthConfig     `mapstructure:"Auth"`
	Cors     *CorsConfig     `mapstructure:"Cors"`
	Swagger  *SwaggerConfig  `mapstructure:"Swagger"`
}

type CorsConfig struct {
	AllowOrigins  []string `mapstructure:"AllowOrigins"`
	AllowMethods  []string `mapstructure:"AllowMethods"`
	AllowHeaders  []string `mapstructure:"AllowHeaders"`
	AllowWildcard bool     `mapstructure:"AllowWildcard"`
}

type SwaggerConfig struct {
	Title       string `mapstructrue:"Title"`
	Description string `mapstructure:"Description"`
	Version     string `mapstructure:"Version"`
	DocUrl      string `mapstructure:"DocUrl"`
}

type HttpConfig struct {
	Host string `mapstructure:"Host" validate:"ipv4"`
	Port int    `mapstructure:"Port" validate:"gte=1,lte=65535"`
}

type DatabaseConfig struct {
	Name         string `mapstructure:"Name"`
	Host         string `mapstructure:"Host"`
	Port         int    `mapstructure:"Port"`
	Username     string `mapstructure:"Username"`
	Password     string `mapstructure:"Password"`
	SslMode      string `mapstructure:"SslMode"`
	TimeZone     string `mapstructure:"TimeZone"`
	MigrationDir string `mapstructure:"MigrationDir"`
}

type AuthConfig struct {
	Enable             string   `mapstructure:"Enable"`
	TokenExpired       int      `mapstructure:"TokenExpired"`
	IgnorePathPrefixes []string `mapstructure:"IgnorePathPrefixes"`
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

func (h HttpConfig) ListenAddr() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

func (d DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		d.Host,
		d.Username,
		d.Password,
		d.Name,
		d.Port,
		d.SslMode,
		d.TimeZone,
	)
}
