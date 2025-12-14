# Debugging Terraform Providers with Delve

Comprehensive guide to debugging Terraform provider code using the Delve debugger.

## Installation

### Install Delve

```bash
# Install latest version
go install github.com/go-delve/delve/cmd/dlv@latest

# Verify installation
dlv version
```

### IDE Integration

**VS Code** - Install the Go extension, then add to `.vscode/launch.json`:

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Debug Acceptance Test",
            "type": "go",
            "request": "launch",
            "mode": "test",
            "program": "${workspaceFolder}/internal/provider",
            "args": [
                "-test.run",
                "TestAccResource_basic",
                "-test.v"
            ],
            "env": {
                "TF_ACC": "1",
                "F5XC_P12_FILE": "/path/to/cert.p12",
                "F5XC_P12_PASSWORD": "password",  // pragma: allowlist secret
                "F5XC_API_URL": "https://api.example.com"
            }
        }
    ]
}
```

**GoLand/IntelliJ** - Right-click test function → Debug

## Basic Debugging Commands

### Starting a Debug Session

```bash
# Debug a specific test
dlv test ./internal/provider/... -- -test.run TestAccResource_basic -test.v

# Debug with TF_ACC
TF_ACC=1 dlv test ./internal/provider/... -- -test.run TestAccResource_basic -test.v

# Debug the provider binary
dlv debug ./main.go

# Attach to running process
dlv attach <pid>
```

### Essential Commands

| Command | Short | Description |
|---------|-------|-------------|
| `break` | `b` | Set breakpoint |
| `continue` | `c` | Continue execution |
| `next` | `n` | Step over (next line) |
| `step` | `s` | Step into function |
| `stepout` | `so` | Step out of function |
| `print` | `p` | Print variable value |
| `locals` | | List local variables |
| `args` | | List function arguments |
| `stack` | `bt` | Print stack trace |
| `goroutines` | `grs` | List goroutines |
| `quit` | `q` | Exit debugger |

### Breakpoint Management

```bash
# Set breakpoint at function
(dlv) break resourceCreate
(dlv) break internal/provider/resource_namespace.go:42

# Set conditional breakpoint
(dlv) break resourceCreate if name == "test-resource"

# List breakpoints
(dlv) breakpoints
(dlv) bp

# Clear breakpoint
(dlv) clear 1

# Clear all breakpoints
(dlv) clearall
```

## Debugging Provider CRUD Operations

### Debug Create Operation

```bash
TF_ACC=1 dlv test ./internal/provider/... -- -test.run TestAccNamespace_basic -test.v

# In dlv:
(dlv) break resourceNamespaceCreate
(dlv) continue

# When breakpoint hits, inspect:
(dlv) args                           # Function arguments
(dlv) print ctx                      # Context value
(dlv) print d.Get("name")           # Resource data
(dlv) print req                      # API request
(dlv) next                           # Step through
```

### Debug Read Operation

```bash
(dlv) break resourceNamespaceRead
(dlv) continue

# Inspect state
(dlv) print d.Id()
(dlv) print d.State()

# Check API response
(dlv) print resp
(dlv) print resp.Metadata.Name
```

### Debug Update Operation

```bash
(dlv) break resourceNamespaceUpdate
(dlv) continue

# Compare old vs new values
(dlv) print d.GetChange("description")
(dlv) print d.HasChange("labels")
```

### Debug Delete Operation

```bash
(dlv) break resourceNamespaceDelete
(dlv) continue

# Verify ID being deleted
(dlv) print d.Id()
```

## Debugging API Client

### HTTP Request/Response Inspection

```bash
(dlv) break (*Client).doRequest
(dlv) continue

# Inspect request
(dlv) print method
(dlv) print path
(dlv) print body

# After response
(dlv) print resp.StatusCode
(dlv) print string(respBody)
```

### Debug Authentication

```bash
(dlv) break (*Client).authenticate
(dlv) continue

# Check credentials
(dlv) print c.config.P12File
(dlv) print c.config.APIUrl

# Inspect certificate loading
(dlv) break loadP12Certificate
```

## Debugging State Management

### Inspect Terraform State

```bash
(dlv) break resourceNamespaceRead
(dlv) continue

# Print all state attributes
(dlv) print d.State().Attributes

# Check specific attribute
(dlv) print d.Get("metadata")
(dlv) print d.Get("spec")

# Check if attribute changed
(dlv) print d.HasChange("description")
```

### Debug State Drift

```bash
# Set breakpoint where state is set
(dlv) break d.Set
(dlv) continue

# Inspect what's being set
(dlv) print key
(dlv) print value

# Compare with API response
(dlv) print apiResponse
```

## Debugging Test Failures

### Common Scenarios

**Test: attribute not found in schema**
```bash
(dlv) break Schema
(dlv) continue
(dlv) print Attributes  # Inspect schema definition
```

**Test: planned value does not match**
```bash
(dlv) break resourceRead
(dlv) continue
(dlv) print d.Get("computed_attr")  # Check what's being read
```

**Test: CheckDestroy failed**
```bash
(dlv) break testAccCheckResourceDestroy
(dlv) continue
(dlv) print s.RootModule().Resources  # List resources
(dlv) print rs.Primary.ID             # Check ID
```

**Test: Import produces diff**
```bash
(dlv) break ImportState
(dlv) continue
(dlv) print d.State().Attributes  # Before import
(dlv) step
(dlv) print d.State().Attributes  # After import
```

### Debug Test Helper Functions

```bash
(dlv) break testAccCheckResourceExists
(dlv) continue

# Inspect state lookup
(dlv) print resourceName
(dlv) print s.RootModule().Resources
(dlv) print rs.Primary.ID
```

## Advanced Debugging Techniques

### Conditional Breakpoints

```bash
# Break only for specific resource
(dlv) break resourceCreate if d.Get("name") == "test-resource"

# Break on specific error
(dlv) break resourceRead if err != nil

# Break on specific status code
(dlv) break parseResponse if resp.StatusCode != 200
```

### Watchpoints

```bash
# Watch variable changes (requires hardware support)
(dlv) watch d.State().Attributes
```

### Goroutine Debugging

```bash
# List all goroutines
(dlv) goroutines

# Switch to specific goroutine
(dlv) goroutine 5

# Print current goroutine
(dlv) goroutine

# Stack trace for all goroutines
(dlv) goroutines -t
```

### Remote Debugging

```bash
# Start delve server
dlv debug --headless --listen=:2345 --api-version=2

# Connect from another terminal
dlv connect :2345
```

## Debugging F5 XC Provider

### Debug Namespace Resource

```bash
TF_ACC=1 F5XC_P12_FILE="/path/to/cert.p12" \
  F5XC_P12_PASSWORD="password" \  # pragma: allowlist secret
  F5XC_API_URL="https://api.example.com" \
  dlv test ./internal/provider/... -- -test.run TestAccNamespaceResource_basic -test.v

(dlv) break internal/provider/namespace_resource.go:resourceNamespaceCreate
(dlv) continue

# Inspect API request
(dlv) print req.Metadata
(dlv) print req.Spec

# Check response
(dlv) step
(dlv) print resp
```

### Debug Certificate Authentication

```bash
(dlv) break loadCertificate
(dlv) continue

# Inspect P12 loading
(dlv) print p12Path
(dlv) print password

# Check certificate parsing
(dlv) step
(dlv) print cert
(dlv) print key
```

### Debug API Errors

```bash
(dlv) break handleAPIError
(dlv) continue

# Inspect error details
(dlv) print resp.StatusCode
(dlv) print body
(dlv) print err
```

## Performance Debugging

### Find Slow Operations

```bash
# Add timing breakpoints
(dlv) break resourceCreate
(dlv) continue
# Note: time
(dlv) next
(dlv) next
# Compare times manually

# Better: use profiling
go test -cpuprofile=cpu.out -run TestAccResource_basic ./internal/provider/...
go tool pprof cpu.out
```

### Memory Debugging

```bash
# Check allocations
(dlv) break runtime.mallocgc
(dlv) continue

# Inspect heap
(dlv) print runtime.memstats
```

## Debugging Tips

### Best Practices

1. **Start narrow**: Debug the smallest reproducible case
2. **Use conditional breakpoints**: Avoid stopping on every iteration
3. **Print intermediate values**: Don't assume, verify
4. **Check goroutine context**: Ensure you're in the right goroutine
5. **Use `step` wisely**: Step into functions you want to inspect
6. **Save session**: Use history for repeated debugging

### Common Pitfalls

1. **Missing TF_ACC**: Tests skip without TF_ACC=1
2. **Wrong package path**: Ensure correct path to test files
3. **Optimized binaries**: Build with `-gcflags="all=-N -l"` for better debugging
4. **Stale binary**: Rebuild before debugging

### Build Without Optimizations

```bash
# Build provider without optimizations for better debugging
go build -gcflags="all=-N -l" -o terraform-provider-example

# Test without optimizations
go test -gcflags="all=-N -l" -c -o provider.test ./internal/provider/...
dlv exec ./provider.test -- -test.run TestAccResource_basic -test.v
```

## Quick Reference Card

```
# Start debugging test
TF_ACC=1 dlv test ./internal/provider/... -- -test.run TestName -test.v

# Common workflow
(dlv) b functionName         # Set breakpoint
(dlv) c                      # Continue to breakpoint
(dlv) n                      # Step over
(dlv) s                      # Step into
(dlv) p variable            # Print variable
(dlv) locals                # Show local vars
(dlv) bt                    # Stack trace
(dlv) q                     # Quit

# Useful inspections
(dlv) p d.Get("name")       # Get resource attribute
(dlv) p d.Id()              # Get resource ID
(dlv) p d.State()           # Get full state
(dlv) p err                 # Check error
(dlv) p resp.StatusCode     # Check HTTP status
```

## Integration with CI/CD

### Debugging CI Failures Locally

```bash
# Reproduce CI environment
export TF_ACC=1
export TF_LOG=DEBUG
export TF_LOG_PATH=/tmp/terraform.log

# Run same test as CI
go test -v -timeout 30m -run TestAccResource_basic ./internal/provider/...

# If fails, debug with delve
dlv test ./internal/provider/... -- -test.run TestAccResource_basic -test.v
```

### Capturing Debug Output

```bash
# Run with trace logging
TF_ACC=1 TF_LOG=TRACE go test -v -run TestAccResource_basic ./internal/provider/... 2>&1 | tee test.log

# Analyze log for debugging clues
grep "API request" test.log
grep "Error" test.log
```
