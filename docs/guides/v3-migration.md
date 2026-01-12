---
page_title: "Guide: Migrating to Provider v3.0.0"
subcategory: "Guides"
description: |-
  Migration guide for upgrading from F5XC Terraform provider v2.x to v3.0.0,
  which introduces enriched API specifications and improved documentation.
---

# Migrating to Provider v3.0.0

This guide covers the migration from F5XC Terraform provider v2.x to v3.0.0. Version 3.0.0 introduces significant improvements to the provider's API specification foundation, bringing enhanced documentation, better categorization, and improved example generation.

## Overview

Provider v3.0.0 represents a major upgrade to the underlying API specification infrastructure:

| Aspect | v2.x | v3.0.0 |
| ------ | ---- | ------ |
| **API Specs** | 269 individual files | 38 domain-organized files |
| **Categories** | Pattern-based | API-defined (`x-f5xc-category`) |
| **Subscription Info** | Static metadata | API-defined (`x-f5xc-requires-tier`) |
| **Examples** | Hardcoded | API-enriched (`x-f5xc-example`) |
| **Descriptions** | Schema-only | Multi-level descriptions |

## What's New

### Improved Resource Categorization

Resources are now categorized based on official F5 API metadata rather than pattern matching. This provides more accurate Terraform Registry navigation:

- **Security**: WAF, DDoS protection, certificates, blindfold encryption
- **Networking**: DNS, CDN, network policies, rate limiting
- **Infrastructure**: Sites, service mesh, cloud infrastructure
- **Platform**: Tenants, identity, authentication, billing
- **Operations**: Observability, statistics, support
- **AI Services**: AI-powered features (preview)

### Enhanced Documentation

Documentation now includes:

- **Subscription badges**: Clear indication when Advanced subscription is required
- **Preview notices**: Warning when features are in preview/beta status
- **Improved descriptions**: Enriched descriptions from API metadata
- **Better examples**: API-sourced example values for attributes

### Specification Foundation

The provider now uses enriched API specifications that include:

- `x-f5xc-category`: Official resource categorization
- `x-f5xc-requires-tier`: Subscription tier requirements
- `x-f5xc-complexity`: Resource complexity indicators
- `x-f5xc-is-preview`: Preview/beta feature flags
- `x-f5xc-example`: Enriched example values

## Migration Steps

### Step 1: Update Provider Version

Update your Terraform configuration to use v3.0.0:

```hcl
terraform {
  required_providers {
    f5xc = {
      source  = "robinmordasiewicz/f5xc"
      version = "~> 3.0"
    }
  }
}
```

### Step 2: Run Terraform Plan

After updating, run `terraform plan` to verify your configuration:

```bash
terraform init -upgrade
terraform plan
```

### Step 3: Review Any Warnings

Version 3.0.0 may display new warnings about:

- **Subscription requirements**: Resources now clearly indicate when Advanced subscription is needed
- **Preview features**: Some features may be marked as preview with potential for change
- **Deprecations**: Any deprecated attributes will be clearly documented

## Breaking Changes

Version 3.0.0 maintains backward compatibility with v2.x configurations. However, be aware of:

### Documentation Changes

- Resource subcategories in the Terraform Registry may change due to improved categorization
- Example configurations may be updated to reflect API-enriched values

### Potential State Changes

No state migration is required. Your existing Terraform state files will continue to work without modification.

## Compatibility Matrix

| Terraform Version | Provider v3.0.0 |
| ----------------- | --------------- |
| 1.8.0+            | Supported       |
| 1.7.x             | Supported       |
| 1.6.x             | Supported       |
| 1.5.x             | Supported       |
| < 1.5.0           | Not supported   |

## Troubleshooting

### Issue: "Subscription Required" Warning

If you see a subscription warning for a resource you're using:

1. Verify your F5 Distributed Cloud subscription tier
2. Contact F5 support if you believe the warning is incorrect
3. Check the resource documentation for subscription requirements

### Issue: Changes Detected After Upgrade

If `terraform plan` shows changes after upgrading:

1. Review the changes carefully - they may be improvements to default values
2. The changes are typically cosmetic (descriptions, examples)
3. Run `terraform apply` to update the state if changes are acceptable

### Issue: Resource Category Changed

If you have automation relying on resource subcategories:

1. Update your automation to use the new category names
2. The new categories align with official F5 API organization

## Getting Help

- **Documentation**: [Provider Documentation](https://registry.terraform.io/providers/robinmordasiewicz/f5xc/latest/docs)
- **Issues**: [GitHub Issues](https://github.com/robinmordasiewicz/terraform-provider-f5xc/issues)
- **F5 Support**: [F5 Distributed Cloud Console](https://console.ves.volterra.io)

## Related Guides

- [Authentication Guide](authentication.md)
- [HTTP Load Balancer Guide](http-loadbalancer.md)
- [Blindfold Encryption Guide](blindfold.md)
