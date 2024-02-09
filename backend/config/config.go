package config

import (
	"github.com/joho/godotenv"
)

type IConfig interface {
	LoadEnv()
}

type Config struct{}

// コンストラクター
// @return IConfig インターフェース
func NewConfig() IConfig {
	return &Config{}
}

// 環境変数読み込み処理
// memo: ビルドする際に、.env.${環境}を.envにしている
func (config Config) LoadEnv() {
	envFile := ".env"
	godotenv.Load(envFile)
}
