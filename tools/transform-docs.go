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

// metadataFields defines the standard F5 XC metadata fields that should be grouped
// under "Metadata Argument Reference" section, following Volterra provider conventions
var metadataFields = map[string]bool{
	"annotations": true,
	"description": true,
	"disable":     true,
	"labels":      true,
	"name":        true,
	"namespace":   true,
}

// uppercaseAcronyms defines acronyms that should always be uppercase
// Based on RFC 4949, IEEE standards, and industry style guides (Google, Microsoft, Apple)
var uppercaseAcronyms = map[string]bool{
	// Networking protocols
	"DNS": true, "HTTP": true, "HTTPS": true, "TCP": true, "UDP": true,
	"TLS": true, "SSL": true, "SSH": true, "FTP": true, "SFTP": true,
	"SMTP": true, "IMAP": true, "POP": true, "LDAP": true, "DHCP": true,
	"ARP": true, "ICMP": true, "SNMP": true, "NTP": true, "SIP": true,
	"RTP": true, "RTSP": true, "QUIC": true, "IP": true, "GRPC": true,
	// Web/API
	"API": true, "URL": true, "URI": true, "REST": true, "SOAP": true,
	"JSON": true, "XML": true, "HTML": true, "CSS": true, "CORS": true,
	"CDN": true, "WAF": true, "JWT": true, "SAML": true,
	// Network infrastructure
	"VPN": true, "NAT": true, "VLAN": true, "BGP": true, "OSPF": true,
	"QOS": true, "MTU": true, "TTL": true, "ACL": true, "CIDR": true,
	"VIP": true, "LB": true, "HA": true, "DR": true,
	// Security
	"PKI": true, "CA": true, "CSR": true, "CRL": true, "OCSP": true,
	"PEM": true, "AES": true, "RSA": true, "SHA": true, "MD5": true,
	"HMAC": true, "MFA": true, "SSO": true, "RBAC": true, "IAM": true,
	"DDOS": true, "DOS": true, "XSS": true, "CSRF": true, "SQL": true,
	// Cloud/Infrastructure
	"AWS": true, "GCP": true, "CPU": true, "RAM": true, "SSD": true,
	"HDD": true, "GPU": true, "RAID": true, "VM": true, "OS": true,
	"SLA": true, "RPO": true, "RTO": true,
	// F5/Volterra specific
	"RE": true, "CE": true, "SPO": true, "SMG": true,
}

// mixedCaseAcronyms defines acronyms with specific mixed-case conventions
// These should be preserved exactly as specified
var mixedCaseAcronyms = map[string]string{
	"mtls":      "mTLS",
	"oauth":     "OAuth",
	"graphql":   "GraphQL",
	"websocket": "WebSocket",
	"iscsi":     "iSCSI",
	"ipv4":      "IPv4",
	"ipv6":      "IPv6",
	"macos":     "macOS",
	"ios":       "iOS",
	"nosql":     "NoSQL",
}

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

	// Process index.md for markdown compliance
	if err := transformIndexDoc("docs/index.md"); err != nil {
		fmt.Fprintf(os.Stderr, "Error transforming docs/index.md: %v\n", err)
	} else {
		fmt.Printf("Transformed: docs/index.md\n")
	}
}

// transformIndexDoc fixes markdown issues in the provider index documentation
func transformIndexDoc(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	text := string(content)

	// Fix bare URLs by wrapping them in backticks (MD034 compliance)
	// Match URLs that are not already in backticks, links, or angle brackets
	bareURLRegex := regexp.MustCompile(`(\s)(https?://[^\s\)\]\x60<>]+)([.\s])`)
	text = bareURLRegex.ReplaceAllString(text, "$1`$2`$3")

	// Normalize multiple blank lines
	text = normalizeBlankLines(text)

	return os.WriteFile(filePath, []byte(text), 0644)
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
		// Handle both original tfplugindocs format (## Schema) and already-transformed format (## Argument Reference)
		if strings.HasPrefix(line, "## Schema") || strings.HasPrefix(line, "## Argument Reference") {
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

	// Collect all attributes first, then group them
	type attrInfo struct {
		name          string
		typeInfo      string
		desc          string
		reqStr        string
		hasNested     bool
		oneOfConstraint string
		rawLine       string
	}

	var metadataAttrs []attrInfo
	var specAttrs []attrInfo
	var readOnlyAttrs []attrInfo
	var oneOfConstraints []string // Collect constraints to print before sections

	inSection := ""
	// Match both tfplugindocs format: `- \`name\` (Type) desc`
	// And already-transformed format: `\`name\` - (Required/Optional) desc`
	attrRegex := regexp.MustCompile(`^- \x60([^\x60]+)\x60 \(([^)]+)\)(.*)$`)
	transformedAttrRegex := regexp.MustCompile("^`([^`]+)` - \\((Required|Optional)\\) (.*)$")

	for i := schemaStart + 1; i < mainSchemaEnd; i++ {
		line := lines[i]

		// Collect OneOf constraints from h6 headings
		if strings.HasPrefix(line, "###### One of") {
			oneOfH6Regex := regexp.MustCompile(`###### One of the arguments from this list "([^"]+)" must be set`)
			if m := oneOfH6Regex.FindStringSubmatch(line); m != nil {
				oneOfConstraints = append(oneOfConstraints, m[1])
			}
			continue
		}

		// Track sections - handle both original tfplugindocs format and already-transformed format
		if strings.HasPrefix(line, "### Required") || strings.HasPrefix(line, "The following arguments are required:") {
			inSection = "required"
			continue
		}
		if strings.HasPrefix(line, "### Optional") || strings.HasPrefix(line, "The following arguments are optional:") {
			inSection = "optional"
			continue
		}
		// Handle already-transformed Metadata/Spec sections
		if strings.HasPrefix(line, "### Metadata Argument Reference") {
			inSection = "optional" // Metadata fields are typically optional except name
			continue
		}
		if strings.HasPrefix(line, "### Spec Argument Reference") {
			inSection = "optional"
			continue
		}
		if strings.HasPrefix(line, "### Read-Only") || strings.HasPrefix(line, "### Attributes Reference") || strings.HasPrefix(line, "In addition to all arguments above") {
			inSection = "readonly"
			continue
		}

		// Parse attribute lines - try both formats
		var name, typeInfo, desc, reqStr string
		var hasNested bool
		var matched bool

		if matches := attrRegex.FindStringSubmatch(line); matches != nil {
			// Original tfplugindocs format: `- \`name\` (Type) desc`
			name = matches[1]
			typeInfo = matches[2]
			desc = strings.TrimSpace(matches[3])
			hasNested = strings.Contains(matches[3], "see [below for nested schema]")
			reqStr = "(Optional)"
			if inSection == "required" {
				reqStr = "(Required)"
			}
			matched = true
		} else if matches := transformedAttrRegex.FindStringSubmatch(line); matches != nil {
			// Already transformed format: `\`name\` - (Required/Optional) desc`
			name = matches[1]
			reqStr = "(" + matches[2] + ")"
			desc = strings.TrimSpace(matches[3])
			// Determine type from description or default to String
			typeInfo = "String"
			if strings.Contains(desc, "(`Block`)") || strings.Contains(desc, "See [") {
				typeInfo = "Block"
				hasNested = strings.Contains(desc, "See [") && strings.Contains(desc, "below")
			} else if strings.Contains(desc, "(`List`)") {
				typeInfo = "List"
			} else if strings.Contains(desc, "(`Map`)") {
				typeInfo = "Map"
			} else if strings.Contains(desc, "(`Bool`)") {
				typeInfo = "Bool"
			} else if strings.Contains(desc, "(`Number`)") {
				typeInfo = "Number"
			}
			// Clean the type annotation from desc
			desc = regexp.MustCompile(`\s*\(\x60(String|Bool|Number|List|Map|Block|Set)\x60\)\.?$`).ReplaceAllString(desc, "")
			desc = strings.TrimSuffix(desc, ".")
			matched = true
		}

		if matched {
			// Extract OneOf constraint if present
			oneOfConstraint := extractOneOfConstraint(desc)
			desc = removeOneOfConstraint(desc)

			// Clean up description
			desc = cleanDescription(desc, name)

			attr := attrInfo{
				name:            name,
				typeInfo:        typeInfo,
				desc:            desc,
				reqStr:          reqStr,
				hasNested:       hasNested,
				oneOfConstraint: oneOfConstraint,
				rawLine:         line,
			}

			// Categorize attribute
			if inSection == "readonly" {
				readOnlyAttrs = append(readOnlyAttrs, attr)
			} else if metadataFields[name] {
				metadataAttrs = append(metadataAttrs, attr)
			} else {
				specAttrs = append(specAttrs, attr)
			}
		}
	}

	// Helper function to write attribute
	printedConstraints := make(map[string]bool)
	writeAttr := func(attr attrInfo) {
		typeStr := extractSimpleType(attr.typeInfo)

		// Write OneOf constraint hint if present and not already printed
		constraintKey := normalizeOneOfKey(attr.oneOfConstraint)
		if attr.oneOfConstraint != "" && !printedConstraints[constraintKey] {
			output.WriteString(fmt.Sprintf("> **Note:** One of the arguments from this list \"%s\" must be set.\n\n", attr.oneOfConstraint))
			printedConstraints[constraintKey] = true
		}

		// Clean any existing "See [X](#x) below" references from description to avoid duplication
		desc := attr.desc
		// Simple approach: remove any "See [...](...) below..." patterns
		seeRefRegex := regexp.MustCompile(`See \[.+?\]\(#.+?\) below[^.]*\.?\s*`)
		desc = seeRefRegex.ReplaceAllString(desc, "")
		desc = strings.TrimSpace(desc)
		desc = strings.TrimSuffix(desc, ".")

		if attr.hasNested {
			anchorName := toAnchorName(attr.name)
			if desc != "" {
				output.WriteString(fmt.Sprintf("`%s` - %s %s. See [%s](#%s) below for details.\n\n",
					attr.name, attr.reqStr, desc, toTitleCase(attr.name), anchorName))
			} else {
				output.WriteString(fmt.Sprintf("`%s` - %s See [%s](#%s) below for details.\n\n",
					attr.name, attr.reqStr, toTitleCase(attr.name), anchorName))
			}
		} else {
			if desc != "" {
				output.WriteString(fmt.Sprintf("`%s` - %s %s (`%s`).\n\n", attr.name, attr.reqStr, desc, typeStr))
			} else {
				output.WriteString(fmt.Sprintf("`%s` - %s (`%s`).\n\n", attr.name, attr.reqStr, typeStr))
			}
		}
	}

	// Write Metadata Argument Reference section (if we have metadata attrs)
	if len(metadataAttrs) > 0 {
		output.WriteString("### Metadata Argument Reference\n\n")
		for _, attr := range metadataAttrs {
			writeAttr(attr)
		}
	}

	// Write Spec Argument Reference section (if we have spec attrs)
	if len(specAttrs) > 0 {
		output.WriteString("### Spec Argument Reference\n\n")
		// Write any collected OneOf constraints at the start of spec section
		for _, constraint := range oneOfConstraints {
			constraintKey := normalizeOneOfKey(constraint)
			if !printedConstraints[constraintKey] {
				output.WriteString(fmt.Sprintf("> **Note:** One of the arguments from this list \"%s\" must be set.\n\n", constraint))
				printedConstraints[constraintKey] = true
			}
		}
		for _, attr := range specAttrs {
			writeAttr(attr)
		}
	}

	// Write Read-Only/Attributes Reference section
	if len(readOnlyAttrs) > 0 {
		output.WriteString("\n### Attributes Reference\n\n")
		output.WriteString("In addition to all arguments above, the following attributes are exported:\n\n")
		for _, attr := range readOnlyAttrs {
			writeAttr(attr)
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
					// Add blank line before anchor for proper markdown formatting
					output.WriteString(fmt.Sprintf("\n<a id=\"%s\"></a>\n\n", anchorID))
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

			// Transform h6 OneOf headings to blockquotes (MD001 compliance) in nested blocks
			if strings.HasPrefix(line, "###### One of") {
				oneOfH6Regex := regexp.MustCompile(`###### One of the arguments from this list "([^"]+)" must be set`)
				if m := oneOfH6Regex.FindStringSubmatch(line); m != nil {
					output.WriteString(fmt.Sprintf("> **Note:** One of the arguments from this list \"%s\" must be set.\n\n", m[1]))
				}
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

	// Normalize multiple consecutive blank lines to single blank lines
	result := normalizeBlankLines(output.String())

	// Final pass: fix any remaining bare URLs not in backticks (MD034 compliance)
	result = fixBareURLs(result)

	return os.WriteFile(filePath, []byte(result), 0644)
}

// fixBareURLs wraps bare URLs in backticks for MD034 compliance
func fixBareURLs(content string) string {
	// Fix incomplete backtick patterns where URL has opening backtick but no closing one
	// Pattern: `https://... (no closing backtick before space or end)
	incompleteBacktickRegex := regexp.MustCompile("`(https?://[^`\\s]+)(\\s)")
	content = incompleteBacktickRegex.ReplaceAllString(content, "`$1`$2")

	// Match URLs not already wrapped in backticks, angle brackets, or parentheses (markdown links)
	// This handles URLs that appear mid-line or at specific positions
	bareURLRegex := regexp.MustCompile("([^`\\(<])\\b(https?://[^\\s\\)\\]`<>]+)")
	content = bareURLRegex.ReplaceAllString(content, "$1`$2`")

	// Also fix www. patterns
	wwwRegex := regexp.MustCompile("([^`\\(<])\\b(www\\.[^\\s\\)\\]`<>]+)")
	content = wwwRegex.ReplaceAllString(content, "$1`$2`")

	return content
}

// normalizeBlankLines removes multiple consecutive blank lines
func normalizeBlankLines(content string) string {
	// Replace 3+ consecutive newlines with 2 (single blank line)
	multiBlankRegex := regexp.MustCompile(`\n{3,}`)
	content = multiBlankRegex.ReplaceAllString(content, "\n\n")

	// Ensure file ends with single newline
	content = strings.TrimRight(content, "\n") + "\n"

	return content
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

	// Fix bare URLs by wrapping in backticks (MD034 compliance)
	// Match URLs not already in backticks
	bareURLRegex := regexp.MustCompile(`(^|[^` + "`" + `])(https?://[^\s\)\]\x60<>]+)`)
	desc = bareURLRegex.ReplaceAllString(desc, "$1`$2`")

	// Fix patterns like www.foo.com that look like URLs
	wwwRegex := regexp.MustCompile(`(^|[^` + "`" + `])(www\.[^\s\)\]\x60<>]+)`)
	desc = wwwRegex.ReplaceAllString(desc, "$1`$2`")

	// Escape false-positive reference link patterns (MD052 compliance)
	// Patterns like [+][country code] or [0-9][smhd] look like markdown reference links but aren't
	// Escape them by adding backslash before the second opening bracket
	falsePosRefRegex := regexp.MustCompile(`(\[[^\]]+\])(\[[^\]]+\])`)
	desc = falsePosRefRegex.ReplaceAllString(desc, "$1\\$2")

	// Clean up whitespace
	desc = strings.TrimSpace(desc)

	// Normalize multiple spaces
	desc = regexp.MustCompile(`\s+`).ReplaceAllString(desc, " ")

	// Remove trailing period (will be added by the output formatting)
	desc = strings.TrimSuffix(desc, ".")

	// Normalize acronym capitalization (e.g., Dns → DNS, Http → HTTP)
	desc = normalizeAcronyms(desc)

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
			// Apply standard title case first
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	result := strings.Join(words, " ")

	// Apply acronym normalization to fix any acronyms
	result = normalizeAcronyms(result)

	return result
}

// normalizeAcronyms corrects acronym capitalization in text
// This function is idempotent - running it multiple times produces the same result
func normalizeAcronyms(text string) string {
	// Build a regex pattern for word boundaries
	// Process each word in the text
	wordRegex := regexp.MustCompile(`\b([A-Za-z0-9]+)\b`)

	return wordRegex.ReplaceAllStringFunc(text, func(word string) string {
		upperWord := strings.ToUpper(word)
		lowerWord := strings.ToLower(word)

		// Check for mixed-case acronyms first (e.g., mTLS, OAuth)
		if corrected, ok := mixedCaseAcronyms[lowerWord]; ok {
			return corrected
		}

		// Check for uppercase acronyms (e.g., DNS, HTTP, TCP)
		if uppercaseAcronyms[upperWord] {
			return upperWord
		}

		// Return original word unchanged
		return word
	})
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
