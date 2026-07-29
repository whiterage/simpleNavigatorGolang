.PHONY: all build run test test-coverage fmt vet clean

BINARY := bin/simple_navigator

all: build

build:
	go build -o $(BINARY) ./cmd/simple_navigator

run: build
	./$(BINARY)

test:
	go test -v -timeout 100s ./...

test-coverage:
	go test -cover ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	rm -rf bin
	rm -f *.dot
	go clean -testcache
