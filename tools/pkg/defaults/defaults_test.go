// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package defaults

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFormatDefaultValue(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"nil", nil, "null"},
		{"true", true, "true"},
		{"false", false, "false"},
		{"integer", float64(0), "0"},
		{"float", float64(3.14), "3.14"},
		{"negative int", float64(-5), "-5"},
		{"empty string", "", `""`},
		{"string", "hello", `"hello"`},
		{"empty array", []interface{}{}, "[]"},
		{"empty map", map[string]interface{}{}, "{}"},
		{"marker block", map[string]interface{}{"_marker": "empty_block"}, "{}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatDefaultValue(tt.input)
			if result != tt.expected {
				t.Errorf("FormatDefaultValue(%v) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFieldPathToTerraform(t *testing.T) {
	tests := []struct {
		jsonPath string
		expected string
	}{
		{"spec.jitter_percent", "jitter_percent"},
		{"spec.http_health_check.use_http2", "http_health_check.use_http2"},
		{"jitter_percent", "jitter_percent"}, // no spec prefix
	}

	for _, tt := range tests {
		t.Run(tt.jsonPath, func(t *testing.T) {
			result := FieldPathToTerraform(tt.jsonPath)
			if result != tt.expected {
				t.Errorf("FieldPathToTerraform(%q) = %q, want %q", tt.jsonPath, result, tt.expected)
			}
		})
	}
}

func TestTerraformToFieldPath(t *testing.T) {
	tests := []struct {
		tfPath   string
		expected string
	}{
		{"jitter_percent", "spec.jitter_percent"},
		{"http_health_check.use_http2", "spec.http_health_check.use_http2"},
	}

	for _, tt := range tests {
		t.Run(tt.tfPath, func(t *testing.T) {
			result := TerraformToFieldPath(tt.tfPath)
			if result != tt.expected {
				t.Errorf("TerraformToFieldPath(%q) = %q, want %q", tt.tfPath, result, tt.expected)
			}
		})
	}
}

func TestStoreLoadAndLookup(t *testing.T) {
	// Create a temporary test file with the actual api-defaults.json structure
	testData := `{
		"version": "1.0.0",
		"generated_at": "2025-01-01T00:00:00Z",
		"api_endpoint": "https://test.api.example.com",
		"total_resources": 1,
		"discovered": 1,
		"skipped": 0,
		"failed": 0,
		"resources": {
			"test_resource": {
				"resource_name": "test_resource",
				"category": "Test",
				"status": "discovered",
				"discovered_at": "2025-01-01T00:00:00Z",
				"defaults": {
					"spec.field1": {
						"path": "spec.field1",
						"default_value": 0,
						"type": "number"
					},
					"spec.field2": {
						"path": "spec.field2",
						"default_value": false,
						"type": "bool"
					},
					"spec.nested.field": {
						"path": "spec.nested.field",
						"default_value": "default",
						"type": "string"
					}
				}
			}
		}
	}`

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test-defaults.json")
	if err := os.WriteFile(tmpFile, []byte(testData), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test loading
	store := &Store{defaults: make(map[string]ResourceDefaults)}
	if err := store.LoadFromFile(tmpFile); err != nil {
		t.Fatalf("LoadFromFile failed: %v", err)
	}

	if !store.IsLoaded() {
		t.Error("IsLoaded() should return true after successful load")
	}

	// Test GetResourceDefaults
	defaults, ok := store.GetResourceDefaults("test_resource")
	if !ok {
		t.Error("GetResourceDefaults(test_resource) should return true")
	}
	if len(defaults) != 3 {
		t.Errorf("Expected 3 defaults, got %d", len(defaults))
	}

	// Test GetDefault
	val, ok := store.GetDefault("test_resource", "spec.field1")
	if !ok {
		t.Error("GetDefault should find spec.field1")
	}
	if val != float64(0) {
		t.Errorf("Expected 0, got %v", val)
	}

	// Test GetDefaultFormatted
	formatted, ok := store.GetDefaultFormatted("test_resource", "spec.field2")
	if !ok {
		t.Error("GetDefaultFormatted should find spec.field2")
	}
	if formatted != "false" {
		t.Errorf("Expected 'false', got %q", formatted)
	}

	// Test GetDefaultByTerraformPath
	val, ok = store.GetDefaultByTerraformPath("test_resource", "field1")
	if !ok {
		t.Error("GetDefaultByTerraformPath should find field1 via spec.field1")
	}
	if val != float64(0) {
		t.Errorf("Expected 0, got %v", val)
	}

	// Test non-existent resource
	_, ok = store.GetResourceDefaults("nonexistent")
	if ok {
		t.Error("GetResourceDefaults should return false for nonexistent resource")
	}

	// Test non-existent field
	_, ok = store.GetDefault("test_resource", "spec.nonexistent")
	if ok {
		t.Error("GetDefault should return false for nonexistent field")
	}

	// Test ListResources
	resources := store.ListResources()
	if len(resources) != 1 || resources[0] != "test_resource" {
		t.Errorf("ListResources unexpected: %v", resources)
	}
}

func TestStoreLoadError(t *testing.T) {
	store := &Store{defaults: make(map[string]ResourceDefaults)}
	err := store.LoadFromFile("/nonexistent/path/file.json")
	if err == nil {
		t.Error("LoadFromFile should fail for nonexistent file")
	}
	if store.IsLoaded() {
		t.Error("IsLoaded should return false after failed load")
	}
}

func TestStoreLoadInvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "invalid.json")
	if err := os.WriteFile(tmpFile, []byte("not valid json"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	store := &Store{defaults: make(map[string]ResourceDefaults)}
	err := store.LoadFromFile(tmpFile)
	if err == nil {
		t.Error("LoadFromFile should fail for invalid JSON")
	}
}
