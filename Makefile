.PHONY: build run

build:
	@go build ./... -o bin/server

run:
	@go run main.go
