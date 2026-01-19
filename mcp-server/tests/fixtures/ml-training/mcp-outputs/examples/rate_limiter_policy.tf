# Complete Example: rate_limiter_policy (Auto-Generated)

⚠️ **Note**: This is an auto-generated example. Verify required fields in documentation.

```hcl
resource "f5xc_rate_limiter_policy" "example" {
  name      = "example-rate-limiter-policy"
  namespace = "default"

  # OneOf group: server_choice - choose one
  any_server {}

}
```

## Block-Type Attributes (use empty block syntax)

- `any_server {}`
- `rules {}`
- `server_name_matcher {}`
- `server_selector {}`

## Simple Attributes

- `annotations = <value>`
- `description = <value>`
- `disable = <value>`
- `id = <value>`
- `labels = <value>`
