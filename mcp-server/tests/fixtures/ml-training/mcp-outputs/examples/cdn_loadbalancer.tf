# Complete Example: cdn_loadbalancer (Auto-Generated)

⚠️ **Note**: This is an auto-generated example. Verify required fields in documentation.

```hcl
resource "f5xc_cdn_loadbalancer" "example" {
  name      = "example-cdn-loadbalancer"
  namespace = "default"

  # OneOf group: service_policy_choice - choose one
  active_service_policies {}

  # OneOf group: rate_limit_choice - choose one
  api_rate_limit {}

  # OneOf group: api_definition_choice - choose one
  api_specification {}

  # OneOf group: waf_choice - choose one
  app_firewall {}

  # OneOf group: bot_defense_choice - choose one
  bot_defense {}

  # OneOf group: challenge_type - choose one
  captcha_challenge {}

  # OneOf group: client_side_defense_choice - choose one
  client_side_defense {}

  # OneOf group: sensitive_data_policy_choice - choose one
  default_sensitive_data_policy {}

  # OneOf group: api_discovery_choice - choose one
  disable_api_discovery {}

  # OneOf group: ip_reputation_choice - choose one
  disable_ip_reputation {}

  # OneOf group: malicious_user_detection_choice - choose one
  disable_malicious_user_detection {}

  # OneOf group: threat_mesh_choice - choose one
  disable_threat_mesh {}

  # OneOf group: loadbalancer_type - choose one
  http {}

  # OneOf group: l7_ddos_auto_mitigation_action - choose one
  l7_ddos_action_block {}

  # OneOf group: slow_ddos_mitigation_choice - choose one
  slow_ddos_mitigation {}

  # OneOf group: user_id_choice - choose one
  user_id_client_ip {}

}
```

## Block-Type Attributes (use empty block syntax)

- `active_service_policies {}`
- `api_rate_limit {}`
- `api_specification {}`
- `app_firewall {}`
- `blocked_clients {}`

## Simple Attributes

- `annotations = <value>`
- `description = <value>`
- `disable = <value>`
- `id = <value>`
- `labels = <value>`
