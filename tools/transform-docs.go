//go:build ignore
// +build ignore

// This tool transforms tfplugindocs-generated documentation into a cleaner,
// Volterra-style format with flat argument references instead of deeply nested schemas.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// subcategoryOverrides provides explicit category assignments for resources
// that don't match any pattern or need a specific override
var subcategoryOverrides = map[string]string{
	// Explicit overrides for resources that don't match patterns well
	"apm":               "Monitoring",
	"crl":               "Certificates",
	"bgp":               "Networking",
	"proxy":             "Networking",
	"tunnel":            "Networking",
	"segment":           "Networking",
	"subnet":            "Networking",
	"fleet":             "Sites",
	"cluster":           "Load Balancing",
	"endpoint":          "Load Balancing",
	"route":             "Load Balancing",
	"healthcheck":       "Load Balancing",
	"origin_pool":       "Load Balancing",
	"virtual_host":      "Load Balancing",
	"discovery":         "Applications",
	"filter_set":        "Applications",
	"policer":           "Service Mesh",
	"quota":             "Organization",
	"contact":           "Organization",
	"role":              "Organization",
	"token":             "Authentication",
	"registration":      "Sites",
	"namespace":         "Organization",
	"data_type":         "Security",
	"data_group":        "BIG-IP Integration",
	"irule":             "BIG-IP Integration",
	"nfv_service":       "Networking",
	"workload":          "Kubernetes",
	"workload_flavor":   "Kubernetes",
	"cminstance":          "Subscriptions",
	"user_identification": "Security",
	"virtual_network":     "Networking",
}

// categoryPatterns defines patterns to auto-categorize resources
// Order matters: more specific patterns should come first
var categoryPatterns = []struct {
	pattern  string
	category string
}{
	// Sites - cloud provider sites and site-related
	{"_vpc_site", "Sites"},
	{"_vnet_site", "Sites"},
	{"_tgw_site", "Sites"},
	{"securemesh_site", "Sites"},
	{"voltstack_site", "Sites"},
	{"virtual_site", "Sites"},
	{"site_mesh", "Sites"},

	// Load Balancing
	{"loadbalancer", "Load Balancing"},
	{"cdn_", "Load Balancing"},
	{"advertise_policy", "Load Balancing"},

	// Security - firewall, policies, WAF
	{"firewall", "Security"},
	{"_policy", "Security"},
	{"_acl", "Security"},
	{"rate_limiter", "Security"},
	{"malicious_user", "Security"},
	{"bot_defense", "Security"},
	{"waf_", "Security"},
	{"protocol_", "Security"},
	{"sensitive_data", "Security"},

	// Networking
	{"network_", "Networking"},
	{"bgp_", "Networking"},
	{"ip_prefix", "Networking"},
	{"cloud_link", "Networking"},
	{"cloud_connect", "Networking"},
	{"_connector", "Networking"},
	{"dc_cluster", "Networking"},
	{"srv6_", "Networking"},
	{"forwarding_", "Networking"},
	{"routing", "Networking"},
	{"nat_", "Networking"},

	// DNS
	{"dns_", "DNS"},

	// Authentication & Secrets
	{"credential", "Authentication"},
	{"authentication", "Authentication"},
	{"oidc_", "Authentication"},
	{"secret_", "Authentication"},

	// Certificates
	{"certificate", "Certificates"},
	{"trusted_ca", "Certificates"},

	// API Security
	{"api_", "API Security"},
	{"app_api", "API Security"},

	// Monitoring & Logging
	{"log_receiver", "Monitoring"},
	{"alert_", "Monitoring"},
	{"report_", "Monitoring"},

	// Organization & Tenants
	{"tenant", "Organization"},
	{"_support", "Organization"},
	{"voltshare", "Organization"},
	{"allowed_", "Organization"},

	// Kubernetes
	{"k8s_", "Kubernetes"},
	{"virtual_k8s", "Kubernetes"},
	{"container_registry", "Kubernetes"},

	// VPN & IPSec
	{"ike", "VPN"},

	// Infrastructure Protection (DDoS)
	{"infraprotect_", "Infrastructure Protection"},

	// Applications
	{"app_setting", "Applications"},
	{"app_type", "Applications"},

	// BIG-IP Integration
	{"bigip_", "BIG-IP Integration"},

	// Cloud Resources
	{"cloud_elastic", "Cloud Resources"},
	{"address_allocator", "Cloud Resources"},
	{"geo_location", "Cloud Resources"},

	// Integrations
	{"ticket_tracking", "Integrations"},
	{"code_base", "Integrations"},
	{"tpm_", "Integrations"},

	// Subscriptions
	{"subscription", "Subscriptions"},

	// Service Mesh (catch remaining mesh-related)
	{"usb_policy", "Service Mesh"},
}

// getSubcategory returns the subcategory for a resource based on filename
// Uses a three-tier approach:
// 1. Check explicit overrides first (for exceptions)
// 2. Apply pattern matching (for automatic categorization)
// 3. Fall back to "Other" to ensure ALL resources are categorized
func getSubcategory(filename string) string {
	base := filepath.Base(filename)
	name := strings.TrimSuffix(base, ".md")

	// 1. Check explicit overrides first
	if cat, ok := subcategoryOverrides[name]; ok {
		return cat
	}

	// 2. Apply pattern matching
	for _, p := range categoryPatterns {
		if strings.Contains(name, p.pattern) {
			return p.category
		}
	}

	// 3. Fall back to "Other" - ensures no uncategorized resources in Registry
	// This prevents the mixed layout issue where uncategorized resources
	// appear under a generic "Resources" section
	return "Other"
}

func main() {
	docsDir := "docs/resources"

	files, err := filepath.Glob(filepath.Join(docsDir, "*.md"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding doc files: %v\n", err)
		os.Exit(1)
	}

	for _, file := range files {
		if err := transformDoc(file); err != nil {
			fmt.Fprintf(os.Stderr, "Error transforming %s: %v\n", file, err)
		} else {
			fmt.Printf("Transformed: %s\n", file)
		}
	}

	// Also process data sources
	dataSourceFiles, err := filepath.Glob(filepath.Join("docs/data-sources", "*.md"))
	if err == nil {
		for _, file := range dataSourceFiles {
			if err := transformDoc(file); err != nil {
				fmt.Fprintf(os.Stderr, "Error transforming %s: %v\n", file, err)
			} else {
				fmt.Printf("Transformed: %s\n", file)
			}
		}
	}
}

func transformDoc(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")

	// Find key sections
	schemaStart := -1
	importStart := -1
	firstNestedAnchor := -1

	for i, line := range lines {
		if strings.HasPrefix(line, "## Schema") {
			schemaStart = i
		}
		if strings.HasPrefix(line, "## Import") {
			importStart = i
		}
		if firstNestedAnchor < 0 && strings.HasPrefix(line, "<a id=\"nestedblock--") {
			firstNestedAnchor = i
		}
	}

	if schemaStart == -1 {
		return nil // No schema section
	}

	var output strings.Builder

	// Get subcategory for this resource
	subcategory := getSubcategory(filePath)

	// Write everything before Schema section, updating subcategory if needed
	for i := 0; i < schemaStart; i++ {
		line := lines[i]
		// Replace empty subcategory with assigned category
		if strings.HasPrefix(line, "subcategory:") && subcategory != "" {
			line = fmt.Sprintf("subcategory: \"%s\"", subcategory)
		}
		output.WriteString(line)
		output.WriteString("\n")
	}

	// Process the Schema section
	output.WriteString("## Argument Reference\n\n")

	// Determine end of main schema (either first nested anchor or import)
	mainSchemaEnd := len(lines)
	if firstNestedAnchor > 0 {
		mainSchemaEnd = firstNestedAnchor
	} else if importStart > 0 {
		mainSchemaEnd = importStart
	}

	// Parse main schema attributes
	inSection := ""
	attrRegex := regexp.MustCompile(`^- \x60([^\x60]+)\x60 \(([^)]+)\)(.*)$`)
	printedConstraints := make(map[string]bool) // Track already printed OneOf constraints

	for i := schemaStart + 1; i < mainSchemaEnd; i++ {
		line := lines[i]

		// Track sections
		if strings.HasPrefix(line, "### Required") {
			output.WriteString("The following arguments are required:\n\n")
			inSection = "required"
			continue
		}
		if strings.HasPrefix(line, "### Optional") {
			output.WriteString("\nThe following arguments are optional:\n\n")
			inSection = "optional"
			continue
		}
		if strings.HasPrefix(line, "### Read-Only") {
			output.WriteString("\n### Attributes Reference\n\n")
			output.WriteString("In addition to all arguments above, the following attributes are exported:\n\n")
			inSection = "readonly"
			continue
		}

		// Transform attribute lines
		if matches := attrRegex.FindStringSubmatch(line); matches != nil {
			name := matches[1]
			typeInfo := matches[2]
			desc := strings.TrimSpace(matches[3])

			// Extract OneOf constraint if present
			oneOfConstraint := extractOneOfConstraint(desc)
			desc = removeOneOfConstraint(desc)

			// Clean up description
			desc = cleanDescription(desc, name)

			// Determine if it has nested block
			hasNested := strings.Contains(matches[3], "see [below for nested schema]")

			// Format the attribute line Volterra-style
			reqStr := "(Optional)"
			if inSection == "required" {
				reqStr = "(Required)"
			}

			typeStr := extractSimpleType(typeInfo)

			// Write OneOf constraint hint if present and not already printed (Volterra-style)
			// Normalize the constraint key to handle different orderings of the same fields
			constraintKey := normalizeOneOfKey(oneOfConstraint)
			if oneOfConstraint != "" && !printedConstraints[constraintKey] {
				output.WriteString(fmt.Sprintf("###### One of the arguments from this list \"%s\" must be set\n\n", oneOfConstraint))
				printedConstraints[constraintKey] = true
			}

			if hasNested {
				anchorName := toAnchorName(name)
				if desc != "" {
					output.WriteString(fmt.Sprintf("`%s` - %s %s. See [%s](#%s) below for details.\n\n",
						name, reqStr, desc, toTitleCase(name), anchorName))
				} else {
					output.WriteString(fmt.Sprintf("`%s` - %s See [%s](#%s) below for details.\n\n",
						name, reqStr, toTitleCase(name), anchorName))
				}
			} else {
				if desc != "" {
					output.WriteString(fmt.Sprintf("`%s` - %s %s (`%s`).\n\n", name, reqStr, desc, typeStr))
				} else {
					output.WriteString(fmt.Sprintf("`%s` - %s (`%s`).\n\n", name, reqStr, typeStr))
				}
			}
		}
	}

	// Process nested blocks if present
	if firstNestedAnchor > 0 {
		output.WriteString("\n---\n\n")

		nestedEnd := importStart
		if nestedEnd <= 0 {
			nestedEnd = len(lines)
		}

		// Transform each nested block
		currentBlockName := ""
		inNestedBlock := false

		for i := firstNestedAnchor; i < nestedEnd; i++ {
			line := lines[i]

			// Detect anchor
			if strings.HasPrefix(line, "<a id=\"nestedblock--") {
				// Extract anchor ID
				anchorRegex := regexp.MustCompile(`<a id="(nestedblock--[^"]+)"></a>`)
				if m := anchorRegex.FindStringSubmatch(line); m != nil {
					anchorID := m[1]
					currentBlockName = strings.ReplaceAll(strings.TrimPrefix(anchorID, "nestedblock--"), "--", ".")
					output.WriteString(fmt.Sprintf("<a id=\"%s\"></a>\n", anchorID))
					inNestedBlock = true
				}
				continue
			}

			// Transform nested schema headers
			if strings.HasPrefix(line, "### Nested Schema for") {
				headerRegex := regexp.MustCompile(`### Nested Schema for \x60([^\x60]+)\x60`)
				if m := headerRegex.FindStringSubmatch(line); m != nil {
					displayName := toTitleCase(strings.ReplaceAll(m[1], ".", " "))
					output.WriteString(fmt.Sprintf("### %s\n\n", displayName))
				}
				continue
			}

			// Skip "Optional:" and "Required:" labels
			if line == "Optional:" || line == "Required:" {
				continue
			}

			// Transform attribute lines in nested blocks
			if inNestedBlock && strings.HasPrefix(line, "- `") {
				if matches := attrRegex.FindStringSubmatch(line); matches != nil {
					name := matches[1]
					typeInfo := matches[2]
					desc := strings.TrimSpace(matches[3])

					// Clean description
					desc = cleanDescription(desc, currentBlockName+"--"+name)

					hasNested := strings.Contains(matches[3], "see [below for nested schema]")
					typeStr := extractSimpleType(typeInfo)

					if hasNested {
						// Extract the nested anchor
						nestedAnchorRegex := regexp.MustCompile(`#(nestedblock--[^)]+)`)
						nestedAnchor := ""
						if am := nestedAnchorRegex.FindStringSubmatch(matches[3]); am != nil {
							nestedAnchor = am[1]
						}

						if nestedAnchor != "" && desc != "" {
							output.WriteString(fmt.Sprintf("`%s` - (Optional) %s. See [%s](#%s) below.\n\n",
								name, desc, toTitleCase(name), nestedAnchor))
						} else if nestedAnchor != "" {
							output.WriteString(fmt.Sprintf("`%s` - (Optional) See [%s](#%s) below.\n\n",
								name, toTitleCase(name), nestedAnchor))
						} else {
							output.WriteString(fmt.Sprintf("`%s` - (Optional) %s (`%s`).\n\n", name, desc, typeStr))
						}
					} else {
						if desc != "" {
							output.WriteString(fmt.Sprintf("`%s` - (Optional) %s (`%s`).\n\n", name, desc, typeStr))
						} else {
							output.WriteString(fmt.Sprintf("`%s` - (Optional) (`%s`).\n\n", name, typeStr))
						}
					}
				}
				continue
			}

			// Skip empty lines between blocks
			if line == "" && inNestedBlock {
				continue
			}
		}
	}

	// Write Import section
	if importStart > 0 {
		output.WriteString("\n")
		for i := importStart; i < len(lines); i++ {
			output.WriteString(lines[i])
			output.WriteString("\n")
		}
	}

	return os.WriteFile(filePath, []byte(output.String()), 0644)
}

func cleanDescription(desc, attrPath string) string {
	// Remove nested schema references
	nestedRefRegex := regexp.MustCompile(`\s*\(see \[below for nested schema\]\([^)]+\)\)`)
	desc = nestedRefRegex.ReplaceAllString(desc, "")

	// Remove ves.io.schema validation annotations that may have passed through
	vesSchemaRegex := regexp.MustCompile(`\s*ves\.io\.schema[^\s]*:\s*\S+`)
	desc = vesSchemaRegex.ReplaceAllString(desc, "")

	// Remove "Required: YES" or "Required: NO" annotations
	requiredRegex := regexp.MustCompile(`\s*Required:\s*(YES|NO)\s*`)
	desc = requiredRegex.ReplaceAllString(desc, " ")

	// Remove "Exclusive with [xxx]" patterns
	exclusiveRegex := regexp.MustCompile(`\s*Exclusive with\s*\[[^\]]*\]\s*`)
	desc = exclusiveRegex.ReplaceAllString(desc, " ")

	// Clean up whitespace
	desc = strings.TrimSpace(desc)

	// Normalize multiple spaces
	desc = regexp.MustCompile(`\s+`).ReplaceAllString(desc, " ")

	// Remove trailing period
	desc = strings.TrimSuffix(desc, ".")

	// Truncate very long descriptions
	if len(desc) > 200 {
		// Find last sentence boundary
		lastPeriod := strings.LastIndex(desc[:200], ".")
		if lastPeriod > 50 {
			desc = desc[:lastPeriod]
		} else {
			desc = desc[:197] + "..."
		}
	}

	return desc
}

func extractSimpleType(typeInfo string) string {
	if strings.Contains(typeInfo, "Block") {
		return "Block"
	}
	if strings.Contains(typeInfo, "List of") {
		return "List"
	}
	if strings.Contains(typeInfo, "List") {
		return "List"
	}
	if strings.Contains(typeInfo, "Set of") {
		return "Set"
	}
	if strings.Contains(typeInfo, "Set") {
		return "Set"
	}
	if strings.Contains(typeInfo, "Map of") {
		return "Map"
	}
	if strings.Contains(typeInfo, "Map") {
		return "Map"
	}
	if strings.Contains(typeInfo, "String") {
		return "String"
	}
	if strings.Contains(typeInfo, "Number") {
		return "Number"
	}
	if strings.Contains(typeInfo, "Bool") {
		return "Bool"
	}
	return "String"
}

func toTitleCase(s string) string {
	// Replace underscores and dots with spaces
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, ".", " ")

	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}

func toAnchorName(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, "_", "-"))
}

// extractOneOfConstraint extracts the [OneOf: field1, field2] constraint from description
func extractOneOfConstraint(desc string) string {
	oneOfRegex := regexp.MustCompile(`\[OneOf:\s*([^\]]+)\]`)
	matches := oneOfRegex.FindStringSubmatch(desc)
	if len(matches) >= 2 {
		return strings.TrimSpace(matches[1])
	}
	return ""
}

// removeOneOfConstraint removes the [OneOf: ...] constraint from description
func removeOneOfConstraint(desc string) string {
	oneOfRegex := regexp.MustCompile(`\[OneOf:\s*[^\]]+\]\s*`)
	return strings.TrimSpace(oneOfRegex.ReplaceAllString(desc, ""))
}

// normalizeOneOfKey creates a normalized key for deduplication by sorting fields
func normalizeOneOfKey(constraint string) string {
	if constraint == "" {
		return ""
	}
	fields := strings.Split(constraint, ", ")
	sort.Strings(fields)
	return strings.Join(fields, ", ")
}
