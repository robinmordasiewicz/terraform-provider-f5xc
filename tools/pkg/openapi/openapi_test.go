package openapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSchema_IsRef(t *testing.T) {
	tests := []struct {
		name     string
		schema   Schema
		expected bool
	}{
		{"with ref", Schema{Ref: "#/components/schemas/Foo"}, true},
		{"without ref", Schema{Type: "string"}, false},
		{"empty ref", Schema{Ref: ""}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.schema.IsRef(); got != tt.expected {
				t.Errorf("Schema.IsRef() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSchema_IsArray(t *testing.T) {
	tests := []struct {
		name     string
		schema   Schema
		expected bool
	}{
		{"array type", Schema{Type: "array"}, true},
		{"object type", Schema{Type: "object"}, false},
		{"string type", Schema{Type: "string"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.schema.IsArray(); got != tt.expected {
				t.Errorf("Schema.IsArray() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSchema_IsObject(t *testing.T) {
	tests := []struct {
		name     string
		schema   Schema
		expected bool
	}{
		{"object type", Schema{Type: "object"}, true},
		{"array type", Schema{Type: "array"}, false},
		{"string type", Schema{Type: "string"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.schema.IsObject(); got != tt.expected {
				t.Errorf("Schema.IsObject() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSchema_IsRequired(t *testing.T) {
	schema := Schema{
		Required: []string{"name", "namespace"},
	}

	tests := []struct {
		property string
		expected bool
	}{
		{"name", true},
		{"namespace", true},
		{"description", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.property, func(t *testing.T) {
			if got := schema.IsRequired(tt.property); got != tt.expected {
				t.Errorf("Schema.IsRequired(%q) = %v, want %v", tt.property, got, tt.expected)
			}
		})
	}
}

func TestSchema_HasProperties(t *testing.T) {
	tests := []struct {
		name     string
		schema   Schema
		expected bool
	}{
		{
			"with properties",
			Schema{Properties: map[string]Schema{"foo": {Type: "string"}}},
			true,
		},
		{
			"empty properties",
			Schema{Properties: map[string]Schema{}},
			false,
		},
		{
			"nil properties",
			Schema{},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.schema.HasProperties(); got != tt.expected {
				t.Errorf("Schema.HasProperties() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGetRefName(t *testing.T) {
	tests := []struct {
		ref      string
		expected string
	}{
		{"#/components/schemas/Namespace", "Namespace"},
		{"#/components/schemas/foo.bar.Baz", "foo.bar.Baz"}, // GetRefName extracts after last /
		{"Namespace", "Namespace"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.ref, func(t *testing.T) {
			if got := GetRefName(tt.ref); got != tt.expected {
				t.Errorf("GetRefName(%q) = %q, want %q", tt.ref, got, tt.expected)
			}
		})
	}
}

func TestParseSpecFilename(t *testing.T) {
	tests := []struct {
		filename     string
		wantResource string
		wantURL      string
		wantErr      bool
	}{
		{
			"docs-cloud-f5-com.1234.public.ves.io.schema.namespace.ves-swagger.json",
			"namespace",
			"namespace",
			false,
		},
		{
			"docs-cloud-f5-com.5678.public.ves.io.schema.views.http_loadbalancer.ves-swagger.json",
			"http_loadbalancer",
			"views-http-loadbalancer",
			false,
		},
		{
			"docs-cloud-f5-com.9999.public.ves.io.schema.aws.vpc_site.ves-swagger.json",
			"vpc_site",
			"aws-vpc-site",
			false,
		},
		{
			"invalid-filename.json",
			"",
			"",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			info, err := ParseSpecFilename(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSpecFilename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if info.ResourceName != tt.wantResource {
				t.Errorf("ParseSpecFilename().ResourceName = %q, want %q", info.ResourceName, tt.wantResource)
			}
			if info.URLPath != tt.wantURL {
				t.Errorf("ParseSpecFilename().URLPath = %q, want %q", info.URLPath, tt.wantURL)
			}
		})
	}
}

func TestSpec_ResolveRef(t *testing.T) {
	spec := &Spec{
		Components: Components{
			Schemas: map[string]Schema{
				"Namespace": {Type: "object", Description: "A namespace"},
			},
		},
	}

	tests := []struct {
		ref     string
		wantErr bool
	}{
		{"#/components/schemas/Namespace", false},
		{"#/components/schemas/NotFound", true},
		{"external/ref", true},
	}

	for _, tt := range tests {
		t.Run(tt.ref, func(t *testing.T) {
			schema, err := spec.ResolveRef(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("Spec.ResolveRef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && schema == nil {
				t.Error("Spec.ResolveRef() returned nil schema")
			}
		})
	}
}

func TestParseFile(t *testing.T) {
	// Create a temporary test file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test-spec.json")

	testJSON := `{
		"openapi": "3.0.0",
		"info": {"title": "Test API", "version": "1.0.0"},
		"paths": {},
		"components": {
			"schemas": {
				"TestSchema": {
					"type": "object",
					"description": "A test schema"
				}
			}
		}
	}`

	if err := os.WriteFile(testFile, []byte(testJSON), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	spec, err := ParseFile(testFile)
	if err != nil {
		t.Fatalf("ParseFile() error = %v", err)
	}

	if spec.OpenAPI != "3.0.0" {
		t.Errorf("ParseFile().OpenAPI = %q, want %q", spec.OpenAPI, "3.0.0")
	}

	if _, ok := spec.Components.Schemas["TestSchema"]; !ok {
		t.Error("ParseFile() missing TestSchema in components")
	}
}

func TestGetAPIDocURL(t *testing.T) {
	pathMap := map[string]string{
		"http_loadbalancer": "views-http-loadbalancer",
		"namespace":         "namespace",
	}

	tests := []struct {
		resource string
		expected string
	}{
		{"http_loadbalancer", "https://docs.cloud.f5.com/docs-v2/api/views-http-loadbalancer"},
		{"namespace", "https://docs.cloud.f5.com/docs-v2/api/namespace"},
		{"unknown", ""},
	}

	for _, tt := range tests {
		t.Run(tt.resource, func(t *testing.T) {
			result := GetAPIDocURL(tt.resource, pathMap)
			if result != tt.expected {
				t.Errorf("GetAPIDocURL(%q) = %q, want %q", tt.resource, result, tt.expected)
			}
		})
	}
}
