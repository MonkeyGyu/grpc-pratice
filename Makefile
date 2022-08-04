# Makefile 

init:
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go mod download

generate:
	go generate ./...
