.PHONY: setup go app

# ローカルバックエンドのビルド
build-backend-local:
	docker compose build

# ローカルバックエンドの起動
up-backend-local:
	docker compose up -d

# ローカルバックエンドの環境変数ファイルをシンボリックリンクで作成
set-up-local-env:
	docker compose run backend ln -s .env.local .env

# ローカル環境のDBマイグレート
migrate-local:
	# マイグレーションファイルのパスは -source://path
	# 接続先の設定 -database DBの種類://host名/database名
	docker compose run backend /go/bin/migrate -source file://usr/local/go/src/app/database/migrations -database postgres://db/db up
