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

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_auth
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

build_auth:
	@echo "Building auth..."
	cd ../auth
	@echo "Done!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## generate proto: generates the proto files
proto_auth:
	@echo "Generating auth proto..."
	protoc --proto_path=pkg/types/proto/ --go_out=pkg/types/proto/pb/auth_pb --go_opt=paths=source_relative \
        --go-grpc_out=pkg/types/proto/pb/auth_pb --go-grpc_opt=paths=source_relative \
        auth.proto
	@echo "Done!"