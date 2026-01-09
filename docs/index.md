---
title: F5 Distributed Cloud Terraform Provider
---

# F5 Distributed Cloud Terraform Provider

<div class="landing-page">

<div class="hero">
  <h2>Manage F5 Distributed Cloud Resources with Terraform</h2>
  <p>Community-driven Terraform provider for F5 Distributed Cloud (F5XC) built using the HashiCorp Terraform Plugin Framework. Implements F5's public OpenAPI specifications to manage F5XC resources via Infrastructure as Code.</p>

  <div class="buttons">
    <a href="getting-started/installation/" class="md-button md-button--primary">Get Started</a>
    <a href="https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest" class="md-button">Terraform Registry</a>
  </div>
</div>

## Key Features

<div class="feature-grid">
  <div class="feature">
    <div class="feature-icon">🚀</div>
    <h3>Complete Resource Coverage</h3>
    <p>Manage HTTP load balancers, namespaces, security policies, and more through Terraform</p>
  </div>

  <div class="feature">
    <div class="feature-icon">🔒</div>
    <h3>Secure Authentication</h3>
    <p>Support for API tokens and P12 certificate authentication</p>
  </div>

  <div class="feature">
    <div class="feature-icon">📖</div>
    <h3>OpenAPI-Based</h3>
    <p>Automatically generated from F5's official OpenAPI specifications</p>
  </div>

  <div class="feature">
    <div class="feature-icon">🧪</div>
    <h3>Well-Tested</h3>
    <p>Comprehensive acceptance tests ensure reliability</p>
  </div>

  <div class="feature">
    <div class="feature-icon">🔄</div>
    <h3>Import Support</h3>
    <p>Import existing F5XC resources into Terraform state</p>
  </div>

  <div class="feature">
    <div class="feature-icon">🛠️</div>
    <h3>Provider Functions</h3>
    <p>Utility functions for Secret Management and more</p>
  </div>
</div>

## Quick Example

```hcl
terraform {
  required_providers {
    f5xc = {
      source  = "robinmordasiewicz/f5xc"
      version = "~> 3.0"
    }
  }
}

provider "f5xc" {
  api_url = "https://console.ves.volterra.io"
  # api_token set via F5XC_API_TOKEN environment variable
}

resource "f5xc_namespace" "example" {
  name = "example-namespace"
}

resource "f5xc_http_loadbalancer" "example" {
  name      = "example-lb"
  namespace = f5xc_namespace.example.name

  domains = ["example.com"]

  https_auto_cert {
    http_redirect       = true
    add_hsts            = false
    no_mtls             = true
    default_header      = true
    enable_path_normalize = true
  }

  default_route_pools {
    pool {
      name      = f5xc_origin_pool.example.name
      namespace = f5xc_namespace.example.name
    }
    weight   = 1
    priority = 1
  }

  advertise_on_public_default_vip = true
}
```

## Getting Started

<div class="highlight-boxes">
  <div class="highlight">
    <h3>📦 Installation</h3>
    <p>Install the provider from the Terraform Registry</p>
    <a href="getting-started/installation/">Learn more →</a>
  </div>

  <div class="highlight">
    <h3>🔐 Authentication</h3>
    <p>Configure API token or P12 certificate authentication</p>
    <a href="getting-started/authentication/">Learn more →</a>
  </div>

  <div class="highlight">
    <h3>⚡ Quick Start</h3>
    <p>Deploy your first F5XC resource in minutes</p>
    <a href="getting-started/quick-start/">Learn more →</a>
  </div>
</div>

## Resources

Explore the available resources and data sources:

- **[Resources](resources/)**: Create and manage F5XC infrastructure
- **[Data Sources](data-sources/)**: Query existing F5XC resources
- **[Functions](functions/)**: Utility functions for advanced use cases
- **[Guides](guides/)**: Step-by-step tutorials for common scenarios

## Community & Support

This is a community-driven provider. Contributions are welcome!

- **GitHub**: [robinmordasiewicz/terraform-provider-f5xc](https://github.com/robinmordasiewicz/terraform-provider-f5xc)
- **Issues**: [Report bugs or request features](https://github.com/robinmordasiewicz/terraform-provider-f5xc/issues)
- **Contributing**: [Development guide](contributing/development/)

---

<div class="footer-note">
Built with ❤️ by the community • <a href="https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest">View on Terraform Registry</a>
</div>

</div>
