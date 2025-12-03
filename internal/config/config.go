package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Database struct {
		Host     string `mapstructure:"host"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		Port     int    `mapstructure:"port"`
		SSLMode  string `mapstructure:"ssl_mode"`
	} `mapstructure:"database"`

	Redis struct {
		Addr string `mapstructure:"addr"`
	} `mapstructure:"redis"`

	JWT struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`

	GoogleCalendar struct {
		ServiceAccountFile string `mapstructure:"service_account_file"`
		CalendarID         string `mapstructure:"calendar_id"`
	} `mapstructure:"google_calendar"`

	Admin struct {
		RegisterSecret string `mapstructure:"register_secret"`
	} `mapstructure:"admin"`
}

var C Config

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Failed read config.yml: %v", err))
	}
	if err := viper.Unmarshal(&C); err != nil {
		panic(fmt.Errorf("Failed unmarshal config.yml: %v", err))
	}

	if C.GoogleCalendar.ServiceAccountFile == "" {
		panic("FATAL: ServiceAccountFile KOSONG! Pastiin config.yml ada dan struct bener.")
	}
}
