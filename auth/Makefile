# note: call scripts from /scripts

# make new_migration MESSAGE_NAME="message name"
new_migration:
	migrate create -ext sql -dir scripts/migration/ -seq $(MESSAGE_NAME)
up_migration:
	migrate -path scripts/migration/ -database "mysql://root:123456@tcp(127.0.0.1:3306)/socialmedia?charset=utf8mb4&parseTime=True&loc=Local" -verbose up
down_migration:
	migrate -path scripts/migration/ -database "mysql://root:123456@tcp(127.0.0.1:3306)/socialmedia?charset=utf8mb4&parseTime=True&loc=Local" -verbose down
proto_auth:
	@echo "Generating auth proto..."
	protoc --proto_path=pkg/types/proto/ --go_out=pkg/types/proto/pb/auth_pb --go_opt=paths=source_relative \
        --go-grpc_out=pkg/types/proto/pb/auth_pb --go-grpc_opt=paths=source_relative \
        auth.proto
	@echo "Done!"

tidy:
	go mod tidy
.PHONY: vendor
vendor:
	go mod vendor -v
docker_clear:
	docker volume rm $(docker volume ls -qf dangling=true) & docker rmi $(docker images -f "dangling=true" -q)

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

build:
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose build --progress plain
	@echo "Docker images built and started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build:
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

gen_swagger:
	swag init -g cmd/main.go
