run:
	go run cmd/main.go

protoc-auth:
	protoc ./pkg/auth-svc/pb/auth.proto --go-grpc_out=./ --go_out=./

protoc-product:
	protoc ./pkg/product-svc/pb/product.proto --go-grpc_out=./ --go_out=./
