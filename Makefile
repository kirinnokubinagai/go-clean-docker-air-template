.PHONY: setup go_app

build backend local:
	cp ./backend/.env.local ./backend/.env \
	&& docker compose build

up backend local:
	cp ./backend/.env.local ./backend/.env \
	&& docker compose up