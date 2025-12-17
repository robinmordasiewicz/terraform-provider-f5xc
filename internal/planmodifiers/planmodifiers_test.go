// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package planmodifiers

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestWarnIfUnknown(t *testing.T) {
	ctx := context.Background()
	modifier := WarnIfUnknown("test_attribute")

	t.Run("description", func(t *testing.T) {
		desc := modifier.Description(ctx)
		if desc == "" {
			t.Error("Description() should return non-empty string")
		}
	})

	t.Run("markdown description", func(t *testing.T) {
		mdDesc := modifier.MarkdownDescription(ctx)
		if mdDesc == "" {
			t.Error("MarkdownDescription() should return non-empty string")
		}
	})

	t.Run("warns when unknown", func(t *testing.T) {
		req := planmodifier.StringRequest{
			PlanValue: types.StringUnknown(),
		}
		resp := &planmodifier.StringResponse{}

		modifier.PlanModifyString(ctx, req, resp)

		if len(resp.Diagnostics.Warnings()) != 1 {
			t.Errorf("Expected 1 warning, got %d", len(resp.Diagnostics.Warnings()))
		}
	})

	t.Run("no warning when known", func(t *testing.T) {
		req := planmodifier.StringRequest{
			PlanValue: types.StringValue("known-value"),
		}
		resp := &planmodifier.StringResponse{}

		modifier.PlanModifyString(ctx, req, resp)

		if len(resp.Diagnostics.Warnings()) != 0 {
			t.Errorf("Expected 0 warnings, got %d", len(resp.Diagnostics.Warnings()))
		}
	})
}

func TestWarnIfUnknownInt64(t *testing.T) {
	ctx := context.Background()
	modifier := WarnIfUnknownInt64("test_int")

	t.Run("description", func(t *testing.T) {
		desc := modifier.Description(ctx)
		if desc == "" {
			t.Error("Description() should return non-empty string")
		}
	})

	t.Run("warns when unknown", func(t *testing.T) {
		req := planmodifier.Int64Request{
			PlanValue: types.Int64Unknown(),
		}
		resp := &planmodifier.Int64Response{}

		modifier.PlanModifyInt64(ctx, req, resp)

		if len(resp.Diagnostics.Warnings()) != 1 {
			t.Errorf("Expected 1 warning, got %d", len(resp.Diagnostics.Warnings()))
		}
	})

	t.Run("no warning when known", func(t *testing.T) {
		req := planmodifier.Int64Request{
			PlanValue: types.Int64Value(42),
		}
		resp := &planmodifier.Int64Response{}

		modifier.PlanModifyInt64(ctx, req, resp)

		if len(resp.Diagnostics.Warnings()) != 0 {
			t.Errorf("Expected 0 warnings, got %d", len(resp.Diagnostics.Warnings()))
		}
	})
}

func TestWarnIfUnknownBool(t *testing.T) {
	ctx := context.Background()
	modifier := WarnIfUnknownBool("test_bool")

	t.Run("description", func(t *testing.T) {
		desc := modifier.Description(ctx)
		if desc == "" {
			t.Error("Description() should return non-empty string")
		}
	})

	t.Run("warns when unknown", func(t *testing.T) {
		req := planmodifier.BoolRequest{
			PlanValue: types.BoolUnknown(),
		}
		resp := &planmodifier.BoolResponse{}

		modifier.PlanModifyBool(ctx, req, resp)

		if len(resp.Diagnostics.Warnings()) != 1 {
			t.Errorf("Expected 1 warning, got %d", len(resp.Diagnostics.Warnings()))
		}
	})

	t.Run("no warning when known", func(t *testing.T) {
		req := planmodifier.BoolRequest{
			PlanValue: types.BoolValue(true),
		}
		resp := &planmodifier.BoolResponse{}

		modifier.PlanModifyBool(ctx, req, resp)

		if len(resp.Diagnostics.Warnings()) != 0 {
			t.Errorf("Expected 0 warnings, got %d", len(resp.Diagnostics.Warnings()))
		}
	})
}

func TestRequiresReplaceIfConfigured(t *testing.T) {
	ctx := context.Background()
	modifier := RequiresReplaceIfConfigured()

	t.Run("description", func(t *testing.T) {
		desc := modifier.Description(ctx)
		if desc == "" {
			t.Error("Description() should return non-empty string")
		}
	})

	t.Run("no replace when config is null", func(t *testing.T) {
		schema := tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"test": tftypes.String,
			},
		}
		stateValue := tftypes.NewValue(schema, map[string]tftypes.Value{
			"test": tftypes.NewValue(tftypes.String, "old-value"),
		})

		req := planmodifier.StringRequest{
			StateValue:  types.StringValue("old-value"),
			PlanValue:   types.StringValue("new-value"),
			ConfigValue: types.StringNull(),
			State: tfsdk.State{
				Raw: stateValue,
			},
		}
		resp := &planmodifier.StringResponse{}

		modifier.PlanModifyString(ctx, req, resp)

		if resp.RequiresReplace {
			t.Error("Should not require replace when config is null")
		}
	})

	t.Run("no replace when state is null (create)", func(t *testing.T) {
		req := planmodifier.StringRequest{
			StateValue:  types.StringNull(),
			PlanValue:   types.StringValue("new-value"),
			ConfigValue: types.StringValue("new-value"),
			State: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{}, nil),
			},
		}
		resp := &planmodifier.StringResponse{}

		modifier.PlanModifyString(ctx, req, resp)

		if resp.RequiresReplace {
			t.Error("Should not require replace during create")
		}
	})

	t.Run("no replace when values equal", func(t *testing.T) {
		schema := tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"test": tftypes.String,
			},
		}
		stateValue := tftypes.NewValue(schema, map[string]tftypes.Value{
			"test": tftypes.NewValue(tftypes.String, "same-value"),
		})

		req := planmodifier.StringRequest{
			StateValue:  types.StringValue("same-value"),
			PlanValue:   types.StringValue("same-value"),
			ConfigValue: types.StringValue("same-value"),
			State: tfsdk.State{
				Raw: stateValue,
			},
		}
		resp := &planmodifier.StringResponse{}

		modifier.PlanModifyString(ctx, req, resp)

		if resp.RequiresReplace {
			t.Error("Should not require replace when values are equal")
		}
	})

	t.Run("requires replace when configured value changes", func(t *testing.T) {
		schema := tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"test": tftypes.String,
			},
		}
		stateValue := tftypes.NewValue(schema, map[string]tftypes.Value{
			"test": tftypes.NewValue(tftypes.String, "old-value"),
		})

		req := planmodifier.StringRequest{
			StateValue:  types.StringValue("old-value"),
			PlanValue:   types.StringValue("new-value"),
			ConfigValue: types.StringValue("new-value"),
			State: tfsdk.State{
				Raw: stateValue,
			},
		}
		resp := &planmodifier.StringResponse{}

		modifier.PlanModifyString(ctx, req, resp)

		if !resp.RequiresReplace {
			t.Error("Should require replace when configured value changes")
		}
	})
}

func TestDefaultValue(t *testing.T) {
	ctx := context.Background()
	defaultVal := "default-value"
	modifier := DefaultValue(defaultVal)

	t.Run("description", func(t *testing.T) {
		desc := modifier.Description(ctx)
		if desc == "" {
			t.Error("Description() should return non-empty string")
		}
	})

	t.Run("sets default when config is null", func(t *testing.T) {
		req := planmodifier.StringRequest{
			ConfigValue: types.StringNull(),
			StateValue:  types.StringNull(),
			PlanValue:   types.StringNull(),
			State: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{}, nil),
			},
		}
		resp := &planmodifier.StringResponse{
			PlanValue: types.StringNull(),
		}

		modifier.PlanModifyString(ctx, req, resp)

		if resp.PlanValue.ValueString() != defaultVal {
			t.Errorf("Expected default value %q, got %q", defaultVal, resp.PlanValue.ValueString())
		}
	})

	t.Run("does not override configured value", func(t *testing.T) {
		configuredVal := "configured-value"
		req := planmodifier.StringRequest{
			ConfigValue: types.StringValue(configuredVal),
			StateValue:  types.StringNull(),
			PlanValue:   types.StringValue(configuredVal),
			State: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{}, nil),
			},
		}
		resp := &planmodifier.StringResponse{
			PlanValue: types.StringValue(configuredVal),
		}

		modifier.PlanModifyString(ctx, req, resp)

		if resp.PlanValue.ValueString() != configuredVal {
			t.Errorf("Should not override configured value, got %q", resp.PlanValue.ValueString())
		}
	})

	t.Run("preserves state value on update", func(t *testing.T) {
		schema := tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"test": tftypes.String,
			},
		}
		stateValue := tftypes.NewValue(schema, map[string]tftypes.Value{
			"test": tftypes.NewValue(tftypes.String, "existing-value"),
		})

		req := planmodifier.StringRequest{
			ConfigValue: types.StringNull(),
			StateValue:  types.StringValue("existing-value"),
			PlanValue:   types.StringNull(),
			State: tfsdk.State{
				Raw: stateValue,
			},
		}
		resp := &planmodifier.StringResponse{
			PlanValue: types.StringNull(),
		}

		modifier.PlanModifyString(ctx, req, resp)

		// Should preserve state value, not override with default
		if resp.PlanValue.ValueString() == defaultVal {
			t.Error("Should not override existing state value with default")
		}
	})
}

func TestPreserveUnknownOnCreate(t *testing.T) {
	ctx := context.Background()
	modifier := PreserveUnknownOnCreate()

	t.Run("description", func(t *testing.T) {
		desc := modifier.Description(ctx)
		if desc == "" {
			t.Error("Description() should return non-empty string")
		}
	})

	t.Run("preserves unknown during create", func(t *testing.T) {
		req := planmodifier.StringRequest{
			StateValue: types.StringNull(),
			PlanValue:  types.StringUnknown(),
			State: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{}, nil),
			},
		}
		resp := &planmodifier.StringResponse{
			PlanValue: types.StringUnknown(),
		}

		modifier.PlanModifyString(ctx, req, resp)

		if !resp.PlanValue.IsUnknown() {
			t.Error("Should preserve unknown value during create")
		}
	})

	t.Run("uses state value during update when plan is unknown", func(t *testing.T) {
		schema := tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"test": tftypes.String,
			},
		}
		stateValue := tftypes.NewValue(schema, map[string]tftypes.Value{
			"test": tftypes.NewValue(tftypes.String, "state-value"),
		})

		req := planmodifier.StringRequest{
			StateValue: types.StringValue("state-value"),
			PlanValue:  types.StringUnknown(),
			State: tfsdk.State{
				Raw: stateValue,
			},
		}
		resp := &planmodifier.StringResponse{
			PlanValue: types.StringUnknown(),
		}

		modifier.PlanModifyString(ctx, req, resp)

		if resp.PlanValue.ValueString() != "state-value" {
			t.Errorf("Should use state value during update, got %q", resp.PlanValue.ValueString())
		}
	})
}

func TestModifierDescriptions(t *testing.T) {
	ctx := context.Background()

	modifiers := []struct {
		name     string
		modifier interface {
			Description(context.Context) string
			MarkdownDescription(context.Context) string
		}
	}{
		{"WarnIfUnknown", WarnIfUnknown("attr")},
		{"WarnIfUnknownInt64", WarnIfUnknownInt64("attr")},
		{"WarnIfUnknownBool", WarnIfUnknownBool("attr")},
		{"RequiresReplaceIfConfigured", RequiresReplaceIfConfigured()},
		{"DefaultValue", DefaultValue("default")},
		{"PreserveUnknownOnCreate", PreserveUnknownOnCreate()},
	}

	for _, m := range modifiers {
		t.Run(m.name, func(t *testing.T) {
			desc := m.modifier.Description(ctx)
			mdDesc := m.modifier.MarkdownDescription(ctx)

			if desc == "" {
				t.Errorf("%s.Description() returned empty string", m.name)
			}
			if mdDesc == "" {
				t.Errorf("%s.MarkdownDescription() returned empty string", m.name)
			}
		})
	}
}
