#!/bin/bash
# Pre-commit hook to lint generated files when generators change
# This catches linting errors early without violating the constitution
#
# Workflow:
# 1. Detect if generator tools are being committed
# 2. Temporarily run generators to produce output
# 3. Lint the generated artifacts
# 4. Restore generated files to original state
# 5. Report linting errors and block commit if needed
#
# This does NOT commit generated files - it only previews what CI/CD will produce
# and validates that it will pass linting.

set -e

# ANSI color codes
RED='\033[0;31m'
YELLOW='\033[1;33m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'
BOLD='\033[1m'
NC='\033[0m' # No Color

# Configuration
SPEC_DIR="docs/specifications/api"
TOOLS_DIR="tools"
PROVIDER_DIR="internal/provider"
CLIENT_DIR="internal/client"

# Generator tools that produce Go code requiring linting
GO_GENERATORS=(
    "tools/generate-all-schemas.go"
)

echo "🔍 Checking for generator tool changes..."

# Get list of staged files
STAGED_FILES=$(git diff --cached --name-only --diff-filter=ACM)

if [ -z "$STAGED_FILES" ]; then
    echo -e "${GREEN}✅ No files staged${NC}"
    exit 0
fi

# Check if any Go generators are being modified
MODIFIED_GO_GENERATORS=()
for generator in "${GO_GENERATORS[@]}"; do
    if echo "$STAGED_FILES" | grep -qE "^${generator}$"; then
        MODIFIED_GO_GENERATORS+=("$generator")
    fi
done

# Also check for spec changes that would trigger regeneration
SPEC_CHANGES=false
if echo "$STAGED_FILES" | grep -qE "^${SPEC_DIR}/"; then
    SPEC_CHANGES=true
fi

# If no generators or specs modified, nothing to preview
if [ ${#MODIFIED_GO_GENERATORS[@]} -eq 0 ] && [ "$SPEC_CHANGES" = false ]; then
    echo -e "${GREEN}✅ No generator tools or specs modified - skipping preview lint${NC}"
    exit 0
fi

# Report what triggered the preview
echo ""
echo -e "${CYAN}╭────────────────────────────────────────────────────────────────────────────╮${NC}"
echo -e "${CYAN}│              GENERATOR CHANGES DETECTED - PREVIEW LINTING                 │${NC}"
echo -e "${CYAN}╰────────────────────────────────────────────────────────────────────────────╯${NC}"
echo ""

if [ ${#MODIFIED_GO_GENERATORS[@]} -gt 0 ]; then
    echo -e "${YELLOW}Modified generators:${NC}"
    for gen in "${MODIFIED_GO_GENERATORS[@]}"; do
        echo "  • $gen"
    done
fi

if [ "$SPEC_CHANGES" = true ]; then
    echo -e "${YELLOW}OpenAPI spec changes detected in:${NC} ${SPEC_DIR}/"
fi
echo ""

# Check if specs exist
if [ ! -d "$SPEC_DIR" ] || ! ls "$SPEC_DIR"/docs-cloud-f5-com.*.ves-swagger.json 1>/dev/null 2>&1; then
    echo -e "${YELLOW}⚠️  No OpenAPI specs found in ${SPEC_DIR}${NC}"
    echo "   Cannot run preview generation without specs."
    echo "   Skipping preview lint - CI/CD will handle generation."
    exit 0
fi

# Check if golangci-lint is available
if ! command -v golangci-lint &> /dev/null; then
    echo -e "${YELLOW}⚠️  golangci-lint not installed${NC}"
    echo "   Install it: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
    echo "   Skipping preview lint - CI/CD will run linting."
    exit 0
fi

echo -e "${BOLD}Running preview generation and lint...${NC}"
echo ""

# Save current state of generated files
TEMP_BACKUP_DIR=$(mktemp -d)
RESTORE_NEEDED=false

# Function to save generated files
save_generated_files() {
    echo "📦 Saving current state of generated files..."

    # Save resource files
    if ls ${PROVIDER_DIR}/*_resource.go 1>/dev/null 2>&1; then
        mkdir -p "$TEMP_BACKUP_DIR/provider"
        cp ${PROVIDER_DIR}/*_resource.go "$TEMP_BACKUP_DIR/provider/" 2>/dev/null || true
        RESTORE_NEEDED=true
    fi

    # Save data source files
    if ls ${PROVIDER_DIR}/*_data_source.go 1>/dev/null 2>&1; then
        mkdir -p "$TEMP_BACKUP_DIR/provider"
        cp ${PROVIDER_DIR}/*_data_source.go "$TEMP_BACKUP_DIR/provider/" 2>/dev/null || true
        RESTORE_NEEDED=true
    fi

    # Save client types (if they exist)
    if ls ${CLIENT_DIR}/*_types.go 1>/dev/null 2>&1; then
        mkdir -p "$TEMP_BACKUP_DIR/client"
        cp ${CLIENT_DIR}/*_types.go "$TEMP_BACKUP_DIR/client/" 2>/dev/null || true
        RESTORE_NEEDED=true
    fi
}

# Function to restore generated files
restore_generated_files() {
    if [ "$RESTORE_NEEDED" = true ]; then
        echo ""
        echo "📦 Restoring original generated files..."

        # Remove newly generated files first
        rm -f "${PROVIDER_DIR}"/*_resource.go 2>/dev/null || true
        rm -f "${PROVIDER_DIR}"/*_data_source.go 2>/dev/null || true
        rm -f "${CLIENT_DIR}"/*_types.go 2>/dev/null || true

        # Restore from backup
        if [ -d "$TEMP_BACKUP_DIR/provider" ]; then
            cp "$TEMP_BACKUP_DIR/provider/"*.go "$PROVIDER_DIR/" 2>/dev/null || true
        fi

        if [ -d "$TEMP_BACKUP_DIR/client" ]; then
            cp "$TEMP_BACKUP_DIR/client/"*.go "$CLIENT_DIR/" 2>/dev/null || true
        fi

        echo -e "${GREEN}✅ Generated files restored to original state${NC}"
    fi

    # Clean up temp directory
    rm -rf "$TEMP_BACKUP_DIR"
}

# Set up trap to ensure cleanup on exit
trap restore_generated_files EXIT

# Save current state
save_generated_files

# Run the generator
echo "🔧 Running generator: generate-all-schemas.go"
echo "   (This is a preview only - generated files will be restored)"
echo ""

GENERATION_OUTPUT=""

if ! GENERATION_OUTPUT=$(go run "${TOOLS_DIR}/generate-all-schemas.go" --spec-dir="$SPEC_DIR" 2>&1); then
    echo -e "${RED}╔══════════════════════════════════════════════════════════════════════════════╗${NC}"
    echo -e "${RED}║                     ❌ GENERATION FAILED                                     ║${NC}"
    echo -e "${RED}╚══════════════════════════════════════════════════════════════════════════════╝${NC}"
    echo ""
    echo -e "${YELLOW}Generator output:${NC}"
    echo "$GENERATION_OUTPUT"
    echo ""
    echo -e "${CYAN}The generator failed to produce output. Fix the generator before committing.${NC}"
    exit 1
fi

echo "   ✅ Generation completed"
echo ""

# Run linting on generated files
echo "🔍 Running golangci-lint on generated files..."
echo ""

LINT_OUTPUT=""
LINT_SUCCESS=true

# Lint only the provider directory where generated files live
# Note: .golangci.yml excludes tools/ but includes internal/
if ! LINT_OUTPUT=$(golangci-lint run --timeout=5m ./internal/provider/... 2>&1); then
    LINT_SUCCESS=false
fi

# Check results
if [ "$LINT_SUCCESS" = true ]; then
    echo -e "${GREEN}╔══════════════════════════════════════════════════════════════════════════════╗${NC}"
    echo -e "${GREEN}║                  ✅ PREVIEW LINT PASSED                                      ║${NC}"
    echo -e "${GREEN}╚══════════════════════════════════════════════════════════════════════════════╝${NC}"
    echo ""
    echo "The generated code will pass linting in CI/CD."
    echo "Generated files have been restored - only your source changes will be committed."
    exit 0
else
    echo -e "${RED}╔══════════════════════════════════════════════════════════════════════════════╗${NC}"
    echo -e "${RED}║                  ❌ PREVIEW LINT FAILED                                       ║${NC}"
    echo -e "${RED}╚══════════════════════════════════════════════════════════════════════════════╝${NC}"
    echo ""
    echo -e "${YELLOW}Linting errors in generated code:${NC}"
    echo -e "${YELLOW}──────────────────────────────────────────────────────────────${NC}"
    echo "$LINT_OUTPUT"
    echo -e "${YELLOW}──────────────────────────────────────────────────────────────${NC}"
    echo ""
    echo -e "${CYAN}╭────────────────────────────────────────────────────────────────────────────╮${NC}"
    echo -e "${CYAN}│                       HOW TO FIX THIS                                      │${NC}"
    echo -e "${CYAN}╰────────────────────────────────────────────────────────────────────────────╯${NC}"
    echo ""
    echo "  The linting errors above will occur when CI/CD regenerates the code."
    echo "  You must fix the generator to produce lint-compliant code."
    echo ""
    echo "  ${BOLD}Steps:${NC}"
    echo "    1. Review the linting errors above"
    echo "    2. Fix the generator: ${BOLD}${TOOLS_DIR}/generate-all-schemas.go${NC}"
    echo "    3. Run this check again: ${BOLD}pre-commit run lint-generated-preview${NC}"
    echo ""
    echo "  ${BOLD}Note:${NC} Generated files have been restored to their original state."
    echo "        Only your generator changes are staged for commit."
    echo ""
    exit 1
fi
