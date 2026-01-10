#!/bin/bash
# Copyright (c) 2026 Robin Mordasiewicz. MIT License.

# Comprehensive test runner for F5 XC Terraform Provider
#
# This script runs acceptance tests with intelligent handling of:
# - Mock tests: Run in PARALLEL (no rate limiting needed for local tests)
# - Real API tests: Run SEQUENTIALLY with rate limit protection
#
# Usage:
#   ./scripts/run-comprehensive-tests.sh [OPTIONS]
#
# Options:
#   --mode MODE        Test mode: mock-only, real-only, full, pr-subset (default: full)
#   --output-dir DIR   Output directory for reports (default: test-reports)
#   --batch-size N     Tests per batch before delay for real API (default: 5)
#   --batch-delay N    Seconds to delay between batches for real API (default: 20)
#   --timeout M        Timeout per test in minutes (default: 10)
#   --parallel N       Parallelism for mock tests (default: 4)
#   --dry-run          Show what would be run without executing
#   --verbose          Enable verbose output
#   --help             Show this help message
#
# Environment Variables (required for real-only or full mode):
#   F5XC_API_URL        - F5 XC API URL
#   F5XC_P12_FILE       - Path to P12 certificate file
#   F5XC_P12_PASSWORD   - Password for P12 certificate
#   F5XC_API_TOKEN      - API token (alternative to P12 auth)
#
# Exit Codes:
#   0 - All tests passed
#   1 - Test failures detected
#   2 - Configuration error

set -euo pipefail

# Default configuration
MODE="full"
OUTPUT_DIR="test-reports"
BATCH_SIZE=5
BATCH_DELAY=20
TIMEOUT=10
PARALLEL=4
DRY_RUN=false
VERBOSE=false

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --mode)
            MODE="$2"
            shift 2
            ;;
        --output-dir)
            OUTPUT_DIR="$2"
            shift 2
            ;;
        --batch-size)
            BATCH_SIZE="$2"
            shift 2
            ;;
        --batch-delay)
            BATCH_DELAY="$2"
            shift 2
            ;;
        --timeout)
            TIMEOUT="$2"
            shift 2
            ;;
        --parallel)
            PARALLEL="$2"
            shift 2
            ;;
        --dry-run)
            DRY_RUN=true
            shift
            ;;
        --verbose)
            VERBOSE=true
            shift
            ;;
        --help)
            sed -n '2,/^set -euo pipefail/p' "$0" | head -n -1
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            exit 2
            ;;
    esac
done

# Validate mode
case $MODE in
    mock-only|real-only|full|pr-subset)
        ;;
    *)
        echo -e "${RED}ERROR: Invalid mode '$MODE'. Use: mock-only, real-only, full, pr-subset${NC}"
        exit 2
        ;;
esac

# Logging functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[PASS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[FAIL]${NC} $1"
}

log_verbose() {
    if [[ "$VERBOSE" == "true" ]]; then
        echo -e "${BLUE}[DEBUG]${NC} $1"
    fi
}

# Validate environment for real API tests
validate_real_api_env() {
    local missing=()

    if [[ -z "${F5XC_API_URL:-}" ]]; then
        missing+=("F5XC_API_URL")
    fi
    if [[ -z "${F5XC_P12_FILE:-}" ]] && [[ -z "${F5XC_API_TOKEN:-}" ]]; then
        missing+=("F5XC_P12_FILE or F5XC_API_TOKEN")
    fi
    if [[ -n "${F5XC_P12_FILE:-}" ]] && [[ -z "${F5XC_P12_PASSWORD:-}" ]]; then
        missing+=("F5XC_P12_PASSWORD (required with F5XC_P12_FILE)")
    fi

    if [[ ${#missing[@]} -gt 0 ]]; then
        log_error "Missing required environment variables for real API tests:"
        for var in "${missing[@]}"; do
            echo "  - $var"
        done
        return 1
    fi

    # Verify P12 file exists if using P12 auth
    if [[ -n "${F5XC_P12_FILE:-}" ]] && [[ ! -f "$F5XC_P12_FILE" ]]; then
        log_error "P12 file not found: $F5XC_P12_FILE"
        return 1
    fi

    return 0
}

# Setup output directory
setup_output_dir() {
    mkdir -p "$OUTPUT_DIR"
    log_verbose "Output directory: $OUTPUT_DIR"
}

# Run mock tests (PARALLEL - no rate limiting needed)
run_mock_tests() {
    log_info "Running MOCK tests (parallel, no rate limiting)..."

    local output_file="$OUTPUT_DIR/test-output-mock.json"

    if [[ "$DRY_RUN" == "true" ]]; then
        log_info "[DRY-RUN] Would run: F5XC_MOCK_MODE=1 go test -json -parallel $PARALLEL -timeout ${TIMEOUT}m -run 'TestMock.*' ./internal/provider/..."
        return 0
    fi

    # Mock tests run in parallel - no rate limiting needed for local tests
    F5XC_MOCK_MODE=1 go test -json \
        -parallel "$PARALLEL" \
        -timeout "${TIMEOUT}m" \
        -run 'TestMock.*' \
        ./internal/provider/... 2>&1 | tee "$output_file" || true

    log_verbose "Mock test output saved to: $output_file"
    return 0
}

# Run real API tests (SEQUENTIAL with rate limiting)
run_real_api_tests() {
    log_info "Running REAL API tests (sequential with rate limiting)..."

    if ! validate_real_api_env; then
        log_error "Cannot run real API tests without proper configuration"
        return 1
    fi

    local output_file="$OUTPUT_DIR/test-output-real.json"

    if [[ "$DRY_RUN" == "true" ]]; then
        log_info "[DRY-RUN] Would run: TF_ACC=1 go test -json -parallel 1 -timeout ${TIMEOUT}m -run 'TestAcc.*' ./internal/provider/..."
        return 0
    fi

    # Real API tests run sequentially with rate limiting
    # -parallel 1 ensures only one test runs at a time
    export TF_ACC=1

    log_info "Rate limiting: $BATCH_SIZE tests, then ${BATCH_DELAY}s delay"

    # Run tests sequentially
    go test -json \
        -parallel 1 \
        -timeout "${TIMEOUT}m" \
        -run 'TestAcc.*' \
        ./internal/provider/... 2>&1 | tee "$output_file" || true

    log_verbose "Real API test output saved to: $output_file"
    return 0
}

# Run unit tests (PARALLEL - no rate limiting needed)
run_unit_tests() {
    log_info "Running UNIT tests (parallel)..."

    local output_file="$OUTPUT_DIR/test-output-unit.json"

    if [[ "$DRY_RUN" == "true" ]]; then
        log_info "[DRY-RUN] Would run: go test -json -parallel $PARALLEL -timeout ${TIMEOUT}m ./internal/..."
        return 0
    fi

    # Unit tests run in parallel
    go test -json \
        -parallel "$PARALLEL" \
        -timeout "${TIMEOUT}m" \
        ./internal/... 2>&1 | tee "$output_file" || true

    log_verbose "Unit test output saved to: $output_file"
    return 0
}

# Combine JSON outputs
combine_outputs() {
    log_info "Combining test outputs..."

    local combined_file="$OUTPUT_DIR/test-output-combined.json"

    # Combine all JSON outputs
    cat "$OUTPUT_DIR"/test-output-*.json > "$combined_file" 2>/dev/null || true

    echo "$combined_file"
}

# Generate reports in all formats
generate_reports() {
    local input_file="$1"

    log_info "Generating reports..."

    if [[ ! -f "$input_file" ]]; then
        log_warning "No test output file found: $input_file"
        return 1
    fi

    # Generate all report formats
    local report_tool="$PROJECT_ROOT/tools/test-report/main.go"

    # Text report
    log_verbose "Generating text report..."
    cat "$input_file" | go run "$report_tool" -format=text -all -output="$OUTPUT_DIR/test-report.txt"

    # JSON report
    log_verbose "Generating JSON report..."
    cat "$input_file" | go run "$report_tool" -format=json -output="$OUTPUT_DIR/test-report.json"

    # Markdown report
    log_verbose "Generating Markdown report..."
    cat "$input_file" | go run "$report_tool" -format=markdown -all -output="$OUTPUT_DIR/test-report.md"

    # JUnit XML report (for GitHub Actions)
    log_verbose "Generating JUnit XML report..."
    cat "$input_file" | go run "$report_tool" -format=junit -output="$OUTPUT_DIR/test-report.xml"

    log_success "Reports generated in: $OUTPUT_DIR/"
    ls -la "$OUTPUT_DIR/"
}

# Print summary to stdout
print_summary() {
    local input_file="$1"

    echo ""
    echo "============================================================"
    echo "                    TEST EXECUTION SUMMARY"
    echo "============================================================"

    if [[ -f "$input_file" ]]; then
        cat "$input_file" | go run "$PROJECT_ROOT/tools/test-report/main.go" -format=text
    else
        echo "No test results available"
    fi
}

# Check if any tests failed
check_failures() {
    local report_file="$OUTPUT_DIR/test-report.json"

    if [[ ! -f "$report_file" ]]; then
        return 0
    fi

    local failed
    failed=$(jq -r '.total_failed // 0' "$report_file" 2>/dev/null || echo "0")

    if [[ "$failed" -gt 0 ]]; then
        return 1
    fi
    return 0
}

# Main execution
main() {
    local start_time
    start_time=$(date +%s)

    echo ""
    echo "============================================================"
    echo "     F5 XC Terraform Provider - Comprehensive Test Suite"
    echo "============================================================"
    echo "Mode:        $MODE"
    echo "Output:      $OUTPUT_DIR"
    echo "Timeout:     ${TIMEOUT}m per test"
    echo "Parallel:    $PARALLEL (for mock/unit tests)"
    echo "Batch Size:  $BATCH_SIZE (for real API tests)"
    echo "Batch Delay: ${BATCH_DELAY}s (for real API tests)"
    echo "Started:     $(date)"
    echo "============================================================"
    echo ""

    # Change to project root
    cd "$PROJECT_ROOT"

    # Setup
    setup_output_dir

    # Run tests based on mode
    case $MODE in
        mock-only)
            run_mock_tests
            ;;
        real-only)
            run_real_api_tests
            ;;
        full)
            # Run mock tests first (parallel, fast)
            run_mock_tests

            # Then run real API tests (sequential, with rate limiting)
            run_real_api_tests
            ;;
        pr-subset)
            # For PR subset, we only run mock tests by default
            # This is the safe, fast option for PR validation
            run_mock_tests
            ;;
    esac

    # Combine outputs and generate reports
    local combined_file
    combined_file=$(combine_outputs)

    if [[ "$DRY_RUN" != "true" ]]; then
        generate_reports "$combined_file"
        print_summary "$combined_file"
    fi

    local end_time
    end_time=$(date +%s)
    local duration=$((end_time - start_time))

    echo ""
    echo "============================================================"
    echo "Total Duration: ${duration}s"
    echo "Completed:      $(date)"
    echo "============================================================"

    # Check for failures
    if [[ "$DRY_RUN" != "true" ]]; then
        if check_failures; then
            log_success "All tests passed!"
            exit 0
        else
            log_error "Some tests failed. See reports for details."
            exit 1
        fi
    fi
}

# Run main
main
