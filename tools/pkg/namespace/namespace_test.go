package namespace

import (
	"testing"
)

func TestForResource(t *testing.T) {
	tests := []struct {
		resourceName   string
		expectedType   Type
		expectedString string
	}{
		// System resources
		{"aws_vpc_site", System, "system"},
		{"azure_vnet_site", System, "system"},
		{"gcp_vpc_site", System, "system"},
		{"namespace", System, "system"},
		{"virtual_network", System, "system"},
		{"cloud_credentials", System, "system"},
		{"k8s_cluster", System, "system"},
		{"bgp", System, "system"},
		{"bgp_asn_set", System, "system"},

		// Shared resources
		{"app_firewall", Shared, "shared"},
		{"service_policy", Shared, "shared"},
		{"certificate", Shared, "shared"},
		{"rate_limiter", Shared, "shared"},
		{"user_identification", Shared, "shared"},
		{"ip_prefix_set", Shared, "shared"},
		{"alert_policy", Shared, "shared"},
		{"api_definition", Shared, "shared"},
		{"policer", Shared, "shared"},

		// Application resources (default)
		{"http_loadbalancer", Application, "staging"},
		{"tcp_loadbalancer", Application, "staging"},
		{"origin_pool", Application, "staging"},
		{"healthcheck", Application, "staging"},
		{"route_table", Application, "staging"},
	}

	for _, tt := range tests {
		t.Run(tt.resourceName, func(t *testing.T) {
			nsType, nsString := ForResource(tt.resourceName)
			if nsType != tt.expectedType {
				t.Errorf("ForResource(%q) type = %v, want %v", tt.resourceName, nsType, tt.expectedType)
			}
			if nsString != tt.expectedString {
				t.Errorf("ForResource(%q) string = %q, want %q", tt.resourceName, nsString, tt.expectedString)
			}
		})
	}
}

func TestTypeString(t *testing.T) {
	tests := []struct {
		nsType   Type
		expected string
	}{
		{System, "system"},
		{Shared, "shared"},
		{Application, "staging"},
		{Type(99), "staging"}, // Unknown defaults to staging
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.nsType.String()
			if result != tt.expected {
				t.Errorf("Type(%d).String() = %q, want %q", tt.nsType, result, tt.expected)
			}
		})
	}
}

func TestForReference(t *testing.T) {
	tests := []struct {
		resourceType string
		expected     string
	}{
		{"aws_vpc_site", "system"},
		{"app_firewall", "shared"},
		{"http_loadbalancer", "staging"},
		{"origin_pool", "staging"},
		{"certificate", "shared"},
	}

	for _, tt := range tests {
		t.Run(tt.resourceType, func(t *testing.T) {
			result := ForReference(tt.resourceType)
			if result != tt.expected {
				t.Errorf("ForReference(%q) = %q, want %q", tt.resourceType, result, tt.expected)
			}
		})
	}
}

func TestIsSystem(t *testing.T) {
	tests := []struct {
		resourceName string
		expected     bool
	}{
		{"aws_vpc_site", true},
		{"namespace", true},
		{"app_firewall", false},
		{"http_loadbalancer", false},
	}

	for _, tt := range tests {
		t.Run(tt.resourceName, func(t *testing.T) {
			result := IsSystem(tt.resourceName)
			if result != tt.expected {
				t.Errorf("IsSystem(%q) = %v, want %v", tt.resourceName, result, tt.expected)
			}
		})
	}
}

func TestIsShared(t *testing.T) {
	tests := []struct {
		resourceName string
		expected     bool
	}{
		{"app_firewall", true},
		{"certificate", true},
		{"aws_vpc_site", false},
		{"http_loadbalancer", false},
	}

	for _, tt := range tests {
		t.Run(tt.resourceName, func(t *testing.T) {
			result := IsShared(tt.resourceName)
			if result != tt.expected {
				t.Errorf("IsShared(%q) = %v, want %v", tt.resourceName, result, tt.expected)
			}
		})
	}
}

func TestIsApplication(t *testing.T) {
	tests := []struct {
		resourceName string
		expected     bool
	}{
		{"http_loadbalancer", true},
		{"origin_pool", true},
		{"aws_vpc_site", false},
		{"app_firewall", false},
	}

	for _, tt := range tests {
		t.Run(tt.resourceName, func(t *testing.T) {
			result := IsApplication(tt.resourceName)
			if result != tt.expected {
				t.Errorf("IsApplication(%q) = %v, want %v", tt.resourceName, result, tt.expected)
			}
		})
	}
}

func TestGetSystemResources(t *testing.T) {
	resources := GetSystemResources()

	// Verify it's a copy (modifying it shouldn't affect the original)
	resources["test_resource"] = true

	// The original should not have the test resource
	if IsSystem("test_resource") {
		t.Error("GetSystemResources() should return a copy, not the original map")
	}

	// Verify some expected entries exist
	expectedResources := []string{"aws_vpc_site", "namespace", "k8s_cluster"}
	for _, name := range expectedResources {
		if !resources[name] {
			t.Errorf("GetSystemResources() missing expected resource %q", name)
		}
	}
}

func TestGetSharedResources(t *testing.T) {
	resources := GetSharedResources()

	// Verify it's a copy (modifying it shouldn't affect the original)
	resources["test_resource"] = true

	// The original should not have the test resource
	if IsShared("test_resource") {
		t.Error("GetSharedResources() should return a copy, not the original map")
	}

	// Verify some expected entries exist
	expectedResources := []string{"app_firewall", "certificate", "rate_limiter"}
	for _, name := range expectedResources {
		if !resources[name] {
			t.Errorf("GetSharedResources() missing expected resource %q", name)
		}
	}
}
