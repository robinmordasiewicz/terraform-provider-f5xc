# Complete Example: service_policy (Auto-Generated)

⚠️ **Note**: This is an auto-generated example. Verify required fields in documentation.

```hcl
resource "f5xc_service_policy" "example" {
  name      = "example-service-policy"
  namespace = "default"

  # OneOf group: rule_choice - choose one
  allow_all_requests {}

  # OneOf group: server_choice - choose one
  any_server {}

}
```

## Block-Type Attributes (use empty block syntax)

- `allow_all_requests {}`
- `allow_list {}`
- `any_server {}`
- `deny_all_requests {}`
- `deny_list {}`

## Simple Attributes

- `annotations = <value>`
- `description = <value>`
- `disable = <value>`
- `id = <value>`
- `labels = <value>`
