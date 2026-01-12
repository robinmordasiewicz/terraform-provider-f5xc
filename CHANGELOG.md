# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Breaking Changes - v3.0.0 Clean Break Release

This is a clean-break pre-release that requires recreating all Terraform-managed resources. This version uses F5 Distributed Cloud API v2 specifications and removes all backwards compatibility with earlier versions.

#### What Changed

- **API Version**: Migrated from F5XC API v1 to API v2 specifications
- **Resource Count**: Reduced from 146 to 98 resources (removed resources without v2 API specs)
- **State Management**: Removed all state upgrade infrastructure
  - No automatic migration from previous versions
  - No schema versioning
  - No private state metadata for drift detection
  - `terraform import` still supported for adopting existing F5XC resources

#### Migration Required

Since this is a pre-release project, users must:

1. **Destroy all existing resources**: Run `terraform destroy` with the previous provider version
2. **Upgrade provider**: Update to v3.0.0 in your Terraform configuration
3. **Reinitialize**: Run `terraform init -upgrade`
4. **Recreate resources**: Run `terraform apply` to create fresh resources

**Note**: This will cause downtime. Plan accordingly.

#### Alternative: Import Existing Resources

If you have existing F5 Distributed Cloud resources (not managed by Terraform), you can import them:

```bash
terraform import f5xc_namespace.example my-namespace
terraform import f5xc_http_loadbalancer.example namespace/loadbalancer-name
```

See the provider documentation for resource-specific import formats.

### Added

- 98 resources based on F5 Distributed Cloud API v2 specifications
- Provider-defined functions:
  - `provider::f5xc::blindfold` - Encrypt plaintext with F5XC blindfold encryption
  - `provider::f5xc::blindfold_file` - Encrypt file contents with F5XC blindfold encryption

### Improved

- Cleaner codebase with ~1,000+ lines of migration code removed
- Simpler maintenance without backwards compatibility infrastructure
- Faster CI without state upgrade test coverage
- Consistent resource schemas based on OpenAPI v2 specifications

### Removed

- 48 resources without F5 Distributed Cloud API v2 specifications
- State upgrade framework (`internal/stateupgraders/`)
- Private state metadata package (`internal/privatestate/`)
- Resource upgrade tool (`tools/upgrade-resources.go`)
- Schema versioning constants
- UpgradeState methods from all resources
