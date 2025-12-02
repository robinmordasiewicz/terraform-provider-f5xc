// Package resource provides resource metadata, categories, and skip lists
// for code generation tools in the F5XC Terraform provider.
package resource

// SkippedResources lists resources that should be skipped during code generation.
// These are typically resources that require special handling or are deprecated.
var SkippedResources = map[string]bool{
	// Deprecated or removed
	"deprecated_resource": true,

	// Special handling required
	"blindfold": true, // Handled specially by provider-defined functions
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

// IsSkipped returns true if the resource should be skipped during generation.
func IsSkipped(resourceName string) bool {
	return SkippedResources[resourceName]
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
