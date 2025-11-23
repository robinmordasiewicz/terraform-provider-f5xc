# Terraform Plugin Framework Best Practices Implementation

This document describes the HashiCorp Terraform Plugin Framework best practices implemented in this provider.

## Overview

The F5 XC Terraform Provider follows HashiCorp's official best practices for provider development. The `namespace_resource.go` serves as the reference implementation demonstrating all patterns.

## Implemented Best Practices

### 1. Configurable Timeouts

**Location:** `internal/timeouts/timeouts.go`

Configurable timeouts allow practitioners to customize CRUD operation durations for different environments and API latencies.

```go
// In resource model
Timeouts timeouts.Value `tfsdk:"timeouts"`

// In schema
Blocks: map[string]schema.Block{
    "timeouts": timeouts.Block(ctx, timeouts.Opts{
        Create: true,
        Read:   true,
        Update: true,
        Delete: true,
    }),
},

// In CRUD methods
createTimeout, diags := data.Timeouts.Create(ctx, inttimeouts.DefaultCreate)
ctx, cancel := context.WithTimeout(ctx, createTimeout)
defer cancel()
```

**Default Timeouts:**
- Standard resources: 10 minutes (create/update/delete), 5 minutes (read)
- Long-running resources (sites, clusters): 30 minutes (create/delete)

### 2. Custom Validators

**Location:** `internal/validators/validators.go`

Custom validators provide early feedback on configuration errors during `terraform validate` and `terraform plan`.

**Available Validators:**
- `NameValidator()` - F5 XC resource names (1-64 chars, lowercase, hyphens)
- `NamespaceValidator()` - Namespace names (includes "system" special case)
- `DomainValidator()` - Domain names with wildcard support
- `LabelKeyValidator()` - Kubernetes-style label keys
- `PortValidator()` - Port numbers (1-65535)
- `NonEmptyStringValidator()` - Non-empty string values

```go
Validators: []validator.String{
    validators.NameValidator(),
},
```

### 3. ModifyPlan for Unknown Value Warnings

**Interface:** `resource.ResourceWithModifyPlan`

Plan modification provides plan-time warnings for:
- Unknown values that will be computed after apply
- Resource destruction notifications (Terraform 1.3+)

```go
func (r *Resource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
    // Destruction warning
    if req.Plan.Raw.IsNull() {
        resp.Diagnostics.AddWarning(
            "Resource Destruction",
            "This will permanently delete the resource from F5 Distributed Cloud.",
        )
        return
    }

    // Unknown value warning
    if req.State.Raw.IsNull() && plan.Name.IsUnknown() {
        resp.Diagnostics.AddWarning(
            "Unknown Name",
            "The name is not yet known. This may affect dependent resources.",
        )
    }
}
```

### 4. StateUpgraders for Schema Versioning

**Interface:** `resource.ResourceWithUpgradeState`

State upgraders handle schema migrations between provider versions, ensuring backward compatibility.

```go
// Schema version constant
const resourceSchemaVersion int64 = 1

// In Schema()
resp.Schema = schema.Schema{
    Version: resourceSchemaVersion,
    // ...
}

// UpgradeState implementation
func (r *Resource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
    return map[int64]resource.StateUpgrader{
        0: {
            PriorSchema: &schema.Schema{...},
            StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
                // Migrate from v0 to current version
            },
        },
    }
}
```

### 5. Private State for API Metadata

**Location:** `internal/privatestate/privatestate.go`

Private state stores API metadata (UID, resource version, etc.) for drift detection without exposing it in Terraform configuration.

```go
// Save metadata after create/update
psd := privatestate.NewPrivateStateData()
psd.SetUID(created.Metadata.UID)
resp.Diagnostics.Append(psd.SaveToPrivateState(ctx, resp)...)

// Load and check for drift in read
psd, _ := privatestate.LoadFromPrivateState(ctx, &req)
if psd.Metadata.UID != "" && apiResource.Metadata.UID != psd.Metadata.UID {
    resp.Diagnostics.AddWarning("Resource Drift Detected", "...")
}
```

### 6. Plan Modifiers

**Location:** `internal/planmodifiers/planmodifiers.go`

Custom plan modifiers for attribute-level behavior:

- `WarnIfUnknown(attributeName)` - Warns when value is unknown during plan
- `RequiresReplaceIfConfigured()` - Requires replacement only if configured
- `DefaultValue(value)` - Sets default value if not configured
- `PreserveUnknownOnCreate()` - Preserves unknown during create, uses state during update

**Built-in modifiers used:**
- `stringplanmodifier.RequiresReplace()` - For immutable fields (name, namespace)
- `stringplanmodifier.UseStateForUnknown()` - For computed fields (id)

### 7. Configuration Validation

**Interface:** `resource.ResourceWithValidateConfig`

Provider-side validation runs before plan/apply for early error detection.

```go
func (r *Resource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
    var data ResourceModel
    resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
    if resp.Diagnostics.HasError() {
        return
    }
    // Custom validation logic
}
```

## File Structure

```
internal/
├── client/           # API client and type definitions
├── planmodifiers/    # Custom plan modifiers
│   └── planmodifiers.go
├── privatestate/     # Private state management
│   └── privatestate.go
├── provider/         # Resources and data sources
│   ├── resource_base.go      # Base resource helpers
│   ├── namespace_resource.go # Reference implementation
│   └── *_resource.go         # Other resources
├── stateupgraders/   # State upgrade utilities
│   └── stateupgraders.go
├── timeouts/         # Timeout configuration
│   └── timeouts.go
└── validators/       # Custom validators
    └── validators.go
```

## Interface Compliance Checklist

Every resource should implement these interfaces:

```go
var (
    _ resource.Resource                   = &MyResource{}
    _ resource.ResourceWithConfigure      = &MyResource{}
    _ resource.ResourceWithImportState    = &MyResource{}
    _ resource.ResourceWithModifyPlan     = &MyResource{}
    _ resource.ResourceWithUpgradeState   = &MyResource{}
    _ resource.ResourceWithValidateConfig = &MyResource{}
)
```

## Additional Recommendations

### Security
- Never store credentials or sensitive information in state
- Use `Sensitive: true` for confidential schema attributes
- Validate inputs to prevent injection attacks

### Testing
- Write acceptance tests for all resources
- Test state upgrades between versions
- Test import functionality

### Documentation
- Document all attributes with clear descriptions
- Provide examples in `examples/` directory
- Use `MarkdownDescription` for rich formatting

### Error Handling
- Return actionable error messages
- Use structured diagnostics (AddError, AddWarning)
- Include relevant context (resource name, API error details)

## References

- [HashiCorp Provider Design Principles](https://developer.hashicorp.com/terraform/plugin/best-practices/hashicorp-provider-design-principles)
- [Terraform Plugin Framework Documentation](https://developer.hashicorp.com/terraform/plugin/framework)
- [State Upgrade Mechanisms](https://developer.hashicorp.com/terraform/plugin/framework/resources/state-upgrade)
- [Plan Modification](https://developer.hashicorp.com/terraform/plugin/framework/resources/plan-modification)
- [Private State](https://developer.hashicorp.com/terraform/plugin/framework/resources/private-state)
- [Configurable Timeouts](https://developer.hashicorp.com/terraform/plugin/framework/resources/timeouts)
