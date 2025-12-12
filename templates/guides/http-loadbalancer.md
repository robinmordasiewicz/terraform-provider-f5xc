---
page_title: "Guide: HTTP Load Balancer with Security Features"
subcategory: "Guides"
description: |-
  Step-by-step guide to deploy a production-ready HTTP Load Balancer with WAF,
  bot defense, and rate limiting using F5 Distributed Cloud and Terraform.
---

# HTTP Load Balancer with Security Features

This guide walks you through deploying a complete HTTP Load Balancer on F5 Distributed Cloud. By the end, you'll have a production-ready load balancer with:

- **Web Application Firewall (WAF)** - Blocks common web attacks (SQLi, XSS, etc.)
- **Bot Defense** - Protects against automated attacks and scrapers
- **Rate Limiting** - Prevents abuse by limiting requests per client
- **JavaScript Challenge** - Client-side bot detection
- **Automatic TLS Certificates** - HTTPS with auto-renewal
- **Health Monitoring** - Active health checks on origin servers
- **Threat Mesh** - Global threat intelligence sharing

## Prerequisites

Before you begin, ensure you have:

- **F5 Distributed Cloud Account** - Sign up at <https://www.f5.com/cloud/products/distributed-cloud-console> if you don't have one
- **API Token** - Generate credentials from the F5XC Console at <https://docs.cloud.f5.com/docs/how-to/user-mgmt/credentials>
- **Terraform >= 1.8** - Download and install from <https://www.terraform.io/downloads>
- **A Domain** - Domain you control for DNS configuration
- **Backend Origin Server** - Your application server accessible from the internet

## Quick Start

### Step 1: Clone the Repository

```bash
git clone https://github.com/robinmordasiewicz/terraform-provider-f5xc.git
cd terraform-provider-f5xc/examples/guides/http-loadbalancer
```

### Step 2: Set Environment Variables

Configure authentication using environment variables. **Never commit credentials to version control.**

```bash
export VES_API_URL="https://your-tenant.console.ves.volterra.io/api"
export VES_API_TOKEN="your-api-token"
```

-> **Tip:** Add these to your shell profile (`~/.bashrc` or `~/.zshrc`) for persistence across terminal sessions.

### Step 3: Configure Your Deployment

```bash
cp terraform.tfvars.example terraform.tfvars
```

Edit `terraform.tfvars` with your values:

```hcl
# Your application's domain
domain = "app.example.com"

# Your backend server
origin_server = "origin.example.com"
origin_port   = 443

# Namespace configuration
namespace_name   = "my-app"
create_namespace = true

# Security features (all enabled by default)
enable_waf           = true
enable_bot_defense   = true
enable_rate_limiting = true
rate_limit_requests  = 100
```

### Step 4: Deploy

```bash
terraform init
terraform plan
terraform apply
```

Review the plan output, then type `yes` to confirm deployment.

### Step 5: Configure DNS

After deployment, Terraform outputs a CNAME target. Create a DNS record:

| Type | Name | Value |
|------|------|-------|
| CNAME | app.example.com | ves-io-app-example-com.ac.vh.ves.io |

~> **Note:** DNS propagation may take up to 48 hours, though typically completes within minutes.

### Step 6: Verify

1. **Wait for TLS provisioning** - Auto-cert typically provisions within 5 minutes
2. **Access your application** - Navigate to `https://your-domain.com`
3. **Check the console** - View traffic and security events in F5 Distributed Cloud Console

## Configuration Options

### Using an Existing Namespace

To deploy into an existing namespace instead of creating a new one:

```hcl
namespace_name   = "my-existing-namespace"
create_namespace = false
```

### Customizing Security Features

Each security feature can be enabled or disabled independently:

```hcl
# Disable WAF for testing (not recommended for production)
enable_waf = false

# Disable bot defense
enable_bot_defense = false

# Disable rate limiting
enable_rate_limiting = false
```

### Adjusting Rate Limits

Fine-tune rate limiting for your application's needs:

```hcl
rate_limit_requests = 200  # requests per minute per client
```

The default burst multiplier is 10x, allowing temporary spikes above the limit.

### Custom Labels

Add labels for organization and filtering:

```hcl
labels = {
  environment = "production"
  team        = "platform"
  cost_center = "engineering"
}
```

## Architecture

This guide creates the following resources:

```
                    ┌─────────────────────────────────────────┐
                    │        F5 Distributed Cloud             │
                    │                                         │
 Users ──────────►  │  ┌─────────────────────────────────┐   │
                    │  │    HTTP Load Balancer           │   │
                    │  │  ┌──────────────────────────┐   │   │
                    │  │  │ • TLS Termination        │   │   │
                    │  │  │ • JavaScript Challenge   │   │   │
                    │  │  │ • WAF (blocking mode)    │   │   │
                    │  │  │ • Bot Defense            │   │   │
                    │  │  │ • Rate Limiting          │   │   │
                    │  │  │ • Threat Mesh            │   │   │
                    │  │  └──────────────────────────┘   │   │
                    │  └──────────────┬──────────────────┘   │
                    │                 │                       │
                    │  ┌──────────────▼──────────────────┐   │
                    │  │         Origin Pool             │   │
                    │  │  ┌──────────────────────────┐   │   │
                    │  │  │ • Health Checks          │   │   │
                    │  │  │ • TLS to Origin          │   │   │
                    │  │  │ • Load Balancing         │   │   │
                    │  │  └──────────────────────────┘   │   │
                    │  └──────────────┬──────────────────┘   │
                    └─────────────────┼───────────────────────┘
                                      │
                                      ▼
                              Your Origin Server
```

### Resources Created

| Resource | Purpose |
|----------|---------|
| `f5xc_namespace` | Isolates resources (optional) |
| `f5xc_healthcheck` | Monitors origin server health |
| `f5xc_origin_pool` | Defines backend servers |
| `f5xc_app_firewall` | WAF configuration |
| `f5xc_rate_limiter` | Rate limiting policy |
| `f5xc_http_loadbalancer` | Main load balancer |

## Troubleshooting

### Certificate Not Provisioning

**Symptom:** HTTPS returns certificate errors after deployment.

**Solutions:**

1. Verify DNS CNAME is correctly configured
2. Wait up to 10 minutes for certificate provisioning
3. Check the Load Balancer status in F5XC Console

### 502 Bad Gateway

**Symptom:** Load balancer returns 502 errors.

**Solutions:**

1. Verify `origin_server` is accessible from the internet
2. Check health check path returns HTTP 200
3. Verify origin port is correct
4. Check origin server TLS configuration

### WAF Blocking Legitimate Traffic

**Symptom:** Valid requests are blocked by WAF.

**Solutions:**

1. Check Security Analytics in F5XC Console
2. Review blocked request details
3. Consider temporarily setting WAF to monitoring mode:

   ```hcl
   enable_waf = false  # Disable in Terraform
   ```

### Rate Limiting Too Aggressive

**Symptom:** Users hitting rate limits during normal usage.

**Solutions:**

1. Increase the rate limit:

   ```hcl
   rate_limit_requests = 500
   ```

2. Review rate limiting events in the console
3. Consider user identification beyond IP address

## Clean Up

To remove all resources created by this guide:

```bash
terraform destroy
```

Type `yes` to confirm destruction.

!> **Warning:** This will immediately remove the load balancer and all associated resources. Traffic to your domain will no longer be handled by F5XC.

## Next Steps

Now that you have a basic HTTP Load Balancer deployed, consider exploring:

- [Origin Pool Resource](../resources/origin_pool) - Add multiple origins for redundancy
- [App Firewall Resource](../resources/app_firewall) - Customize WAF rules
- [Service Policy Resource](../resources/service_policy) - Add custom access control
- [TCP Load Balancer Resource](../resources/tcp_loadbalancer) - For non-HTTP applications

## Support

- **Provider Documentation:** [F5XC Provider](../index)
- **F5 Documentation:** [F5 Distributed Cloud Docs](https://docs.cloud.f5.com/)
- **Issues:** [GitHub Issues](https://github.com/robinmordasiewicz/terraform-provider-f5xc/issues)
