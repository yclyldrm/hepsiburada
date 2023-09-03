package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig(filename string) {
	if err := godotenv.Load(filename); err != nil {
		panic(err)
	}
}

func GetFromEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	panic("unsupported environment variable")
}
