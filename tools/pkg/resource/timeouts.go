// Package resource provides resource metadata, categories, and skip lists
// for code generation tools in the F5XC Terraform provider.
package resource

// LongRunningResources lists resources that have longer default timeouts (30 minutes).
// These match LongRunningResourceTypes in internal/timeouts/timeouts.go.
var LongRunningResources = map[string]bool{
	"aws_vpc_site":       true,
	"azure_vnet_site":    true,
	"gcp_vpc_site":       true,
	"aws_tgw_site":       true,
	"voltstack_site":     true,
	"securemesh_site":    true,
	"securemesh_site_v2": true,
	"k8s_cluster":        true,
	"virtual_k8s":        true,
}

// DefaultTimeout is the default timeout for most resources (10 minutes).
const DefaultTimeout = "10m"

// LongTimeout is the timeout for long-running resources (30 minutes).
const LongTimeout = "30m"

// IsLongRunning returns true if the resource has longer default timeouts.
func IsLongRunning(resourceName string) bool {
	return LongRunningResources[resourceName]
}

// GetTimeout returns the appropriate timeout string for a resource.
func GetTimeout(resourceName string) string {
	if IsLongRunning(resourceName) {
		return LongTimeout
	}
	return DefaultTimeout
}
