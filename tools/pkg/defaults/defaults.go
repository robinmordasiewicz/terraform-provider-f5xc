// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package defaults provides API-discovered default value lookup for F5XC resources.
// These defaults are discovered by comparing API request/response payloads to find
// values that the API applies automatically when not specified in the request.
package defaults

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// DefaultValue represents an API-discovered default for a specific field.
type DefaultValue struct {
	Path         string      `json:"path"`
	DefaultValue interface{} `json:"default_value"`
	Type         string      `json:"type"`
}

// ResourceDefaults maps field paths to their discovered defaults for a resource.
type ResourceDefaults map[string]DefaultValue

// ResourceEntry represents a single resource entry in the api-defaults.json file.
type ResourceEntry struct {
	ResourceName string           `json:"resource_name"`
	Category     string           `json:"category"`
	Status       string           `json:"status"`
	DiscoveredAt string           `json:"discovered_at,omitempty"`
	FailureReason string          `json:"failure_reason,omitempty"`
	Defaults     ResourceDefaults `json:"defaults,omitempty"`
}

// APIDefaultsFile represents the top-level structure of api-defaults.json.
type APIDefaultsFile struct {
	Version        string                   `json:"version"`
	GeneratedAt    string                   `json:"generated_at"`
	APIEndpoint    string                   `json:"api_endpoint"`
	TotalResources int                      `json:"total_resources"`
	Discovered     int                      `json:"discovered"`
	Skipped        int                      `json:"skipped"`
	Failed         int                      `json:"failed"`
	Resources      map[string]ResourceEntry `json:"resources"`
}

// Store holds all discovered defaults indexed by resource name.
type Store struct {
	mu        sync.RWMutex
	defaults  map[string]ResourceDefaults
	loaded    bool
	loadError error
}

var (
	globalStore *Store
	once        sync.Once
)

// GetStore returns the global defaults store, loading from api-defaults.json if needed.
func GetStore() *Store {
	once.Do(func() {
		globalStore = &Store{
			defaults: make(map[string]ResourceDefaults),
		}
	})
	return globalStore
}

// LoadFromFile loads defaults from the specified JSON file.
func (s *Store) LoadFromFile(path string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := os.ReadFile(path)
	if err != nil {
		s.loadError = fmt.Errorf("failed to read defaults file: %w", err)
		return s.loadError
	}

	var apiFile APIDefaultsFile
	if err := json.Unmarshal(data, &apiFile); err != nil {
		s.loadError = fmt.Errorf("failed to parse defaults file: %w", err)
		return s.loadError
	}

	// Extract defaults from the nested structure
	s.defaults = make(map[string]ResourceDefaults)
	for name, entry := range apiFile.Resources {
		if entry.Status == "discovered" && entry.Defaults != nil {
			s.defaults[name] = entry.Defaults
		}
	}

	s.loaded = true
	s.loadError = nil
	return nil
}

// LoadFromDefaultPath loads defaults from the standard api-defaults.json location.
// It searches relative to the executable and common project paths.
func (s *Store) LoadFromDefaultPath() error {
	// Try common locations
	searchPaths := []string{
		"tools/api-defaults.json",
		"../tools/api-defaults.json",
		"../../tools/api-defaults.json",
	}

	// Also try relative to executable
	if execPath, err := os.Executable(); err == nil {
		execDir := filepath.Dir(execPath)
		searchPaths = append(searchPaths,
			filepath.Join(execDir, "api-defaults.json"),
			filepath.Join(execDir, "..", "tools", "api-defaults.json"),
		)
	}

	for _, path := range searchPaths {
		if _, err := os.Stat(path); err == nil {
			return s.LoadFromFile(path)
		}
	}

	return fmt.Errorf("api-defaults.json not found in search paths: %v", searchPaths)
}

// IsLoaded returns whether defaults have been successfully loaded.
func (s *Store) IsLoaded() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.loaded
}

// GetResourceDefaults returns all defaults for a specific resource.
func (s *Store) GetResourceDefaults(resourceName string) (ResourceDefaults, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	defaults, ok := s.defaults[resourceName]
	return defaults, ok
}

// GetDefault returns the default value for a specific resource field.
// The fieldPath should match the JSON path format (e.g., "spec.jitter_percent").
func (s *Store) GetDefault(resourceName, fieldPath string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	resourceDefaults, ok := s.defaults[resourceName]
	if !ok {
		return nil, false
	}

	defaultVal, ok := resourceDefaults[fieldPath]
	if !ok {
		return nil, false
	}

	return defaultVal.DefaultValue, true
}

// GetDefaultFormatted returns the default value formatted as a string suitable
// for documentation (e.g., "false", "0", "[]", "{}").
func (s *Store) GetDefaultFormatted(resourceName, fieldPath string) (string, bool) {
	val, ok := s.GetDefault(resourceName, fieldPath)
	if !ok {
		return "", false
	}

	return FormatDefaultValue(val), true
}

// FormatDefaultValue converts a default value to its documentation string format.
func FormatDefaultValue(val interface{}) string {
	if val == nil {
		return "null"
	}

	switch v := val.(type) {
	case bool:
		if v {
			return "true"
		}
		return "false"
	case float64:
		// JSON numbers are float64
		if v == float64(int64(v)) {
			return fmt.Sprintf("%d", int64(v))
		}
		return fmt.Sprintf("%v", v)
	case string:
		if v == "" {
			return `""`
		}
		return fmt.Sprintf(`"%s"`, v)
	case []interface{}:
		if len(v) == 0 {
			return "[]"
		}
		// For non-empty arrays, serialize to JSON
		data, _ := json.Marshal(v)
		return string(data)
	case map[string]interface{}:
		if len(v) == 0 {
			return "{}"
		}
		// Check for marker blocks (e.g., {"_marker": "empty_block"})
		if marker, ok := v["_marker"].(string); ok && marker == "empty_block" {
			return "{}"
		}
		// For non-empty maps, serialize to JSON
		data, _ := json.Marshal(v)
		return string(data)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// ListResources returns all resource names that have discovered defaults.
func (s *Store) ListResources() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	resources := make([]string, 0, len(s.defaults))
	for name := range s.defaults {
		resources = append(resources, name)
	}
	return resources
}

// FieldPathToTerraform converts a JSON path (e.g., "spec.http_health_check.use_http2")
// to a Terraform attribute path format for documentation matching.
func FieldPathToTerraform(jsonPath string) string {
	// Remove "spec." prefix if present (Terraform attributes don't have this)
	path := strings.TrimPrefix(jsonPath, "spec.")
	return path
}

// TerraformToFieldPath converts a Terraform attribute path to JSON path format.
func TerraformToFieldPath(tfPath string) string {
	// Add "spec." prefix for API field lookup
	return "spec." + tfPath
}

// GetDefaultByTerraformPath looks up a default using Terraform attribute path format.
func (s *Store) GetDefaultByTerraformPath(resourceName, tfPath string) (interface{}, bool) {
	// Try direct path first
	if val, ok := s.GetDefault(resourceName, tfPath); ok {
		return val, true
	}

	// Try with spec. prefix
	return s.GetDefault(resourceName, TerraformToFieldPath(tfPath))
}

// GetDefaultFormattedByTerraformPath returns formatted default using Terraform path.
func (s *Store) GetDefaultFormattedByTerraformPath(resourceName, tfPath string) (string, bool) {
	val, ok := s.GetDefaultByTerraformPath(resourceName, tfPath)
	if !ok {
		return "", false
	}
	return FormatDefaultValue(val), true
}
