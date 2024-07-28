.PHONY: build
build:
	go build ./...
all: build

.PHONY: deps.update
deps.update:
	go get -u
	go mod tidy