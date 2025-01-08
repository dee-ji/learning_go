# A Makefile that implements clean and that removes binaries and anything created by go build
.DEFAULT_GOAL := build

.PHONY:fmt vet build clean
fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build
clean:
	rm hello_world
