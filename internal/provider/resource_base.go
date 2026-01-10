// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package provider implements base resource functionality for F5 XC Terraform resources
// following HashiCorp Terraform Plugin Framework best practices.
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5xc/terraform-provider-f5xc/internal/client"
	"github.com/f5xc/terraform-provider-f5xc/internal/privatestate"
	inttimeouts "github.com/f5xc/terraform-provider-f5xc/internal/timeouts"
)

// ResourceType categorizes resources for timeout configuration
type ResourceType int

const (
	// ResourceTypeStandard is for resources with quick API operations
	ResourceTypeStandard ResourceType = iota
	// ResourceTypeLongRunning is for resources involving infrastructure provisioning
	ResourceTypeLongRunning
	// ResourceTypeCritical is for resources requiring extra time for safety
	ResourceTypeCritical
)

// SchemaVersion is the current schema version for all resources
// Increment this when making breaking schema changes across all resources
const SchemaVersion int64 = 1

// BaseResourceConfig holds configuration for base resource behavior
type BaseResourceConfig struct {
	// ResourceTypeName is the Terraform resource type name (e.g., "f5xc_namespace")
	ResourceTypeName string

	// ResourceType categorizes the resource for timeout configuration
	ResourceType ResourceType

	// SupportsTimeouts indicates if the resource supports configurable timeouts
	SupportsTimeouts bool

	// SupportsPrivateState indicates if the resource uses private state for metadata
	SupportsPrivateState bool

	// SupportsStateUpgrade indicates if the resource supports state upgrades
	SupportsStateUpgrade bool
}

// BaseResourceModel contains common fields for all F5 XC resources
type BaseResourceModel struct {
	ID          types.String   `tfsdk:"id"`
	Name        types.String   `tfsdk:"name"`
	Namespace   types.String   `tfsdk:"namespace"`
	Labels      types.Map      `tfsdk:"labels"`
	Annotations types.Map      `tfsdk:"annotations"`
	Timeouts    timeouts.Value `tfsdk:"timeouts"`
}

// GetTimeoutsBlock returns the timeouts schema block based on resource type
func GetTimeoutsBlock(ctx context.Context, rt ResourceType) schema.Block {
	switch rt {
	case ResourceTypeLongRunning:
		return inttimeouts.LongRunningBlock()
	case ResourceTypeCritical:
		return inttimeouts.Block(inttimeouts.ConfigForResourceType(inttimeouts.Critical))
	default:
		return inttimeouts.StandardBlock()
	}
}

// CommonSchemaAttributes returns the common schema attributes for all resources
func CommonSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			MarkdownDescription: "Unique identifier for the resource.",
			Computed:            true,
		},
	}
}

// ModifyPlanForUnknownValues adds warnings for unknown values during plan
// This implements the ResourceWithModifyPlan interface recommendation
func ModifyPlanForUnknownValues(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse, resourceName string) {
	// Skip if resource is being destroyed
	if req.Plan.Raw.IsNull() {
		return
	}

	// Check if this is a create operation (no prior state)
	if req.State.Raw.IsNull() {
		// During create, computed values will be unknown - this is expected
		resp.Diagnostics.AddWarning(
			fmt.Sprintf("Creating %s Resource", resourceName),
			"Some computed values will be determined after the resource is created.",
		)
	}
}

// ModifyPlanForDestruction adds warnings for resource destruction
// This follows Terraform 1.3+ best practices for destroy plan diagnostics
func ModifyPlanForDestruction(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse, resourceName string, hasCustomDestroyBehavior bool) {
	// Check if the entire plan is null (resource is planned for destruction)
	if !req.Plan.Raw.IsNull() {
		return
	}

	if hasCustomDestroyBehavior {
		resp.Diagnostics.AddWarning(
			"Resource Destruction Considerations",
			fmt.Sprintf("Destroying the %s resource will remove it from both Terraform state and F5 Distributed Cloud. "+
				"This action cannot be undone. Ensure any dependent resources are updated accordingly.", resourceName),
		)
	}
}

// SaveAPIMetadataToPrivateState saves API metadata to private state for drift detection
func SaveAPIMetadataToPrivateState(ctx context.Context, resp interface{}, metadata *client.Metadata) diag.Diagnostics {
	psd := privatestate.NewPrivateStateData()

	if metadata.UID != "" {
		psd.SetUID(metadata.UID)
	}

	return psd.SaveToPrivateState(ctx, resp)
}

// LoadAPIMetadataFromPrivateState loads API metadata from private state
func LoadAPIMetadataFromPrivateState(ctx context.Context, req interface{}) (*privatestate.PrivateStateData, diag.Diagnostics) {
	return privatestate.LoadFromPrivateState(ctx, req)
}

// GetLabelsFromModel extracts labels from a types.Map to map[string]string
func GetLabelsFromModel(ctx context.Context, labelsMap types.Map) (map[string]string, diag.Diagnostics) {
	var diags diag.Diagnostics

	if labelsMap.IsNull() || labelsMap.IsUnknown() {
		return nil, diags
	}

	labels := make(map[string]string)
	diags.Append(labelsMap.ElementsAs(ctx, &labels, false)...)
	return labels, diags
}

// GetAnnotationsFromModel extracts annotations from a types.Map to map[string]string
func GetAnnotationsFromModel(ctx context.Context, annotationsMap types.Map) (map[string]string, diag.Diagnostics) {
	var diags diag.Diagnostics

	if annotationsMap.IsNull() || annotationsMap.IsUnknown() {
		return nil, diags
	}

	annotations := make(map[string]string)
	diags.Append(annotationsMap.ElementsAs(ctx, &annotations, false)...)
	return annotations, diags
}

// SetLabelsToModel converts map[string]string to types.Map for the model
func SetLabelsToModel(ctx context.Context, labels map[string]string) (types.Map, diag.Diagnostics) {
	if len(labels) == 0 {
		return types.MapNull(types.StringType), nil
	}
	return types.MapValueFrom(ctx, types.StringType, labels)
}

// SetAnnotationsToModel converts map[string]string to types.Map for the model
func SetAnnotationsToModel(ctx context.Context, annotations map[string]string) (types.Map, diag.Diagnostics) {
	if len(annotations) == 0 {
		return types.MapNull(types.StringType), nil
	}
	return types.MapValueFrom(ctx, types.StringType, annotations)
}

// StateUpgraderV0ToV1 provides a generic state upgrader from v0 to v1
// This handles the common case of adding timeouts block to existing resources
func StateUpgraderV0ToV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	// The v0 to v1 upgrade primarily adds the timeouts block
	// Since timeouts are optional with defaults, no data migration is needed
	// The framework will handle null timeouts values automatically
}

// CommonV0Schema returns the schema for version 0 (without timeouts)
func CommonV0Schema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"namespace": schema.StringAttribute{
				Required: true,
			},
			"labels": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"annotations": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

// ValidateResourceConfig validates common resource configuration
func ValidateResourceConfig(ctx context.Context, config interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// Add common validation logic here
	return diags
}

// IsLongRunningResource returns true if the resource type is long-running
func IsLongRunningResource(resourceType string) bool {
	return inttimeouts.IsLongRunning(resourceType)
}

// GetResourceTypeFromName determines the resource type from the resource name
func GetResourceTypeFromName(resourceName string) ResourceType {
	if inttimeouts.IsLongRunning(resourceName) {
		return ResourceTypeLongRunning
	}
	return ResourceTypeStandard
}
