# name app
APP_NAME = server

# docker run --name mysql -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=Thanh20@$ mysql

# test benchmark connection database
# go test -bench=. -benchmem

run:
	go run ./cmd/${APP_NAME}/