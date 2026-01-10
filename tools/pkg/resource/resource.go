// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package resource provides resource metadata, categories, and skip lists
// for code generation tools in the F5XC Terraform provider.
package resource

// SkipReason documents why a resource should be skipped.
type SkipReason struct {
	Reason      string // Human-readable reason
	Category    string // Category of skip: "deprecated", "special", "credentials", "infrastructure", "permissions", "premium"
	SkipGenerate bool  // Skip during code generation
	SkipAPITest  bool  // Skip during API testing/discovery
}

// SkippedResources lists resources that should be skipped during code generation and/or API testing.
// Each entry includes a documented reason and which operations to skip.
var SkippedResources = map[string]SkipReason{
	// ============================================================================
	// Special Handling Required
	// ============================================================================
	"blindfold": {
		Reason:      "Handled specially by provider-defined functions",
		Category:    "special",
		SkipGenerate: true,
		SkipAPITest:  true,
	},

	// ============================================================================
	// Cloud Provider Sites - Require External Credentials
	// ============================================================================
	"aws_vpc_site": {
		Reason:      "Requires AWS credentials and VPC infrastructure",
		Category:    "credentials",
		SkipGenerate: false,
		SkipAPITest:  true,
	},
	"azure_vnet_site": {
		Reason:      "Requires Azure credentials and VNet infrastructure",
		Category:    "credentials",
		SkipGenerate: false,
		SkipAPITest:  true,
	},
	"gcp_vpc_site": {
		Reason:      "Requires GCP credentials and VPC infrastructure",
		Category:    "credentials",
		SkipGenerate: false,
		SkipAPITest:  true,
	},
	"aws_tgw_site": {
		Reason:      "Requires AWS credentials and Transit Gateway",
		Category:    "credentials",
		SkipGenerate: false,
		SkipAPITest:  true,
	},

	// ============================================================================
	// Cloud Credentials - Require External Secrets
	// ============================================================================
	"cloud_credentials": {
		Reason:      "Requires cloud provider API secrets",
		Category:    "credentials",
		SkipGenerate: false,
		SkipAPITest:  true,
	},

	// ============================================================================
	// Tenant Operations - Require Special Permissions
	// ============================================================================
	"tenant_configuration": {
		Reason:      "Requires tenant administrator privileges",
		Category:    "permissions",
		SkipGenerate: false,
		SkipAPITest:  true,
	},
	"child_tenant": {
		Reason:      "Requires MSP (Multi-Service Provider) tenant",
		Category:    "permissions",
		SkipGenerate: false,
		SkipAPITest:  true,
	},
	"child_tenant_manager": {
		Reason:      "Requires MSP tenant administrator",
		Category:    "permissions",
		SkipGenerate: false,
		SkipAPITest:  true,
	},
	"managed_tenant": {
		Reason:      "Requires MSP tenant",
		Category:    "permissions",
		SkipGenerate: false,
		SkipAPITest:  true,
	},
	"allowed_tenant": {
		Reason:      "Requires cross-tenant configuration",
		Category:    "permissions",
		SkipGenerate: false,
		SkipAPITest:  true,
	},

	// ============================================================================
	// Physical Infrastructure - Require Hardware
	// ============================================================================
	"securemesh_site": {
		Reason:      "Requires physical Secure Mesh hardware",
		Category:    "infrastructure",
		SkipGenerate: false,
		SkipAPITest:  true,
	},
	"securemesh_site_v2": {
		Reason:      "Requires physical Secure Mesh hardware v2",
		Category:    "infrastructure",
		SkipGenerate: false,
		SkipAPITest:  true,
	},
	"voltstack_site": {
		Reason:      "Requires physical VoltStack hardware",
		Category:    "infrastructure",
		SkipGenerate: false,
		SkipAPITest:  true,
	},
	"registration": {
		Reason:      "Requires physical site registration token",
		Category:    "infrastructure",
		SkipGenerate: false,
		SkipAPITest:  true,
	},

	// ============================================================================
	// Infrastructure Dependencies - Require Existing Resources
	// ============================================================================
	"fleet": {
		Reason:      "Requires existing site infrastructure",
		Category:    "infrastructure",
		SkipGenerate: false,
		SkipAPITest:  true,
	},
	"workload": {
		Reason:      "Requires existing Kubernetes infrastructure",
		Category:    "infrastructure",
		SkipGenerate: false,
		SkipAPITest:  true,
	},

	// ============================================================================
	// Premium/Licensed Features
	// ============================================================================
	"cminstance": {
		Reason:      "Requires Cloud Manager subscription",
		Category:    "premium",
		SkipGenerate: false,
		SkipAPITest:  true,
	},
	"addon_subscription": {
		Reason:      "Requires active subscription",
		Category:    "premium",
		SkipGenerate: false,
		SkipAPITest:  true,
	},
}

// ManuallyMaintainedFiles lists files in internal/provider that are manually maintained
// and should not be overwritten by code generation.
var ManuallyMaintainedFiles = map[string]bool{
	"provider.go":                  true,
	"functions_registration.go":   true,
}

// ManuallyMaintainedDirs lists directories that contain manually maintained code.
var ManuallyMaintainedDirs = []string{
	"internal/functions",
	"internal/blindfold",
}

// IsSkipped returns true if the resource should be skipped during code generation.
func IsSkipped(resourceName string) bool {
	if reason, ok := SkippedResources[resourceName]; ok {
		return reason.SkipGenerate
	}
	return false
}

// IsSkippedForAPITest returns true if the resource should be skipped during API testing/discovery.
func IsSkippedForAPITest(resourceName string) bool {
	if reason, ok := SkippedResources[resourceName]; ok {
		return reason.SkipAPITest
	}
	return false
}

// GetSkipReason returns the skip reason for a resource, or empty string if not skipped.
func GetSkipReason(resourceName string) string {
	if reason, ok := SkippedResources[resourceName]; ok {
		return reason.Reason
	}
	return ""
}

// IsManuallyMaintained returns true if the file is manually maintained.
func IsManuallyMaintained(filename string) bool {
	return ManuallyMaintainedFiles[filename]
}

// ResourceInfo contains metadata about a resource.
type ResourceInfo struct {
	Name        string
	Category    string
	Namespace   string
	IsLongRunning bool
}

// GetInfo returns comprehensive information about a resource.
func GetInfo(resourceName string) ResourceInfo {
	return ResourceInfo{
		Name:          resourceName,
		Category:      GetCategory(resourceName),
		IsLongRunning: IsLongRunning(resourceName),
	}
}
