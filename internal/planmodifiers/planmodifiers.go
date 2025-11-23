// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

// Package planmodifiers provides custom plan modifiers for F5 XC Terraform resources
// following Terraform Plugin Framework best practices.
package planmodifiers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// WarnIfUnknown returns a plan modifier that adds a warning when a value is unknown during plan.
// This helps users understand that certain values can only be determined after apply.
func WarnIfUnknown(attributeName string) planmodifier.String {
	return &warnIfUnknownModifier{
		attributeName: attributeName,
	}
}

type warnIfUnknownModifier struct {
	attributeName string
}

func (m *warnIfUnknownModifier) Description(ctx context.Context) string {
	return fmt.Sprintf("Warns when %s is unknown during plan", m.attributeName)
}

func (m *warnIfUnknownModifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("Warns when `%s` is unknown during plan", m.attributeName)
}

func (m *warnIfUnknownModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Only warn if the value is unknown and this is a create or update
	if req.PlanValue.IsUnknown() {
		resp.Diagnostics.AddWarning(
			fmt.Sprintf("Unknown Value for %s", m.attributeName),
			fmt.Sprintf("The value of %s will be determined after apply. This may affect dependent resources.", m.attributeName),
		)
	}
}

// WarnIfUnknownInt64 returns a plan modifier for int64 attributes
func WarnIfUnknownInt64(attributeName string) planmodifier.Int64 {
	return &warnIfUnknownInt64Modifier{
		attributeName: attributeName,
	}
}

type warnIfUnknownInt64Modifier struct {
	attributeName string
}

func (m *warnIfUnknownInt64Modifier) Description(ctx context.Context) string {
	return fmt.Sprintf("Warns when %s is unknown during plan", m.attributeName)
}

func (m *warnIfUnknownInt64Modifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("Warns when `%s` is unknown during plan", m.attributeName)
}

func (m *warnIfUnknownInt64Modifier) PlanModifyInt64(ctx context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	if req.PlanValue.IsUnknown() {
		resp.Diagnostics.AddWarning(
			fmt.Sprintf("Unknown Value for %s", m.attributeName),
			fmt.Sprintf("The value of %s will be determined after apply.", m.attributeName),
		)
	}
}

// WarnIfUnknownBool returns a plan modifier for bool attributes
func WarnIfUnknownBool(attributeName string) planmodifier.Bool {
	return &warnIfUnknownBoolModifier{
		attributeName: attributeName,
	}
}

type warnIfUnknownBoolModifier struct {
	attributeName string
}

func (m *warnIfUnknownBoolModifier) Description(ctx context.Context) string {
	return fmt.Sprintf("Warns when %s is unknown during plan", m.attributeName)
}

func (m *warnIfUnknownBoolModifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("Warns when `%s` is unknown during plan", m.attributeName)
}

func (m *warnIfUnknownBoolModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if req.PlanValue.IsUnknown() {
		resp.Diagnostics.AddWarning(
			fmt.Sprintf("Unknown Value for %s", m.attributeName),
			fmt.Sprintf("The value of %s will be determined after apply.", m.attributeName),
		)
	}
}

// RequiresReplaceIfConfigured returns a plan modifier that requires replacement
// only if the attribute is configured (not null)
func RequiresReplaceIfConfigured() planmodifier.String {
	return &requiresReplaceIfConfiguredModifier{}
}

type requiresReplaceIfConfiguredModifier struct{}

func (m *requiresReplaceIfConfiguredModifier) Description(ctx context.Context) string {
	return "Requires replacement if the value is configured and changed"
}

func (m *requiresReplaceIfConfiguredModifier) MarkdownDescription(ctx context.Context) string {
	return "Requires replacement if the value is configured and changed"
}

func (m *requiresReplaceIfConfiguredModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Don't require replace if:
	// - resource is being created
	// - config value is null (attribute not configured)
	// - values are equal
	if req.State.Raw.IsNull() || req.ConfigValue.IsNull() {
		return
	}

	if req.StateValue.Equal(req.PlanValue) {
		return
	}

	resp.RequiresReplace = true
}

// DefaultValue returns a plan modifier that sets a default value if not configured
func DefaultValue(defaultValue string) planmodifier.String {
	return &defaultValueModifier{
		defaultValue: defaultValue,
	}
}

type defaultValueModifier struct {
	defaultValue string
}

func (m *defaultValueModifier) Description(ctx context.Context) string {
	return fmt.Sprintf("Defaults to %s if not configured", m.defaultValue)
}

func (m *defaultValueModifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("Defaults to `%s` if not configured", m.defaultValue)
}

func (m *defaultValueModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Only set default if configuration value is null
	if !req.ConfigValue.IsNull() {
		return
	}

	// Don't override with default if we're updating and the state has a value
	if !req.StateValue.IsNull() && !req.State.Raw.IsNull() {
		return
	}

	resp.PlanValue = types.StringValue(m.defaultValue)
}

// PreserveUnknownOnCreate returns a plan modifier that keeps unknown values during create
// This is useful for computed values that need to be fetched from the API
func PreserveUnknownOnCreate() planmodifier.String {
	return &preserveUnknownOnCreateModifier{}
}

type preserveUnknownOnCreateModifier struct{}

func (m *preserveUnknownOnCreateModifier) Description(ctx context.Context) string {
	return "Preserves unknown value during create, uses state value during update"
}

func (m *preserveUnknownOnCreateModifier) MarkdownDescription(ctx context.Context) string {
	return "Preserves unknown value during create, uses state value during update"
}

func (m *preserveUnknownOnCreateModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// During create, keep the unknown value
	if req.State.Raw.IsNull() {
		return
	}

	// During update, use the state value if plan is unknown
	if resp.PlanValue.IsUnknown() {
		resp.PlanValue = req.StateValue
	}
}
