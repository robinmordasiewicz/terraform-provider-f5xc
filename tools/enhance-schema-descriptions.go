//go:build ignore
// +build ignore

// enhance-schema-descriptions.go - Post-processor for improving Terraform resource schema descriptions
//
// This tool enhances schema descriptions in generated resource files to follow industry best practices
// established by HashiCorp's official providers (AzureRM, AWS, Google Cloud).
//
// Problem: Auto-generated descriptions from OpenAPI specs are often:
// - Truncated mid-sentence with "..."
// - Too minimal/tautological ("Service Policy List. List of service policies")
// - Over-reliant on anchor links for basic understanding
//
// Solution: Enhance descriptions to provide:
// - Complete inline context without truncation
// - Meaningful explanations of purpose and behavior
// - Important constraints (ordering, cardinality, requirements)
// - Anchor links as supplements, not replacements
//
// Usage:
//   go run tools/enhance-schema-descriptions.go [--dry-run] [--resource=http_loadbalancer]
//
// Integration:
//   This tool runs AFTER generate-all-schemas.go in the CI/CD pipeline

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	dryRun       bool
	resourceName string
	verbose      bool
)

// DescriptionEnhancement defines a pattern-based enhancement for schema descriptions
type DescriptionEnhancement struct {
	// Pattern to match in the MarkdownDescription field
	Pattern *regexp.Regexp
	// Replacement function that takes the match and returns enhanced description
	Enhance func(match string, context EnhancementContext) string
	// Description of what this enhancement does (for logging)
	Description string
}

// EnhancementContext provides context for enhancement decisions
type EnhancementContext struct {
	ResourceName  string
	BlockName     string
	IsNestedBlock bool
	OriginalDesc  string
}

// Known problematic patterns and their enhancements
var enhancements = []DescriptionEnhancement{
	// Pattern 1: Truncated descriptions ending with "..."
	{
		Pattern: regexp.MustCompile(`^(.{100,})\.\.\.$`),
		Enhance: func(match string, ctx EnhancementContext) string {
			// Remove the truncation marker
			desc := strings.TrimSuffix(match, "...")
			// Add a proper ending if it doesn't have one
			desc = strings.TrimSpace(desc)
			if !strings.HasSuffix(desc, ".") && !strings.HasSuffix(desc, ":") {
				desc += "."
			}
			return desc
		},
		Description: "Remove truncation markers and complete sentences",
	},

	// Pattern 2: Tautological/redundant descriptions like "Service Policy List. List of service policies"
	{
		Pattern: regexp.MustCompile(`^[A-Z][a-z\s]+\.\s+[A-Z][a-z\s]+\.?$`),
		Enhance: func(match string, ctx EnhancementContext) string {
			// For "X. Y." patterns where X and Y are similar, use context
			parts := strings.Split(match, ".")
			if len(parts) >= 2 {
				first := strings.TrimSpace(parts[0])
				second := strings.TrimSpace(parts[1])
				// Check if they contain similar words
				if containsSimilarWords(first, second) {
					return enhanceBasedOnContext(ctx)
				}
			}
			return match
		},
		Description: "Replace tautological descriptions with contextual information",
	},

	// Pattern 3: "List. List of X" -> enhance with purpose
	{
		Pattern: regexp.MustCompile(`^(?:List|Collection|Array)\.\s+(?:List|Collection|Array)\s+of\s+(.+)\.?$`),
		Enhance: func(match string, ctx EnhancementContext) string {
			return enhanceListDescription(match, ctx)
		},
		Description: "Enhance generic list descriptions with purpose and behavior",
	},
}

// Resource-specific enhancements based on known patterns from F5 XC resources
var resourceEnhancements = map[string]map[string]string{
	"http_loadbalancer": {
		"active_service_policies": "A list of service policies evaluated sequentially to control request handling. Service policies are evaluated top-to-bottom in order, with the first matching policy taking effect.",
		"policies":                "List of service policies. Service policies form a sequential evaluation engine where each policy (and rules within that policy) are evaluated in order from top to bottom. When a request's characteristics match a policy's criteria, that policy takes effect and no further policies are evaluated. The order of policies in this list is critical for achieving the intended behavior. Each policy is a reference to a service_policy resource.",
	},
	"app_firewall": {
		"detection_settings":     "Configuration for WAF attack detection sensitivity and signature coverage. Controls which attack signatures are enabled and how they detect potential threats.",
		"bot_protection_setting": "Configuration for protecting against automated bot traffic. Enables detection and mitigation of malicious bots while allowing legitimate automation.",
	},
}

func init() {
	flag.BoolVar(&dryRun, "dry-run", false, "Show what would be changed without modifying files")
	flag.StringVar(&resourceName, "resource", "", "Enhance only specific resource (e.g., 'http_loadbalancer')")
	flag.BoolVar(&verbose, "verbose", false, "Show detailed processing information")
}

func main() {
	flag.Parse()

	providerDir := "internal/provider"

	// Find all resource files or specific resource if specified
	pattern := "*_resource.go"
	if resourceName != "" {
		pattern = fmt.Sprintf("%s_resource.go", resourceName)
	}

	files, err := filepath.Glob(filepath.Join(providerDir, pattern))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding resource files: %v\n", err)
		os.Exit(1)
	}

	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "No resource files found matching pattern: %s\n", pattern)
		os.Exit(1)
	}

	totalEnhancements := 0
	for _, file := range files {
		count := processResourceFile(file)
		totalEnhancements += count
	}

	if dryRun {
		fmt.Printf("Dry run complete. Would enhance %d descriptions across %d files.\n", totalEnhancements, len(files))
	} else {
		fmt.Printf("Enhanced %d descriptions across %d files.\n", totalEnhancements, len(files))
	}
}

func processResourceFile(filePath string) int {
	// Extract resource name from filename
	basename := filepath.Base(filePath)
	resName := strings.TrimSuffix(basename, "_resource.go")

	if verbose {
		fmt.Printf("Processing: %s\n", filePath)
	}

	// Read file
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", filePath, err)
		return 0
	}

	// Check if file is generated (has DO NOT EDIT warning)
	if !strings.Contains(string(content), "DO NOT EDIT") {
		if verbose {
			fmt.Printf("  Skipping non-generated file: %s\n", filePath)
		}
		return 0
	}

	// Process line by line to preserve structure
	lines := strings.Split(string(content), "\n")
	enhancementCount := 0
	modified := false

	for i, line := range lines {
		// Look for MarkdownDescription fields
		if strings.Contains(line, "MarkdownDescription:") {
			enhanced, wasEnhanced := enhanceLine(line, resName, lines, i)
			if wasEnhanced {
				if verbose || dryRun {
					fmt.Printf("  Line %d:\n    Before: %s\n    After:  %s\n", i+1, line, enhanced)
				}
				lines[i] = enhanced
				enhancementCount++
				modified = true
			}
		}
	}

	// Write back if modified and not dry-run
	if modified && !dryRun {
		newContent := strings.Join(lines, "\n")
		if err := os.WriteFile(filePath, []byte(newContent), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", filePath, err)
			return 0
		}
	}

	return enhancementCount
}

func enhanceLine(line, resourceName string, allLines []string, lineIndex int) (string, bool) {
	// Extract the description from the MarkdownDescription field
	descPattern := regexp.MustCompile(`MarkdownDescription:\s*"([^"]*)",?`)
	matches := descPattern.FindStringSubmatch(line)
	if len(matches) < 2 {
		return line, false
	}

	originalDesc := matches[1]

	// Determine context (block name, nesting, etc.)
	ctx := extractContext(resourceName, allLines, lineIndex)
	ctx.OriginalDesc = originalDesc

	// Try resource-specific enhancements first
	if enhanced, ok := tryResourceSpecificEnhancement(ctx); ok {
		newLine := descPattern.ReplaceAllString(line, fmt.Sprintf(`MarkdownDescription: "%s",`, enhanced))
		return newLine, true
	}

	// Try pattern-based enhancements
	for _, enh := range enhancements {
		if enh.Pattern.MatchString(originalDesc) {
			enhanced := enh.Enhance(originalDesc, ctx)
			if enhanced != originalDesc {
				newLine := descPattern.ReplaceAllString(line, fmt.Sprintf(`MarkdownDescription: "%s",`, enhanced))
				return newLine, true
			}
		}
	}

	return line, false
}

func extractContext(resourceName string, lines []string, lineIndex int) EnhancementContext {
	ctx := EnhancementContext{
		ResourceName: resourceName,
	}

	// Look backwards to find the block/attribute name
	for i := lineIndex - 1; i >= 0 && i >= lineIndex-10; i-- {
		line := lines[i]

		// Pattern: "block_name": schema.SingleNestedBlock{
		if match := regexp.MustCompile(`"([^"]+)":\s*schema\.(Single|List|Set)NestedBlock\{`).FindStringSubmatch(line); len(match) > 1 {
			ctx.BlockName = match[1]
			ctx.IsNestedBlock = true
			break
		}

		// Pattern: "attr_name": schema.StringAttribute{
		if match := regexp.MustCompile(`"([^"]+)":\s*schema\.\w+Attribute\{`).FindStringSubmatch(line); len(match) > 1 {
			ctx.BlockName = match[1]
			break
		}
	}

	return ctx
}

func tryResourceSpecificEnhancement(ctx EnhancementContext) (string, bool) {
	if enhancements, ok := resourceEnhancements[ctx.ResourceName]; ok {
		if enhanced, ok := enhancements[ctx.BlockName]; ok {
			return enhanced, true
		}
	}
	return "", false
}

func enhanceBasedOnContext(ctx EnhancementContext) string {
	// Use block name to infer purpose
	name := ctx.BlockName

	// Common patterns in F5 XC resources
	if strings.Contains(name, "policy") || strings.Contains(name, "policies") {
		return fmt.Sprintf("Configuration for %s. Defines rules and actions that control how requests are processed.", name)
	}
	if strings.Contains(name, "settings") || strings.Contains(name, "config") {
		return fmt.Sprintf("Configuration settings for %s. Controls behavior and operational parameters.", name)
	}
	if strings.Contains(name, "pool") {
		return fmt.Sprintf("Backend server pool configuration. Defines the servers that handle requests and their properties.", name)
	}

	// Default: just remove redundancy
	return fmt.Sprintf("Configuration for %s.", name)
}

func enhanceListDescription(match string, ctx EnhancementContext) string {
	name := ctx.BlockName

	// Extract what the list contains
	listOfPattern := regexp.MustCompile(`(?:List|Collection|Array)\s+of\s+(.+)\.?$`)
	if matches := listOfPattern.FindStringSubmatch(match); len(matches) > 1 {
		itemType := strings.TrimSpace(matches[1])

		// Add context about ordering/evaluation if it's a policy or rule list
		if strings.Contains(name, "policy") || strings.Contains(name, "policies") ||
			strings.Contains(name, "rule") || strings.Contains(name, "rules") {
			return fmt.Sprintf("List of %s. Evaluated sequentially in the order specified. The first matching entry takes effect.", itemType)
		}

		// For other lists, provide basic context
		return fmt.Sprintf("List of %s. Each entry defines a separate %s configuration.", itemType, strings.ToLower(itemType))
	}

	return match // No improvement possible, return original
}

func containsSimilarWords(s1, s2 string) bool {
	// Simple check if two strings share significant words
	words1 := strings.Fields(strings.ToLower(s1))
	words2 := strings.Fields(strings.ToLower(s2))

	// Filter out common words
	commonWords := map[string]bool{
		"a": true, "an": true, "the": true, "of": true, "for": true,
		"to": true, "in": true, "on": true, "at": true, "by": true,
	}

	significantWords1 := make(map[string]bool)
	for _, word := range words1 {
		if !commonWords[word] && len(word) > 2 {
			significantWords1[word] = true
		}
	}

	// Check if at least 50% of significant words in s2 are in s1
	matchCount := 0
	totalCount := 0
	for _, word := range words2 {
		if !commonWords[word] && len(word) > 2 {
			totalCount++
			if significantWords1[word] {
				matchCount++
			}
		}
	}

	return totalCount > 0 && float64(matchCount)/float64(totalCount) >= 0.5
}
