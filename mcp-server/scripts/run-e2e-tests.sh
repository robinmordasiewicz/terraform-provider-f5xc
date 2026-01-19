#!/bin/bash
# Copyright (c) 2026 Robin Mordasiewicz. MIT License.

#
# F5XC Terraform MCP Server - E2E Test Runner
#
# This script runs the complete E2E determinism validation suite.
#
# Prerequisites:
#   - Node.js 18+
#   - Terraform CLI
#   - F5XC API credentials
#
# Usage:
#   # Set credentials
#   export F5XC_API_URL="https://f5-amer-ent.console.ves.volterra.io"
#   export F5XC_API_TOKEN="your-api-token"
#
#   # Run all E2E tests
#   ./scripts/run-e2e-tests.sh
#
#   # Run only determinism tests (no apply/destroy)
#   ./scripts/run-e2e-tests.sh --determinism-only
#
#   # Run only full lifecycle tests (plan/apply/destroy)
#   ./scripts/run-e2e-tests.sh --lifecycle-only
#
#   # Generate ML training matrix
#   ./scripts/run-e2e-tests.sh --generate-matrix
#

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MCP_SERVER_DIR="$(dirname "$SCRIPT_DIR")"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check prerequisites
check_prerequisites() {
    log_info "Checking prerequisites..."

    # Check Node.js
    if ! command -v node &> /dev/null; then
        log_error "Node.js is not installed"
        exit 1
    fi
    NODE_VERSION=$(node --version)
    log_info "Node.js: $NODE_VERSION"

    # Check npm
    if ! command -v npm &> /dev/null; then
        log_error "npm is not installed"
        exit 1
    fi
    NPM_VERSION=$(npm --version)
    log_info "npm: $NPM_VERSION"

    # Check Terraform
    if ! command -v terraform &> /dev/null; then
        log_warning "Terraform CLI is not installed - some tests will be skipped"
    else
        TF_VERSION=$(terraform version -json | jq -r '.terraform_version')
        log_info "Terraform: $TF_VERSION"
    fi

    # Check credentials (support both F5XC_API_* and F5XC_TEST_API_* naming)
    if [ -z "$F5XC_API_URL" ] && [ -z "$F5XC_TEST_API_URL" ]; then
        log_warning "F5XC_API_URL not set - E2E tests will be skipped"
        log_info "Set with: export F5XC_API_URL=\"https://f5-amer-ent.console.ves.volterra.io\""
    else
        log_info "F5XC_API_URL: ${F5XC_API_URL:-$F5XC_TEST_API_URL}"
    fi

    if [ -z "$F5XC_API_TOKEN" ] && [ -z "$F5XC_TEST_API_TOKEN" ]; then
        log_warning "F5XC_API_TOKEN not set - E2E tests will be skipped"
    else
        log_info "F5XC_API_TOKEN: [set]"
    fi

    log_success "Prerequisites check complete"
}

# Build MCP server
build_server() {
    log_info "Building MCP server..."
    cd "$MCP_SERVER_DIR"
    npm ci
    npm run build
    npm run typecheck
    log_success "Build complete"
}

# Run unit tests
run_unit_tests() {
    log_info "Running unit tests..."
    cd "$MCP_SERVER_DIR"
    npm run test -- tests/unit/ --reporter=verbose
    log_success "Unit tests complete"
}

# Run determinism tests (terraform validate only, no apply)
run_determinism_tests() {
    log_info "Running determinism tests..."
    cd "$MCP_SERVER_DIR"
    npm run test:e2e:determinism -- --reporter=verbose
    log_success "Determinism tests complete"
}

# Run full lifecycle tests (plan/apply/destroy)
run_lifecycle_tests() {
    log_info "Running full lifecycle tests..."
    cd "$MCP_SERVER_DIR"
    npm run test:e2e:resource-creation -- --reporter=verbose
    log_success "Lifecycle tests complete"
}

# Generate ML training matrix
generate_ml_matrix() {
    log_info "Generating ML training matrix..."
    cd "$MCP_SERVER_DIR"
    npx tsx tests/fixtures/ml-training/generate-matrix.ts
    log_success "ML training matrix generated"
}

# Test MCP server JSON-RPC
test_mcp_jsonrpc() {
    log_info "Testing MCP server JSON-RPC..."
    cd "$MCP_SERVER_DIR"
    RESPONSE=$(echo '{"jsonrpc":"2.0","id":1,"method":"tools/list"}' | timeout 10 node dist/index.js 2>/dev/null)
    if echo "$RESPONSE" | jq -e '.result.tools' > /dev/null 2>&1; then
        TOOL_COUNT=$(echo "$RESPONSE" | jq '.result.tools | length')
        log_success "MCP server responds with $TOOL_COUNT tools"
    else
        log_error "MCP server JSON-RPC test failed"
        exit 1
    fi
}

# Print usage
print_usage() {
    echo "F5XC Terraform MCP Server - E2E Test Runner"
    echo ""
    echo "Usage: $0 [options]"
    echo ""
    echo "Options:"
    echo "  --all              Run all tests (default)"
    echo "  --unit-only        Run only unit tests"
    echo "  --determinism-only Run only determinism tests (no apply/destroy)"
    echo "  --lifecycle-only   Run only full lifecycle tests (plan/apply/destroy)"
    echo "  --generate-matrix  Generate ML training matrix"
    echo "  --skip-build       Skip npm build step"
    echo "  --help             Show this help message"
    echo ""
    echo "Environment Variables:"
    echo "  F5XC_API_URL   F5XC API URL (required for E2E tests)"
    echo "  F5XC_API_TOKEN F5XC API token (required for E2E tests)"
    echo ""
    echo "Examples:"
    echo "  # Run all tests with credentials"
    echo "  export F5XC_API_URL=\"https://f5-amer-ent.console.ves.volterra.io\""
    echo "  export F5XC_API_TOKEN=\"your-token\""
    echo "  $0"
    echo ""
    echo "  # Run only unit tests (no credentials needed)"
    echo "  $0 --unit-only"
}

# Main
main() {
    local run_all=false
    local run_unit=false
    local run_determinism=false
    local run_lifecycle=false
    local generate_matrix=false
    local skip_build=false

    # Parse arguments
    while [[ $# -gt 0 ]]; do
        case $1 in
            --all)
                run_all=true
                shift
                ;;
            --unit-only)
                run_unit=true
                shift
                ;;
            --determinism-only)
                run_determinism=true
                shift
                ;;
            --lifecycle-only)
                run_lifecycle=true
                shift
                ;;
            --generate-matrix)
                generate_matrix=true
                shift
                ;;
            --skip-build)
                skip_build=true
                shift
                ;;
            --help)
                print_usage
                exit 0
                ;;
            *)
                log_error "Unknown option: $1"
                print_usage
                exit 1
                ;;
        esac
    done

    # Default to run all if no specific option given
    if [ "$run_unit" = false ] && [ "$run_determinism" = false ] && [ "$run_lifecycle" = false ] && [ "$generate_matrix" = false ]; then
        run_all=true
    fi

    echo ""
    echo "=============================================="
    echo "F5XC Terraform MCP Server - E2E Test Suite"
    echo "=============================================="
    echo ""

    check_prerequisites

    if [ "$skip_build" = false ]; then
        build_server
    fi

    test_mcp_jsonrpc

    if [ "$run_all" = true ] || [ "$run_unit" = true ]; then
        run_unit_tests
    fi

    if [ "$run_all" = true ] || [ "$run_determinism" = true ]; then
        run_determinism_tests
    fi

    if [ "$run_all" = true ] || [ "$run_lifecycle" = true ]; then
        run_lifecycle_tests
    fi

    if [ "$generate_matrix" = true ]; then
        generate_ml_matrix
    fi

    echo ""
    echo "=============================================="
    log_success "All tests complete!"
    echo "=============================================="
}

main "$@"
