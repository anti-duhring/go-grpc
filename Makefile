OS=linux
OS_ARCH=amd64

GATEWAY_NAME=gateway
GATEWAY_PATH=cmd/gateway

ECOMMERCE_NAME=ecommerce
ECOMMERCE_PATH=cmd/ecommerce

docs:
	swag init -g ./$(ECOMMERCE_PATH)/main.go -o ./internal/docs

generate_grpc:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	internal/wallet/wallet.proto

run_gateway:
	go run cmd/gateway/main.go

build_gateway:
	rm -f $(GATEWAY_NAME)
	env GOS=$(OS) GOARCH=$(OS_ARCH) go build -o $(GATEWAY_NAME) $(GATEWAY_PATH)/main.go

run_ecommerce:
	go run cmd/ecommerce/main.go

build_ecommerce:
	rm -f $(ECOMMERCE_NAME)
	env GOS=$(OS) GOARCH=$(OS_ARCH) go build -o $(ECOMMERCE_NAME) $(ECOMMERCE_PATH)/main.go

build: build_gateway build_ecommerce

dockerbuild: build
	docker compose down
	docker compose build
	docker compose up -d