# Complete Example: TCP Load Balancer (Basic Pattern)

Creates a TCP load balancer in F5 Distributed Cloud.

## Required Resources

```hcl
resource "f5xc_tcp_loadbalancer" "example" {
  name      = "my-tcp-lb"
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

  # Listen port
  listen_port = 8080

  # Advertise configuration - use empty block syntax
  advertise_on_public_default_vip {}
}
```

## Key Syntax Notes

1. **origin_pools_weights**: References existing origin pool
2. **advertise_on_public_default_vip**: Use empty block syntax {}
3. **dns_volterra_managed**: Enable F5XC-managed DNS
4. **listen_port**: TCP port for load balancer to listen on

## Important Notes

- Origin pool must exist before creating TCP load balancer
- Use weight and priority for traffic distribution
- Empty block syntax for advertise configuration choices
