.PHONY: all build clean run test tidy

# Default target
all: build

# Build the binary
build:
	go build -o bin/goodfirstgo ./cmd/goodfirstgo

# Install globally
install:
	go install ./cmd/goodfirstgo

# Run the CLI
run: tidy build
	./bin/goodfirstgo --help

# Development run with default flags
run-dev:
	go run ./cmd/goodfirstgo --language go --label good-first-issue

# Tidy dependencies
tidy:
	go mod tidy

# Test
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Lint (requires golangci-lint)
lint:
	golangci-lint run

# Generate docs (if using godoc)
docs:
	go doc ./...

# Check formatting
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

