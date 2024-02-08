package config

import (
	"github.com/joho/godotenv"
)

func LoadEnv() {
	envFile := ".env"
	godotenv.Load(envFile)
}
