run: build
	@./bin/chatapp

build:
	@go build -o bin/chatapp

test:
	@go test -v ./...
