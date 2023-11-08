package config

import (
	"github.com/joho/godotenv"
)

func loadEnv() {
	envFile := ".env"
	godotenv.Load(envFile)
}
