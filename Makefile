.PHONY: build clean lint tool help

all: build

build:
	#@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o xdance ./cmd/
	go build -v -o xdance ./cmd/main.go

clean:
	@rm -rf xdance

lint:
	@golint ./...

tool:
	@go vet .
	@gofmt -w .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"

