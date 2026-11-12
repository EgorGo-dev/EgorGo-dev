.PHONY: generate build run up down

generate:
	cd 6_1 && protoc -I . --go_out=. --go-grpc_out=. --grpc-gateway_out=. order.proto

build:
	go build -o bin/server ./6_1/server

run:
	go run ./6_1/server

up:
	docker compose -f 6_1/docker-compose.yml up --build

down:
	docker compose -f 6_1/docker-compose.yml down