mysql:
	docker network create my_network
	docker run -d --rm --name mysql -v mysql_data:/var/lib/mysql --network my_network -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=lostandfound -p 3306:3306 mysql:9.0.1 --max_connections=200

run:
	docker build -t aadarshnaik/go-authn:v1-amd64 .
	docker run -d --rm --name authn -p 9090:9090 --network my_network --env-file envr_var.env aadarshnaik/go-authn:v1-amd64

network-del:
	docker network rm my_network

.Phony: build run mysql network-del