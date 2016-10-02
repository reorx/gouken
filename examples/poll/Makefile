all: proto

.PHONY: proto
proto:
	protoc ./proto/*.proto --go_out=plugins=grpc:.

run:
	go run cmd/rpcserver/main.go
