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

// subcategoryMap assigns resources to logical categories for better organization
var subcategoryMap = map[string]string{
	// Sites and Infrastructure
	"aws_vpc_site": "Sites", "azure_vnet_site": "Sites", "gcp_vpc_site": "Sites",
	"aws_tgw_site": "Sites", "securemesh_site_v2": "Sites", "virtual_site": "Sites",
	"fleet": "Sites", "k8s_cluster": "Sites",

	// Load Balancing
	"http_loadbalancer": "Load Balancing", "tcp_loadbalancer": "Load Balancing",
	"cdn_loadbalancer": "Load Balancing", "dns_load_balancer": "Load Balancing",
	"origin_pool": "Load Balancing", "healthcheck": "Load Balancing",
	"endpoint": "Load Balancing", "cluster": "Load Balancing",

	// Security
	"app_firewall": "Security", "enhanced_firewall_policy": "Security",
	"network_policy": "Security", "network_firewall": "Security",
	"forward_proxy_policy": "Security", "service_policy": "Security",
	"fast_acl": "Security", "fast_acl_rule": "Security", "rate_limiter": "Security",
	"malicious_user_mitigation": "Security", "bot_defense_app_infrastructure": "Security",

	// Networking
	"network_connector": "Networking", "network_interface": "Networking",
	"virtual_network": "Networking", "nat_policy": "Networking",
	"bgp": "Networking", "bgp_asn_set": "Networking", "bgp_routing_policy": "Networking",
	"ip_prefix_set": "Networking", "cloud_link": "Networking", "cloud_connect": "Networking",

	// DNS
	"dns_zone": "DNS", "dns_domain": "DNS", "dns_lb_pool": "DNS",
	"dns_lb_health_check": "DNS",

	// Authentication & Credentials
	"api_credential": "Authentication", "cloud_credentials": "Authentication",
	"authentication": "Authentication", "oidc_provider": "Authentication",
	"certificate": "Certificates", "certificate_chain": "Certificates", "crl": "Certificates",

	// API Security
	"api_definition": "API Security", "api_discovery": "API Security",
	"app_api_group": "API Security", "api_crawler": "API Security",
	"api_testing": "API Security",

	// Monitoring & Logging
	"log_receiver": "Monitoring", "global_log_receiver": "Monitoring",
	"alert_policy": "Monitoring", "alert_receiver": "Monitoring",

	// Namespace & Organization
	"namespace": "Organization", "managed_tenant": "Organization",
	"child_tenant": "Organization",
}

// getSubcategory returns the subcategory for a resource based on filename
func getSubcategory(filename string) string {
	base := filepath.Base(filename)
	name := strings.TrimSuffix(base, ".md")

	if cat, ok := subcategoryMap[name]; ok {
		return cat
	}
	return "" // Default empty subcategory for uncategorized resources
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
