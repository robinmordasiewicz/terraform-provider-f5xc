# Parallel Resource Test Development

Spawn multiple terraform-test-developer agents in parallel to develop tests for the specified resources.

## Usage

```
/test-resources resource1 resource2 resource3 ...
```

## Instructions

1. Parse the resource names from the arguments: $ARGUMENTS
2. For EACH resource, spawn a terraform-test-developer agent in PARALLEL (all Task calls in ONE message)
3. Each agent should:
   - First invoke Skill(terraform-provider-testing)
   - Read the resource at internal/provider/{resource}_resource.go
   - Create comprehensive tests at internal/provider/{resource}_resource_test.go
   - Run tests and report results

## Agent Prompt Template

For each resource, use this prompt:

```
Develop and run acceptance tests for the {resource_name} resource in this Terraform provider.

FIRST: Invoke Skill(terraform-provider-testing) to load all testing patterns and standards.

THEN:
1. Read internal/provider/{resource_name}_resource.go to understand the schema
2. Check if internal/provider/{resource_name}_resource_test.go exists
3. If exists, review and enhance; if not, create new test file
4. Implement these test functions using modern assertion patterns:
   - TestAcc{PascalCaseName}Resource_basic
   - TestAcc{PascalCaseName}Resource_update (if applicable)
   - Import test step within _basic
5. Use custom namespace (f5xc_namespace.test), NOT system namespace
6. Use ConfigPlanChecks and ConfigStateChecks with knownvalue assertions
7. Run: TF_ACC=1 go test -v -timeout 15m -run "TestAcc{PascalCaseName}" ./internal/provider/...
8. Report results including any failures and their root causes

IMPORTANT: If tests fail due to schema or generator issues, identify the root cause but do NOT skip tests.
```

## Parallel Execution

**CRITICAL**: Spawn ALL agents in a SINGLE message to maximize parallelism.

Example for 3 resources:
- Task 1: terraform-test-developer for namespace_resource
- Task 2: terraform-test-developer for healthcheck_resource
- Task 3: terraform-test-developer for origin_pool_resource

All three run simultaneously.

## Result Aggregation

After all agents complete, summarize:
- Total resources processed
- Tests created/updated
- Pass/fail status for each
- Any issues requiring manual follow-up
