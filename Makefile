.PHONY: build clean test release version setup-homebrew

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

# Setup Homebrew tap (requires GitHub repo access)
setup-homebrew:
	@echo "IMPORTANT: Before running this target, you need to create a Personal Access Token"
	@echo "with 'repo' permissions and add it as a secret in your GitHub repository:"
	@echo "1. Go to GitHub → Settings → Developer settings → Personal access tokens"
	@echo "2. Create a new token with 'repo' access"
	@echo "3. Go to your repository → Settings → Secrets → Actions"
	@echo "4. Add a new secret named HOMEBREW_TAP_TOKEN with your token value"
	@echo "\nThen you can trigger the workflow manually via GitHub Actions UI"
	@echo "or using this command:"
	@echo "\ncurl -X POST \\"
	@echo "  -H \"Accept: application/vnd.github.v3+json\" \\"
	@echo "  -H \"Authorization: token YOUR_PAT_HERE\" \\"
	@echo "  https://api.github.com/repos/$(GITHUB_USER)/cpw/actions/workflows/setup_tap.yml/dispatches \\"
	@echo "  -d '{\"ref\":\"main\",\"inputs\":{\"create_repo\":\"true\"}}'"
	@echo "\nRequires GITHUB_USER environment variable to be set."
	@echo "Example: GITHUB_USER=mxvsh make setup-homebrew"

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
	@echo "  make setup-homebrew    - Setup Homebrew tap repository (requires GitHub token)"
	@echo "  make help              - Show this help message" 