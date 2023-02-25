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
	go mod init github.com/indiependente/go-proverbs-server
	go mod tidy

.PHONY: deps
deps:
	go mod download

.PHONY: docker
docker:
	docker build . -t indiependente/proverber-server

.PHONY: run
run:
	bin/service

generate: generate/proto
generate/proto:
	buf generate

lint:
	buf lint
