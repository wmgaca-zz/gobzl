.PHONY: all
DEFAULT_TARGET: all

all:
	go fmt
	go vet
	go build
