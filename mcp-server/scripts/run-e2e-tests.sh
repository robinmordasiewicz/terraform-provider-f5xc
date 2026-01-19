#!/usr/bin/env bash
# Copyright (c) 2026 Robin Mordasiewicz. MIT License.

# E2E Determinism Test Runner
# Validates MCP server produces deterministic, production-ready Terraform
#
# Prerequisites:
#   - F5XC_API_URL and F5XC_API_TOKEN environment variables
#   - Terraform CLI installed
#   - Node.js 18+
#
# Usage:
#   ./scripts/run-e2e-tests.sh [--verbose] [--skip-build]

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Flags
VERBOSE=false
SKIP_BUILD=false

# Parse arguments
while [[ $# -gt 0 ]]; do
  case $1 in
    --verbose|-v)
      VERBOSE=true
      shift
      ;;
    --skip-build)
      SKIP_BUILD=true
      shift
      ;;
    --help|-h)
      echo "Usage: $0 [--verbose] [--skip-build]"
      echo ""
      echo "Options:"
      echo "  --verbose, -v    Show detailed test output"
      echo "  --skip-build     Skip npm build step"
      echo "  --help, -h       Show this help message"
      exit 0
      ;;
    *)
      echo "Unknown option: $1"
      exit 1
      ;;
  esac
done

log_info() {
  echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
  echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warn() {
  echo -e "${YELLOW}[WARN]${NC} $1"
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
  NODE_VERSION=$(node --version | cut -d'v' -f2 | cut -d'.' -f1)
  if [[ $NODE_VERSION -lt 18 ]]; then
    log_error "Node.js 18+ required, found $(node --version)"
    exit 1
  fi
  log_success "Node.js $(node --version)"

  # Check npm
  if ! command -v npm &> /dev/null; then
    log_error "npm is not installed"
    exit 1
  fi
  log_success "npm $(npm --version)"

  # Check Terraform
  if ! command -v terraform &> /dev/null; then
    log_warn "Terraform CLI not installed - some tests will be skipped"
  else
    log_success "Terraform $(terraform version -json | jq -r '.terraform_version')"
  fi

  # Check F5XC credentials
  if [[ -z "${F5XC_API_URL:-}" ]] && [[ -z "${F5XC_TEST_API_URL:-}" ]]; then
    log_warn "F5XC_API_URL not set - E2E tests requiring API will be skipped"
  else
    local api_url="${F5XC_API_URL:-${F5XC_TEST_API_URL}}"
    log_success "F5XC_API_URL: ${api_url}"
  fi

  if [[ -z "${F5XC_API_TOKEN:-}" ]] && [[ -z "${F5XC_TEST_API_TOKEN:-}" ]]; then
    log_warn "F5XC_API_TOKEN not set - E2E tests requiring API will be skipped"
  else
    log_success "F5XC_API_TOKEN: ***configured***"
  fi

  echo ""
}

# Build project
build_project() {
  if [[ "$SKIP_BUILD" == "true" ]]; then
    log_info "Skipping build (--skip-build flag)"
    return
  fi

  log_info "Building MCP server..."
  cd "$PROJECT_DIR"

  npm ci --silent
  npm run build

  log_success "Build complete"
  echo ""
}

# Run unit tests
run_unit_tests() {
  log_info "Running unit tests..."
  cd "$PROJECT_DIR"

  if [[ "$VERBOSE" == "true" ]]; then
    npm run test -- tests/unit/ --reporter=verbose
  else
    npm run test -- tests/unit/ --reporter=dot
  fi

  log_success "Unit tests passed"
  echo ""
}

# Run determinism tests
run_determinism_tests() {
  log_info "Running determinism tests..."
  cd "$PROJECT_DIR"

  if [[ "$VERBOSE" == "true" ]]; then
    npm run test:e2e:determinism -- --reporter=verbose
  else
    npm run test:e2e:determinism -- --reporter=default
  fi

  log_success "Determinism tests passed"
  echo ""
}

# Run resource creation tests (if credentials available)
run_resource_creation_tests() {
  local api_url="${F5XC_API_URL:-${F5XC_TEST_API_URL:-}}"
  local api_token="${F5XC_API_TOKEN:-${F5XC_TEST_API_TOKEN:-}}"

  if [[ -z "$api_url" ]] || [[ -z "$api_token" ]]; then
    log_warn "Skipping resource creation tests (no credentials)"
    return
  fi

  log_info "Running resource creation tests..."
  cd "$PROJECT_DIR"

  if [[ "$VERBOSE" == "true" ]]; then
    npm run test:e2e:resource-creation -- --reporter=verbose
  else
    npm run test:e2e:resource-creation -- --reporter=default
  fi

  log_success "Resource creation tests passed"
  echo ""
}

# Generate ML training matrix
generate_ml_matrix() {
  log_info "Generating ML training matrix..."
  cd "$PROJECT_DIR"

  if [[ -f "package.json" ]] && grep -q "generate:ml-matrix" package.json; then
    npm run generate:ml-matrix
    log_success "ML training matrix generated"
  else
    log_warn "generate:ml-matrix script not found, skipping"
  fi
  echo ""
}

# Main execution
main() {
  echo ""
  echo "=========================================="
  echo " F5XC Terraform MCP - E2E Test Suite"
  echo "=========================================="
  echo ""

  check_prerequisites
  build_project
  run_unit_tests
  run_determinism_tests
  run_resource_creation_tests
  generate_ml_matrix

  echo "=========================================="
  log_success "All E2E tests completed successfully!"
  echo "=========================================="
}

main "$@"
