package naming

import (
	"testing"
)

func TestToTitleCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"http_load_balancer", "HTTP Load Balancer"},
		{"dns_zone", "DNS Zone"},
		{"aws_vpc_site", "AWS VPC Site"},
		{"tcp_loadbalancer", "TCP Loadbalancer"},
		{"origin_pool", "Origin Pool"},
		{"app_firewall", "App Firewall"},
		{"api_definition", "API Definition"},
		{"mtls_certificate", "mTLS Certificate"},
		{"oauth_provider", "OAuth Provider"},
		{"ipv4_address", "IPv4 Address"},
		{"bgp_asn_set", "BGP Asn Set"}, // ASN intentionally not uppercased for backward compatibility
		{"k8s_cluster", "K8S Cluster"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ToTitleCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToTitleCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToTitleCaseFromAnchor(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"http-load-balancer", "HTTP Load Balancer"},
		{"dns-zone", "DNS Zone"},
		{"aws-vpc-site", "AWS VPC Site"},
		{"mtls-certificate", "mTLS Certificate"},
		{"oauth-provider", "OAuth Provider"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ToTitleCaseFromAnchor(tt.input)
			if result != tt.expected {
				t.Errorf("ToTitleCaseFromAnchor(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HTTPLoadBalancer", "http_load_balancer"},
		{"DNSZone", "dns_zone"},
		{"AWSVPCSite", "awsvpc_site"},
		{"OriginPool", "origin_pool"},
		{"AppFirewall", "app_firewall"},
		{"SimpleTest", "simple_test"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ToSnakeCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToSnakeCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToHumanName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"http_loadbalancer", "HTTP Loadbalancer"},
		{"dns_zone", "DNS Zone"},
		{"aws_vpc_site", "AWS VPC Site"},
		{"origin_pool", "Origin Pool"},
		{"app_firewall", "App Firewall"},
		{"mtls_config", "mTLS Config"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ToHumanName(tt.input)
			if result != tt.expected {
				t.Errorf("ToHumanName(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToHumanReadableName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"http_loadbalancer", "HTTP Load Balancer"},   // compound word with spaces
		{"tcp_loadbalancer", "TCP Load Balancer"},     // compound word with spaces
		{"udp_loadbalancer", "UDP Load Balancer"},     // compound word with spaces
		{"cdn_loadbalancer", "CDN Load Balancer"},     // compound word with spaces
		{"dns_zone", "DNS Zone"},                      // acronym + word
		{"aws_vpc_site", "AWS VPC Site"},              // multiple acronyms
		{"origin_pool", "Origin Pool"},                // regular words
		{"app_firewall", "App Firewall"},              // regular words
		{"mtls_config", "mTLS Config"},                // mixed case acronym
		{"oauth_provider", "OAuth Provider"},          // mixed case acronym
		{"bigip_server", "BIG-IP Server"},             // compound word with special formatting
		{"websocket_connection", "WebSocket Connection"}, // mixed case compound word
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ToHumanReadableName(tt.input)
			if result != tt.expected {
				t.Errorf("ToHumanReadableName(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToAnchorName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"http_load_balancer", "http-load-balancer"},
		{"DNS_Zone", "dns-zone"},
		{"aws_VPC_site", "aws-vpc-site"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ToAnchorName(tt.input)
			if result != tt.expected {
				t.Errorf("ToAnchorName(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestNormalizeAcronyms(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"http load balancer", "HTTP load balancer"},
		{"Dns Zone", "DNS Zone"},
		{"Aws vpc Site", "AWS VPC Site"},
		{"mtls certificate", "mTLS certificate"},
		{"Oauth Provider", "OAuth Provider"},
		{"ipv4 address", "IPv4 address"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := NormalizeAcronyms(tt.input)
			if result != tt.expected {
				t.Errorf("NormalizeAcronyms(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStartsWithVowel(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"apple", true},
		{"elephant", true},
		{"ice", true},
		{"orange", true},
		{"umbrella", true},
		{"Apple", true},
		{"banana", false},
		{"car", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := StartsWithVowel(tt.input)
			if result != tt.expected {
				t.Errorf("StartsWithVowel(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToResourceTypeName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"http_loadbalancer", "HTTPLoadBalancer"}, // compound word: loadbalancer -> LoadBalancer
		{"dns_zone", "DNSZone"},
		{"aws_vpc_site", "AWSVPCSite"},
		{"origin_pool", "OriginPool"},
		{"app_firewall", "AppFirewall"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ToResourceTypeName(tt.input)
			if result != tt.expected {
				t.Errorf("ToResourceTypeName(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsUppercaseAcronym(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"HTTP", true},
		{"http", true},
		{"Http", true},
		{"DNS", true},
		{"aws", true},
		{"foo", false},
		{"bar", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := IsUppercaseAcronym(tt.input)
			if result != tt.expected {
				t.Errorf("IsUppercaseAcronym(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetMixedCaseAcronym(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"mtls", "mTLS"},
		{"MTLS", "mTLS"},
		{"oauth", "OAuth"},
		{"OAUTH", "OAuth"},
		{"ipv4", "IPv4"},
		{"foo", ""},
		{"http", ""}, // HTTP is uppercase, not mixed
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := GetMixedCaseAcronym(tt.input)
			if result != tt.expected {
				t.Errorf("GetMixedCaseAcronym(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
