//go:build ignore
// +build ignore

// This tool generates an AI-friendly constraint manifest from F5 XC OpenAPI specifications.
// The manifest enables AI tools to deterministically create valid Terraform configurations
// without trial-and-error iteration.
//
// Usage:
//
//	go run tools/generate-ai-manifest.go
//	go run tools/generate-ai-manifest.go -output docs/ai-manifest.json
//	go run tools/generate-ai-manifest.go -verbose
//	go run tools/generate-ai-manifest.go -validate
//
// Output:
//
//	Creates docs/ai-manifest.json with:
//	- OneOf constraint groups (mutually exclusive fields)
//	- Default values for attributes
//	- Resource dependencies and ordering
//	- Common usage patterns and AI hints
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/f5xc/terraform-provider-f5xc/tools/pkg/defaults"
	"github.com/f5xc/terraform-provider-f5xc/tools/pkg/manifest"
)

// Version of the manifest generator
const generatorVersion = "1.0.0"

func main() {
	// Command-line flags
	outputFile := flag.String("output", "docs/ai-manifest.json", "Output file path for the AI manifest")
	specDir := flag.String("specs", "docs/specifications/api", "Directory containing OpenAPI spec files")
	defaultsFile := flag.String("defaults", "tools/api-defaults.json", "Path to the API defaults file")
	providerVersion := flag.String("version", "", "Provider version to include in manifest")
	verbose := flag.Bool("verbose", false, "Enable verbose output")
	validate := flag.Bool("validate", false, "Validate existing manifest without regenerating")
	dryRun := flag.Bool("dry-run", false, "Generate manifest but don't write to file")

	flag.Parse()

	if *validate {
		if err := validateManifest(*outputFile); err != nil {
			fmt.Fprintf(os.Stderr, "Validation failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Manifest validation successful")
		return
	}

	// Verify spec directory exists
	if _, err := os.Stat(*specDir); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: Spec directory not found: %s\n", *specDir)
		os.Exit(1)
	}

	if *verbose {
		fmt.Printf("AI Manifest Generator v%s\n", generatorVersion)
		fmt.Printf("Spec directory: %s\n", *specDir)
		fmt.Printf("Output file: %s\n", *outputFile)
		fmt.Println("---")
	}

	// Step 1: Extract constraints from OpenAPI specs
	if *verbose {
		fmt.Println("Step 1: Extracting constraints from OpenAPI specifications...")
	}

	extractor := manifest.NewExtractor(*specDir, *verbose)
	resources, err := extractor.ExtractFromAllSpecs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error extracting from specs: %v\n", err)
		os.Exit(1)
	}

	if *verbose {
		fmt.Printf("Extracted constraints from %d resources\n", len(resources))
	}

	// Step 2: Convert to AI manifest format
	if *verbose {
		fmt.Println("Step 2: Converting to AI manifest format...")
	}

	aiManifest := manifest.ConvertToManifest(resources, *providerVersion)
	aiManifest.GeneratedAt = time.Now().UTC()

	// Step 3: Merge API-discovered defaults if available
	if *verbose {
		fmt.Println("Step 3: Merging API-discovered defaults...")
	}

	if _, err := os.Stat(*defaultsFile); err == nil {
		if err := mergeDefaults(aiManifest, *defaultsFile, *verbose); err != nil {
			if *verbose {
				fmt.Printf("Warning: Could not merge defaults: %v\n", err)
			}
		}
	} else if *verbose {
		fmt.Printf("Defaults file not found: %s (skipping)\n", *defaultsFile)
	}

	// Step 4: Generate statistics
	stats := generateStats(aiManifest)
	if *verbose {
		fmt.Println("Step 4: Generating statistics...")
		fmt.Printf("  Total resources: %d\n", stats.TotalResources)
		fmt.Printf("  Resources with OneOf constraints: %d\n", stats.ResourcesWithOneOf)
		fmt.Printf("  Total OneOf groups: %d\n", stats.TotalOneOfGroups)
		fmt.Printf("  Resources with defaults: %d\n", stats.ResourcesWithDefaults)
		fmt.Printf("  Categories: %d\n", stats.Categories)
	}

	// Step 5: Write manifest to file
	if *dryRun {
		if *verbose {
			fmt.Println("Step 5: Dry run - not writing to file")
		}
		// Print to stdout instead
		output, err := json.MarshalIndent(aiManifest, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling manifest: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(output))
		return
	}

	if *verbose {
		fmt.Printf("Step 5: Writing manifest to %s...\n", *outputFile)
	}

	// Ensure output directory exists
	outputDir := filepath.Dir(*outputFile)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Marshal with pretty printing
	output, err := json.MarshalIndent(aiManifest, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling manifest: %v\n", err)
		os.Exit(1)
	}

	// Write to file
	if err := os.WriteFile(*outputFile, output, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing manifest: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated AI manifest: %s\n", *outputFile)
	fmt.Printf("  Resources: %d\n", aiManifest.TotalResources)
	fmt.Printf("  OneOf constraint groups: %d\n", stats.TotalOneOfGroups)
	fmt.Printf("  Categories: %d\n", stats.Categories)
}

// ManifestStats holds statistics about the generated manifest
type ManifestStats struct {
	TotalResources        int
	ResourcesWithOneOf    int
	TotalOneOfGroups      int
	ResourcesWithDefaults int
	Categories            int
}

// generateStats calculates statistics about the manifest
func generateStats(m *manifest.AIManifest) ManifestStats {
	stats := ManifestStats{
		TotalResources: len(m.Resources),
		Categories:     len(m.Categories),
	}

	for _, res := range m.Resources {
		if res.Constraints != nil && len(res.Constraints.OneOfGroups) > 0 {
			stats.ResourcesWithOneOf++
			stats.TotalOneOfGroups += len(res.Constraints.OneOfGroups)
		}
		if len(res.Defaults) > 0 {
			stats.ResourcesWithDefaults++
		}
	}

	return stats
}

// mergeDefaults merges API-discovered defaults into the manifest
func mergeDefaults(m *manifest.AIManifest, defaultsFile string, verbose bool) error {
	store := defaults.GetStore()
	if err := store.LoadFromFile(defaultsFile); err != nil {
		return fmt.Errorf("failed to load defaults: %w", err)
	}

	merged := 0
	for _, resourceName := range store.ListResources() {
		if res, ok := m.Resources[resourceName]; ok {
			resourceDefaults, exists := store.GetResourceDefaults(resourceName)
			if !exists {
				continue
			}
			if res.Defaults == nil {
				res.Defaults = make(map[string]interface{})
			}
			for fieldPath, defaultVal := range resourceDefaults {
				res.Defaults[fieldPath] = defaultVal.DefaultValue
				merged++
			}
		}
	}

	if verbose {
		fmt.Printf("  Merged %d default values from %s\n", merged, defaultsFile)
	}

	return nil
}

// validateManifest validates an existing manifest file
func validateManifest(filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read manifest: %w", err)
	}

	var m manifest.AIManifest
	if err := json.Unmarshal(data, &m); err != nil {
		return fmt.Errorf("failed to parse manifest: %w", err)
	}

	// Validate required fields
	if m.Version == "" {
		return fmt.Errorf("manifest missing version")
	}
	if m.Provider == "" {
		return fmt.Errorf("manifest missing provider name")
	}
	if len(m.Resources) == 0 {
		return fmt.Errorf("manifest contains no resources")
	}

	// Validate each resource
	for name, res := range m.Resources {
		if res.Name == "" {
			return fmt.Errorf("resource %s missing name", name)
		}
		if res.Constraints == nil {
			return fmt.Errorf("resource %s missing constraints", name)
		}

		// Validate OneOf groups
		for i, oneOf := range res.Constraints.OneOfGroups {
			if oneOf.GroupID == "" {
				return fmt.Errorf("resource %s OneOf group %d missing group_id", name, i)
			}
			if len(oneOf.Fields) == 0 {
				return fmt.Errorf("resource %s OneOf group %s has no fields", name, oneOf.GroupID)
			}
		}
	}

	// Print validation summary
	fmt.Printf("Validated manifest: %s\n", filepath)
	fmt.Printf("  Version: %s\n", m.Version)
	fmt.Printf("  Provider: %s\n", m.Provider)
	fmt.Printf("  Resources: %d\n", len(m.Resources))

	return nil
}
