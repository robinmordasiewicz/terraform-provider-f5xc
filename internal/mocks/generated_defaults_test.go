// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package mocks

import (
	"testing"
)

func TestApplyDiscoveredDefaults(t *testing.T) {
	tests := []struct {
		name         string
		resourceType string
		inputSpec    map[string]interface{}
		checkFunc    func(t *testing.T, spec map[string]interface{})
	}{
		{
			name:         "policer defaults",
			resourceType: "policer",
			inputSpec: map[string]interface{}{
				"burst_size":                 10000,
				"committed_information_rate": 10000,
			},
			checkFunc: func(t *testing.T, spec map[string]interface{}) {
				if spec["policer_mode"] != "POLICER_MODE_NOT_SHARED" {
					t.Errorf("expected policer_mode=POLICER_MODE_NOT_SHARED, got %v", spec["policer_mode"])
				}
				if spec["policer_type"] != "POLICER_SINGLE_RATE_TWO_COLOR" {
					t.Errorf("expected policer_type=POLICER_SINGLE_RATE_TWO_COLOR, got %v", spec["policer_type"])
				}
			},
		},
		{
			name:         "policer does not override existing values",
			resourceType: "policer",
			inputSpec: map[string]interface{}{
				"policer_mode": "POLICER_MODE_SHARED",
				"policer_type": "POLICER_TWO_RATE_THREE_COLOR",
			},
			checkFunc: func(t *testing.T, spec map[string]interface{}) {
				// Should not override existing values
				if spec["policer_mode"] != "POLICER_MODE_SHARED" {
					t.Errorf("expected policer_mode=POLICER_MODE_SHARED (preserved), got %v", spec["policer_mode"])
				}
				if spec["policer_type"] != "POLICER_TWO_RATE_THREE_COLOR" {
					t.Errorf("expected policer_type=POLICER_TWO_RATE_THREE_COLOR (preserved), got %v", spec["policer_type"])
				}
			},
		},
		{
			name:         "healthcheck defaults",
			resourceType: "healthcheck",
			inputSpec: map[string]interface{}{
				"healthy_threshold":   3,
				"unhealthy_threshold": 2,
				"interval":            10,
				"timeout":             1,
				"http_health_check": map[string]interface{}{
					"path": "/healthz",
				},
			},
			checkFunc: func(t *testing.T, spec map[string]interface{}) {
				// Top-level defaults
				if spec["jitter"] != 0 {
					t.Errorf("expected jitter=0, got %v", spec["jitter"])
				}
				if spec["jitter_percent"] != 0 {
					t.Errorf("expected jitter_percent=0, got %v", spec["jitter_percent"])
				}

				// Nested defaults
				httpCheck, ok := spec["http_health_check"].(map[string]interface{})
				if !ok {
					t.Fatal("http_health_check not found or wrong type")
				}
				if httpCheck["use_http2"] != false {
					t.Errorf("expected use_http2=false, got %v", httpCheck["use_http2"])
				}
				if _, hasHeaders := httpCheck["headers"]; !hasHeaders {
					t.Error("expected headers to be set")
				}
			},
		},
		{
			name:         "unknown resource type",
			resourceType: "unknown_resource",
			inputSpec: map[string]interface{}{
				"some_field": "value",
			},
			checkFunc: func(t *testing.T, spec map[string]interface{}) {
				// Should not modify spec for unknown resources
				if len(spec) != 1 {
					t.Errorf("expected spec to have 1 field, got %d", len(spec))
				}
				if spec["some_field"] != "value" {
					t.Errorf("expected some_field=value, got %v", spec["some_field"])
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ApplyDiscoveredDefaults(tt.inputSpec, tt.resourceType)
			tt.checkFunc(t, tt.inputSpec)
		})
	}
}

func TestGetResourcesWithDefaults(t *testing.T) {
	resources := GetResourcesWithDefaults()
	if len(resources) == 0 {
		t.Error("expected at least one resource with defaults")
	}

	// Verify policer is in the list
	found := false
	for _, r := range resources {
		if r == "policer" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected policer to be in resources with defaults")
	}
}

func TestHasDiscoveredDefaults(t *testing.T) {
	tests := []struct {
		resourceType string
		expected     bool
	}{
		{"policer", true},
		{"policers", true},
		{"healthcheck", true},
		{"healthchecks", true},
		{"ip_prefix_set", true},
		{"ip_prefix_sets", true},
		{"unknown", false},
		{"namespace", false},
	}

	for _, tt := range tests {
		t.Run(tt.resourceType, func(t *testing.T) {
			result := HasDiscoveredDefaults(tt.resourceType)
			if result != tt.expected {
				t.Errorf("HasDiscoveredDefaults(%q) = %v, expected %v", tt.resourceType, result, tt.expected)
			}
		})
	}
}
