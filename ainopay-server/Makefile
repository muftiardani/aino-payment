.PHONY: run migrate seed build test clean

# Run the server
run:
	go run cmd/server/main.go

# Build the server
build:
	go build -o bin/server cmd/server/main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod download
	go mod tidy

# Format code
fmt:
	go fmt ./...
