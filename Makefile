generate_grpc_code:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	internal/invoicer/invoicer.proto

run_gateway:
	go run cmd/gateway/main.go

run_ecommerce:
	go run cmd/ecommerce/main.go