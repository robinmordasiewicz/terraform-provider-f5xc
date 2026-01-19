# Complete Example: App Firewall (WAF)

Creates a Web Application Firewall policy.

```hcl
resource "f5xc_app_firewall" "example" {
  name      = "my-waf-policy"
  namespace = "default"

  # Enforcement mode - select ONE (use empty block syntax)
  # Option 1: Blocking mode (actively blocks attacks)
  blocking {}

  # Option 2: Monitoring mode (logs only, no blocking)
  # monitoring {}

  # Anonymization setting - select ONE
  default_anonymization {}
  # OR: custom_anonymization { ... }
  # OR: disable_anonymization {}

  # Bot protection settings - select ONE
  default_bot_setting {}
  # OR: custom_bot_protection_setting { ... }

  # Detection settings - select ONE
  default_detection_settings {}
  # OR: detection_settings { ... }
}
```

## Key OneOf Groups

### enforcement_mode_choice
- `blocking {}` - Actively block attacks
- `monitoring {}` - Detection only, no blocking

### anonymization_setting
- `default_anonymization {}` - Use default settings
- `custom_anonymization { ... }` - Custom configuration
- `disable_anonymization {}` - No anonymization

## Important

All OneOf options use **empty block syntax** `{}`, not boolean assignment.
