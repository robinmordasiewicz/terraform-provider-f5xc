# Complete Example: app_api_group (Auto-Generated)

⚠️ **Note**: This is an auto-generated example. Verify required fields in documentation.

```hcl
resource "f5xc_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "default"

  # OneOf group: scope_choice - choose one
  bigip_virtual_server {}

}
```

## Block-Type Attributes (use empty block syntax)

- `bigip_virtual_server {}`
- `cdn_loadbalancer {}`
- `elements {}`
- `http_loadbalancer {}`

## Simple Attributes

- `annotations = <value>`
- `description = <value>`
- `disable = <value>`
- `id = <value>`
- `labels = <value>`
