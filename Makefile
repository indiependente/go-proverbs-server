all: clean build run

ifneq ($(OS),Windows_NT)
    OS := $(shell sh -c 'uname -s 2>/dev/null')
endif

ifeq ($(OS),Linux)
    LD_FLAGS = -ldflags="-s -w"
endif

.PHONY: build
build:
	CGO_ENABLED=0 go build $(LD_FLAGS) -o bin/service

.PHONY: clean
clean:
	rm -rf bin/*

.PHONY: deps-init
deps-init:
	rm -f go.mod go.sum
	go mod init
	go mod tidy

.PHONY: deps
deps:
	go mod download

.PHONY: docker
docker:
	docker build . -t indiependente/proverber-server

.PHONY: proto
proto: install-grpc-gateway
	protoc -I/usr/local/include -I. \
	-I$(GOPATH)/src \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:. proto/service.proto ; \
	protoc -I/usr/local/include -I. \
	-I$$GOPATH/src \
	-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:. proto/service.proto ; \
	mkdir -p ./rpc ; \
	mv ./proto/*.go ./rpc/ ; \
	mkdir -p ./swagger ; \
	protoc -I/usr/local/include -I. \
	-I$$GOPATH/src \
	-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--swagger_out=logtostderr=true:./swagger proto/service.proto ; \
	mv ./swagger/proto/service.swagger.json ./swagger/service.swagger.json ; \
	rm -rf ./swagger/proto ; \



.PHONY: run
run:
	bin/service

.PHONY: install-grpc-gateway
install-grpc-gateway:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go