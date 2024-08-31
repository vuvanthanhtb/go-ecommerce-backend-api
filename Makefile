# name app
APP_NAME = server

# docker run --name mysql -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=Thanh20@$ mysql

# test benchmark connection database
# go test -bench=. -benchmem

dev:
	go run ./cmd/$(APP_NAME)

run:
	docker compose up -d && go run ./cmd/$(APP_NAME)

kill:
	docker compose kill

up:
	docker compose up -d

down:
	docker compose down

.PHONY: run

.PHONY: air