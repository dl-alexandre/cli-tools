.PHONY: all test coverage lint build clean help

# Default target
all: test

# Run all tests
test:
	go test -v ./...

# Run tests with race detection
test-race:
	go test -race ./...

# Run tests with coverage
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run benchmarks
bench:
	go test -bench=. -benchmem ./...

# Run linter
lint:
	golangci-lint run

# Build all packages
build:
	go build ./...

# Build example
build-example:
	cd example && go build -o example-cli .

# Clean build artifacts
clean:
	rm -f coverage.out coverage.html
	rm -f example/example-cli

# Download dependencies
deps:
	go mod download
	go mod tidy

# Update dependencies
update:
	go get -u ./...
	go mod tidy

# Verify module
verify:
	go mod verify

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Check for vulnerabilities
vuln:
	go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...

# Show help
help:
	@echo "Available targets:"
	@echo "  test         - Run all tests"
	@echo "  test-race    - Run tests with race detection"
	@echo "  coverage     - Generate test coverage report"
	@echo "  bench        - Run benchmarks"
	@echo "  lint         - Run linter"
	@echo "  build        - Build all packages"
	@echo "  build-example- Build example CLI"
	@echo "  clean        - Clean build artifacts"
	@echo "  deps         - Download and verify dependencies"
	@echo "  update       - Update dependencies"
	@echo "  verify       - Verify module checksums"
	@echo "  fmt          - Format code"
	@echo "  vet          - Vet code"
	@echo "  vuln         - Check for vulnerabilities"
	@echo "  help         - Show this help message"
