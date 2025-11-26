# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Claude Constitution: GitHub Workflow & Automation Rules

### Core Principle: Issue-First Development

**ALL work MUST begin with a GitHub Issue. No exceptions.**

This repository follows a strict Issue → Branch → PR workflow. Every code change must be traceable to a documented issue that describes what problem is being solved and why.

---

## Part 1: GitHub Workflow Rules

### Rule 0: Always Create an Issue First (MANDATORY)

**Before writing ANY code or creating ANY branch, create a GitHub Issue.**

```bash
# ✅ CORRECT workflow
gh issue create --title "feat: Add user authentication" --body "Description of the feature..."
# Note the issue number (e.g., #42)
gh issue develop 42 --checkout  # Creates branch linked to issue
# ... make changes ...
git commit -m "feat(auth): implement login flow

Closes #42"
```

```bash
# ❌ INCORRECT workflow (what I was doing wrong)
git checkout -b feature/some-feature  # NO! No issue created first
# ... make changes ...
git commit -m "add feature"  # NO! No issue reference
gh pr create  # NO! PR not linked to any issue
```

**Issue Types and Templates:**
- `bug` - Something isn't working correctly
- `feature` - New functionality request
- `docs` - Documentation improvements
- `refactor` - Code restructuring without behavior change
- `chore` - Maintenance tasks, dependency updates

### Rule 1: Branch Naming Convention

**Branch names MUST include the issue number and follow this pattern:**

```
<type>/<issue-number>-<short-description>
```

| Type | Purpose | Example |
|------|---------|---------|
| `feature/` | New functionality | `feature/42-add-user-auth` |
| `bugfix/` | Bug corrections | `bugfix/57-fix-login-crash` |
| `docs/` | Documentation changes | `docs/63-update-readme` |
| `refactor/` | Code restructuring | `refactor/71-cleanup-api-client` |
| `hotfix/` | Urgent production fixes | `hotfix/89-security-patch` |

**Creating branches from issues:**
```bash
# Preferred: Use GitHub CLI to create linked branch
gh issue develop 42 --checkout

# Alternative: Manual creation with proper naming
git checkout -b feature/42-add-user-auth
```

### Rule 2: Conventional Commits (MANDATORY)

**All commit messages MUST follow the Conventional Commits specification.**

Format:
```
<type>(<optional scope>): <description>

[optional body]

[optional footer(s)]
```

**Commit Types:**
| Type | Description | Version Bump |
|------|-------------|--------------|
| `feat` | New feature | Minor |
| `fix` | Bug fix | Patch |
| `docs` | Documentation only | None |
| `style` | Formatting, no code change | None |
| `refactor` | Code restructuring | None |
| `perf` | Performance improvement | Patch |
| `test` | Adding/fixing tests | None |
| `chore` | Maintenance tasks | None |
| `ci` | CI/CD changes | None |

**Examples:**
```bash
# Feature with scope
git commit -m "feat(auth): add OAuth2 login support

Implements Google and GitHub OAuth providers.

Closes #42"

# Bug fix
git commit -m "fix(api): resolve null pointer in user lookup

The user service was not checking for nil responses
from the database layer.

Fixes #57"

# Breaking change (note the !)
git commit -m "feat(api)!: change authentication endpoint

BREAKING CHANGE: The /auth endpoint now requires
a JSON body instead of form data.

Closes #63"
```

### Rule 3: Pull Request Requirements

**Every PR MUST:**

1. **Reference an issue** using closing keywords in the description
2. **Have a descriptive title** following Conventional Commits format
3. **Include a summary** of changes and testing performed

**Closing Keywords** (use in PR description):
- `Closes #42` - Closes the issue when PR merges
- `Fixes #42` - Same as Closes
- `Resolves #42` - Same as Closes

**PR Description Template:**
```markdown
## Summary
Brief description of changes

## Related Issue
Closes #42

## Changes Made
- Change 1
- Change 2

## Testing
- [ ] Unit tests pass
- [ ] Manual testing completed

## Screenshots (if applicable)
```

**Creating PRs:**
```bash
# Create PR linked to issue
gh pr create --title "feat(auth): add OAuth2 login support" --body "$(cat <<'EOF'
## Summary
Implements OAuth2 authentication with Google and GitHub providers.

## Related Issue
Closes #42

## Changes Made
- Added OAuth2 provider configuration
- Implemented callback handlers
- Added user session management

## Testing
- [x] Unit tests pass
- [x] Manual testing with Google OAuth

🤖 Generated with [Claude Code](https://claude.com/claude-code)
EOF
)"
```

### Rule 4: Issue-Branch-PR Traceability

**Maintain complete traceability:**

```
Issue #42: "Add user authentication"
    ↓
Branch: feature/42-add-user-auth
    ↓
Commits: feat(auth): implement login flow (Closes #42)
    ↓
PR #43: "feat(auth): add OAuth2 login support" (Closes #42)
    ↓
Merge → Issue #42 auto-closes
```

**Verification checklist before PR:**
- [ ] Issue exists and describes the work
- [ ] Branch name includes issue number
- [ ] All commits reference the issue
- [ ] PR description uses closing keyword

---

## Part 2: CI/CD Automation Rules

### Core Principle: Automation First

**BEFORE making any changes, analyze `.github/workflows/` to understand existing automation patterns.**

This repository uses CI/CD automation extensively. Respect the automation - do not bypass it.

### Rule 1: Never Commit Generated Files

**Generated/derived files are managed by CI/CD workflows, not manual commits.**

| Directory/File | Generated By | Workflow |
|----------------|--------------|----------|
| `docs/resources/*.md` | `tfplugindocs` + `transform-docs.go` | `docs.yml` |
| `docs/data-sources/*.md` | `tfplugindocs` + `transform-docs.go` | `docs.yml` |
| `examples/resources/*/*.tf` | `generate-examples.go` | `docs.yml` |
| `examples/data-sources/*/*.tf` | `generate-examples.go` | `docs.yml` |
| `docs/api/*.md` | `generate-api-docs.go` | `pages.yml` |
| `docs/nav-api.yml` | `generate-api-docs.go` | `pages.yml` |
| `internal/provider/*_resource.go` | `generate-all-schemas.go` | `generate.yml` |
| `internal/provider/*_data_source.go` | `generate-all-schemas.go` | `generate.yml` |
| `site/` | `mkdocs build` | `pages.yml` |

**Correct behavior**: Commit only source/tool changes. Let workflows generate artifacts.

**Important**: Example files (`examples/`) are generated from `tools/generate-examples.go` and contain content derived from schema analysis. Never commit example files manually - modify the generator tool and let CI regenerate them.

### Rule 2: Understand the Workflow Chain

```
Source Change → Push to Branch → CI Workflow Runs → Auto-PR Created
```

**Workflows in this repository:**

| Workflow | Trigger | Action | Creates PR? |
|----------|---------|--------|-------------|
| `docs.yml` | Changes to `internal/`, `docs/`, `examples/`, `templates/`, `tools/` | Regenerates examples and provider docs | ✅ Yes |
| `generate.yml` | Changes to `docs/specifications/api/` | Regenerates provider code | ✅ Yes |
| `sync-openapi.yml` | Scheduled (daily) | Syncs OpenAPI specs from F5 | ✅ Yes |
| `pages.yml` | Changes to `docs/`, `mkdocs.yml` | Builds & deploys MkDocs | ❌ Direct deploy |
| `auto-tag.yml` | Push to main (non-docs) | Creates version tag | ❌ Creates tag |
| `release.yml` | Tag push (`v*`) | Publishes release | ❌ Creates release |
| `lint.yml` | All pushes/PRs | Runs tests & linting | ❌ Status checks |

### Rule 3: One Concern Per Commit

When modifying generator tools (e.g., `tools/transform-docs.go`):

✅ **Correct**:
```bash
git add tools/transform-docs.go
git commit -m "feat: add Metadata/Spec grouping to doc transformer"
git push
# Let docs.yml workflow create PR with generated docs
```

❌ **Incorrect**:
```bash
go run tools/transform-docs.go  # DON'T manually run generators
git add tools/transform-docs.go docs/  # DON'T commit generated files
git commit -m "feat: add grouping + regenerate all docs"
```

### Rule 4: Check .gitignore for Generated Patterns

Before committing, verify files aren't in `.gitignore`:

```bash
# Generated documentation (rebuilt at deploy time)
docs/api/
docs/nav-api.yml
site/
```

If a file pattern is in `.gitignore`, it's generated - don't commit it manually.

### Rule 5: Respect Auto-Generated PRs

When you see PRs from `github-actions` bot with labels `automated`, `documentation`:
- These are legitimate CI-generated changes
- Review them for correctness
- Merge or close based on content quality
- Do NOT duplicate their work manually

### Rule 6: Pre-Flight Checklist

Before creating any PR, verify:

**GitHub Workflow:**
- [ ] GitHub Issue exists for this work
- [ ] Branch name includes issue number (e.g., `feature/42-description`)
- [ ] All commits reference the issue number
- [ ] PR description includes closing keyword (`Closes #42`)

**Automation Compliance:**
- [ ] Analyzed relevant workflows in `.github/workflows/`
- [ ] Checked `.gitignore` for generated file patterns
- [ ] Committing ONLY source code, not generated artifacts
- [ ] Not bypassing automation that handles this task
- [ ] PR contains single concern (not mixed source + generated)

### Rule 7: Automation Security and Token Management

**All automated workflows use the `AUTO_MERGE_TOKEN` for PR creation to enable workflow triggers.**

**Why this matters:**
- PRs created with `GITHUB_TOKEN` **do not trigger workflows** (GitHub security feature)
- Our automation requires status checks to run on bot-created PRs
- The `AUTO_MERGE_TOKEN` is a Personal Access Token (PAT) that enables this

**Security practices:**
- See `.github/AUTOMATION_SECURITY.md` for complete token management procedures
- Token rotation required every 90 days
- Token scopes: `repo`, `workflow` (minimum required)
- Token access: This repository only (fine-grained PAT)

**Affected workflows:**
- `docs.yml` (line 75): Creates PRs with regenerated documentation
- `generate.yml` (line 98): Creates PRs with regenerated provider code
- `sync-openapi.yml` (lines 236, 277): Creates PRs with updated OpenAPI specs

**If automation fails:**
1. Check if `AUTO_MERGE_TOKEN` has expired
2. Follow token rotation procedure in `.github/AUTOMATION_SECURITY.md`
3. Verify workflow logs for authentication errors

---

## Part 3: Constitution Enforcement

### Automated Enforcement Mechanisms

The repository includes multiple layers of enforcement to prevent violations of the constitution rules.

#### Layer 1: Pre-Commit Hooks (Local)

**Location**: `.pre-commit-config.yaml` + `scripts/check-no-generated-files.sh`

**What it does**: Blocks commits containing generated files before they reach the repository

**Setup** (one-time):
```bash
# Install pre-commit
pip install pre-commit

# Install hooks
pre-commit install

# Test (optional)
pre-commit run --all-files
```

**How it works**:
```bash
# Attempt to commit generated file
git add docs/resources/namespace.md
git commit -m "docs: update namespace"

# ❌ Blocked by pre-commit hook:
# ERROR: Attempting to commit generated files
# - docs/resources/namespace.md
#
# CLAUDE.md Rule 1: Never Commit Generated Files
```

**Patterns checked**:
- `docs/resources/*.md` - Provider resource documentation
- `docs/data-sources/*.md` - Provider data source documentation
- `docs/api/*.md` - API reference documentation
- `docs/nav-api.yml` - API navigation structure
- `site/` - Built MkDocs site
- `internal/provider/*_resource.go` - Generated resource implementations
- `internal/provider/*_data_source.go` - Generated data source implementations

#### Layer 2: CI/CD Checks (Remote)

**Location**: `.github/workflows/lint.yml` → `check-constitution` job

**What it does**: Fails PRs that contain generated files, preventing merge to main

**How it works**:
```yaml
# Runs automatically on every PR
- Compares PR branch against base branch
- Identifies changed files
- Checks against generated file patterns
- Fails CI if violations detected
```

**Result**:
- ❌ PR status check fails with clear error message
- 🚫 PR cannot be merged until violations are removed
- 📝 Instructions provided for proper fix

#### Layer 3: Branch Protection (GitHub Settings)

**Recommended settings** (requires repo admin):
```yaml
Branch: main
  ✅ Require status checks to pass before merging
  ✅ Require "Constitution Check" to pass
  ✅ Require "Lint" to pass
  ✅ Require "Build and Test" to pass
  ✅ Include administrators (no exceptions!)
```

### Enforcement Examples

#### ✅ Correct Workflow (Passes All Checks)
```bash
# 1. Modify generator tool
vim tools/transform-docs.go

# 2. Pre-commit runs automatically
git add tools/transform-docs.go
git commit -m "feat: improve doc formatting"
# ✅ Pre-commit: No generated files detected

# 3. Push and create PR
git push
gh pr create

# 4. CI checks run
# ✅ Constitution Check: No generated files in PR
# ✅ Lint: Code passes linting
# ✅ Build and Test: All tests pass

# 5. Merge PR
# 6. docs.yml workflow triggers automatically
# 7. Auto-PR created with regenerated docs
# 8. Review and merge auto-PR
```

#### ❌ Incorrect Workflow (Blocked by Enforcement)
```bash
# 1. Manually run generator and commit results
go run tools/transform-docs.go
git add tools/transform-docs.go docs/
git commit -m "feat: add grouping"

# ❌ Pre-commit blocks:
# ERROR: Attempting to commit generated files
# - docs/resources/namespace.md
# - docs/resources/http_loadbalancer.md
# ... (142 more files)

# Fix: Unstage generated files
git restore --staged docs/

# Commit only source code
git commit -m "feat: add grouping"
# ✅ Pre-commit: No generated files detected
```

### Bypassing Enforcement (Emergency Only)

**NEVER bypass enforcement in normal development!**

In genuine emergencies only (e.g., fixing broken CI):

```bash
# Skip pre-commit hooks (requires conscious decision)
git commit --no-verify -m "emergency: fix broken workflow"

# Note: CI checks will still fail if generated files present!
```

**Important**: Even with `--no-verify`, the PR will fail CI checks and cannot be merged. This provides defense-in-depth.

### Testing Enforcement

To verify enforcement is working:

```bash
# Test pre-commit hook
echo "test" >> docs/resources/test.md
git add docs/resources/test.md
git commit -m "test"
# Should fail with clear error message

# Clean up
git restore --staged docs/resources/test.md
rm docs/resources/test.md
```

### Maintenance

**Updating enforcement patterns**:

When adding new generated file types:

1. Update `scripts/check-no-generated-files.sh` (GENERATED_PATTERNS)
2. Update `.github/workflows/lint.yml` (check-constitution job)
3. Update this documentation (Part 3)
4. Test enforcement with new patterns

---

## Part 4: Self-Correction Protocol

### When I Violate These Rules

If I create a branch or PR without following these rules:

1. **Stop immediately** - Do not continue with incorrect workflow
2. **Create the missing issue** - Document what work is being done
3. **Rename or recreate branch** - Include issue number
4. **Update PR** - Add issue reference with closing keyword
5. **Document the violation** - Learn from the mistake

### Behavior Analysis Triggers

I should analyze my GitHub behavior when:
- Creating a branch without `gh issue develop`
- PR description missing `Closes #`, `Fixes #`, or `Resolves #`
- Branch name missing issue number
- Commit messages not following Conventional Commits
- Manually running generators and committing output

## Project Overview

Community-driven Terraform provider for F5 Distributed Cloud (F5XC) built using the HashiCorp Terraform Plugin Framework. The provider implements F5's public OpenAPI specifications to manage F5XC resources via Terraform.

## Build & Development Commands

```bash
# Build the provider binary
go build -o terraform-provider-f5xc

# Run all tests
go test ./...

# Run acceptance tests (requires F5XC_API_TOKEN)
TF_ACC=1 go test ./... -v -timeout 120m

# Generate documentation (requires terraform CLI)
go generate ./...

# Install locally for testing
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/robinmordasiewicz/f5xc/0.1.0/darwin_arm64
cp terraform-provider-f5xc ~/.terraform.d/plugins/registry.terraform.io/robinmordasiewicz/f5xc/0.1.0/darwin_arm64/
```

## Architecture

### Directory Structure

- `main.go` - Provider entry point with version injection via goreleaser
- `internal/provider/` - All Terraform resources and data sources
- `internal/client/` - F5XC API client (HTTP client with CRUD operations and type definitions)
- `tools/` - Code generation utilities for scaffolding new resources from OpenAPI specs
- `docs/` - Auto-generated provider documentation for Terraform Registry
- `examples/` - Example Terraform configurations

### Resource Implementation Pattern

Each resource follows the Terraform Plugin Framework pattern:

1. **Resource struct** implementing `resource.Resource`, `resource.ResourceWithConfigure`, `resource.ResourceWithImportState`
2. **Model struct** with `tfsdk` tags for state management
3. **CRUD methods**: `Create`, `Read`, `Update`, `Delete`
4. **Registration** in `provider.go` via `Resources()` function

Example reference: `internal/provider/namespace_resource.go`

### Client Architecture

`internal/client/client.go` contains:

- HTTP client wrapper for F5XC API (`Client` struct)
- Type definitions for all F5XC resources (e.g., `Namespace`, `HTTPLoadBalancer`)
- CRUD methods for each resource type following pattern: `Create{Resource}`, `Get{Resource}`, `Update{Resource}`, `Delete{Resource}`

API authentication uses Bearer token: `Authorization: APIToken {token}`

### Code Generation

The `tools/` directory contains generators for scaffolding resources from F5 OpenAPI specs:

- `generate-resources.go` - Generates resource files from OpenAPI specs
- `generate-client-types.go` - Generates client type definitions
- `batch-generate.sh` - Batch generation script

## Environment Variables

- `F5XC_API_TOKEN` - Required API token for F5 Distributed Cloud
- `F5XC_API_URL` - Optional API URL (defaults to `https://console.ves.volterra.io/api`)
- `TF_ACC=1` - Enable acceptance tests

## Key Dependencies

- `github.com/hashicorp/terraform-plugin-framework` - Core Terraform provider SDK
- `github.com/hashicorp/terraform-plugin-log` - Structured logging

## Release Process

Releases are automated via GoReleaser on tag push (`v*`). The workflow:

1. Builds cross-platform binaries
2. Signs checksums with GPG
3. Publishes to GitHub Releases with Terraform Registry manifest
