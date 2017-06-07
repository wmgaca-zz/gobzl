.PHONY: all
DEFAULT_TARGET: all

build:
	go fmt
	go vet
	go build

test:
	go test

all: build test
