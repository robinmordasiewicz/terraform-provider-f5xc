# Self-Hosted GitHub Actions Runner (Containerized)

Run the GitHub Actions runner in a Docker container with all dependencies pre-installed.

## What's Included

- GitHub Actions Runner 2.321.0
- Go 1.23.3
- Terraform 1.9.8
- GitHub CLI (gh)
- curl, jq, git

## Quick Start

### 1. Configure Environment

```bash
cd .github-runner-docker
cp .env.example .env
```

Edit `.env` with your values:
- `GITHUB_REPOSITORY` - Your repo (e.g., `owner/repo`)
- `GITHUB_TOKEN` - Personal Access Token with `repo` scope
- `VES_API_URL` - F5 XC API URL
- `VES_API_TOKEN` - F5 XC API Token

### 2. Build and Start

```bash
docker-compose up -d --build
```

### 3. View Logs

```bash
docker-compose logs -f
```

### 4. Stop Runner

```bash
docker-compose down
```

The runner will automatically unregister from GitHub on graceful shutdown.

## Getting a GitHub Token

1. Go to [GitHub Settings → Tokens](https://github.com/settings/tokens)
2. Generate new token (classic)
3. Select scope: `repo` (Full control of private repositories)
4. Copy and save the token

## Getting F5 XC API Token

1. Log in to F5 XC Console
2. Go to Administration → Personal Management → Credentials
3. Add Credentials → API Token
4. Copy and save the token

## Architecture Support

The container supports both `amd64` and `arm64` architectures. Docker will automatically select the correct binaries during build.

## Resource Requirements

- **CPU**: 2 cores recommended
- **Memory**: 4GB recommended
- **Disk**: 10GB for container and work directory

## Troubleshooting

### Runner won't register

1. Check `GITHUB_TOKEN` has `repo` scope
2. Verify you have admin access to the repository
3. Check logs: `docker-compose logs`

### Tests fail with missing tools

The container includes Go, Terraform, and common tools. If a tool is missing:

1. Add it to the Dockerfile
2. Rebuild: `docker-compose up -d --build`

### Runner appears offline

1. Check container is running: `docker ps`
2. Check logs for errors: `docker-compose logs`
3. Restart: `docker-compose restart`

## Advanced Configuration

### Custom Labels

```bash
RUNNER_LABELS=self-hosted,Linux,X64,production
```

### Multiple Runners

Create multiple `.env` files and compose files for scaling:

```bash
# runner1.env, runner2.env, etc.
docker-compose -f docker-compose.yml --env-file runner1.env up -d
```

### Persistent Caching

The work directory is persisted in a Docker volume (`f5xc-runner-work`) for caching Go modules and build artifacts between runs.
