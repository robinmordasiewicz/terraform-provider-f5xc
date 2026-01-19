# Complete Example: Origin Pool with Healthcheck

Creates an origin pool with healthcheck reference and multiple configuration options.

```hcl
# Optional: Create a healthcheck resource first
# (healthcheck is a REFERENCE BLOCK - only name/namespace/tenant are valid)
resource "f5xc_healthcheck" "example" {
  name      = "my-healthcheck"
  namespace = "default"

  # Health check type - select ONE
  http_health_check {
    use_origin_server_name {}
    path = "/health"
  }
  # OR: tcp_health_check {}
  # OR: grpc_health_check { ... }

  healthy_threshold   = 2
  unhealthy_threshold = 3
  interval            = 15
  timeout             = 5
}

resource "f5xc_origin_pool" "example" {
  name      = "my-origin-pool"
  namespace = "default"

  # TLS Choice - select ONE (use empty block syntax)
  # Option 1: No TLS (plain HTTP to origin)
  no_tls {}

  # Option 2: Use TLS (uncomment and configure)
  # use_tls {
  #   # Nested TLS configuration
  #   skip_server_verification {}
  #   # OR
  #   # use_server_verification {}
  # }

  # Port Choice - select ONE
  # Option 1: Automatic port (80 for HTTP, 443 for HTTPS)
  automatic_port {}

  # Option 2: Same as endpoint port
  # same_as_endpoint_port {}

  # Option 3: Explicit port number (use attribute, not block)
  # port = 8080

  # Origin servers (at least one required)
  origin_servers {
    # Public DNS name
    public_name {
      dns_name = "api.example.com"
    }
  }

  # Healthcheck - REFERENCE BLOCK (only name/namespace/tenant!)
  # This references the separate f5xc_healthcheck resource above
  # DO NOT put configuration parameters here (interval, timeout, etc.)
  healthcheck {
    name      = f5xc_healthcheck.example.name
    namespace = "default"
  }

  # Load balancing algorithm (attribute, not block)
  loadbalancer_algorithm = "ROUND_ROBIN"
}
```

## Reference Blocks vs Inline Blocks

**IMPORTANT**: The `healthcheck` field is a **reference block**, NOT an inline configuration block.

### Reference Blocks (only accept name/namespace/tenant)
- `healthcheck {}` - References a separate `f5xc_healthcheck` resource
- Configuration parameters (interval, timeout, path, etc.) go in the separate resource

### Common Error (WRONG):
```hcl
# DO NOT DO THIS - healthcheck is a reference, not inline config
healthcheck {
  interval_seconds = 30  # ERROR!
  http_request {}        # ERROR!
}
```

## OneOf Groups Explained

### tls_choice
- `no_tls {}` - Plain HTTP to origin (use empty block)
- `use_tls { ... }` - TLS to origin (use block with nested config)

### port_choice
- `automatic_port {}` - Auto-select based on TLS (use empty block)
- `same_as_endpoint_port {}` - Use endpoint's port (use empty block)
- `port = 8080` - Explicit port (use attribute assignment)

## Key Syntax Pattern

Block-type fields like `no_tls`, `automatic_port` use `{}` syntax.
Attribute fields like `port` use `= value` syntax.
Reference fields like `healthcheck` only accept `name`/`namespace`/`tenant`.
