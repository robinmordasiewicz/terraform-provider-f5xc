# Makefile for terraform-provider-f5xc
# Automated build, test, and code generation

BINARY_NAME=terraform-provider-f5xc
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

# Directories
TOOLS_DIR=tools
PROVIDER_DIR=internal/provider
CLIENT_DIR=internal/client
DOCS_DIR=docs
SPEC_DIR?=/tmp

# Go commands
GO=go
GOFMT=gofmt
GOLINT=golangci-lint

.PHONY: all build test lint fmt clean clean-generated regenerate generate docs install help sweep sweep-dry-run testacc

# Default target
all: generate build lint test docs

# Help
help:
	@echo "terraform-provider-f5xc Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make              - Generate, build, lint, test, and generate docs"
	@echo "  make build        - Build the provider binary"
	@echo "  make test         - Run tests"
	@echo "  make testacc      - Run acceptance tests (requires F5XC credentials)"
	@echo "  make lint         - Run linters"
	@echo "  make fmt          - Format Go code"
	@echo "  make generate     - Generate resources from OpenAPI specs"
	@echo "  make docs         - Generate Terraform documentation"
	@echo "  make clean        - Remove build artifacts"
	@echo "  make install      - Install provider locally"
	@echo "  make sweep        - Clean up orphaned test resources (requires F5XC credentials)"
	@echo "  make sweep-resource RESOURCE=f5xc_namespace - Sweep specific resource type"
	@echo ""
	@echo "Environment Variables:"
	@echo "  SPEC_DIR          - Directory containing OpenAPI specs (default: /tmp)"
	@echo "  F5XC_SPEC_DIR     - Alternative env var for spec directory"
	@echo ""
	@echo "For acceptance tests and sweepers, set one of:"
	@echo "  F5XC_API_URL + F5XC_API_P12_FILE + F5XC_P12_PASSWORD (P12 auth)"
	@echo "  F5XC_API_URL + F5XC_API_CERT + F5XC_API_KEY (PEM auth)"
	@echo "  F5XC_API_URL + F5XC_API_TOKEN (Token auth)"

# Build the provider
build:
	@echo "Building $(BINARY_NAME)..."
	$(GO) build -ldflags="-X main.version=$(VERSION)" -o $(BINARY_NAME) .

# Run tests
test:
	@echo "Running tests..."
	$(GO) test -v -race ./internal/...

# Run linters
lint:
	@echo "Running linters..."
	$(GOLINT) run --timeout=5m ./internal/... .

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) -s -w .

# Generate resources from OpenAPI specs
generate: generate-schemas
	@echo "Generation complete"

generate-schemas:
	@echo "Generating schemas from OpenAPI specs..."
	@if [ -d "$(SPEC_DIR)" ] && ls $(SPEC_DIR)/docs-cloud-f5-com.*.ves-swagger.json 1>/dev/null 2>&1; then \
		$(GO) run $(TOOLS_DIR)/generate-all-schemas.go --spec-dir=$(SPEC_DIR); \
	else \
		echo "No OpenAPI specs found in $(SPEC_DIR). Skipping generation."; \
		echo "To generate, download specs to $(SPEC_DIR) or set SPEC_DIR"; \
	fi

# Generate Terraform documentation
docs:
	@echo "Generating Terraform documentation..."
	@if command -v tfplugindocs >/dev/null 2>&1; then \
		tfplugindocs generate; \
	else \
		echo "tfplugindocs not installed. Installing..."; \
		$(GO) install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest; \
		tfplugindocs generate; \
	fi
	@echo "Transforming documentation to Volterra-style format..."
	$(GO) run $(TOOLS_DIR)/transform-docs.go

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f $(BINARY_NAME)
	rm -rf dist/

# Clean all generated files (for full regeneration)
clean-generated:
	@echo "Cleaning generated files..."
	rm -f $(PROVIDER_DIR)/*_resource.go
	rm -f $(PROVIDER_DIR)/*_data_source.go
	rm -f $(PROVIDER_DIR)/provider.go
	rm -f $(CLIENT_DIR)/*_types.go
	@echo "Generated files cleaned. Run 'make generate' to regenerate."

# Full clean rebuild from specs
regenerate: clean-generated generate
	@echo "Full regeneration complete"

# Install provider locally for testing
install: build
	@echo "Installing provider locally..."
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/f5xc/f5xc/$(VERSION)/$(GOOS)_$(GOARCH)
	cp $(BINARY_NAME) ~/.terraform.d/plugins/registry.terraform.io/f5xc/f5xc/$(VERSION)/$(GOOS)_$(GOARCH)/

# Acceptance testing and cleanup
testacc:
	@echo "Running acceptance tests..."
	TF_ACC=1 $(GO) test -v -timeout 120m ./internal/provider/...

# Sweep test resources - clean up orphaned resources from failed tests
# Usage: make sweep
# Environment variables required:
#   - F5XC_API_URL: F5 XC API URL
#   - F5XC_API_P12_FILE and F5XC_P12_PASSWORD (for P12 auth)
#   - OR F5XC_API_CERT and F5XC_API_KEY (for PEM auth)
#   - OR F5XC_API_TOKEN (for token auth)
sweep:
	@echo "Sweeping test resources with prefix 'tf-acc-test-' or 'tf-test-'..."
	@echo "This will delete all matching resources from F5 XC."
	@echo ""
	TF_ACC=1 $(GO) test ./internal/acctest -v -sweep=all -timeout 30m

# Sweep specific resource type
# Usage: make sweep-resource RESOURCE=f5xc_namespace
sweep-resource:
	@if [ -z "$(RESOURCE)" ]; then \
		echo "Error: RESOURCE variable not set"; \
		echo "Usage: make sweep-resource RESOURCE=f5xc_namespace"; \
		exit 1; \
	fi
	@echo "Sweeping $(RESOURCE) resources..."
	TF_ACC=1 $(GO) test ./internal/acctest -v -sweep=$(RESOURCE) -timeout 30m

# CI targets
.PHONY: ci ci-lint ci-test ci-build ci-generate

ci: ci-generate ci-build ci-lint ci-test

ci-generate:
	@echo "CI: Generating schemas (if specs available)..."
	@if [ -d "$(SPEC_DIR)" ] && ls $(SPEC_DIR)/docs-cloud-f5-com.*.ves-swagger.json 1>/dev/null 2>&1; then \
		$(GO) run $(TOOLS_DIR)/generate-all-schemas.go --spec-dir=$(SPEC_DIR); \
	fi

ci-build:
	@echo "CI: Building..."
	$(GO) build -v .

ci-lint:
	@echo "CI: Linting..."
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GOLINT) run --timeout=5m ./internal/... .

ci-test:
	@echo "CI: Testing..."
	$(GO) test -v -race ./internal/...

# Release preparation
.PHONY: release-prep
release-prep: fmt lint test docs
	@echo "Release preparation complete"
	@echo "Ensure all changes are committed before tagging"

# Verify no uncommitted generated changes
.PHONY: verify-generate
verify-generate: generate
	@echo "Verifying no uncommitted changes from generation..."
	@if [ -n "$$(git status --porcelain $(PROVIDER_DIR) $(CLIENT_DIR))" ]; then \
		echo "Error: Generated files have uncommitted changes"; \
		git status --porcelain $(PROVIDER_DIR) $(CLIENT_DIR); \
		exit 1; \
	fi
	@echo "All generated files are up to date"
