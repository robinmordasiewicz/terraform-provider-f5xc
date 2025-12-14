# HTTP Load Balancer with Full Security Features

Deploy a production-ready HTTP Load Balancer on F5 Distributed Cloud with:

- **Web Application Firewall (WAF)** - Blocks SQLi, XSS, and other web attacks
- **Bot Defense** - Protects against automated attacks and scrapers
- **Rate Limiting** - Prevents abuse by limiting requests per client
- **JavaScript Challenge** - Client-side bot detection
- **Automatic TLS Certificates** - HTTPS with auto-renewal
- **Health Monitoring** - Active health checks on origin servers

## Prerequisites

- [Terraform](https://www.terraform.io/downloads) >= 1.8
- F5 Distributed Cloud account with API credentials
- A domain you control (for DNS CNAME configuration)
- A backend origin server accessible from F5 Distributed Cloud

## Quick Start

### 1. Set Environment Variables

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io"
export F5XC_API_TOKEN="your-api-token"
```

### 2. Configure Your Deployment

```bash
cp terraform.tfvars.example terraform.tfvars
```

Edit `terraform.tfvars` with your values:
- `domain` - Your application's domain name
- `origin_server` - Your backend server's DNS name or IP
- Optionally customize security features

### 3. Deploy

```bash
terraform init
terraform plan
terraform apply
```

### 4. Configure DNS

After deployment, create a CNAME record pointing your domain to the F5 Distributed Cloud edge:

```
app.example.com → ves-io-app-example-com.ac.vh.ves.io
```

### 5. Verify

1. Wait a few minutes for TLS certificate provisioning
2. Access your application at `https://your-domain.com`
3. Check the F5 Distributed Cloud Console for traffic and security events

## Configuration Options

### Using an Existing Namespace

To deploy into an existing namespace instead of creating a new one:

```hcl
namespace_name   = "my-existing-namespace"
create_namespace = false
```

### Disabling Security Features

For testing, you can disable individual security features:

```hcl
enable_waf           = false
enable_bot_defense   = false
enable_rate_limiting = false
```

### Adjusting Rate Limits

Change the rate limiting threshold:

```hcl
rate_limit_requests = 200  # requests per minute
```

## Clean Up

To remove all created resources:

```bash
terraform destroy
```

## Resources Created

| Resource | Description |
|----------|-------------|
| `f5xc_namespace` | Namespace for resource isolation (optional) |
| `f5xc_healthcheck` | HTTP health check for origin pool |
| `f5xc_origin_pool` | Backend server pool configuration |
| `f5xc_app_firewall` | Web Application Firewall policy |
| `f5xc_rate_limiter` | Rate limiting configuration |
| `f5xc_http_loadbalancer` | HTTP load balancer with all features |

## Full Documentation

For detailed documentation including architecture diagrams and troubleshooting, see:
[HTTP Load Balancer Guide](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest/docs/guides/http-loadbalancer)

## Support

- [F5XC Provider Documentation](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest/docs)
- [F5 Distributed Cloud Documentation](https://docs.cloud.f5.com/)
- [GitHub Issues](https://github.com/robinmordasiewicz/terraform-provider-f5xc/issues)
