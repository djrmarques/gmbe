VERSION := $(shell cat .version | tr "\n" " " | tr " " "")

.PHONY: build
build: 
	go build -ldflags "-X github.com/djrmarques/gmbe/cmd.version=$(VERSION)"
