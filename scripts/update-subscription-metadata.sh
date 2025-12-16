#!/bin/bash
# =============================================================================
# Update Subscription Tier Metadata
# =============================================================================
# This script runs the subscription metadata generator to update
# tools/subscription-tiers.json with the latest tier information from the
# F5 XC Catalog API.
#
# The generated file is committed to the repository and changes trigger
# documentation regeneration in CI/CD.
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

    # Run the generator
    echo "Running subscription metadata generator..."
    cd "${REPO_ROOT}"

    if ! go run "${GENERATOR}"; then
        echo -e "${RED}ERROR: Generator failed${NC}"
        exit 1
    fi

    # Check if file was created/updated
    if [[ ! -f "${METADATA_FILE}" ]]; then
        echo -e "${RED}ERROR: Generator did not create ${METADATA_FILE}${NC}"
        exit 1
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
            echo "Changes detected - please commit tools/subscription-tiers.json"

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
            exit 0
        fi
    fi
}

main "$@"
