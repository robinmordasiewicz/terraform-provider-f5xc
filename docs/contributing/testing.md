# Testing

The provider includes comprehensive testing to ensure reliability.

## Unit Tests

Run all unit tests:

```bash
go test ./...
```

Run tests for a specific package:

```bash
go test ./internal/provider
```

Run with verbose output:

```bash
go test ./... -v
```

## Acceptance Tests

Acceptance tests run against a real F5 Distributed Cloud tenant.

!!! warning "Required Credentials"
    Acceptance tests require valid F5XC credentials via `F5XC_API_TOKEN` or `F5XC_P12_FILE`.

Run all acceptance tests:

```bash
TF_ACC=1 go test ./... -v -timeout 120m
```

Run tests for a specific resource:

```bash
TF_ACC=1 go test ./internal/provider -v -run TestAccNamespaceResource
```

## Test Conventions

- Unit test files: `*_test.go`
- Acceptance tests: Prefixed with `TestAcc`
- Timeout: 120 minutes for acceptance tests
- Cleanup: Tests should clean up created resources

## CI/CD Testing

All pull requests run:

- Unit tests
- Linting (golangci-lint)
- Build verification
- Constitution checks

See `.github/workflows/ci.yml` for details.

## See Also

- [Development guide](development.md)
- [CI/CD automation](../../CLAUDE.md#part-2-cicd-automation-rules)
