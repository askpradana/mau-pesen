package config

import "github.com/spf13/viper"

type Config struct {
	Server   struct{ Port int }
	Database struct {
		Host, User, Password, Name string
		Port                       int
		SSLMode                    string
	}
	Redis          struct{ Addr string }
	JWT            struct{ Secret string }
	GoogleCalendar struct{ ServiceAccountFile, CalendarID string }
	Admin          struct {
		RegisterSecret string `mapstructure:"register_secret"`
	}
}

var C Config

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	viper.Unmarshal(&C)
}
