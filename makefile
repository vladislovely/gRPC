.PHONY: generate-proto docker-build docker-build-server docker-build-client docker-up-and-build docker-up

APP_SERVER_NAME := server
APP_CLIENT_NAME := client

generate-proto:
	protoc --proto_path=proto --go_out=internal/gen --go-grpc_out=internal/gen proto/*.proto

docker-build-server:
	docker compose build $(APP_SERVER_NAME)

docker-build-client:
	docker compose build $(APP_CLIENT_NAME)

docker-up-and-build: docker-build-client docker-build-server
	docker compose up -d

docker-up:
	docker compose up -d

lint:
	golangci-lint run --config ./golangci.yml ./...

format:
	golines --max-len=100 -w ./