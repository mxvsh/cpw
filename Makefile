.PHONY: build clean test release version

# Version information
VERSION ?= dev
COMMIT_SHA ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS := -X main.Version=$(VERSION) -X main.CommitSHA=$(COMMIT_SHA) -X main.BuildDate=$(BUILD_DATE)

# Default target
all: build

# Build the binary
build:
	go build -ldflags "$(LDFLAGS)" -o cpw

# Clean build artifacts
clean:
	rm -f cpw
	rm -f cpw-*

# Run tests
test:
	go test ./...

# Show version information
version:
	@echo "Version: $(VERSION)"
	@echo "Commit: $(COMMIT_SHA)"
	@echo "Build Date: $(BUILD_DATE)"

# Create a new release (Usage: make release version=v1.0.0)
release:
	@if [ -z "$(version)" ]; then \
		echo "Error: version parameter is required. Use 'make release version=v1.0.0'"; \
		exit 1; \
	fi
	git tag -a $(version) -m "Release $(version)"
	git push origin $(version)
	@echo "Tag $(version) pushed. GitHub Actions will now build and release the binaries."

# Build with specific version (Usage: make VERSION=v1.0.0 build)
build-version:
	@echo "Building version: $(VERSION)"
	@$(MAKE) build VERSION=$(VERSION)

# Show help
help:
	@echo "Available targets:"
	@echo "  make build             - Build the binary with dev version"
	@echo "  make build VERSION=v1.0.0 - Build with specific version"
	@echo "  make clean             - Remove build artifacts"
	@echo "  make test              - Run tests"
	@echo "  make version           - Display version information"
	@echo "  make release version=v1.0.0 - Create and push a new release tag"
	@echo "  make help              - Show this help message" 