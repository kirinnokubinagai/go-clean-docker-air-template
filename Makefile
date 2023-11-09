.PHONY: setup go app

######################################
### local環境 
######################################
# ローカルバックエンドのビルド
local-build:
	docker compose build

# ローカルバックエンドの起動
local-up:
	docker compose up -d

# ローカルバックエンドの環境変数ファイルをシンボリックリンクで作成
local-set-up-backend-env:
	docker compose exec backend ln -s .env.local .env

# ローカル環境のDBマイグレーションファイルの作成
local-create-migration:
	docker compose exec backend /go/bin/migrate create -ext sql -dir /usr/local/go/src/app/database/migrations -seq ${FILENAME}
	
# ローカル環境のDBマイグレート
#
# ※ マイグレーションファイルのパス --path path
# ※ 接続先の設定 --database DBのドライバー user:password@host名:ポート/database名
local-migrate-up:
	docker compose exec backend /go/bin/migrate --path /usr/local/go/src/app/database/migrations --database 'postgresql://postgres:password@db:5432/db?sslmode=disable' -verbose up

# ローカル環境のDBマイグレートのロールバック
#
# ※ マイグレーションファイルのパス --path path
# ※ 接続先の設定 --database DBのドライバー user:password@host名:ポート/database名
local-migrate-down:
	docker compose exec backend /go/bin/migrate --path /usr/local/go/src/app/database/migrations --database 'postgresql://postgres:password@db:5432/db?sslmode=disable' -verbose down