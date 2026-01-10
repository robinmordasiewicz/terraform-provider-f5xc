// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package naming provides consistent case conversion and acronym handling
// for code generation tools in the F5XC Terraform provider.
package naming

import (
	"regexp"
	"strings"
)

// ExampleNameReplacements defines exact match replacements for example names
// found in F5XC OpenAPI specifications. These patterns should be normalized
// to follow the provider's "example-" naming convention.
// Key: original pattern (lowercase for case-insensitive matching)
// Value: replacement pattern
var ExampleNameReplacements = map[string]string{
	// From OpenAPI specs in docs/specifications/api/
	"my-bucket":     "example-bucket",
	"my-log-bucket": "example-log-bucket",
	"my-resources":  "example-resources",
	"my-tsig-key":   "example-tsig-key",
	"my-website":    "example-website",
	// Additional common patterns
	"my-file":      "example-file",
	"my-key":       "example-key",
	"my-ns":        "example-ns",
	"my-namespace": "example-namespace",
	"my-app":       "example-app",
	"my-config":    "example-config",
	"my-secret":    "example-secret",
	"my-cert":      "example-cert",
	"my-policy":    "example-policy",
	"my-pool":      "example-pool",
	"my-origin":    "example-origin",
	"my-site":      "example-site",
	"my-cluster":   "example-cluster",
}

// ExampleDomainReplacements defines domain patterns to normalize.
// These appear in URL examples within OpenAPI specifications.
var ExampleDomainReplacements = map[string]string{
	"my.example.com":   "example.example.com",
	"my.tenant.domain": "example.tenant.domain",
}

// examplePrefixPattern matches "my-" or "my_" followed by alphanumeric/hyphen/underscore chars.
// This is used as a fallback for patterns not in the exact replacement maps.
var examplePrefixPattern = regexp.MustCompile(`(?i)\bmy[-_]([a-zA-Z0-9_-]+)`)

// NormalizeExampleNames transforms example values from F5 internal conventions
// ("my-*") to the provider standard convention ("example-*").
// This function is idempotent - running it multiple times produces the same result.
//
// The function applies transformations in order:
// 1. Domain replacements (my.example.com -> example.example.com)
// 2. Exact name replacements (my-bucket -> example-bucket)
// 3. Prefix-based replacements for patterns not in exact matches
func NormalizeExampleNames(text string) string {
	if text == "" {
		return text
	}

	// Apply domain replacements first (most specific)
	for original, replacement := range ExampleDomainReplacements {
		// Case-insensitive replacement
		pattern := regexp.MustCompile(`(?i)` + regexp.QuoteMeta(original))
		text = pattern.ReplaceAllString(text, replacement)
	}

	// Apply exact name replacements (case-insensitive, word boundary)
	for original, replacement := range ExampleNameReplacements {
		// Match at word boundaries to avoid partial replacements
		pattern := regexp.MustCompile(`(?i)\b` + regexp.QuoteMeta(original) + `\b`)
		text = pattern.ReplaceAllString(text, replacement)
	}

	// Apply prefix-based replacement for patterns not already handled
	// This catches patterns like "my-custom-name" -> "example-custom-name"
	text = examplePrefixPattern.ReplaceAllStringFunc(text, func(match string) string {
		// Check if this pattern was already replaced by exact match
		lower := strings.ToLower(match)
		if _, exists := ExampleNameReplacements[lower]; exists {
			return match // Already handled by exact replacement
		}

		// Determine the separator used (- or _)
		var sep string
		if strings.Contains(match, "_") {
			sep = "_"
		} else {
			sep = "-"
		}

		// Extract the suffix after my- or my_
		parts := strings.SplitN(lower, sep, 2)
		if len(parts) == 2 {
			return "example" + sep + parts[1]
		}
		return match
	})

	return text
}

// NormalizeExampleInDescription applies example name normalization to description text.
// This is a convenience wrapper for use in documentation generation.
func NormalizeExampleInDescription(desc string) string {
	return NormalizeExampleNames(desc)
}

// NormalizeXVesExample normalizes an x-ves-example value from the OpenAPI spec.
// This specifically handles the example values embedded in the specification.
func NormalizeXVesExample(example string) string {
	return NormalizeExampleNames(example)
}
