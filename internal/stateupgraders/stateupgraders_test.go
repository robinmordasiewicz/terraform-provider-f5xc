// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package stateupgraders

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func TestSchemaVersionConstant(t *testing.T) {
	if SchemaVersion < 1 {
		t.Errorf("SchemaVersion = %d, expected >= 1", SchemaVersion)
	}
}

func TestNewUpgradeResourceState(t *testing.T) {
	upgrader := NewUpgradeResourceState(2)

	if upgrader.CurrentVersion != 2 {
		t.Errorf("CurrentVersion = %d, expected 2", upgrader.CurrentVersion)
	}
	if upgrader.UpgradeFuncs == nil {
		t.Error("UpgradeFuncs should not be nil")
	}
}

func TestRegisterUpgrade(t *testing.T) {
	upgrader := NewUpgradeResourceState(2)

	// Register an upgrade function
	upgrader.RegisterUpgrade(0, func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
		// Test function body
	})

	if _, exists := upgrader.UpgradeFuncs[0]; !exists {
		t.Error("RegisterUpgrade should add function to UpgradeFuncs map")
	}
}

func TestStateUpgraders(t *testing.T) {
	upgrader := NewUpgradeResourceState(2)

	// Register upgrade functions for versions 0 and 1
	upgrader.RegisterUpgrade(0, func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {})
	upgrader.RegisterUpgrade(1, func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {})

	stateUpgraders := upgrader.StateUpgraders()

	if len(stateUpgraders) != 2 {
		t.Errorf("StateUpgraders() returned %d upgraders, expected 2", len(stateUpgraders))
	}

	if _, exists := stateUpgraders[0]; !exists {
		t.Error("StateUpgraders() should include version 0")
	}
	if _, exists := stateUpgraders[1]; !exists {
		t.Error("StateUpgraders() should include version 1")
	}
}

func TestNewBaseResourceUpgrader(t *testing.T) {
	upgrader := NewBaseResourceUpgrader("namespace")

	if upgrader.ResourceType != "namespace" {
		t.Errorf("ResourceType = %q, expected namespace", upgrader.ResourceType)
	}
	if upgrader.Upgrader == nil {
		t.Error("Upgrader should not be nil")
	}
	if upgrader.Upgrader.CurrentVersion != SchemaVersion {
		t.Errorf("Upgrader.CurrentVersion = %d, expected %d", upgrader.Upgrader.CurrentVersion, SchemaVersion)
	}
}

func TestMigrateLabelsFormat(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]string
	}{
		{
			name:     "empty labels",
			input:    map[string]interface{}{},
			expected: map[string]string{},
		},
		{
			name: "string values",
			input: map[string]interface{}{
				"app":  "myapp",
				"env":  "prod",
				"team": "platform",
			},
			expected: map[string]string{
				"app":  "myapp",
				"env":  "prod",
				"team": "platform",
			},
		},
		{
			name: "non-string values ignored",
			input: map[string]interface{}{
				"app":   "myapp",
				"count": 42, // non-string, should be ignored
			},
			expected: map[string]string{
				"app": "myapp",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MigrateLabelsFormat(tt.input)

			if len(result) != len(tt.expected) {
				t.Errorf("MigrateLabelsFormat() returned %d labels, expected %d", len(result), len(tt.expected))
			}

			for k, v := range tt.expected {
				if result[k] != v {
					t.Errorf("MigrateLabelsFormat()[%q] = %q, expected %q", k, result[k], v)
				}
			}
		})
	}
}

func TestAddDefaultIfMissing(t *testing.T) {
	tests := []struct {
		name         string
		initial      map[string]interface{}
		key          string
		defaultValue interface{}
		expected     interface{}
	}{
		{
			name:         "add missing key",
			initial:      map[string]interface{}{},
			key:          "new_field",
			defaultValue: "default",
			expected:     "default",
		},
		{
			name: "don't override existing",
			initial: map[string]interface{}{
				"existing": "original",
			},
			key:          "existing",
			defaultValue: "default",
			expected:     "original",
		},
		{
			name:         "add int default",
			initial:      map[string]interface{}{},
			key:          "count",
			defaultValue: 10,
			expected:     10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddDefaultIfMissing(tt.initial, tt.key, tt.defaultValue)

			if tt.initial[tt.key] != tt.expected {
				t.Errorf("AddDefaultIfMissing() set %q = %v, expected %v", tt.key, tt.initial[tt.key], tt.expected)
			}
		})
	}
}

func TestRemoveDeprecatedField(t *testing.T) {
	tests := []struct {
		name    string
		initial map[string]interface{}
		key     string
		exists  bool
	}{
		{
			name: "remove existing field",
			initial: map[string]interface{}{
				"old_field": "value",
				"keep":      "this",
			},
			key:    "old_field",
			exists: false,
		},
		{
			name: "remove non-existent field (no-op)",
			initial: map[string]interface{}{
				"keep": "this",
			},
			key:    "missing",
			exists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveDeprecatedField(tt.initial, tt.key)

			_, exists := tt.initial[tt.key]
			if exists != tt.exists {
				t.Errorf("RemoveDeprecatedField() field %q exists = %v, expected %v", tt.key, exists, tt.exists)
			}
		})
	}
}

func TestRenameField(t *testing.T) {
	tests := []struct {
		name          string
		initial       map[string]interface{}
		oldKey        string
		newKey        string
		expectedValue interface{}
		oldExists     bool
		newExists     bool
	}{
		{
			name: "rename existing field",
			initial: map[string]interface{}{
				"old_name": "value",
			},
			oldKey:        "old_name",
			newKey:        "new_name",
			expectedValue: "value",
			oldExists:     false,
			newExists:     true,
		},
		{
			name: "rename non-existent field (no-op)",
			initial: map[string]interface{}{
				"other": "value",
			},
			oldKey:    "missing",
			newKey:    "new_name",
			oldExists: false,
			newExists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RenameField(tt.initial, tt.oldKey, tt.newKey)

			_, oldExists := tt.initial[tt.oldKey]
			if oldExists != tt.oldExists {
				t.Errorf("RenameField() old key %q exists = %v, expected %v", tt.oldKey, oldExists, tt.oldExists)
			}

			newVal, newExists := tt.initial[tt.newKey]
			if newExists != tt.newExists {
				t.Errorf("RenameField() new key %q exists = %v, expected %v", tt.newKey, newExists, tt.newExists)
			}

			if tt.newExists && newVal != tt.expectedValue {
				t.Errorf("RenameField() new value = %v, expected %v", newVal, tt.expectedValue)
			}
		})
	}
}

func TestTransformField(t *testing.T) {
	tests := []struct {
		name          string
		initial       map[string]interface{}
		key           string
		transform     func(interface{}) interface{}
		expectedValue interface{}
	}{
		{
			name: "transform string to uppercase",
			initial: map[string]interface{}{
				"name": "lowercase",
			},
			key: "name",
			transform: func(v interface{}) interface{} {
				if _, ok := v.(string); ok {
					return "UPPERCASE"
				}
				return v
			},
			expectedValue: "UPPERCASE",
		},
		{
			name: "transform int",
			initial: map[string]interface{}{
				"count": 5,
			},
			key: "count",
			transform: func(v interface{}) interface{} {
				if i, ok := v.(int); ok {
					return i * 2
				}
				return v
			},
			expectedValue: 10,
		},
		{
			name: "transform non-existent field (no-op)",
			initial: map[string]interface{}{
				"other": "value",
			},
			key: "missing",
			transform: func(v interface{}) interface{} {
				return "transformed"
			},
			expectedValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TransformField(tt.initial, tt.key, tt.transform)

			if tt.expectedValue != nil {
				if tt.initial[tt.key] != tt.expectedValue {
					t.Errorf("TransformField() value = %v, expected %v", tt.initial[tt.key], tt.expectedValue)
				}
			}
		})
	}
}

func TestCommonV0Schema(t *testing.T) {
	schema := CommonV0Schema()

	expectedFields := []string{"id", "name", "namespace", "labels", "annotations"}

	for _, field := range expectedFields {
		if _, exists := schema[field]; !exists {
			t.Errorf("CommonV0Schema() missing expected field %q", field)
		}
	}
}
