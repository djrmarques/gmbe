VERSION := $(shell cat .version)

.PHONY: build
build: 
	go build -ldflags "-X github.com/djrmarques/gmbe/cmd.version=$(VERSION)"
