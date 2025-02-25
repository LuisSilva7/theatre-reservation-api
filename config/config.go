package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBPort     uint32
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetUint32("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		JWTSecret:  viper.GetString("JWT_SECRET"),
	}

	return config, nil
}
