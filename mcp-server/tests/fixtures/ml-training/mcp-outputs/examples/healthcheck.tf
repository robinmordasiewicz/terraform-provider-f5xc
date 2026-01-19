# Complete Example: healthcheck (Auto-Generated)

⚠️ **Note**: This is an auto-generated example. Verify required fields in documentation.

```hcl
resource "f5xc_healthcheck" "example" {
  name      = "example-healthcheck"
  namespace = "default"

  # OneOf group: health_check - choose one
  http_health_check {}

}
```

## Block-Type Attributes (use empty block syntax)

- `http_health_check {}`
- `tcp_health_check {}`
- `udp_icmp_health_check {}`

## Simple Attributes

- `annotations = <value>`
- `description = <value>`
- `disable = <value>`
- `healthy_threshold = <value>`
- `id = <value>`
