// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// calculate-terraform-version.go
//
// This tool dynamically calculates the minimum required Terraform version
// based on features implemented in this provider.
//
// Usage:
//   go run tools/calculate-terraform-version.go [--update-templates]
//
// Version Detection Rules:
//   - Protocol 6.0 (terraform-plugin-framework): Terraform 1.0+
//   - Provider-defined functions: Terraform 1.8+
//   - Write-only attributes: Terraform 1.11+ (future)

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// VersionRequirement represents a feature and its minimum Terraform version
type VersionRequirement struct {
	Feature    string
	MinVersion string
	Detected   bool
}

// VersionInfo contains the calculated version information
type VersionInfo struct {
	MinimumVersion string               `json:"minimum_version"`
	Requirements   []VersionRequirement `json:"requirements"`
}

func main() {
	updateTemplates := flag.Bool("update-templates", false, "Update template files with calculated version")
	outputJSON := flag.Bool("json", false, "Output version info as JSON")
	flag.Parse()

	// Find project root (look for go.mod)
	projectRoot, err := findProjectRoot()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding project root: %v\n", err)
		os.Exit(1)
	}

	// Calculate minimum version
	versionInfo := calculateMinimumVersion(projectRoot)

	if *outputJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(versionInfo); err != nil {
			fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
			os.Exit(1)
		}
		return
	}

	fmt.Printf("Minimum Terraform Version: %s\n", versionInfo.MinimumVersion)
	fmt.Println("\nDetected Features:")
	for _, req := range versionInfo.Requirements {
		status := "not detected"
		if req.Detected {
			status = "detected"
		}
		fmt.Printf("  - %s (requires %s): %s\n", req.Feature, req.MinVersion, status)
	}

	if *updateTemplates {
		if err := updateTemplateFiles(projectRoot, versionInfo.MinimumVersion); err != nil {
			fmt.Fprintf(os.Stderr, "Error updating templates: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("\nTemplates updated successfully.")
	}
}

func findProjectRoot() (string, error) {
	// Start from current directory and look for go.mod
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find project root (no go.mod found)")
		}
		dir = parent
	}
}

func calculateMinimumVersion(projectRoot string) VersionInfo {
	requirements := []VersionRequirement{
		{
			Feature:    "Protocol 6.0 (terraform-plugin-framework)",
			MinVersion: "1.0",
			Detected:   true, // Always true for this provider
		},
		{
			Feature:    "Provider-defined functions",
			MinVersion: "1.8",
			Detected:   hasProviderFunctions(projectRoot),
		},
		// Future: Add more feature detection
		// {
		// 	Feature:    "Write-only attributes",
		// 	MinVersion: "1.11",
		// 	Detected:   hasWriteOnlyAttributes(projectRoot),
		// },
	}

	// Find the maximum required version
	minVersion := "1.0"
	for _, req := range requirements {
		if req.Detected && compareVersions(req.MinVersion, minVersion) > 0 {
			minVersion = req.MinVersion
		}
	}

	return VersionInfo{
		MinimumVersion: minVersion,
		Requirements:   requirements,
	}
}

// hasProviderFunctions checks if the provider implements provider-defined functions
func hasProviderFunctions(projectRoot string) bool {
	functionsDir := filepath.Join(projectRoot, "internal", "functions")

	// Check if directory exists
	info, err := os.Stat(functionsDir)
	if err != nil || !info.IsDir() {
		return false
	}

	// Check if there are any .go files (excluding test files)
	entries, err := os.ReadDir(functionsDir)
	if err != nil {
		return false
	}

	for _, entry := range entries {
		name := entry.Name()
		if strings.HasSuffix(name, ".go") && !strings.HasSuffix(name, "_test.go") {
			return true
		}
	}

	return false
}

// compareVersions compares two version strings (e.g., "1.8" vs "1.0")
// Returns: -1 if v1 < v2, 0 if v1 == v2, 1 if v1 > v2
func compareVersions(v1, v2 string) int {
	parts1 := strings.Split(v1, ".")
	parts2 := strings.Split(v2, ".")

	maxLen := len(parts1)
	if len(parts2) > maxLen {
		maxLen = len(parts2)
	}

	for i := 0; i < maxLen; i++ {
		var n1, n2 int
		if i < len(parts1) {
			fmt.Sscanf(parts1[i], "%d", &n1)
		}
		if i < len(parts2) {
			fmt.Sscanf(parts2[i], "%d", &n2)
		}

		if n1 < n2 {
			return -1
		}
		if n1 > n2 {
			return 1
		}
	}

	return 0
}

// updateTemplateFiles updates version references in template files
func updateTemplateFiles(projectRoot, version string) error {
	// Update functions.md.tmpl
	functionsTemplate := filepath.Join(projectRoot, "templates", "functions.md.tmpl")
	if err := updateVersionInFile(functionsTemplate, version); err != nil {
		return fmt.Errorf("updating functions template: %w", err)
	}

	// Update index.md.tmpl if it has version references
	indexTemplate := filepath.Join(projectRoot, "templates", "index.md.tmpl")
	if _, err := os.Stat(indexTemplate); err == nil {
		if err := updateVersionInFile(indexTemplate, version); err != nil {
			return fmt.Errorf("updating index template: %w", err)
		}
	}

	return nil
}

// updateVersionInFile updates Terraform version references in a file
func updateVersionInFile(filePath, version string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	original := string(content)
	updated := original

	// Pattern 1: "Terraform X.Y.Z or later" or "Terraform X.Y or later"
	re1 := regexp.MustCompile(`(Terraform\s+)\d+\.\d+(?:\.\d+)?(\s+or later)`)
	updated = re1.ReplaceAllString(updated, "${1}"+version+"${2}")

	// Pattern 2: "requires Terraform X.Y.Z" or "requires Terraform X.Y"
	re2 := regexp.MustCompile(`(requires\s+Terraform\s+)\d+\.\d+(?:\.\d+)?`)
	updated = re2.ReplaceAllString(updated, "${1}"+version)

	// Pattern 3: ">= X.Y" or ">= X.Y.Z" version constraints in table contexts
	// Format: "| terraform | >= X.Y |"
	if strings.Contains(updated, "| terraform |") {
		re3Table := regexp.MustCompile(`(\|\s*terraform\s*\|\s*>=\s*)\d+\.\d+(?:\.\d+)?(\s*\|)`)
		updated = re3Table.ReplaceAllString(updated, "${1}"+version+"${2}")
	}

	// Pattern 4: "Terraform X.Y or later" in notes (without "requires")
	re4 := regexp.MustCompile(`(Terraform\s+)\d+\.\d+(?:\.\d+)?(\s+or\s+later)`)
	updated = re4.ReplaceAllString(updated, "${1}"+version+"${2}")

	if updated != original {
		if err := os.WriteFile(filePath, []byte(updated), 0644); err != nil {
			return err
		}
		fmt.Printf("Updated: %s\n", filePath)
	} else {
		fmt.Printf("No changes needed: %s\n", filePath)
	}

	return nil
}
