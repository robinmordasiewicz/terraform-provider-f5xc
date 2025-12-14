#!/bin/bash
# Setup Self-Hosted GitHub Actions Runner
#
# This script sets up a self-hosted GitHub Actions runner for running
# acceptance tests against the F5 XC API.
#
# Two modes available:
#   Native:    Installs runner directly on host (requires Go, etc.)
#   Container: Runs runner in Docker (recommended - fully self-contained)
#
# Usage:
#   ./scripts/setup-self-hosted-runner.sh [OPTIONS]
#
# Options:
#   --container         Use Docker container mode (recommended)
#   -y, --yes           Auto-confirm all prompts
#   --skip-secrets      Skip configuring GitHub secrets (use if already set)
#   --skip-start        Skip starting the runner after setup
#   --runner-name NAME  Set runner name (default: hostname-f5xc)
#
# Environment Variables:
#   F5XC_API_URL    - F5 XC API URL (used in non-interactive mode)
#   F5XC_API_TOKEN  - F5 XC API Token (used in non-interactive mode)
#   GITHUB_TOKEN    - GitHub PAT for container mode (repo scope required)
#
# Container Mode Requirements:
#   - Docker and docker-compose installed
#   - GitHub Personal Access Token with 'repo' scope
#
# Native Mode Requirements:
#   - GitHub CLI (gh) authenticated with repo admin access
#   - Go 1.23+ installed
#   - curl, jq installed

set -euo pipefail

# Command line options
AUTO_YES=false
SKIP_SECRETS=false
SKIP_START=false
RUNNER_NAME_ARG=""
CONTAINER_MODE=false

# Parse arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --container)
            CONTAINER_MODE=true
            shift
            ;;
        -y|--yes)
            AUTO_YES=true
            shift
            ;;
        --skip-secrets)  # pragma: allowlist secret
            SKIP_SECRETS=true
            shift
            ;;
        --skip-start)
            SKIP_START=true
            shift
            ;;
        --runner-name)
            RUNNER_NAME_ARG="$2"
            shift 2
            ;;
        --help|-h)
            sed -n '2,/^set -euo pipefail/p' "$0" | head -n -1
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            echo "Use --help for usage information"
            exit 1
            ;;
    esac
done

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'
BOLD='\033[1m'

# Directories
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"
RUNNER_DIR="${PROJECT_ROOT}/.github-runner"

# Logging
log_info() { echo -e "${BLUE}[INFO]${NC} $1"; }
log_success() { echo -e "${GREEN}[OK]${NC} $1"; }
log_warning() { echo -e "${YELLOW}[WARN]${NC} $1"; }
log_error() { echo -e "${RED}[ERROR]${NC} $1"; }
log_step() { echo -e "\n${CYAN}${BOLD}==> $1${NC}"; }

# Check command exists
check_command() {
    if command -v "$1" &> /dev/null; then
        log_success "$1 found"
        return 0
    else
        log_error "$1 not found"
        return 1
    fi
}

# Prompt for input
prompt_input() {
    local prompt="$1"
    local default="$2"
    local var_name="$3"
    local hide_input="${4:-false}"  # pragma: allowlist secret

    if [[ "$hide_input" == "true" ]]; then  # pragma: allowlist secret
        echo -en "${CYAN}$prompt${NC}"
        [[ -n "$default" ]] && echo -en " [****]: " || echo -en ": "
        read -rs value
        echo ""
    else
        echo -en "${CYAN}$prompt${NC}"
        [[ -n "$default" ]] && echo -en " [$default]: " || echo -en ": "
        read -r value
    fi

    [[ -z "$value" ]] && value="$default"
    eval "$var_name='$value'"
}

# Confirm prompt
confirm() {
    if [[ "$AUTO_YES" == "true" ]]; then
        echo -e "${YELLOW}$1 (y/N): ${NC}y (auto)"
        return 0
    fi
    echo -en "${YELLOW}$1 (y/N): ${NC}"
    read -r response
    [[ "$response" =~ ^[Yy]$ ]]
}

# Header
print_header() {
    echo ""
    echo -e "${BOLD}════════════════════════════════════════════════════════════${NC}"
    echo -e "${BOLD}   Self-Hosted GitHub Actions Runner Setup${NC}"
    echo -e "${BOLD}════════════════════════════════════════════════════════════${NC}"
    echo ""
}

# Check prerequisites
check_prerequisites() {
    log_step "Checking prerequisites"

    local missing=()

    check_command "gh" || missing+=("gh (GitHub CLI)")
    check_command "go" || missing+=("go (Go 1.23+)")
    check_command "curl" || missing+=("curl")
    check_command "jq" || missing+=("jq")

    # Check gh auth
    if command -v gh &> /dev/null; then
        if gh auth status &> /dev/null; then
            log_success "GitHub CLI authenticated"
        else
            missing+=("gh auth (run 'gh auth login')")
        fi
    fi

    if [[ ${#missing[@]} -gt 0 ]]; then
        echo ""
        log_error "Missing prerequisites:"
        for item in "${missing[@]}"; do
            echo "  - $item"
        done
        exit 1
    fi
}

# Get repository info
get_repo_info() {
    log_step "Detecting repository"

    cd "$PROJECT_ROOT"

    REPO_URL=$(git remote get-url origin 2>/dev/null || echo "")
    if [[ -z "$REPO_URL" ]]; then
        log_error "Could not detect git remote"
        exit 1
    fi

    if [[ "$REPO_URL" =~ github\.com[:/]([^/]+)/([^/.]+) ]]; then
        REPO_OWNER="${BASH_REMATCH[1]}"
        REPO_NAME="${BASH_REMATCH[2]}"
        REPO_FULL="${REPO_OWNER}/${REPO_NAME}"
        log_success "Repository: $REPO_FULL"
    else
        log_error "Could not parse repository URL"
        exit 1
    fi
}

# Configure credentials
configure_credentials() {
    if [[ "$SKIP_SECRETS" == "true" ]]; then  # pragma: allowlist secret
        log_step "Skipping credential configuration (--skip-secrets)"
        return 0
    fi

    log_step "Configuring F5 XC API credentials"

    # Use environment variables if available (non-interactive mode)
    if [[ -n "${F5XC_API_URL:-}" ]] && [[ -n "${F5XC_API_TOKEN:-}" ]]; then
        log_info "Using credentials from environment variables"
    else
        echo ""
        echo "To get an API token from F5 XC Console:"
        echo "  1. Administration → Personal Management → Credentials"
        echo "  2. Add Credentials → API Token"
        echo "  3. Copy the token"
        echo ""

        prompt_input "F5 XC API URL" "https://console.ves.volterra.io" "F5XC_API_URL" "false"
        prompt_input "F5 XC API Token" "" "F5XC_API_TOKEN" "true"
    fi

    if [[ -z "${F5XC_API_TOKEN:-}" ]]; then
        log_error "API Token is required (set F5XC_API_TOKEN or provide interactively)"
        exit 1
    fi

    log_info "Setting GitHub secrets..."
    echo "$F5XC_API_URL" | gh secret set F5XC_API_URL --repo "$REPO_FULL"
    echo "$F5XC_API_TOKEN" | gh secret set F5XC_API_TOKEN --repo "$REPO_FULL"
    log_success "Secrets configured"
}

# Detect platform
detect_platform() {
    local os arch
    os=$(uname -s | tr '[:upper:]' '[:lower:]')
    arch=$(uname -m)

    case "$os" in
        linux) RUNNER_OS="linux" ;;
        darwin) RUNNER_OS="osx" ;;
        *) log_error "Unsupported OS: $os"; exit 1 ;;
    esac

    case "$arch" in
        x86_64|amd64) RUNNER_ARCH="x64" ;;
        arm64|aarch64) RUNNER_ARCH="arm64" ;;
        *) log_error "Unsupported architecture: $arch"; exit 1 ;;
    esac

    log_info "Platform: ${RUNNER_OS}-${RUNNER_ARCH}"
}

# Install runner
install_runner() {
    log_step "Installing GitHub Actions runner"

    detect_platform

    # Get latest version
    RUNNER_VERSION=$(curl -s https://api.github.com/repos/actions/runner/releases/latest | jq -r '.tag_name' | sed 's/v//')
    [[ -z "$RUNNER_VERSION" || "$RUNNER_VERSION" == "null" ]] && RUNNER_VERSION="2.311.0"
    log_info "Runner version: $RUNNER_VERSION"

    mkdir -p "$RUNNER_DIR"
    cd "$RUNNER_DIR"

    # Skip if already installed
    if [[ -f "run.sh" ]]; then
        log_success "Runner already installed"
        return 0
    fi

    local runner_file="actions-runner-${RUNNER_OS}-${RUNNER_ARCH}-${RUNNER_VERSION}.tar.gz"
    local runner_url="https://github.com/actions/runner/releases/download/v${RUNNER_VERSION}/${runner_file}"

    log_info "Downloading runner..."
    curl -sL "$runner_url" -o "$runner_file"

    log_info "Extracting..."
    tar -xzf "$runner_file"
    rm -f "$runner_file"

    log_success "Runner installed to: $RUNNER_DIR"
}

# Configure runner
configure_runner() {
    log_step "Configuring runner"

    cd "$RUNNER_DIR"

    # Get registration token
    log_info "Getting registration token..."
    RUNNER_TOKEN=$(gh api --method POST \
        -H "Accept: application/vnd.github+json" \
        "/repos/${REPO_FULL}/actions/runners/registration-token" \
        --jq '.token' 2>/dev/null)

    if [[ -z "$RUNNER_TOKEN" || "$RUNNER_TOKEN" == "null" ]]; then
        log_error "Could not get registration token (need admin access)"
        exit 1
    fi

    # Remove existing config if present
    if [[ -f ".runner" ]]; then
        log_info "Removing existing configuration..."
        ./config.sh remove --token "$RUNNER_TOKEN" 2>/dev/null || true
    fi

    # Configure
    local runner_name="${HOSTNAME:-$(hostname)}-f5xc"
    if [[ -n "$RUNNER_NAME_ARG" ]]; then
        RUNNER_NAME="$RUNNER_NAME_ARG"
        log_info "Using runner name: $RUNNER_NAME"
    elif [[ "$AUTO_YES" == "true" ]]; then
        RUNNER_NAME="$runner_name"
        log_info "Using default runner name: $RUNNER_NAME"
    else
        prompt_input "Runner name" "$runner_name" "RUNNER_NAME" "false"
    fi

    log_info "Registering runner..."
    ./config.sh \
        --url "https://github.com/${REPO_FULL}" \
        --token "$RUNNER_TOKEN" \
        --name "$RUNNER_NAME" \
        --labels "self-hosted,${RUNNER_OS},${RUNNER_ARCH}" \
        --work "_work" \
        --unattended \
        --replace

    log_success "Runner registered"
}

# Start runner
start_runner() {
    if [[ "$SKIP_START" == "true" ]]; then
        log_step "Skipping runner start (--skip-start)"
        log_info "To start manually: cd $RUNNER_DIR && ./run.sh"
        return 0
    fi

    log_step "Starting runner"

    cd "$RUNNER_DIR"

    if [[ "$AUTO_YES" == "true" ]]; then
        # In auto mode, start in background
        START_OPTION="2"
        log_info "Auto mode: starting in background"
    else
        echo ""
        echo "Run options:"
        echo "  1) Foreground - Interactive (Ctrl+C to stop)"
        echo "  2) Background - Runs in background"
        echo "  3) Service    - Install as system service"
        echo "  4) Skip       - Don't start now"
        echo ""

        prompt_input "Choose option" "1" "START_OPTION" "false"
    fi

    case "$START_OPTION" in
        1)
            log_info "Starting in foreground (Ctrl+C to stop)..."
            ./run.sh
            ;;
        2)
            nohup ./run.sh > runner.log 2>&1 &
            echo $! > runner.pid
            log_success "Started in background (PID: $(cat runner.pid))"
            log_info "Logs: tail -f $RUNNER_DIR/runner.log"
            ;;
        3)
            if [[ "$RUNNER_OS" == "linux" ]]; then
                sudo ./svc.sh install && sudo ./svc.sh start
            else
                ./svc.sh install && ./svc.sh start
            fi
            log_success "Service installed and started"
            ;;
        4)
            log_info "To start manually: cd $RUNNER_DIR && ./run.sh"
            ;;
    esac
}

# Summary
print_summary() {
    echo ""
    echo -e "${BOLD}════════════════════════════════════════════════════════════${NC}"
    echo -e "${GREEN}${BOLD}   Setup Complete${NC}"
    echo -e "${BOLD}════════════════════════════════════════════════════════════${NC}"
    echo ""
    echo "  Repository:  $REPO_FULL"
    echo "  Runner:      $RUNNER_DIR"
    [[ -n "${F5XC_API_URL:-}" ]] && echo "  API URL:     $F5XC_API_URL"
    echo ""
    echo "  Trigger tests:"
    echo "    gh workflow run acceptance-tests.yml -f mode=full"
    echo ""
}

# ═══════════════════════════════════════════════════════════════════════════
# Container Mode Functions
# ═══════════════════════════════════════════════════════════════════════════

check_container_prerequisites() {
    log_step "Checking container prerequisites"

    local missing=()

    if command -v docker &> /dev/null; then
        log_success "docker found"
    else
        missing+=("docker")
    fi

    if command -v docker-compose &> /dev/null || docker compose version &> /dev/null 2>&1; then
        log_success "docker-compose found"
    else
        missing+=("docker-compose")
    fi

    check_command "gh" || missing+=("gh (GitHub CLI)")

    # Check gh auth
    if command -v gh &> /dev/null; then
        if gh auth status &> /dev/null; then
            log_success "GitHub CLI authenticated"
        else
            missing+=("gh auth (run 'gh auth login')")
        fi
    fi

    if [[ ${#missing[@]} -gt 0 ]]; then
        echo ""
        log_error "Missing prerequisites:"
        for item in "${missing[@]}"; do
            echo "  - $item"
        done
        exit 1
    fi
}

configure_container_env() {
    log_step "Configuring container environment"

    local env_file="${PROJECT_ROOT}/.github-runner-docker/.env"

    # Get GitHub token for container auth
    local github_token="${GITHUB_TOKEN:-}"
    if [[ -z "$github_token" ]]; then
        echo ""
        echo "Container mode requires a GitHub Personal Access Token with 'repo' scope."
        echo "Create one at: https://github.com/settings/tokens"
        echo ""
        prompt_input "GitHub Personal Access Token" "" "github_token" "true"
    fi

    if [[ -z "$github_token" ]]; then
        log_error "GitHub token is required for container mode"
        exit 1
    fi

    # Get F5XC credentials if not skipping secrets
    local f5xc_url="${F5XC_API_URL:-}"
    local f5xc_token="${F5XC_API_TOKEN:-}"

    if [[ "$SKIP_SECRETS" != "true" ]]; then  # pragma: allowlist secret
        if [[ -z "$f5xc_url" ]] || [[ -z "$f5xc_token" ]]; then
            echo ""
            echo "To get an API token from F5 XC Console:"
            echo "  1. Administration → Personal Management → Credentials"
            echo "  2. Add Credentials → API Token"
            echo ""
            [[ -z "$f5xc_url" ]] && prompt_input "F5 XC API URL" "https://console.ves.volterra.io" "f5xc_url" "false"
            [[ -z "$f5xc_token" ]] && prompt_input "F5 XC API Token" "" "f5xc_token" "true"
        fi

        # Set GitHub secrets
        log_info "Setting GitHub secrets..."
        echo "$f5xc_url" | gh secret set F5XC_API_URL --repo "$REPO_FULL"
        echo "$f5xc_token" | gh secret set F5XC_API_TOKEN --repo "$REPO_FULL"
        log_success "GitHub secrets configured"
    fi

    # Create .env file
    local runner_name="${RUNNER_NAME_ARG:-f5xc-container-runner}"

    log_info "Creating container configuration..."
    cat > "$env_file" << EOF
# Auto-generated by setup-self-hosted-runner.sh
GITHUB_REPOSITORY=${REPO_FULL}
GITHUB_TOKEN=${github_token}
RUNNER_NAME=${runner_name}
RUNNER_LABELS=self-hosted,Linux,X64,container
F5XC_API_URL=${f5xc_url:-}
F5XC_API_TOKEN=${f5xc_token:-}
EOF

    chmod 600 "$env_file"
    log_success "Container configuration saved to .env"
}

# Get the docker compose command (v2 or v1)
get_compose_cmd() {
    if docker compose version &> /dev/null 2>&1; then
        echo "docker compose"
    else
        echo "docker-compose"
    fi
}

start_container_runner() {
    local compose_cmd
    compose_cmd=$(get_compose_cmd)

    if [[ "$SKIP_START" == "true" ]]; then
        log_step "Skipping container start (--skip-start)"
        log_info "To start manually:"
        echo "  cd ${PROJECT_ROOT}/.github-runner-docker"
        echo "  $compose_cmd up -d --build"
        return 0
    fi

    log_step "Building and starting container runner"

    cd "${PROJECT_ROOT}/.github-runner-docker"

    log_info "Building container (this may take a few minutes)..."
    $compose_cmd build

    log_info "Starting runner container..."
    $compose_cmd up -d

    sleep 3

    if $compose_cmd ps | grep -q "Up\|running"; then
        log_success "Container runner started successfully"
        log_info "View logs: $compose_cmd logs -f"
        log_info "Stop runner: $compose_cmd down"
    else
        log_error "Container failed to start"
        $compose_cmd logs
        exit 1
    fi
}

print_container_summary() {
    echo ""
    echo -e "${BOLD}════════════════════════════════════════════════════════════${NC}"
    echo -e "${GREEN}${BOLD}   Container Runner Setup Complete${NC}"
    echo -e "${BOLD}════════════════════════════════════════════════════════════${NC}"
    echo ""
    echo "  Repository:  $REPO_FULL"
    echo "  Runner:      Docker container (f5xc-github-runner)"
    [[ -n "${F5XC_API_URL:-}" ]] && echo "  API URL:     $F5XC_API_URL"
    echo ""
    echo "  Commands:"
    echo "    View logs:    cd .github-runner-docker && docker-compose logs -f"
    echo "    Stop runner:  cd .github-runner-docker && docker-compose down"
    echo "    Restart:      cd .github-runner-docker && docker-compose restart"
    echo ""
    echo "  Trigger tests:"
    echo "    gh workflow run acceptance-tests.yml -f mode=full"
    echo ""
}

# ═══════════════════════════════════════════════════════════════════════════
# Main Entry Points
# ═══════════════════════════════════════════════════════════════════════════

main_container() {
    print_header
    echo -e "${CYAN}Mode: Container (Docker)${NC}"
    echo ""

    check_container_prerequisites
    get_repo_info

    if ! confirm "Set up containerized runner for $REPO_FULL?"; then
        exit 0
    fi

    configure_container_env
    start_container_runner
    print_container_summary
}

main_native() {
    print_header
    echo -e "${CYAN}Mode: Native${NC}"
    echo ""

    check_prerequisites
    get_repo_info

    if ! confirm "Set up self-hosted runner for $REPO_FULL?"; then
        exit 0
    fi

    configure_credentials
    install_runner
    configure_runner
    start_runner
    print_summary
}

# Main dispatcher
main() {
    if [[ "$CONTAINER_MODE" == "true" ]]; then
        main_container
    else
        main_native
    fi
}

main "$@"
