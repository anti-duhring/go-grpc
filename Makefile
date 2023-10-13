generate_grpc:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	internal/invoicer/invoicer.proto

gateway:
	go run cmd/gateway/main.go

ecommerce:
	go run cmd/ecommerce/main.go