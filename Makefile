VERSION := $(shell cat .version | tr -d "\n"  | tr -d " ")

.PHONY: build
build: 
	go build -ldflags "-X github.com/djrmarques/gmbe/cmd.version=$(VERSION)"

.PHONY: test
test:
	go test ./...
