ifneq (,$(wildcard ./.env))
	include .env
	export
endif

.PHONY: test

PROJECT_NAME := $(shell basename "$(PWD)")

all: build run

build:
	go build -o ./build/bin/$(PROJECT_NAME) ./cmd/app

run:
	rm -rf ../data && ln -s "$(realpath ./)" ../data
	go run ./cmd/app

clean:
	go mod tidy

test:
	go test ./... --race -coverpkg=./...  --coverprofile=coverage.out

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

generate-mocks:
	mockgen -source=internal/domain/services/listservice.go -destination=internal/domain/services/listservice_mock.go -package=services
