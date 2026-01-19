# Complete Example: api_testing (Auto-Generated)

⚠️ **Note**: This is an auto-generated example. Verify required fields in documentation.

```hcl
resource "f5xc_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "default"

  # OneOf group: frequency_choice - choose one
  every_day {}

}
```

## Block-Type Attributes (use empty block syntax)

- `domains {}`
- `every_day {}`
- `every_month {}`
- `every_week {}`

## Simple Attributes

- `annotations = <value>`
- `custom_header_value = <value>`
- `description = <value>`
- `disable = <value>`
- `id = <value>`
