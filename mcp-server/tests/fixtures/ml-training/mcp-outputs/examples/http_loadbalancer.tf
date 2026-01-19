# Complete Example: HTTP Load Balancer (Basic Pattern)

Creates an HTTP load balancer with an origin pool pointing to httpbin.org.

## Required Resources

```hcl
# 1. Create the origin pool
resource "f5xc_origin_pool" "example" {
  name      = "httpbin-pool"
  namespace = "default"

  # TLS choice - use empty block syntax
  no_tls {}

  # Port choice - use empty block syntax
  automatic_port {}

  origin_servers {
    public_name {
      dns_name = "httpbin.org"
    }
  }

  loadbalancer_algorithm = "ROUND_ROBIN"
}

# 2. Create the HTTP load balancer
resource "f5xc_http_loadbalancer" "example" {
  name      = "httpbin-lb"
  namespace = "default"
  domains   = ["httpbin.example.com"]

  # Load balancer type - use empty block syntax
  https_auto_cert {
    http_redirect = true
    add_hsts      = false
  }

  # Advertising - use empty block syntax
  advertise_on_public_default_vip {}

  # Origin pool reference
  default_route_pools {
    pool {
      name      = f5xc_origin_pool.example.name
      namespace = "default"
    }
  }

  # Security options - use empty block syntax
  disable_waf {}
  no_challenge {}
  disable_rate_limit {}
  no_service_policies {}
  disable_bot_defense {}
  disable_api_definition {}
  disable_api_discovery {}

  # Load balancing algorithm - use empty block syntax
  round_robin {}
}
```

## Key Syntax Notes

1. **Empty blocks for OneOf choices**: `no_tls {}`, `advertise_on_public_default_vip {}`
2. **Nested blocks with values**: `https_auto_cert { http_redirect = true }`
3. **Reference syntax**: `f5xc_origin_pool.example.name`

## Common Mistakes to Avoid

```hcl
# WRONG - Do not use boolean assignment for these fields!
no_tls = true                          # ERROR: use no_tls {}
advertise_on_public_default_vip = true # ERROR: use advertise_on_public_default_vip {}
round_robin = true                     # ERROR: use round_robin {}
```
