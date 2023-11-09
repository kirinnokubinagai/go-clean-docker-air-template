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
