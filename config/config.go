package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env файл не найден, используем переменные окружения")
	}
	key := os.Getenv("KEY")
	if key == "" {
		log.Fatal("KEY не найден в .env или переменных окружения")
	}
	return &Config{Key: key}
} 