# Complete Example: service_policy_rule (Auto-Generated)

⚠️ **Note**: This is an auto-generated example. Verify required fields in documentation.

```hcl
resource "f5xc_service_policy_rule" "example" {
  name      = "example-service-policy-rule"
  namespace = "default"

  # OneOf group: asn_choice - choose one
  any_asn {}

  # OneOf group: client_choice - choose one
  any_client {}

  # OneOf group: ip_choice - choose one
  any_ip {}

  # OneOf group: tls_fingerprint_choice - choose one
  ja4_tls_fingerprint {}

}
```

## Block-Type Attributes (use empty block syntax)

- `any_asn {}`
- `any_client {}`
- `any_ip {}`
- `api_group_matcher {}`
- `arg_matchers {}`

## Simple Attributes

- `action = <value>`
- `annotations = <value>`
- `client_name = <value>`
- `description = <value>`
- `disable = <value>`
