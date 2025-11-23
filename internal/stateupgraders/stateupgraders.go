// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

// Package stateupgraders provides state upgrade functionality for F5 XC Terraform resources.
// StateUpgraders handle schema evolution when resource schemas change between provider versions.
package stateupgraders

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SchemaVersion tracks the current schema version for resources.
// Increment this when making breaking schema changes.
const SchemaVersion int64 = 1

// UpgradeResourceState provides the standard implementation of StateUpgrader
// for F5 XC resources that tracks schema versions.
type UpgradeResourceState struct {
	// CurrentVersion is the version of the current schema
	CurrentVersion int64

	// UpgradeFuncs maps version numbers to upgrade functions
	UpgradeFuncs map[int64]UpgradeFunc
}

// UpgradeFunc defines the signature for state upgrade functions
type UpgradeFunc func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse)

// NewUpgradeResourceState creates a new state upgrader with the specified version
func NewUpgradeResourceState(currentVersion int64) *UpgradeResourceState {
	return &UpgradeResourceState{
		CurrentVersion: currentVersion,
		UpgradeFuncs:   make(map[int64]UpgradeFunc),
	}
}

// RegisterUpgrade registers an upgrade function for a specific version
func (u *UpgradeResourceState) RegisterUpgrade(fromVersion int64, fn UpgradeFunc) {
	u.UpgradeFuncs[fromVersion] = fn
}

// StateUpgraders returns the map of version to StateUpgrader for the resource
func (u *UpgradeResourceState) StateUpgraders() map[int64]resource.StateUpgrader {
	upgraders := make(map[int64]resource.StateUpgrader)

	for version, fn := range u.UpgradeFuncs {
		// Capture fn in closure
		upgradeFn := fn
		upgraders[version] = resource.StateUpgrader{
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				upgradeFn(ctx, req, resp)
			},
		}
	}

	return upgraders
}

// BaseResourceUpgrader provides common state upgrade logic for F5 XC resources
type BaseResourceUpgrader struct {
	ResourceType string
	Upgrader     *UpgradeResourceState
}

// NewBaseResourceUpgrader creates a new base upgrader for a resource type
func NewBaseResourceUpgrader(resourceType string) *BaseResourceUpgrader {
	return &BaseResourceUpgrader{
		ResourceType: resourceType,
		Upgrader:     NewUpgradeResourceState(SchemaVersion),
	}
}

// UpgradeStateV0ToV1 provides a template for upgrading from version 0 to 1
// This is a common upgrade path when adding new required fields or restructuring
func UpgradeStateV0ToV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	// Example: Get the raw JSON state
	rawStateJSON := req.RawState.JSON

	// Parse the JSON to a map for manipulation
	var rawState map[string]interface{}
	if len(rawStateJSON) > 0 {
		if err := json.Unmarshal(rawStateJSON, &rawState); err != nil {
			resp.Diagnostics.AddError(
				"Error Upgrading State",
				"Could not unmarshal prior state: "+err.Error(),
			)
			return
		}
	}

	// Add default values for new fields
	// Example: rawState["new_field"] = "default_value"

	// You would typically create a new state struct and set it:
	// newState := MyResourceModelV1{...}
	// resp.State.Set(ctx, newState)
}

// CommonV0Schema returns the schema for version 0 of common resource fields
// Use this when defining prior schema versions for state upgraders
func CommonV0Schema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
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
	}
}

// MigrateLabelsFormat handles migration of labels from old format to new format
func MigrateLabelsFormat(oldLabels map[string]interface{}) map[string]string {
	newLabels := make(map[string]string)
	for k, v := range oldLabels {
		if str, ok := v.(string); ok {
			newLabels[k] = str
		}
	}
	return newLabels
}

// AddDefaultIfMissing adds a default value to the raw state if the key doesn't exist
func AddDefaultIfMissing(rawState map[string]interface{}, key string, defaultValue interface{}) {
	if _, exists := rawState[key]; !exists {
		rawState[key] = defaultValue
	}
}

// RemoveDeprecatedField removes a field from the raw state during upgrade
func RemoveDeprecatedField(rawState map[string]interface{}, key string) {
	delete(rawState, key)
}

// RenameField renames a field in the raw state during upgrade
func RenameField(rawState map[string]interface{}, oldKey, newKey string) {
	if value, exists := rawState[oldKey]; exists {
		rawState[newKey] = value
		delete(rawState, oldKey)
	}
}

// TransformField applies a transformation function to a field during upgrade
func TransformField(rawState map[string]interface{}, key string, transform func(interface{}) interface{}) {
	if value, exists := rawState[key]; exists {
		rawState[key] = transform(value)
	}
}

// Example upgrader registration for a resource:
//
// func (r *MyResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
//     return map[int64]resource.StateUpgrader{
//         0: {
//             PriorSchema: &schema.Schema{
//                 Attributes: stateupgraders.CommonV0Schema(),
//             },
//             StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
//                 // Perform the upgrade
//                 stateupgraders.UpgradeStateV0ToV1(ctx, req, resp)
//             },
//         },
//     }
// }
