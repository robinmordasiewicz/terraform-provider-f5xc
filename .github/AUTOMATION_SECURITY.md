# GitHub Automation Security

## Personal Access Token Management

This repository uses a Personal Access Token (PAT) for automated workflow triggers. This document describes the security practices and maintenance procedures.

## Why a PAT is Required

GitHub's `GITHUB_TOKEN` has a security limitation: PRs created using `GITHUB_TOKEN` **do not trigger workflows**. This prevents infinite workflow loops but blocks our automation requirements.

Our workflows need to:
1. Generate code/documentation changes automatically
2. Create PRs with those changes
3. **Trigger status checks** on those PRs
4. Auto-merge when checks pass

To enable step 3, we use a Personal Access Token stored as `AUTO_MERGE_TOKEN`.

## Current Token Configuration

### Token Scopes Required

The `AUTO_MERGE_TOKEN` PAT requires these scopes:

| Scope | Purpose | Workflows Using It |
|-------|---------|-------------------|
| `repo` | Create PRs and enable auto-merge | All automated workflows |
| `workflow` | Trigger workflows on bot-created PRs | All automated workflows |

### Token Storage

```bash
# Token stored as repository secret
Secret Name: AUTO_MERGE_TOKEN
Created: 2025-11-24
Owner: @robinmordasiewicz
```

### Workflows Using AUTO_MERGE_TOKEN

1. **`.github/workflows/docs.yml`**
   - Line 75: `token: ${{ secrets.AUTO_MERGE_TOKEN }}`
   - Purpose: Create PRs with regenerated documentation

2. **`.github/workflows/generate.yml`**
   - Line 98: `token: ${{ secrets.AUTO_MERGE_TOKEN }}`
   - Purpose: Create PRs with regenerated provider code

3. **`.github/workflows/sync-openapi.yml`**
   - Line 236: `github-token: ${{ secrets.AUTO_MERGE_TOKEN }}`
   - Line 277: `GH_TOKEN: ${{ secrets.AUTO_MERGE_TOKEN }}`
   - Purpose: Create PRs with updated OpenAPI specs

## Security Best Practices

### 1. Token Rotation Schedule

**Recommended**: Rotate the PAT every 90 days

```bash
# Rotation procedure:
1. Generate new fine-grained PAT at https://github.com/settings/tokens?type=beta
2. Configure scopes: repo, workflow
3. Set repository access: robinmordasiewicz/terraform-provider-f5xc only
4. Update repository secret:
   echo "new_token_value" | gh secret set AUTO_MERGE_TOKEN
5. Test with workflow trigger:
   gh workflow run docs.yml
6. Verify PR creation and auto-merge
7. Revoke old token
```

### 2. Token Scope Minimization

✅ **Correct**: Fine-grained PAT with repository-specific access
❌ **Incorrect**: Classic PAT with all-repos access

The current token should be configured as:
- **Type**: Fine-grained personal access token
- **Repository access**: Only this repository
- **Permissions**:
  - Contents: Read and write
  - Pull requests: Read and write
  - Workflows: Read and write

### 3. Secret Protection

- ✅ Token stored as GitHub repository secret (encrypted at rest)
- ✅ Token never committed to repository
- ✅ Token not visible in workflow logs
- ✅ Token scoped to minimum required permissions
- ✅ Token limited to single repository

### 4. Monitoring and Alerts

Monitor for:
- Unexpected workflow failures
- Unauthorized repository access
- Token expiration warnings
- Unusual PR creation patterns

## Token Rotation Procedure

### Step 1: Generate New Token

1. Visit: https://github.com/settings/tokens?type=beta
2. Click "Generate new token"
3. Configure token:
   - **Name**: `terraform-provider-f5xc-automation`
   - **Expiration**: 90 days
   - **Repository access**: Only select repositories → `robinmordasiewicz/terraform-provider-f5xc`
   - **Permissions**:
     - Contents: Read and write
     - Pull requests: Read and write
     - Workflows: Read and write
4. Click "Generate token"
5. **Copy token immediately** (shown only once)

### Step 2: Update Repository Secret

```bash
# Method 1: Using GitHub CLI
echo "github_pat_NEW_TOKEN_HERE" | gh secret set AUTO_MERGE_TOKEN

# Method 2: Via GitHub UI
# Navigate to: Settings → Secrets and variables → Actions
# Click "AUTO_MERGE_TOKEN" → "Update secret"
# Paste new token value
```

### Step 3: Validate New Token

```bash
# Trigger a workflow to test
gh workflow run docs.yml

# Wait for workflow to complete
sleep 30

# Check if PR was created
gh pr list --label automated

# Verify checks are running
gh pr checks <PR_NUMBER>

# Verify auto-merge is enabled
gh pr view <PR_NUMBER> --json autoMergeRequest
```

### Step 4: Revoke Old Token

1. Visit: https://github.com/settings/tokens
2. Find the old token
3. Click "Revoke"
4. Confirm revocation

## Emergency Procedures

### Token Compromised

If the token is accidentally exposed:

```bash
# 1. Immediately revoke the token
# Visit: https://github.com/settings/tokens → Revoke

# 2. Generate new token (see Step 1 above)

# 3. Update secret immediately
echo "github_pat_NEW_TOKEN" | gh secret set AUTO_MERGE_TOKEN

# 4. Audit recent repository activity
gh api repos/robinmordasiewicz/terraform-provider-f5xc/events

# 5. Review recent PRs for suspicious changes
gh pr list --state all --limit 20

# 6. Check workflow runs
gh run list --limit 20
```

### Token Expired

If automated PRs stop being created:

```bash
# 1. Check last workflow run
gh run list --workflow=docs.yml --limit 1

# 2. Check workflow logs for auth errors
gh run view <RUN_ID> --log

# 3. Generate new token and update secret (see rotation procedure)

# 4. Re-run failed workflow
gh run rerun <RUN_ID>
```

## Validation Checklist

After rotating the token, verify:

- [ ] New token has correct scopes (repo, workflow)
- [ ] Repository secret `AUTO_MERGE_TOKEN` updated
- [ ] Test workflow triggered successfully: `gh workflow run docs.yml`
- [ ] PR created with auto-merge enabled
- [ ] Status checks triggered on bot-created PR
- [ ] PR auto-merged after checks passed
- [ ] Old token revoked at https://github.com/settings/tokens

## Alternative Approaches Considered

### GitHub App

**Pros**:
- More granular permissions
- Can be installed per-repository
- Tokens auto-rotate
- Better audit trail

**Cons**:
- More complex setup
- Requires GitHub App creation and installation
- Overkill for single-repository automation

**Decision**: PAT chosen for simplicity. Consider GitHub App if expanding to multiple repositories.

### Classic Personal Access Token

**Pros**:
- Simple to create
- Long-lived

**Cons**:
- ❌ All-or-nothing repository access
- ❌ Broader permissions than needed
- ❌ Less secure than fine-grained PAT

**Decision**: Fine-grained PAT preferred for better security.

## References

- [GitHub: Creating a fine-grained personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-fine-grained-personal-access-token)
- [GitHub: Using secrets in workflows](https://docs.github.com/en/actions/security-guides/using-secrets-in-github-actions)
- [GitHub: Automatic token authentication](https://docs.github.com/en/actions/security-guides/automatic-token-authentication#permissions-for-the-github_token)
- [GitHub: Events that trigger workflows](https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows#pull_request)

## Maintenance Schedule

| Task | Frequency | Last Completed | Next Due |
|------|-----------|----------------|----------|
| Token rotation | Every 90 days | 2025-11-24 | 2026-02-24 |
| Security audit | Quarterly | 2025-11-24 | 2026-02-24 |
| Workflow testing | After rotation | 2025-11-24 | After next rotation |
| Documentation review | Annually | 2025-11-24 | 2026-11-24 |

## Contact

For questions or issues with automated workflows:
- Repository owner: @robinmordasiewicz
- Last updated: 2025-11-24
