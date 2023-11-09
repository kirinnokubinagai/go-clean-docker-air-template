.PHONY: setup go app

build backend local:
	docker compose build

up backend local:
	docker compose up -d

set up local env:
	docker compose run backend ln -s .env.local .env
