.PHONY: test build

test:
	go test ./...

build:
	mkdir -p build
	go build -o build/labda ./cli
