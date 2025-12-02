// Package openapi provides types and utilities for parsing OpenAPI specifications
// for the F5XC Terraform provider code generation tools.
package openapi

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ParseFile reads and parses an OpenAPI specification file.
func ParseFile(path string) (*Spec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	var spec Spec
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("parsing JSON: %w", err)
	}

	return &spec, nil
}

// ResolveRef resolves a $ref string to the actual schema.
// Only supports local references (#/components/schemas/SchemaName).
func (s *Spec) ResolveRef(ref string) (*Schema, error) {
	if !strings.HasPrefix(ref, "#/components/schemas/") {
		return nil, fmt.Errorf("unsupported ref format: %s", ref)
	}

	schemaName := strings.TrimPrefix(ref, "#/components/schemas/")
	schema, ok := s.Components.Schemas[schemaName]
	if !ok {
		return nil, fmt.Errorf("schema not found: %s", schemaName)
	}

	return &schema, nil
}

// GetRefName extracts the schema name from a $ref string.
func GetRefName(ref string) string {
	if ref == "" {
		return ""
	}
	parts := strings.Split(ref, "/")
	return parts[len(parts)-1]
}

// SpecFileInfo contains information extracted from an OpenAPI spec filename.
type SpecFileInfo struct {
	FilePath     string
	ResourceName string
	SchemaPath   string // e.g., "views.http_loadbalancer"
	URLPath      string // e.g., "views-http-loadbalancer"
}

// ParseSpecFilename extracts resource information from an OpenAPI spec filename.
// Pattern: docs-cloud-f5-com.XXXX.public.ves.io.schema.{path}.ves-swagger.json
func ParseSpecFilename(filename string) (*SpecFileInfo, error) {
	specRegex := regexp.MustCompile(`docs-cloud-f5-com\.\d+\.public\.ves\.io\.schema\.(.+)\.ves-swagger\.json`)

	base := filepath.Base(filename)
	matches := specRegex.FindStringSubmatch(base)
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("filename does not match expected pattern: %s", filename)
	}

	schemaPath := matches[1]
	parts := strings.Split(schemaPath, ".")
	resourceName := parts[len(parts)-1]

	// Convert schema path to URL format: dots -> hyphens, underscores -> hyphens
	urlPath := strings.ReplaceAll(schemaPath, ".", "-")
	urlPath = strings.ReplaceAll(urlPath, "_", "-")

	return &SpecFileInfo{
		FilePath:     filename,
		ResourceName: resourceName,
		SchemaPath:   schemaPath,
		URLPath:      urlPath,
	}, nil
}

// FindSpecFiles finds all OpenAPI spec files in a directory.
func FindSpecFiles(dir string) ([]string, error) {
	pattern := filepath.Join(dir, "*.json")
	return filepath.Glob(pattern)
}

// BuildResourceAPIPathMap scans the OpenAPI spec directory and builds a mapping
// from resource names to their API documentation paths.
// Example: "http_loadbalancer" -> "views-http-loadbalancer"
func BuildResourceAPIPathMap(specDir string) (map[string]string, error) {
	files, err := FindSpecFiles(specDir)
	if err != nil {
		return nil, fmt.Errorf("scanning spec directory: %w", err)
	}

	pathMap := make(map[string]string)
	for _, file := range files {
		info, err := ParseSpecFilename(file)
		if err != nil {
			// Skip files that don't match the pattern
			continue
		}
		pathMap[info.ResourceName] = info.URLPath
	}

	return pathMap, nil
}

// GetAPIDocURL returns the F5 API documentation URL for a resource.
func GetAPIDocURL(resourceName string, pathMap map[string]string) string {
	if urlPath, ok := pathMap[resourceName]; ok {
		return fmt.Sprintf("https://docs.cloud.f5.com/docs-v2/api/%s", urlPath)
	}
	return ""
}
