package config

import (
	"github.com/joho/godotenv"
)

type IConfig interface {
	LoadEnv()
}

type Config struct{}

func NewConfig() IConfig {
	return &Config{}
}

// 環境変数読み込み処理
func (config Config) LoadEnv() {
	envFile := ".env"
	godotenv.Load(envFile)
}
