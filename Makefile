.PHONY: build
build:
	go build -v ./cmd/apiserver && go build -v ./cmd/data-migrator

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
