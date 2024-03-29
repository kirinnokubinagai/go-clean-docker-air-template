package main

import (
	"go_app/config"
	"go_app/registry"
	"go_app/server"

	"github.com/sirupsen/logrus"
)

/*
サーバーの起動、DB接続、DIコンテナの作成
*/
func main() {
	config.NewConfig().LoadEnv()

	container, err := registry.NewContainer().BuildContainer()

	if err != nil {
		logrus.Fatalln(err)
	}

	err = container.Invoke(func(server server.IServer) {
		err := server.GetEngine().Run(":8080")
		if err != nil {
			logrus.Fatalln(err)
		}
	})
	if err != nil {
		logrus.Fatalln(err)
	}
}
