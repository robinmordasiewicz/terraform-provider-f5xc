package resource

import (
	"testing"
)

func TestGetCategory(t *testing.T) {
	tests := []struct {
		resourceName string
		expected     string
	}{
		// Override tests
		{"apm", "Monitoring"},
		{"crl", "Certificates"},
		{"bgp", "Networking"},
		{"namespace", "Organization"},
		{"data_group", "BIG-IP Integration"},
		{"user_identification", "Security"},

		// Pattern matching tests
		{"aws_vpc_site", "Sites"},
		{"azure_vnet_site", "Sites"},
		{"gcp_vpc_site", "Sites"},
		{"http_loadbalancer", "Load Balancing"},
		{"tcp_loadbalancer", "Load Balancing"},
		{"app_firewall", "Security"},
		{"service_policy", "Security"},
		{"rate_limiter", "Security"},
		{"network_connector", "Networking"},
		{"bgp_asn_set", "Networking"},
		{"dns_zone", "DNS"},
		{"cloud_credentials", "Authentication"},
		{"certificate_chain", "Certificates"},
		{"api_definition", "API Security"},
		{"log_receiver", "Monitoring"},
		{"tenant_configuration", "Organization"},
		{"k8s_cluster", "Kubernetes"},
		{"virtual_k8s", "Kubernetes"},

		// Default to Uncategorized
		{"unknown_resource", "Uncategorized"},
	}

	for _, tt := range tests {
		t.Run(tt.resourceName, func(t *testing.T) {
			result := GetCategory(tt.resourceName)
			if result != tt.expected {
				t.Errorf("GetCategory(%q) = %q, want %q", tt.resourceName, result, tt.expected)
			}
		})
	}
}

func TestIsLongRunning(t *testing.T) {
	tests := []struct {
		resourceName string
		expected     bool
	}{
		{"aws_vpc_site", true},
		{"azure_vnet_site", true},
		{"gcp_vpc_site", true},
		{"k8s_cluster", true},
		{"virtual_k8s", true},
		{"http_loadbalancer", false},
		{"origin_pool", false},
		{"namespace", false},
	}

	for _, tt := range tests {
		t.Run(tt.resourceName, func(t *testing.T) {
			result := IsLongRunning(tt.resourceName)
			if result != tt.expected {
				t.Errorf("IsLongRunning(%q) = %v, want %v", tt.resourceName, result, tt.expected)
			}
		})
	}
}

func TestGetTimeout(t *testing.T) {
	tests := []struct {
		resourceName string
		expected     string
	}{
		{"aws_vpc_site", "30m"},
		{"k8s_cluster", "30m"},
		{"http_loadbalancer", "10m"},
		{"origin_pool", "10m"},
	}

	for _, tt := range tests {
		t.Run(tt.resourceName, func(t *testing.T) {
			result := GetTimeout(tt.resourceName)
			if result != tt.expected {
				t.Errorf("GetTimeout(%q) = %q, want %q", tt.resourceName, result, tt.expected)
			}
		})
	}
}

func TestIsSkipped(t *testing.T) {
	tests := []struct {
		resourceName string
		expected     bool
	}{
		{"blindfold", true},
		{"http_loadbalancer", false},
		{"origin_pool", false},
	}

	for _, tt := range tests {
		t.Run(tt.resourceName, func(t *testing.T) {
			result := IsSkipped(tt.resourceName)
			if result != tt.expected {
				t.Errorf("IsSkipped(%q) = %v, want %v", tt.resourceName, result, tt.expected)
			}
		})
	}
}

func TestIsManuallyMaintained(t *testing.T) {
	tests := []struct {
		filename string
		expected bool
	}{
		{"provider.go", true},
		{"functions_registration.go", true},
		{"http_loadbalancer_resource.go", false},
		{"namespace_resource.go", false},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			result := IsManuallyMaintained(tt.filename)
			if result != tt.expected {
				t.Errorf("IsManuallyMaintained(%q) = %v, want %v", tt.filename, result, tt.expected)
			}
		})
	}
}

func TestGetInfo(t *testing.T) {
	info := GetInfo("aws_vpc_site")

	if info.Name != "aws_vpc_site" {
		t.Errorf("GetInfo(aws_vpc_site).Name = %q, want %q", info.Name, "aws_vpc_site")
	}
	if info.Category != "Sites" {
		t.Errorf("GetInfo(aws_vpc_site).Category = %q, want %q", info.Category, "Sites")
	}
	if !info.IsLongRunning {
		t.Error("GetInfo(aws_vpc_site).IsLongRunning = false, want true")
	}
}

func TestAllCategories(t *testing.T) {
	categories := AllCategories()

	// Should have multiple categories
	if len(categories) < 10 {
		t.Errorf("AllCategories() returned %d categories, expected at least 10", len(categories))
	}

	// Check for expected categories
	expectedCategories := []string{"Sites", "Load Balancing", "Security", "Networking", "DNS"}
	for _, expected := range expectedCategories {
		found := false
		for _, cat := range categories {
			if cat == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("AllCategories() missing expected category %q", expected)
		}
	}
}
