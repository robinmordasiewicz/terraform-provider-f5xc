#!/bin/bash
# =============================================================================
# Update Subscription Tier Metadata
# =============================================================================
# This script runs the subscription metadata generator to update
# tools/subscription-tiers.json with the latest tier information from the
# F5 XC Catalog API.
#
# When run as a pre-commit hook:
# - Fetches latest subscription tier data from F5 XC API (requires VPN)
# - If data has changed, automatically adds the file to the current commit
# - Skips gracefully if credentials are unavailable
#
# The generated file triggers documentation regeneration in CI/CD when changed.
#
# Required Environment Variables:
#   F5XC_API_URL      - F5 XC API URL (e.g., https://console.ves.volterra.io)
#   F5XC_API_TOKEN    - API token (preferred for automation)
#   OR
#   F5XC_P12_FILE     - Path to P12 certificate file
#   F5XC_P12_PASSWORD - Password for P12 file
#
# Usage:
#   ./scripts/update-subscription-metadata.sh [--check]
#
# Options:
#   --check  Only check if updates are needed, don't modify files (for CI)
# =============================================================================

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"
METADATA_FILE="${REPO_ROOT}/tools/subscription-tiers.json"
GENERATOR="${REPO_ROOT}/tools/generate-subscription-metadata.go"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

check_mode=false
if [[ "${1:-}" == "--check" ]]; then
    check_mode=true
fi

# Check if credentials are available
has_credentials() {
    if [[ -n "${F5XC_API_TOKEN:-}" && -n "${F5XC_API_URL:-}" ]]; then
        return 0
    elif [[ -n "${F5XC_P12_FILE:-}" && -n "${F5XC_P12_PASSWORD:-}" && -n "${F5XC_API_URL:-}" ]]; then
        return 0
    fi
    return 1
}

# Main logic
main() {
    echo "=== Subscription Tier Metadata Update ==="

    # Check for credentials
    if ! has_credentials; then
        echo -e "${YELLOW}SKIP: F5XC credentials not configured${NC}"
        echo "Set F5XC_API_URL and either F5XC_API_TOKEN or F5XC_P12_FILE+F5XC_P12_PASSWORD"
        echo "to enable subscription tier metadata generation."
        exit 0
    fi

    # Check if generator exists
    if [[ ! -f "${GENERATOR}" ]]; then
        echo -e "${RED}ERROR: Generator not found: ${GENERATOR}${NC}"
        exit 1
    fi

    # Store current file hash (if exists)
    old_hash=""
    if [[ -f "${METADATA_FILE}" ]]; then
        old_hash=$(md5sum "${METADATA_FILE}" 2>/dev/null | cut -d' ' -f1 || echo "")
    fi

    # Run the generator with timeout (advisory hook - should not block commits)
    echo "Running subscription metadata generator..."
    cd "${REPO_ROOT}"

    # Use timeout to prevent blocking on network issues (30 second limit)
    # On timeout or network errors, skip gracefully since this is advisory
    TIMEOUT_CMD=""
    if command -v timeout &> /dev/null; then
        TIMEOUT_CMD="timeout 30"
    elif command -v gtimeout &> /dev/null; then
        TIMEOUT_CMD="gtimeout 30"  # macOS with coreutils
    fi

    if [[ -n "${TIMEOUT_CMD}" ]]; then
        if ! ${TIMEOUT_CMD} go run "${GENERATOR}" 2>&1; then
            exit_code=$?
            if [[ ${exit_code} -eq 124 ]]; then
                echo -e "${YELLOW}SKIP: Generator timed out (network issues)${NC}"
                echo "This is an advisory hook - continuing without update."
                exit 0
            fi
            # Check for common network error patterns
            echo -e "${YELLOW}SKIP: Generator failed (likely network issues)${NC}"
            echo "This is an advisory hook - continuing without update."
            exit 0
        fi
    else
        if ! go run "${GENERATOR}" 2>&1; then
            echo -e "${YELLOW}SKIP: Generator failed${NC}"
            echo "This is an advisory hook - continuing without update."
            exit 0
        fi
    fi

    # Check if file was created/updated
    if [[ ! -f "${METADATA_FILE}" ]]; then
        echo -e "${YELLOW}SKIP: Generator did not create ${METADATA_FILE}${NC}"
        echo "This is an advisory hook - continuing without update."
        exit 0
    fi

    # Compare hashes
    new_hash=$(md5sum "${METADATA_FILE}" 2>/dev/null | cut -d' ' -f1 || echo "")

    if [[ "${old_hash}" == "${new_hash}" ]]; then
        echo -e "${GREEN}No changes detected in subscription tier metadata${NC}"
        exit 0
    else
        if [[ "${check_mode}" == "true" ]]; then
            echo -e "${YELLOW}Changes detected in subscription tier metadata!${NC}"
            echo "Run: ./scripts/update-subscription-metadata.sh"
            echo "Then commit the updated tools/subscription-tiers.json"
            exit 1
        else
            echo -e "${GREEN}Updated subscription tier metadata${NC}"

            # Show summary of changes
            if command -v jq &> /dev/null; then
                echo ""
                echo "Summary:"
                services=$(jq '.services | length' "${METADATA_FILE}" 2>/dev/null || echo "?")
                resources=$(jq '.resources | length' "${METADATA_FILE}" 2>/dev/null || echo "?")
                advanced=$(jq '[.resources | to_entries[] | select(.value.minimum_tier == "ADVANCED")] | length' "${METADATA_FILE}" 2>/dev/null || echo "?")
                echo "  Services: ${services}"
                echo "  Resources: ${resources}"
                echo "  Advanced tier resources: ${advanced}"
            fi

            # Auto-add to current commit (for pre-commit hook integration)
            if git rev-parse --git-dir > /dev/null 2>&1; then
                git add "${METADATA_FILE}"
                echo -e "${GREEN}Added ${METADATA_FILE} to commit${NC}"
            fi
            exit 0
        fi
    fi
}

main "$@"
