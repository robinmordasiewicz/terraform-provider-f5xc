<a href="https://terraform.io">
    <img src="docs/images/tf.png" alt="Terraform logo" title="Terraform" align="left" height="50" />
</a>

# Terraform Provider for F5 Distributed Cloud (F5XC)

> Open Source Community Provider

This is an independent, community-driven Terraform provider for F5 Distributed Cloud (F5XC), built entirely from public F5 API documentation using the official Terraform Plugin Framework.

## 🎯 Project Goals

- **Vendor Independence**: No dependency on proprietary code or tools
- **Community Governance**: Transparent, community-driven development
- **Open Source**: Mozilla Public License (MPL) 2.0
- **Feature Parity**: Comprehensive coverage of F5 Distributed Cloud services
- **Quality**: Production-ready with comprehensive testing and documentation

## 📚 Documentation

- **Provider Documentation**: [Terraform Registry](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest/docs)
- **API Reference**: [F5 Distributed Cloud API](https://docs.cloud.f5.com/)
- **Examples**: [examples/](examples/)

## 🚀 Quick Start

### Installation

Add the provider to your Terraform configuration:

```terraform
terraform {
  required_providers {
    f5xc = {
      source  = "robinmordasiewicz/f5xc"
      version = "~> 0.1"
    }
  }
}

provider "f5xc" {
  # API token can be set via F5XC_API_TOKEN environment variable
  api_token = var.f5xc_api_token

  # API URL defaults to https://console.ves.volterra.io/api
  api_url = "https://console.ves.volterra.io/api"
}
```

### Authentication

Set your F5 Distributed Cloud API token as an environment variable:

```bash
export F5XC_API_TOKEN="your-api-token-here"
```

Or configure it directly in the provider block (not recommended for production).

### Example Usage

Create a namespace:

```terraform
resource "f5xc_namespace" "example" {
  name        = "production"
  description = "Production environment namespace"

  labels = {
    environment = "production"
    managed-by  = "terraform"
  }
}
```

Reference an existing namespace:

```terraform
data "f5xc_namespace" "shared" {
  name = "shared"
}

output "shared_namespace_id" {
  value = data.f5xc_namespace.shared.id
}
```

## 📦 Available Resources

### Current (v0.1.0 - Proof of Concept)

- `f5xc_namespace` - Namespace management

### Planned (v0.2.0+)

- `f5xc_http_loadbalancer` - HTTP load balancer
- `f5xc_origin_pool` - Backend server pools
- `f5xc_healthcheck` - Health monitoring
- `f5xc_certificate` - TLS certificate management
- `f5xc_api_credential` - API authentication
- `f5xc_cloud_credentials` - Cloud provider integration
- `f5xc_network_connector` - Network connectivity
- `f5xc_app_firewall` - Web application firewall
- `f5xc_rate_limiter` - Rate limiting policies
- ...and 250+ more resources

See [roadmap](docs/ROADMAP.md) for full resource coverage plan.

## 🏗️ Development

### Requirements

- [Go](https://golang.org/doc/install) >= 1.22
- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [F5 Distributed Cloud Account](https://www.f5.com/cloud)

### Building from Source

```bash
# Clone repository
git clone https://github.com/f5xc/terraform-provider-f5xc.git
cd terraform-provider-f5xc

# Install dependencies
go mod download

# Build provider
go build -o terraform-provider-f5xc

# Install locally for testing
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/robinmordasiewicz/f5xc/0.1.0/darwin_arm64
cp terraform-provider-f5xc ~/.terraform.d/plugins/registry.terraform.io/robinmordasiewicz/f5xc/0.1.0/darwin_arm64/
```

### Running Tests

```bash
# Unit tests
go test ./...

# Acceptance tests (requires F5XC_API_TOKEN)
TF_ACC=1 go test ./... -v -timeout 120m
```

### Generating Documentation

```bash
go generate ./...
```

## 🤝 Contributing

We welcome contributions from the community! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### How to Contribute

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Code of Conduct

This project adheres to the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md).

## 📋 Roadmap

See [ROADMAP.md](docs/ROADMAP.md) for detailed development timeline and feature plans.

### Short-term (v0.1-0.2)

- ✅ Provider scaffold and namespace resource
- 🔄 HTTP load balancer, origin pools, health checks
- 🔄 Certificate and credential management

### Medium-term (v0.3-0.5)

- Network connectivity and security resources
- API security and bot defense
- Comprehensive testing and documentation

### Long-term (v1.0+)

- Full feature parity with proprietary provider
- Linux Foundation / CNCF governance
- Production adoption and community growth

## 📜 License

This project is licensed under the Mozilla Public License 2.0 - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- **HashiCorp**: For the excellent Terraform Plugin Framework
- **F5**: For providing public OpenAPI documentation
- **OpenTofu Community**: For demonstrating community-driven infrastructure-as-code
- **Contributors**: Thank you to all contributors who help make this project better!

## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/f5xc/terraform-provider-f5xc/issues)
- **Discussions**: [GitHub Discussions](https://github.com/f5xc/terraform-provider-f5xc/discussions)
- **Documentation**: [Official Docs](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest/docs)

## ⚖️ Legal Notice

This is an independent, community-driven project built from public F5 API documentation. It is not affiliated with, endorsed by, or supported by F5 Networks, Inc.

The provider implements the F5 Distributed Cloud API as documented in F5's public documentation. All API interactions follow F5's published specifications and terms of service.

---

> Made with ❤️ by the community, for the community
