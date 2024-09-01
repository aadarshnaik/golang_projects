mysql:
	docker network create my_network
	docker run -d --name mysql -v mysql_data:/var/lib/mysql --network my_network -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=lostandfound -p 3306:3306 mysql:9.0.1

build:
	docker build -t aadarshnaik/go-authn:v1 .

run:
	docker run -d --rm --name authn -p 9090:9090 --network my_network --env-file envr_var.env aadarshnaik/go-authn:v1

.Phony: build run mysql