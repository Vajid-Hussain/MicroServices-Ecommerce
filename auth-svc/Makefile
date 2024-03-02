run:
	go run ./cmd/main.go

protoc:
	protoc ./pkg/pb/auth.proto --go_out=./ --go-grpc_out=./
	