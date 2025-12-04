# Go Provider Development Patterns

Deep dive into idiomatic Go patterns for building reliable, performant, and maintainable Terraform providers.

## Core Go Principles for Providers

### Interface-Based Design

Terraform providers benefit from Go's interface system for testability and flexibility.

```go
// ✅ Accept interfaces for flexibility
type ResourceClient interface {
    Create(ctx context.Context, req CreateRequest) (*Resource, error)
    Read(ctx context.Context, id string) (*Resource, error)
    Update(ctx context.Context, id string, req UpdateRequest) (*Resource, error)
    Delete(ctx context.Context, id string) error
}

// ✅ Return concrete types for clarity
func NewClient(config *Config) *Client {
    return &Client{
        httpClient: &http.Client{Timeout: 30 * time.Second},
        baseURL:    config.BaseURL,
    }
}

// ✅ Small, focused interfaces for mocking
type ResourceGetter interface {
    GetResource(ctx context.Context, id string) (*Resource, error)
}

type ResourceCreator interface {
    CreateResource(ctx context.Context, req CreateRequest) (*Resource, error)
}
```

### Composition Over Inheritance

```go
// ✅ Compose functionality through embedding
type Client struct {
    httpClient *http.Client
    baseURL    string
    logger     *slog.Logger
}

type ResourceService struct {
    client *Client  // Composition
}

func (s *ResourceService) Create(ctx context.Context, req CreateRequest) (*Resource, error) {
    return s.client.doCreate(ctx, "/resources", req)
}

// ✅ Embed interfaces for selective implementation
type MockClient struct {
    ResourceGetter  // Only implement what you need
    ResourceCreator
}
```

## Context Propagation

### Context in All API Operations

Every API call should accept and propagate context for cancellation and timeouts.

```go
// ✅ Always accept context as first parameter
func (c *Client) GetResource(ctx context.Context, id string) (*Resource, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL+"/resources/"+id, nil)
    if err != nil {
        return nil, fmt.Errorf("create request: %w", err)
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("execute request: %w", err)
    }
    defer resp.Body.Close()

    // ... handle response
}

// ✅ Propagate context through provider CRUD
func resourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    client := m.(*Client)

    resource, err := client.CreateResource(ctx, buildCreateRequest(d))
    if err != nil {
        return diag.FromErr(err)
    }

    d.SetId(resource.ID)
    return resourceRead(ctx, d, m)
}
```

### Context with Timeout

```go
// ✅ Add timeout for long operations
func resourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    client := m.(*Client)

    // Add timeout for this operation
    ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
    defer cancel()

    resource, err := client.CreateResource(ctx, buildCreateRequest(d))
    if err != nil {
        if ctx.Err() == context.DeadlineExceeded {
            return diag.Errorf("operation timed out after 5 minutes")
        }
        return diag.FromErr(err)
    }

    d.SetId(resource.ID)
    return resourceRead(ctx, d, m)
}
```

## Error Handling Excellence

### Error Wrapping

```go
// ✅ Wrap errors with context at each level
func (c *Client) GetResource(ctx context.Context, id string) (*Resource, error) {
    resp, err := c.doRequest(ctx, "GET", "/resources/"+id, nil)
    if err != nil {
        return nil, fmt.Errorf("get resource %s: %w", id, err)
    }

    var resource Resource
    if err := json.Unmarshal(resp, &resource); err != nil {
        return nil, fmt.Errorf("parse resource %s response: %w", id, err)
    }

    return &resource, nil
}

// ✅ Check wrapped errors
func resourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    client := m.(*Client)

    resource, err := client.GetResource(ctx, d.Id())
    if err != nil {
        var apiErr *APIError
        if errors.As(err, &apiErr) && apiErr.IsNotFound() {
            d.SetId("")  // Resource was deleted externally
            return nil
        }
        return diag.Errorf("reading resource %s: %s", d.Id(), err)
    }

    // ... set attributes
    return nil
}
```

### Sentinel Errors

```go
// ✅ Define sentinel errors for known conditions
var (
    ErrNotFound        = errors.New("resource not found")
    ErrConflict        = errors.New("resource conflict")
    ErrInvalidRequest  = errors.New("invalid request")
    ErrUnauthorized    = errors.New("unauthorized")
    ErrRateLimited     = errors.New("rate limited")
)

// ✅ Use sentinel errors for control flow
func (c *Client) DeleteResource(ctx context.Context, id string) error {
    resp, err := c.doRequest(ctx, "DELETE", "/resources/"+id, nil)
    if err != nil {
        return fmt.Errorf("delete resource %s: %w", id, err)
    }

    switch resp.StatusCode {
    case http.StatusNoContent, http.StatusOK:
        return nil
    case http.StatusNotFound:
        return ErrNotFound  // Already deleted, not an error
    default:
        return fmt.Errorf("unexpected status %d", resp.StatusCode)
    }
}

// ✅ Check sentinel errors
func resourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    client := m.(*Client)

    err := client.DeleteResource(ctx, d.Id())
    if err != nil {
        if errors.Is(err, ErrNotFound) {
            return nil  // Already deleted
        }
        return diag.Errorf("deleting resource: %s", err)
    }

    return nil
}
```

### Custom Error Types

```go
// ✅ Custom error type with behavior
type APIError struct {
    StatusCode int    `json:"status_code"`
    Message    string `json:"message"`
    RequestID  string `json:"request_id"`
    Details    []struct {
        Field   string `json:"field"`
        Message string `json:"message"`
    } `json:"details,omitempty"`
}

func (e *APIError) Error() string {
    return fmt.Sprintf("API error %d: %s (request: %s)", e.StatusCode, e.Message, e.RequestID)
}

func (e *APIError) IsNotFound() bool {
    return e.StatusCode == http.StatusNotFound
}

func (e *APIError) IsConflict() bool {
    return e.StatusCode == http.StatusConflict
}

func (e *APIError) IsRetryable() bool {
    return e.StatusCode == http.StatusTooManyRequests ||
           e.StatusCode == http.StatusServiceUnavailable ||
           e.StatusCode >= 500
}

// ✅ Parse API errors
func (c *Client) parseError(resp *http.Response) error {
    var apiErr APIError
    if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
        return fmt.Errorf("status %d: failed to parse error", resp.StatusCode)
    }
    return &apiErr
}
```

## Concurrency Patterns

### Goroutine Safety in Providers

```go
// ✅ Thread-safe client with mutex for shared state
type Client struct {
    mu         sync.RWMutex
    httpClient *http.Client
    token      string
    tokenExp   time.Time
}

func (c *Client) getToken() string {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.token
}

func (c *Client) refreshToken(ctx context.Context) error {
    c.mu.Lock()
    defer c.mu.Unlock()

    // Check if another goroutine already refreshed
    if time.Now().Before(c.tokenExp) {
        return nil
    }

    newToken, exp, err := c.doTokenRefresh(ctx)
    if err != nil {
        return err
    }

    c.token = newToken
    c.tokenExp = exp
    return nil
}
```

### Parallel Resource Operations

```go
// ✅ Parallel operations with error group
func (c *Client) GetResources(ctx context.Context, ids []string) ([]*Resource, error) {
    g, ctx := errgroup.WithContext(ctx)
    resources := make([]*Resource, len(ids))

    for i, id := range ids {
        i, id := i, id  // capture loop variables
        g.Go(func() error {
            resource, err := c.GetResource(ctx, id)
            if err != nil {
                return fmt.Errorf("get resource %s: %w", id, err)
            }
            resources[i] = resource
            return nil
        })
    }

    if err := g.Wait(); err != nil {
        return nil, err
    }

    return resources, nil
}
```

### Rate Limiting

```go
// ✅ Rate limiter for API calls
type RateLimitedClient struct {
    client  *Client
    limiter *rate.Limiter
}

func NewRateLimitedClient(client *Client, rps float64) *RateLimitedClient {
    return &RateLimitedClient{
        client:  client,
        limiter: rate.NewLimiter(rate.Limit(rps), 1),
    }
}

func (c *RateLimitedClient) GetResource(ctx context.Context, id string) (*Resource, error) {
    if err := c.limiter.Wait(ctx); err != nil {
        return nil, fmt.Errorf("rate limit wait: %w", err)
    }
    return c.client.GetResource(ctx, id)
}
```

## Memory Management

### Efficient Slice Usage

```go
// ✅ Pre-allocate slices when size is known
func flattenResources(resources []*Resource) []interface{} {
    result := make([]interface{}, 0, len(resources))  // Pre-allocate
    for _, r := range resources {
        result = append(result, map[string]interface{}{
            "id":   r.ID,
            "name": r.Name,
        })
    }
    return result
}

// ✅ Avoid unnecessary allocations
func expandStringList(v []interface{}) []string {
    if len(v) == 0 {
        return nil  // Return nil, not empty slice
    }
    result := make([]string, 0, len(v))
    for _, item := range v {
        if s, ok := item.(string); ok && s != "" {
            result = append(result, s)
        }
    }
    return result
}
```

### Object Pooling

```go
// ✅ Pool for frequently allocated objects
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
    buf := bufferPool.Get().(*bytes.Buffer)
    buf.Reset()
    defer bufferPool.Put(buf)

    if body != nil {
        if err := json.NewEncoder(buf).Encode(body); err != nil {
            return nil, fmt.Errorf("encode body: %w", err)
        }
    }

    // ... make request
}
```

### Efficient String Building

```go
// ✅ Use strings.Builder for concatenation
func buildResourcePath(namespace, resourceType, name string) string {
    var b strings.Builder
    b.Grow(len("/config/namespaces/") + len(namespace) + 1 + len(resourceType) + 1 + len(name))
    b.WriteString("/config/namespaces/")
    b.WriteString(namespace)
    b.WriteByte('/')
    b.WriteString(resourceType)
    b.WriteByte('/')
    b.WriteString(name)
    return b.String()
}
```

## Testing Patterns

### Table-Driven Tests

```go
// ✅ Comprehensive table-driven test
func TestExpandStringList(t *testing.T) {
    t.Parallel()

    testCases := map[string]struct {
        input    []interface{}
        expected []string
    }{
        "empty": {
            input:    []interface{}{},
            expected: nil,
        },
        "single": {
            input:    []interface{}{"a"},
            expected: []string{"a"},
        },
        "multiple": {
            input:    []interface{}{"a", "b", "c"},
            expected: []string{"a", "b", "c"},
        },
        "with_nil": {
            input:    []interface{}{"a", nil, "c"},
            expected: []string{"a", "c"},
        },
        "with_empty_string": {
            input:    []interface{}{"a", "", "c"},
            expected: []string{"a", "c"},
        },
    }

    for name, tc := range testCases {
        tc := tc
        t.Run(name, func(t *testing.T) {
            t.Parallel()
            result := expandStringList(tc.input)
            if !reflect.DeepEqual(result, tc.expected) {
                t.Errorf("expected %v, got %v", tc.expected, result)
            }
        })
    }
}
```

### Mock Interfaces

```go
// ✅ Mock implementation for testing
type MockResourceClient struct {
    GetResourceFunc    func(ctx context.Context, id string) (*Resource, error)
    CreateResourceFunc func(ctx context.Context, req CreateRequest) (*Resource, error)
    UpdateResourceFunc func(ctx context.Context, id string, req UpdateRequest) (*Resource, error)
    DeleteResourceFunc func(ctx context.Context, id string) error
}

func (m *MockResourceClient) GetResource(ctx context.Context, id string) (*Resource, error) {
    if m.GetResourceFunc != nil {
        return m.GetResourceFunc(ctx, id)
    }
    return nil, errors.New("GetResourceFunc not implemented")
}

// Usage in tests
func TestResourceRead(t *testing.T) {
    mock := &MockResourceClient{
        GetResourceFunc: func(ctx context.Context, id string) (*Resource, error) {
            return &Resource{ID: id, Name: "test"}, nil
        },
    }

    // ... use mock in test
}
```

### Test Helpers

```go
// ✅ Helper function with t.Helper()
func assertNoError(t *testing.T, err error) {
    t.Helper()
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
}

func assertEqual(t *testing.T, expected, actual interface{}) {
    t.Helper()
    if !reflect.DeepEqual(expected, actual) {
        t.Errorf("expected %v, got %v", expected, actual)
    }
}

// ✅ Setup/teardown with t.Cleanup()
func setupTestResource(t *testing.T, client *Client) *Resource {
    t.Helper()

    resource, err := client.CreateResource(context.Background(), CreateRequest{
        Name: "test-" + acctest.RandString(8),
    })
    assertNoError(t, err)

    t.Cleanup(func() {
        _ = client.DeleteResource(context.Background(), resource.ID)
    })

    return resource
}
```

## Performance Optimization

### Benchmarking

```go
// ✅ Benchmark with memory allocation tracking
func BenchmarkFlattenResources(b *testing.B) {
    resources := make([]*Resource, 100)
    for i := range resources {
        resources[i] = &Resource{
            ID:   fmt.Sprintf("id-%d", i),
            Name: fmt.Sprintf("name-%d", i),
        }
    }

    b.ResetTimer()
    b.ReportAllocs()

    for i := 0; i < b.N; i++ {
        _ = flattenResources(resources)
    }
}

// ✅ Benchmark comparison
func BenchmarkStringConcat(b *testing.B) {
    b.Run("plus", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            _ = "prefix" + "middle" + "suffix"
        }
    })

    b.Run("sprintf", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            _ = fmt.Sprintf("%s%s%s", "prefix", "middle", "suffix")
        }
    })

    b.Run("builder", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            var b strings.Builder
            b.WriteString("prefix")
            b.WriteString("middle")
            b.WriteString("suffix")
            _ = b.String()
        }
    })
}
```

### CPU Profiling

```bash
# Generate CPU profile
go test -cpuprofile=cpu.out -bench=BenchmarkFlattenResources ./internal/provider/...

# Analyze profile
go tool pprof cpu.out
(pprof) top
(pprof) list flattenResources
(pprof) web  # Opens visualization in browser
```

### Memory Profiling

```bash
# Generate memory profile
go test -memprofile=mem.out -bench=BenchmarkFlattenResources ./internal/provider/...

# Analyze allocations
go tool pprof mem.out
(pprof) top --alloc_space
(pprof) list flattenResources
```

## Build and Tooling

### go generate

```go
// ✅ Generate code from specifications
//go:generate go run ../tools/generate-schema.go

// ✅ Generate mocks
//go:generate mockgen -destination=mock_client.go -package=provider . ResourceClient
```

### Module Best Practices

```go
// go.mod
module github.com/example/terraform-provider-example

go 1.21

require (
    github.com/hashicorp/terraform-plugin-framework v1.5.0
    github.com/hashicorp/terraform-plugin-testing v1.6.0
)

// Pin indirect dependencies if needed
require (
    golang.org/x/sync v0.5.0 // indirect
)
```

### Build Tags

```go
// ✅ Build tags for platform-specific code
//go:build !windows
// +build !windows

package provider

// Unix-specific implementation

// ✅ Build tags for tests
//go:build integration
// +build integration

package provider_test

// Integration tests that require external services
```

## Logging Best Practices

### Structured Logging with slog

```go
// ✅ Use structured logging
func (c *Client) GetResource(ctx context.Context, id string) (*Resource, error) {
    c.logger.Debug("fetching resource",
        slog.String("id", id),
        slog.String("endpoint", c.baseURL),
    )

    resource, err := c.doGet(ctx, "/resources/"+id)
    if err != nil {
        c.logger.Error("failed to fetch resource",
            slog.String("id", id),
            slog.Any("error", err),
        )
        return nil, err
    }

    c.logger.Debug("fetched resource",
        slog.String("id", id),
        slog.String("name", resource.Name),
    )

    return resource, nil
}
```

### Terraform Plugin Logging

```go
// ✅ Use tflog for Terraform-aware logging
import "github.com/hashicorp/terraform-plugin-log/tflog"

func resourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    tflog.Debug(ctx, "Creating resource", map[string]interface{}{
        "name": d.Get("name"),
    })

    // ... create logic

    tflog.Info(ctx, "Created resource", map[string]interface{}{
        "id": d.Id(),
    })

    return nil
}
```

## Summary: Go Provider Checklist

### Before Writing Code
- [ ] Design interfaces for testability
- [ ] Plan error handling strategy
- [ ] Consider concurrency requirements

### Implementation
- [ ] Accept context as first parameter
- [ ] Wrap errors with context
- [ ] Use sentinel errors for known conditions
- [ ] Pre-allocate slices when size is known
- [ ] Use sync primitives for shared state

### Testing
- [ ] Write table-driven tests
- [ ] Create mock implementations
- [ ] Use t.Helper() in helper functions
- [ ] Run benchmarks for critical paths
- [ ] Profile memory and CPU

### Quality
- [ ] Run gofmt
- [ ] Run golangci-lint
- [ ] Check test coverage
- [ ] Run race detector: `go test -race`
