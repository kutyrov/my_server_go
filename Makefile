.PHONY: build
build:
		go build -v ./cmd/server

.PHONY: test
test:
		go test -v -race -timeout 5s ./...

/DEFAULT_GOAL := build