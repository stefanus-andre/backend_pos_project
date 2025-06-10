package Config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	viper.AutomaticEnv()

	config := &Config{
		DBHost: viper.GetString("DB_HOST"),
		DBPort: viper.GetString("DB_PORT"),
		DBUser: viper.GetString("DB_USER"),
		DBPass: viper.GetString("DB_PASSWORD"),
		DBName: viper.GetString("DB_NAME"),
	}
	return config, nil
}
