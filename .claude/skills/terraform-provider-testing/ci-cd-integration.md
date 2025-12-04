# CI/CD Integration for Terraform Testing

Complete guide to integrating Terraform tests into CI/CD pipelines.

## GitHub Actions

### Native Terraform Test Workflow

```yaml
# .github/workflows/terraform-test.yml
name: Terraform Tests

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

env:
  TF_VERSION: "1.7.0"

jobs:
  validate:
    name: Validate Configuration
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Setup Terraform
        uses: hashicorp/setup-terraform@v3
        with:
          terraform_version: ${{ env.TF_VERSION }}

      - name: Terraform Format Check
        run: terraform fmt -check -recursive

      - name: Terraform Init
        run: terraform init -backend=false

      - name: Terraform Validate
        run: terraform validate

  unit-tests:
    name: Unit Tests (Plan-Only)
    runs-on: ubuntu-latest
    needs: validate
    steps:
      - uses: actions/checkout@v4

      - name: Setup Terraform
        uses: hashicorp/setup-terraform@v3
        with:
          terraform_version: ${{ env.TF_VERSION }}

      - name: Terraform Init
        run: terraform init -backend=false

      - name: Run Unit Tests (with mocks)
        run: terraform test -filter=tests/unit.tftest.hcl

  integration-tests:
    name: Integration Tests
    runs-on: ubuntu-latest
    needs: unit-tests
    if: github.event_name == 'push' && github.ref == 'refs/heads/main'
    env:
      AWS_ACCESS_KEY_ID: ${{ secrets.AWS_ACCESS_KEY_ID }}
      AWS_SECRET_ACCESS_KEY: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
      AWS_REGION: us-west-2
    steps:
      - uses: actions/checkout@v4

      - name: Setup Terraform
        uses: hashicorp/setup-terraform@v3
        with:
          terraform_version: ${{ env.TF_VERSION }}

      - name: Terraform Init
        run: terraform init

      - name: Run Integration Tests
        run: terraform test -filter=tests/integration.tftest.hcl

      - name: Run All Tests (verbose)
        run: terraform test -verbose
```

### Provider Acceptance Test Workflow

```yaml
# .github/workflows/acceptance-tests.yml
name: Acceptance Tests

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]
  schedule:
    - cron: '0 2 * * *'  # Nightly at 2 AM

env:
  GO_VERSION: "1.21"

jobs:
  build:
    name: Build
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: ${{ env.GO_VERSION }}

      - name: Build
        run: go build -v ./...

      - name: Unit Tests
        run: go test -v ./...

  lint:
    name: Lint
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: ${{ env.GO_VERSION }}

      - name: golangci-lint
        uses: golangci-lint/golangci-lint-action@v3
        with:
          version: latest

  acceptance-tests:
    name: Acceptance Tests
    runs-on: ubuntu-latest
    needs: [build, lint]
    timeout-minutes: 120
    strategy:
      fail-fast: false
      matrix:
        terraform: ['1.5.*', '1.6.*', '1.7.*']
    env:
      TF_ACC: "1"
      EXAMPLE_API_KEY: ${{ secrets.EXAMPLE_API_KEY }}
      EXAMPLE_API_ENDPOINT: ${{ secrets.EXAMPLE_API_ENDPOINT }}
    steps:
      - uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: ${{ env.GO_VERSION }}

      - name: Setup Terraform
        uses: hashicorp/setup-terraform@v3
        with:
          terraform_version: ${{ matrix.terraform }}
          terraform_wrapper: false

      - name: Get Terraform Version
        run: terraform version

      - name: Run Acceptance Tests
        run: go test -v -cover -timeout 120m ./internal/...
        env:
          TF_ACC_TERRAFORM_VERSION: ${{ matrix.terraform }}

  acceptance-tests-short:
    name: Acceptance Tests (Short)
    runs-on: ubuntu-latest
    needs: [build, lint]
    if: github.event_name == 'pull_request'
    timeout-minutes: 30
    env:
      TF_ACC: "1"
      EXAMPLE_API_KEY: ${{ secrets.EXAMPLE_API_KEY }}
    steps:
      - uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: ${{ env.GO_VERSION }}

      - name: Setup Terraform
        uses: hashicorp/setup-terraform@v3
        with:
          terraform_wrapper: false

      - name: Run Short Tests
        run: go test -v -short -timeout 30m ./internal/...

  sweeper:
    name: Sweeper
    runs-on: ubuntu-latest
    needs: acceptance-tests
    if: always() && needs.acceptance-tests.result != 'skipped'
    env:
      EXAMPLE_API_KEY: ${{ secrets.EXAMPLE_API_KEY }}
    steps:
      - uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: ${{ env.GO_VERSION }}

      - name: Run Sweeper
        run: go test ./internal/... -sweep=us-west-2 -v
```

### Module Test Workflow with Multiple Providers

```yaml
# .github/workflows/module-tests.yml
name: Module Tests

on:
  push:
    branches: [main]
  pull_request:

jobs:
  test:
    name: Test (${{ matrix.provider }})
    runs-on: ubuntu-latest
    strategy:
      fail-fast: false
      matrix:
        provider: [aws, azure, gcp]
        include:
          - provider: aws
            test_filter: tests/aws.tftest.hcl
            env_prefix: AWS
          - provider: azure
            test_filter: tests/azure.tftest.hcl
            env_prefix: ARM
          - provider: gcp
            test_filter: tests/gcp.tftest.hcl
            env_prefix: GOOGLE
    env:
      TF_VERSION: "1.7.0"
    steps:
      - uses: actions/checkout@v4

      - name: Setup Terraform
        uses: hashicorp/setup-terraform@v3
        with:
          terraform_version: ${{ env.TF_VERSION }}

      - name: Terraform Init
        run: terraform init -backend=false

      - name: Run Tests with Mocks
        run: terraform test -filter=${{ matrix.test_filter }}
```

## GitLab CI

### Basic Pipeline

```yaml
# .gitlab-ci.yml
stages:
  - validate
  - test
  - acceptance

variables:
  TF_VERSION: "1.7.0"
  GO_VERSION: "1.21"

.terraform-job:
  image: hashicorp/terraform:$TF_VERSION
  before_script:
    - terraform version

.go-job:
  image: golang:$GO_VERSION
  before_script:
    - go version

validate:
  extends: .terraform-job
  stage: validate
  script:
    - terraform fmt -check -recursive
    - terraform init -backend=false
    - terraform validate

unit-tests:
  extends: .terraform-job
  stage: test
  script:
    - terraform init -backend=false
    - terraform test -filter=tests/unit.tftest.hcl
  needs: [validate]

integration-tests:
  extends: .terraform-job
  stage: acceptance
  script:
    - terraform init
    - terraform test
  only:
    - main
  variables:
    AWS_ACCESS_KEY_ID: $AWS_ACCESS_KEY_ID
    AWS_SECRET_ACCESS_KEY: $AWS_SECRET_ACCESS_KEY

acceptance-tests:
  extends: .go-job
  stage: acceptance
  script:
    - go test -v -timeout 120m ./internal/...
  variables:
    TF_ACC: "1"
    EXAMPLE_API_KEY: $EXAMPLE_API_KEY
  only:
    - main
  timeout: 2h
```

## Azure DevOps

### Pipeline Configuration

```yaml
# azure-pipelines.yml
trigger:
  branches:
    include:
      - main

pool:
  vmImage: 'ubuntu-latest'

variables:
  TF_VERSION: '1.7.0'
  GO_VERSION: '1.21'

stages:
  - stage: Validate
    jobs:
      - job: Validate
        steps:
          - task: TerraformInstaller@0
            inputs:
              terraformVersion: $(TF_VERSION)

          - script: |
              terraform fmt -check -recursive
              terraform init -backend=false
              terraform validate
            displayName: 'Validate Terraform'

  - stage: Test
    dependsOn: Validate
    jobs:
      - job: UnitTests
        steps:
          - task: TerraformInstaller@0
            inputs:
              terraformVersion: $(TF_VERSION)

          - script: |
              terraform init -backend=false
              terraform test -filter=tests/unit.tftest.hcl
            displayName: 'Run Unit Tests'

      - job: AcceptanceTests
        condition: and(succeeded(), eq(variables['Build.SourceBranch'], 'refs/heads/main'))
        timeoutInMinutes: 120
        steps:
          - task: GoTool@0
            inputs:
              version: $(GO_VERSION)

          - task: TerraformInstaller@0
            inputs:
              terraformVersion: $(TF_VERSION)

          - script: |
              export TF_ACC=1
              go test -v -timeout 120m ./internal/...
            displayName: 'Run Acceptance Tests'
            env:
              EXAMPLE_API_KEY: $(EXAMPLE_API_KEY)
```

## Jenkins

### Jenkinsfile

```groovy
// Jenkinsfile
pipeline {
    agent any

    environment {
        TF_VERSION = '1.7.0'
        GO_VERSION = '1.21'
        TF_IN_AUTOMATION = 'true'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Setup') {
            parallel {
                stage('Setup Terraform') {
                    steps {
                        sh '''
                            wget https://releases.hashicorp.com/terraform/${TF_VERSION}/terraform_${TF_VERSION}_linux_amd64.zip
                            unzip terraform_${TF_VERSION}_linux_amd64.zip
                            mv terraform /usr/local/bin/
                        '''
                    }
                }
                stage('Setup Go') {
                    steps {
                        sh '''
                            wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz
                            tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
                            export PATH=$PATH:/usr/local/go/bin
                        '''
                    }
                }
            }
        }

        stage('Validate') {
            steps {
                sh '''
                    terraform fmt -check -recursive
                    terraform init -backend=false
                    terraform validate
                '''
            }
        }

        stage('Unit Tests') {
            steps {
                sh '''
                    terraform init -backend=false
                    terraform test -filter=tests/unit.tftest.hcl
                '''
            }
        }

        stage('Acceptance Tests') {
            when {
                branch 'main'
            }
            environment {
                TF_ACC = '1'
                EXAMPLE_API_KEY = credentials('example-api-key')
            }
            steps {
                sh 'go test -v -timeout 120m ./internal/...'
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
```

## Test Result Reporting

### GitHub Actions with Test Reports

```yaml
- name: Run Tests with JSON Output
  run: terraform test -json > test-results.json
  continue-on-error: true

- name: Process Test Results
  uses: actions/github-script@v7
  if: always()
  with:
    script: |
      const fs = require('fs');
      const results = fs.readFileSync('test-results.json', 'utf8')
        .split('\n')
        .filter(line => line.trim())
        .map(line => JSON.parse(line));

      let summary = '## Terraform Test Results\n\n';
      let passed = 0;
      let failed = 0;

      results.forEach(result => {
        if (result['@level'] === 'info' && result['@message']) {
          if (result['@message'].includes('pass')) passed++;
          if (result['@message'].includes('fail')) failed++;
          summary += `- ${result['@message']}\n`;
        }
      });

      summary += `\n**Total: ${passed} passed, ${failed} failed**`;

      await github.rest.issues.createComment({
        issue_number: context.issue.number,
        owner: context.repo.owner,
        repo: context.repo.repo,
        body: summary
      });
```

### Go Test Coverage Report

```yaml
- name: Run Tests with Coverage
  run: |
    go test -v -coverprofile=coverage.out -covermode=atomic ./internal/...
    go tool cover -html=coverage.out -o coverage.html

- name: Upload Coverage
  uses: codecov/codecov-action@v3
  with:
    files: ./coverage.out
    flags: unittests
    name: codecov-umbrella
```

## Secrets Management

### GitHub Actions Secrets

Required secrets for provider testing:

```
# AWS Provider
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY

# Azure Provider
ARM_CLIENT_ID
ARM_CLIENT_SECRET
ARM_SUBSCRIPTION_ID
ARM_TENANT_ID

# GCP Provider
GOOGLE_CREDENTIALS  # Service account JSON

# Custom Provider
EXAMPLE_API_KEY
EXAMPLE_API_ENDPOINT
```

### Using OIDC for Cloud Authentication

```yaml
# AWS OIDC Authentication
jobs:
  test:
    permissions:
      id-token: write
      contents: read
    steps:
      - name: Configure AWS Credentials
        uses: aws-actions/configure-aws-credentials@v4
        with:
          role-to-assume: arn:aws:iam::123456789012:role/GitHubActionsRole
          aws-region: us-west-2

      - name: Run Tests
        run: terraform test
```

## Parallel Test Execution

### Matrix Strategy for Parallel Tests

```yaml
jobs:
  test:
    strategy:
      fail-fast: false
      max-parallel: 4
      matrix:
        test-group:
          - tests/group1.tftest.hcl
          - tests/group2.tftest.hcl
          - tests/group3.tftest.hcl
          - tests/group4.tftest.hcl
    steps:
      - name: Run Test Group
        run: terraform test -filter=${{ matrix.test-group }}
```

### Go Test Parallelism

```yaml
- name: Run Parallel Tests
  run: |
    go test -v -parallel 4 -timeout 120m ./internal/...
  env:
    TF_ACC: "1"
    TF_ACC_TERRAFORM_PATH: /usr/local/bin/terraform
```

## Best Practices

### 1. Test Organization

- **Unit tests** on every PR (fast, no infrastructure)
- **Integration tests** on main branch merges
- **Full acceptance tests** nightly or on release branches

### 2. Timeout Configuration

```yaml
timeout-minutes: 120  # For acceptance tests
timeout-minutes: 30   # For short/unit tests
```

### 3. Failure Handling

```yaml
- name: Run Tests
  id: test
  run: terraform test
  continue-on-error: true

- name: Check Test Results
  if: steps.test.outcome == 'failure'
  run: |
    echo "Tests failed. Running sweeper..."
    go test ./internal/... -sweep=us-west-2
```

### 4. Caching

```yaml
- name: Cache Go Modules
  uses: actions/cache@v3
  with:
    path: ~/go/pkg/mod
    key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
    restore-keys: |
      ${{ runner.os }}-go-

- name: Cache Terraform Plugins
  uses: actions/cache@v3
  with:
    path: ~/.terraform.d/plugin-cache
    key: ${{ runner.os }}-terraform-${{ hashFiles('**/.terraform.lock.hcl') }}
```

### 5. Cost Control

- Use mocks for unit tests (no cloud costs)
- Run full acceptance tests only on main/release branches
- Implement sweepers to clean up orphaned resources
- Set resource quotas in test accounts
- Use spot/preemptible instances where possible
