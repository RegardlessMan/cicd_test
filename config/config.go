package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error read config file #v", err)
	}

	AppConfig = &Config{}

	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decoed into struct #v", err)
	}

	// 初始化数据库Db
	initDb()
}
