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

# LLM Enrichment Configuration
OLLAMA_URL?=http://localhost:11434
OLLAMA_MODEL?=qwen2.5-coder:7b
ENRICHMENT_CACHE=$(TOOLS_DIR)/enriched-descriptions-cache.json

# Go commands
GO=go
GOFMT=gofmt
GOLINT=golangci-lint

.PHONY: all build test lint fmt clean clean-generated regenerate generate enrich-descriptions docs install help sweep sweep-dry-run testacc testacc-mock testacc-real testacc-all test-report test-comprehensive test-comprehensive-mock test-comprehensive-real test-pr-subset

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
	@echo "  make lint         - Run linters"
	@echo "  make fmt          - Format Go code"
	@echo "  make generate     - Generate resources from OpenAPI specs"
	@echo "  make enrich-descriptions - Enrich descriptions using local LLM (ollama)"
	@echo "  make docs         - Generate Terraform documentation"
	@echo "  make clean        - Remove build artifacts"
	@echo "  make install      - Install provider locally"
	@echo ""
	@echo "Acceptance Testing (Categorized):"
	@echo "  make testacc      - Run all acceptance tests (requires F5XC credentials)"
	@echo "  make testacc-real - Run REAL API tests only (TestAcc* prefix)"
	@echo "  make testacc-mock - Run MOCK API tests only (TestMock* prefix)"
	@echo "  make testacc-all  - Run both real and mock tests with report"
	@echo "  make test-report  - Generate test report from last test run"
	@echo ""
	@echo "Comprehensive Testing (CI/CD):"
	@echo "  make test-comprehensive      - Full test suite with professional reports"
	@echo "  make test-comprehensive-mock - Mock tests only (parallel, fast)"
	@echo "  make test-comprehensive-real - Real API tests only (sequential, rate-limited)"
	@echo "  make test-pr-subset          - PR validation (mock tests only)"
	@echo ""
	@echo "Test Categories:"
	@echo "  REAL_API (TestAcc*) - Tests against real F5 XC API endpoints"
	@echo "  MOCK_API (TestMock*) - Tests against local mock server"
	@echo "  UNIT (Test*) - Unit tests without external dependencies"
	@echo ""
	@echo "Test Resource Cleanup:"
	@echo "  make sweep        - Clean up ALL orphaned test resources (prefix-based)"
	@echo "                      WARNING: Deletes any resource with tf-acc-test-* or tf-test-* prefix"
	@echo "                      Use only when no other users are running tests on the same tenant"
	@echo "  make sweep-resource RESOURCE=f5xc_namespace - Sweep specific resource type"
	@echo ""
	@echo "  For SAFE multi-user cleanup, use CleanupTracked() in your test code:"
	@echo "    defer acctest.CleanupTracked()  // Only deletes resources THIS test created"
	@echo ""
	@echo "API Default Discovery (Issue #327):"
	@echo "  make discover-defaults  - Discover API defaults for all resources"
	@echo "  make discover-defaults-resource RESOURCE=xxx - Discover for specific resource"
	@echo "  make validate-defaults  - Validate stored defaults against current API"
	@echo "  make generate-mock-fixtures - Generate mock fixtures from defaults"
	@echo "  make discover-all       - Full pipeline: discover → generate → test"
	@echo ""
	@echo "Environment Variables:"
	@echo "  TF_ACC=1           - Enable real acceptance tests"
	@echo "  F5XC_MOCK_MODE=1   - Enable mock server tests"
	@echo "  SPEC_DIR           - Directory containing OpenAPI specs (default: /tmp)"
	@echo "  F5XC_SPEC_DIR      - Alternative env var for spec directory"
	@echo "  OLLAMA_URL         - Ollama server URL (default: http://localhost:11434)"
	@echo "  OLLAMA_MODEL       - LLM model for description enrichment (default: qwen2.5-coder:7b)"
	@echo ""
	@echo "For real acceptance tests, set one of:"
	@echo "  F5XC_API_URL + F5XC_P12_FILE + F5XC_P12_PASSWORD (P12 auth)"
	@echo "  F5XC_API_URL + F5XC_CERT + F5XC_KEY (PEM auth)"
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

# Enrich descriptions using local LLM (ollama)
# This improves documentation quality by using AI to rewrite poor OpenAPI descriptions
# Requirements: ollama installed and running with qwen2.5-coder:7b model
# The cache is hash-based: only regenerates when schema content changes
enrich-descriptions:
	@echo "Enriching descriptions using LLM..."
	@if [ -d "$(SPEC_DIR)" ] && ls $(SPEC_DIR)/docs-cloud-f5-com.*.ves-swagger.json 1>/dev/null 2>&1; then \
		$(GO) run $(TOOLS_DIR)/enrich-descriptions.go \
			--spec-dir=$(SPEC_DIR) \
			--ollama-url=$(OLLAMA_URL) \
			--model=$(OLLAMA_MODEL) \
			--cache-file=$(ENRICHMENT_CACHE); \
	else \
		echo "No OpenAPI specs found in $(SPEC_DIR). Skipping enrichment."; \
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
# WARNING: This will delete ALL resources matching test prefixes, including
# resources created by other users. For safe multi-user cleanup, use
# CleanupTracked() in your test code instead.
#
# Usage: make sweep
# Environment variables required:
#   - F5XC_API_URL: F5 XC API URL
#   - F5XC_P12_FILE and F5XC_P12_PASSWORD (for P12 auth)
#   - OR F5XC_CERT and F5XC_KEY (for PEM auth)
#   - OR F5XC_API_TOKEN (for token auth)
sweep:
	@echo "⚠️  WARNING: Prefix-based sweep - will delete ALL test resources!"
	@echo "Sweeping resources with prefix 'tf-acc-test-' or 'tf-test-'..."
	@echo "This may delete resources created by other users on the same tenant."
	@echo ""
	@echo "For SAFE multi-user cleanup, use CleanupTracked() in your tests."
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

# =============================================================================
# Categorized Acceptance Tests
# =============================================================================

# Run REAL API tests only (TestAcc* prefix)
# These tests require F5XC credentials and run against the real API
testacc-real:
	@echo "Running REAL API acceptance tests (TestAcc*)..."
	@echo "Category: REAL_API - Tests against real F5 XC API endpoints"
	@echo ""
	TF_ACC=1 $(GO) test -v -timeout 120m ./internal/provider/... -run "^TestAcc" 2>&1 | tee .test-output-real.txt
	@echo ""
	@echo "Test output saved to .test-output-real.txt"

# Run MOCK API tests only (TestMock* prefix)
# These tests use the mock server and don't require real credentials
testacc-mock:
	@echo "Running MOCK API acceptance tests (TestMock*)..."
	@echo "Category: MOCK_API - Tests against local mock server"
	@echo ""
	F5XC_MOCK_MODE=1 $(GO) test -v -timeout 30m ./internal/provider/... -run "^TestMock" 2>&1 | tee .test-output-mock.txt
	@echo ""
	@echo "Test output saved to .test-output-mock.txt"

# Run both real and mock tests with JSON output and generate report
testacc-all:
	@echo "Running ALL acceptance tests (Real + Mock) with categorized report..."
	@echo ""
	@echo "========================================================================"
	@echo "PHASE 1: MOCK API TESTS (no credentials required)"
	@echo "========================================================================"
	F5XC_MOCK_MODE=1 $(GO) test -json -timeout 30m ./internal/provider/... -run "^TestMock" 2>&1 | tee .test-json-mock.txt | $(GO) run $(TOOLS_DIR)/test-report/main.go || true
	@echo ""
	@echo "========================================================================"
	@echo "PHASE 2: REAL API TESTS (requires credentials)"
	@echo "========================================================================"
	@if [ -n "$$F5XC_API_URL" ]; then \
		TF_ACC=1 $(GO) test -json -timeout 120m ./internal/provider/... -run "^TestAcc" 2>&1 | tee .test-json-real.txt | $(GO) run $(TOOLS_DIR)/test-report/main.go || true; \
	else \
		echo "⚠️  Skipping real API tests: F5XC_API_URL not set"; \
	fi
	@echo ""
	@echo "========================================================================"
	@echo "COMBINED REPORT"
	@echo "========================================================================"
	@cat .test-json-mock.txt .test-json-real.txt 2>/dev/null | $(GO) run $(TOOLS_DIR)/test-report/main.go || echo "No test data to report"

# Generate a test report from JSON test output
# Usage: go test -json ./... > test-output.json && make test-report
test-report:
	@echo "Generating test report..."
	@if [ -f ".test-json-mock.txt" ] || [ -f ".test-json-real.txt" ]; then \
		cat .test-json-mock.txt .test-json-real.txt 2>/dev/null | $(GO) run $(TOOLS_DIR)/test-report/main.go; \
	else \
		echo "No test output files found. Run tests with:"; \
		echo "  make testacc-all"; \
		echo "Or manually:"; \
		echo "  go test -json ./internal/provider/... | go run tools/test-report/main.go"; \
	fi

# Generate markdown test report
test-report-md:
	@echo "Generating markdown test report..."
	@cat .test-json-mock.txt .test-json-real.txt 2>/dev/null | $(GO) run $(TOOLS_DIR)/test-report/main.go -format=markdown -output=test-report.md
	@echo "Report saved to test-report.md"

# Clean test output files
clean-test-output:
	@echo "Cleaning test output files..."
	rm -f .test-output-*.txt .test-json-*.txt test-report.md test-report.json

# =============================================================================
# API Default Discovery (Issue #327)
# =============================================================================
# These targets help discover and maintain API default values for resources.
# The discovery process requires VPN access to the F5 XC staging environment.
#
# Local Development:
#   1. Connect to VPN
#   2. Set environment variables (F5XC_API_URL, F5XC_P12_FILE, F5XC_P12_PASSWORD)
#   3. Run: make discover-defaults
#
# CI/CD Note:
#   The discover-defaults.yml workflow uses self-hosted runners with VPN access.
#   Public GitHub runners cannot access the staging environment.

.PHONY: discover-defaults discover-defaults-resource validate-defaults generate-mock-fixtures

# Discover API defaults for all resources
# Requires: VPN access + F5XC_API_URL + F5XC_P12_FILE + F5XC_P12_PASSWORD
discover-defaults:
	@echo "Discovering API defaults for all resources..."
	@echo "This requires VPN access to the F5 XC staging environment."
	@echo ""
	@if [ -z "$$F5XC_API_URL" ]; then \
		echo "Error: F5XC_API_URL not set"; \
		echo "Set F5XC_API_URL, F5XC_P12_FILE, and F5XC_P12_PASSWORD"; \
		exit 1; \
	fi
	$(GO) run $(TOOLS_DIR)/discover-defaults.go -all
	@echo ""
	@echo "Defaults saved to $(TOOLS_DIR)/api-defaults.json"
	@echo "Run 'make generate-mock-fixtures' to update mock fixtures"

# Discover API defaults for a specific resource
# Usage: make discover-defaults-resource RESOURCE=namespace
discover-defaults-resource:
	@if [ -z "$(RESOURCE)" ]; then \
		echo "Error: RESOURCE variable not set"; \
		echo "Usage: make discover-defaults-resource RESOURCE=namespace"; \
		exit 1; \
	fi
	@if [ -z "$$F5XC_API_URL" ]; then \
		echo "Error: F5XC_API_URL not set"; \
		exit 1; \
	fi
	$(GO) run $(TOOLS_DIR)/discover-defaults.go -resource=$(RESOURCE)

# Validate stored defaults against current API
# Compares tools/api-defaults.json against live API responses
validate-defaults:
	@echo "Validating stored API defaults..."
	@if [ -z "$$F5XC_API_URL" ]; then \
		echo "Error: F5XC_API_URL not set"; \
		exit 1; \
	fi
	$(GO) run $(TOOLS_DIR)/discover-defaults.go -validate
	@echo "Validation complete"

# Generate mock fixtures from discovered defaults
# This updates internal/mocks/generated_defaults.go
generate-mock-fixtures:
	@echo "Generating mock fixtures from API defaults..."
	@if [ ! -f "$(TOOLS_DIR)/api-defaults.json" ]; then \
		echo "Error: $(TOOLS_DIR)/api-defaults.json not found"; \
		echo "Run 'make discover-defaults' first"; \
		exit 1; \
	fi
	$(GO) run $(TOOLS_DIR)/generate-mock-fixtures.go
	$(GOFMT) -s -w internal/mocks/generated_defaults.go
	@echo ""
	@echo "Mock fixtures updated. Run 'make test' to verify."

# Full discovery pipeline: discover → generate fixtures → test
discover-all: discover-defaults generate-mock-fixtures test
	@echo ""
	@echo "Full discovery pipeline complete"
	@echo "Review changes and commit if satisfactory"

# =============================================================================
# Comprehensive Testing (CI/CD Ready)
# =============================================================================
# These targets use the comprehensive test runner script that produces
# professional reports in multiple formats (Text, JSON, Markdown, JUnit XML).
#
# Key differences from testacc-* targets:
#   - Mock tests run in PARALLEL (no rate limiting - local tests)
#   - Real API tests run SEQUENTIAL with rate limiting (API protection)
#   - Generates JUnit XML for GitHub Actions test UI
#   - Detects transient errors (rate limit, timeout, connection)
#   - Categorizes skip reasons
#   - Tracks slowest tests for optimization

# Run full comprehensive test suite (mock + real)
# Mock tests run parallel, real API tests run sequential with rate limiting
test-comprehensive:
	@echo "Running comprehensive test suite..."
	./scripts/run-comprehensive-tests.sh --mode full

# Run mock tests only - PARALLEL (fast, no rate limiting needed)
test-comprehensive-mock:
	@echo "Running comprehensive mock tests (parallel)..."
	./scripts/run-comprehensive-tests.sh --mode mock-only

# Run real API tests only - SEQUENTIAL with rate limiting
# Requires: F5XC_API_URL, F5XC_P12_FILE, F5XC_P12_PASSWORD (or F5XC_API_TOKEN)
test-comprehensive-real:
	@echo "Running comprehensive real API tests (sequential)..."
	./scripts/run-comprehensive-tests.sh --mode real-only

# Run PR subset tests - mock tests only for PR validation
test-pr-subset:
	@echo "Running PR subset tests (mock only)..."
	./scripts/run-comprehensive-tests.sh --mode pr-subset

# Clean comprehensive test reports
clean-test-reports:
	@echo "Cleaning test reports..."
	rm -rf test-reports/
