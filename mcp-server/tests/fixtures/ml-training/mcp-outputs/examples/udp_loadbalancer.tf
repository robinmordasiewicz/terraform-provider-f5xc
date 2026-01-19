# Complete Example: UDP Load Balancer (Basic Pattern)

Creates a UDP load balancer in F5 Distributed Cloud.

## Required Resources

```hcl
resource "f5xc_udp_loadbalancer" "example" {
  name      = "my-udp-lb"
  namespace = "default"

  # DNS configuration
  dns_volterra_managed = true

  # Origin pool reference
  origin_pools_weights {
    pool {
      name      = "my-origin-pool"
      namespace = "default"
    }
    weight   = 1
    priority = 1
  }

  # Port range (single port or range like "53" or "5000-5010")
  port_ranges = "53"

  # Advertise configuration - use empty block syntax
  advertise_on_public_default_vip {}
}
```

## Key Syntax Notes

1. **origin_pools_weights**: References existing origin pool
2. **advertise_on_public_default_vip**: Use empty block syntax {}
3. **dns_volterra_managed**: Enable F5XC-managed DNS
4. **port_ranges**: Single port ("53") or range ("5000-5010")

## Important Notes

- Origin pool must exist before creating UDP load balancer
- Use weight and priority for traffic distribution
- Empty block syntax for advertise configuration choices
- Port ranges support single port or range format
