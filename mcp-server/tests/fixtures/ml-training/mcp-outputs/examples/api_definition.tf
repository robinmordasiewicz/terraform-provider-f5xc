# Complete Example: api_definition (Auto-Generated)

⚠️ **Note**: This is an auto-generated example. Verify required fields in documentation.

```hcl
resource "f5xc_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "default"

  # OneOf group: schema_updates_strategy - choose one
  mixed_schema_origin {}

}
```

## Block-Type Attributes (use empty block syntax)

- `api_inventory_exclusion_list {}`
- `api_inventory_inclusion_list {}`
- `mixed_schema_origin {}`
- `non_api_endpoints {}`
- `strict_schema_origin {}`

## Simple Attributes

- `annotations = <value>`
- `description = <value>`
- `disable = <value>`
- `id = <value>`
- `labels = <value>`
