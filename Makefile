.PHONY: serve subscribe

build:
	@go build ./... -o bin/server

run:
	@go run main.go
