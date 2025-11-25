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

// Terraform Registry limits
// NOTE: Testing confirmed NO hard limits on file size, H2 headers, bold sections, or code blocks.
// The previous truncation issue was caused by complex nesting with H3 headings, NOT count limits.
// Solution: Convert H3 headings to bold text for nested blocks (single-page rendering).
const (
	// File size limit - documents exceeding this will be truncated
	maxDocSizeBytes = 500 * 1024 // 500KB

	// These thresholds are for documentation only - H3→bold conversion prevents truncation
	maxSafeH3Headings = 60
	warnH3Headings    = 50
)

// docWarning tracks documentation files that may have issues
type docWarning struct {
	path         string
	sizeKB       float64
	h3Count      int
	isOversized  bool // exceeds 500KB
	willTruncate bool // has too many H3 headings
}

// NOTE: nestedBlockInfo and guidePageInfo structs removed - single-page rendering mode
// eliminates the need for page splitting functionality

func main() {
	docsDir := "docs/resources"
	var docWarnings []docWarning

	files, err := filepath.Glob(filepath.Join(docsDir, "*.md"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding doc files: %v\n", err)
		os.Exit(1)
	}

	for _, file := range files {
		// Skip _nested_blocks files - they will be deleted since we now use single-page rendering
		if strings.Contains(file, "_nested_blocks") {
			// Delete existing nested_blocks files (cleanup from previous pagination approach)
			if err := os.Remove(file); err == nil {
				fmt.Printf("Removed (single-page mode): %s\n", file)
			}
			continue
		}

		if err := transformDoc(file); err != nil {
			fmt.Fprintf(os.Stderr, "Error transforming %s: %v\n", file, err)
		} else {
			fmt.Printf("Transformed: %s\n", file)
		}

		// Check for potential Registry issues after transformation
		if warn := checkDocLimits(file); warn != nil {
			docWarnings = append(docWarnings, *warn)
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
			// Check for potential Registry issues after transformation
			if warn := checkDocLimits(file); warn != nil {
				docWarnings = append(docWarnings, *warn)
			}
		}
	}

	// Process index.md for markdown compliance
	if err := transformIndexDoc("docs/index.md"); err != nil {
		fmt.Fprintf(os.Stderr, "Error transforming docs/index.md: %v\n", err)
	} else {
		fmt.Printf("Transformed: docs/index.md\n")
	}

	// Report documents with potential issues
	reportDocWarnings(docWarnings)
}

// checkDocLimits validates document against Terraform Registry limits
// Returns docWarning if file exceeds size or heading limits, nil otherwise
func checkDocLimits(filePath string) *docWarning {
	info, err := os.Stat(filePath)
	if err != nil {
		return nil
	}

	// Read file to count headings
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil
	}

	h3Count := 0
	for _, line := range strings.Split(string(content), "\n") {
		if strings.HasPrefix(line, "### ") {
			h3Count++
		}
	}

	isOversized := info.Size() > maxDocSizeBytes
	willTruncate := h3Count > warnH3Headings

	// Only return warning if there's an issue
	if isOversized || willTruncate {
		return &docWarning{
			path:         filePath,
			sizeKB:       float64(info.Size()) / 1024,
			h3Count:      h3Count,
			isOversized:  isOversized,
			willTruncate: willTruncate,
		}
	}

	return nil
}

// reportDocWarnings outputs warnings about documents that may have Registry issues
func reportDocWarnings(warnings []docWarning) {
	if len(warnings) == 0 {
		return
	}

	// Separate by issue type
	var oversized, truncated []docWarning
	for _, w := range warnings {
		if w.isOversized {
			oversized = append(oversized, w)
		}
		if w.willTruncate {
			truncated = append(truncated, w)
		}
	}

	fmt.Fprintf(os.Stderr, "\n")

	// Report oversized files (500KB limit)
	if len(oversized) > 0 {
		fmt.Fprintf(os.Stderr, "⛔ ERROR: %d document(s) exceed Terraform Registry 500KB storage limit:\n", len(oversized))
		for _, doc := range oversized {
			fmt.Fprintf(os.Stderr, "   • %s: %.1fKB\n", doc.path, doc.sizeKB)
		}
		fmt.Fprintf(os.Stderr, "   Reference: https://developer.hashicorp.com/terraform/registry/providers/docs#storage-limits\n\n")
	}

	// Report files that will have rendering truncated
	if len(truncated) > 0 {
		fmt.Fprintf(os.Stderr, "⚠️  WARNING: %d document(s) exceed %d H3 headings (Registry truncates ~65 headings):\n", len(truncated), warnH3Headings)
		fmt.Fprintf(os.Stderr, "   These documents will have content truncated when displayed in the Registry.\n")
		fmt.Fprintf(os.Stderr, "   Dead anchor links for truncated sections have been automatically removed.\n\n")
		for _, doc := range truncated {
			fmt.Fprintf(os.Stderr, "   • %s: %d H3 headings (%.1fKB)\n", doc.path, doc.h3Count, doc.sizeKB)
		}
		fmt.Fprintf(os.Stderr, "\n")
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

// attrInfo holds attribute information for documentation generation
type attrInfo struct {
	name            string
	typeInfo        string
	desc            string
	reqStr          string
	hasNested       bool
	oneOfConstraint string
	rawLine         string
}

// extractOneOfFromGoFile reads the resource Go file and extracts [OneOf: ...] markers
// from MarkdownDescription fields, returning a map of attribute name to constraint
func extractOneOfFromGoFile(mdFilePath string) map[string]string {
	// Convert docs path to resource path
	// Example: docs/resources/http_loadbalancer.md -> internal/provider/http_loadbalancer_resource.go
	baseName := filepath.Base(mdFilePath)
	resourceName := strings.TrimSuffix(baseName, ".md")
	goFilePath := filepath.Join("internal/provider", resourceName+"_resource.go")

	content, err := os.ReadFile(goFilePath)
	if err != nil {
		// If resource file doesn't exist, return empty map
		return make(map[string]string)
	}

	text := string(content)
	constraintMap := make(map[string]string)

	// Find all MarkdownDescription fields with [OneOf: ...] markers
	// Pattern: "attribute_name": schema.SingleNestedBlock{
	//           MarkdownDescription: "[OneOf: attr1, attr2, attr3] description...",
	oneOfRegex := regexp.MustCompile(`"([^"]+)":\s*schema\.\w+\{\s*MarkdownDescription:\s*"\[OneOf:\s*([^\]]+)\]`)
	matches := oneOfRegex.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		if len(match) >= 3 {
			constraintList := match[2]

			// Add this constraint to all attributes mentioned in the list
			attrNames := strings.Split(constraintList, ", ")
			for _, name := range attrNames {
				constraintMap[strings.TrimSpace(name)] = constraintList
			}
		}
	}

	return constraintMap
}

// propagateOneOfConstraints ensures all attributes mentioned in a OneOf constraint
// have that constraint set, even if only one of them has the [OneOf: ...] marker.
// This is necessary because in the schema, only one property typically has the marker.
func propagateOneOfConstraints(attrs *[]attrInfo, constraintMap map[string]string) {
	// Apply the constraints from the Go file to all matching attributes
	for i := range *attrs {
		if constraint, found := constraintMap[(*attrs)[i].name]; found {
			(*attrs)[i].oneOfConstraint = constraint
		}
	}
}

// groupAndSortAttributes groups mutually exclusive attributes (with same OneOf constraint)
// together and sorts them properly for display in documentation.
//
// Algorithm:
// 1. Separate attributes into OneOf groups and standalone attributes
// 2. Sort OneOf groups by the first attribute name in each group (alphabetically)
// 3. Within each OneOf group, maintain the order specified in the constraint
// 4. Merge OneOf groups with standalone attributes in sorted order
func groupAndSortAttributes(attrs []attrInfo) []attrInfo {
	// Group attributes by their normalized OneOf constraint
	oneOfGroups := make(map[string][]attrInfo)
	var standaloneAttrs []attrInfo

	for _, attr := range attrs {
		if attr.oneOfConstraint != "" {
			key := normalizeOneOfKey(attr.oneOfConstraint)
			oneOfGroups[key] = append(oneOfGroups[key], attr)
		} else {
			standaloneAttrs = append(standaloneAttrs, attr)
		}
	}

	// Sort standalone attributes alphabetically
	sort.Slice(standaloneAttrs, func(i, j int) bool {
		return standaloneAttrs[i].name < standaloneAttrs[j].name
	})

	// Process each OneOf group
	type oneOfGroup struct {
		constraintKey  string
		constraintText string
		attrs          []attrInfo
		firstAttrName  string
	}

	var groups []oneOfGroup
	for key, groupAttrs := range oneOfGroups {
		if len(groupAttrs) == 0 {
			continue
		}

		// Parse the constraint to get the expected order
		constraintText := groupAttrs[0].oneOfConstraint
		constraintOrder := strings.Split(constraintText, ", ")

		// Sort attributes within the group according to constraint order
		sort.Slice(groupAttrs, func(i, j int) bool {
			nameI := groupAttrs[i].name
			nameJ := groupAttrs[j].name

			// Find positions in constraint order
			posI := indexOf(constraintOrder, nameI)
			posJ := indexOf(constraintOrder, nameJ)

			// If both found, sort by constraint order
			if posI != -1 && posJ != -1 {
				return posI < posJ
			}

			// If one found and one not, found one comes first
			if posI != -1 {
				return true
			}
			if posJ != -1 {
				return false
			}

			// If neither found, sort alphabetically
			return nameI < nameJ
		})

		groups = append(groups, oneOfGroup{
			constraintKey:  key,
			constraintText: constraintText,
			attrs:          groupAttrs,
			firstAttrName:  groupAttrs[0].name,
		})
	}

	// Sort groups by first attribute name
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].firstAttrName < groups[j].firstAttrName
	})

	// Merge groups and standalone attrs in sorted order
	var result []attrInfo
	groupIdx := 0
	standaloneIdx := 0

	for groupIdx < len(groups) || standaloneIdx < len(standaloneAttrs) {
		// If we've exhausted groups, add remaining standalone attrs
		if groupIdx >= len(groups) {
			result = append(result, standaloneAttrs[standaloneIdx:]...)
			break
		}

		// If we've exhausted standalone attrs, add remaining groups
		if standaloneIdx >= len(standaloneAttrs) {
			for _, group := range groups[groupIdx:] {
				result = append(result, group.attrs...)
			}
			break
		}

		// Compare group's first attr with next standalone attr
		group := groups[groupIdx]
		standaloneAttr := standaloneAttrs[standaloneIdx]

		if group.firstAttrName < standaloneAttr.name {
			// Add the entire group
			result = append(result, group.attrs...)
			groupIdx++
		} else {
			// Add the standalone attr
			result = append(result, standaloneAttr)
			standaloneIdx++
		}
	}

	return result
}

// indexOf returns the index of target in slice, or -1 if not found
func indexOf(slice []string, target string) int {
	for i, s := range slice {
		if s == target {
			return i
		}
	}
	return -1
}

// convertNestedBlockAnchor converts nestedblock--xxx_yyy to xxx-yyy format
func convertNestedBlockAnchor(nestedPath string) string {
	// nestedPath is like "no_service_policies" or "active_service_policies--policies"
	// Convert underscores to hyphens, and -- to single hyphen
	result := strings.ToLower(nestedPath)
	result = strings.ReplaceAll(result, "_", "-")
	result = strings.ReplaceAll(result, "--", "-")
	return result
}

// transformAnchorsOnly handles already-transformed files by:
// 1. Converting any remaining nestedblock anchor IDs to simplified format
// 2. Removing empty sections (anchor + header with no content)
// 3. Removing "See...below" links that point to truly empty sections
// Note: Single-page rendering mode - no external nested_blocks files
func transformAnchorsOnly(filePath string, content string) error {
	lines := strings.Split(content, "\n")

	// First pass: identify which anchors have content in this file
	anchorsWithContent := make(map[string]bool)
	allAnchors := make(map[string]bool)
	anchorRegex := regexp.MustCompile(`<a id="([^"]+)"></a>`)
	attrLineRegex := regexp.MustCompile("^`[^`]+` - \\((Required|Optional)\\)")
	var currentAnchor string

	for _, line := range lines {
		if m := anchorRegex.FindStringSubmatch(line); m != nil {
			currentAnchor = m[1]
			allAnchors[currentAnchor] = true
		} else if currentAnchor != "" && attrLineRegex.MatchString(line) {
			anchorsWithContent[currentAnchor] = true
		}
	}

	// Second pass: filter output
	var output strings.Builder
	skipUntilNextAnchor := false
	seeRefRegex := regexp.MustCompile(`\s*See \[([^\]]+)\]\(#([^)]+)\) below[^.]*\.?`)

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		// Check if this is an anchor line
		if m := anchorRegex.FindStringSubmatch(line); m != nil {
			anchorName := m[1]

			// Convert any nestedblock format anchors
			if strings.HasPrefix(anchorName, "nestedblock--") {
				anchorName = convertNestedBlockAnchor(strings.TrimPrefix(anchorName, "nestedblock--"))
				line = fmt.Sprintf(`<a id="%s"></a>`, anchorName)
			}

			// Skip empty anchors and their headers
			if !anchorsWithContent[anchorName] {
				skipUntilNextAnchor = true
				// Also skip the next line if it's a header
				if i+1 < len(lines) && strings.HasPrefix(lines[i+1], "###") {
					i++ // Skip the header line
				}
				// Also skip following blank lines
				for i+1 < len(lines) && strings.TrimSpace(lines[i+1]) == "" {
					i++
				}
				continue
			}

			skipUntilNextAnchor = false
			output.WriteString(line)
			output.WriteString("\n")
			continue
		}

		// Skip content of empty blocks
		if skipUntilNextAnchor {
			continue
		}

		// Process attribute lines to update or remove anchor links
		// Single-page mode: all anchors are in same file, remove links to empty anchors
		if attrLineRegex.MatchString(line) {
			if m := seeRefRegex.FindStringSubmatch(line); m != nil {
				anchorRef := m[2]

				if anchorsWithContent[anchorRef] {
					// Anchor exists in this file with content - keep link as is
				} else {
					// Anchor doesn't exist or is empty - remove the link
					line = seeRefRegex.ReplaceAllString(line, "")
					line = strings.TrimSpace(line)
					// Ensure proper ending punctuation
					if !strings.HasSuffix(line, ".") && !strings.HasSuffix(line, ")") {
						line = line + "."
					}
				}
			}
		}

		output.WriteString(line)
		output.WriteString("\n")
	}

	// Normalize blank lines and write
	result := normalizeBlankLines(output.String())
	return os.WriteFile(filePath, []byte(result), 0644)
}

// NOTE: convertNestedBlocksHeadings function removed - single-page mode handles H3→bold inline

func transformDoc(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	contentStr := string(content)

	// Skip nested_blocks files - they are deleted in single-page mode
	if strings.Contains(filePath, "_nested_blocks") {
		return nil
	}

	// Check if file is already transformed (has "## Argument Reference" instead of "## Schema")
	// For already-transformed files, just do anchor/link conversion without full restructuring
	if strings.Contains(contentStr, "## Argument Reference") && !strings.Contains(contentStr, "## Schema") {
		return transformAnchorsOnly(filePath, contentStr)
	}

	lines := strings.Split(contentStr, "\n")

	// Maximum number of H3 headings that Terraform Registry reliably renders
	// Registry truncates large pages around heading ~66, so we use 60 as safe threshold
	const maxSafeHeadings = 60

	// Find key sections and collect all anchor names, tracking which have content
	schemaStart := -1
	importStart := -1
	firstNestedAnchor := -1
	existingAnchors := make(map[string]bool)
	anchorsWithContent := make(map[string]bool) // track anchors that have actual attributes
	safeAnchors := make(map[string]bool)        // track anchors within safe rendering range
	// Match both original tfplugindocs format (nestedblock--xxx) and already-transformed format (xxx-yyy)
	anchorRegexOriginal := regexp.MustCompile(`<a id="nestedblock--([^"]+)"></a>`)
	anchorRegexSimplified := regexp.MustCompile(`<a id="([a-z0-9-]+)"></a>`)
	// Attribute line patterns to detect content
	attrLineRegex := regexp.MustCompile(`^- \x60[^\x60]+\x60 \([^)]+\)`)
	transformedAttrLineRegex := regexp.MustCompile("^`[^`]+` - \\((Required|Optional)\\)")

	// First pass: identify all anchors and which ones have content
	// Also count H3 headings to determine safe rendering range
	var currentAnchorName string
	h3HeadingCount := 0
	for i, line := range lines {
		// Handle both original tfplugindocs format (## Schema) and already-transformed format (## Argument Reference)
		if strings.HasPrefix(line, "## Schema") || strings.HasPrefix(line, "## Argument Reference") {
			schemaStart = i
		}
		if strings.HasPrefix(line, "## Import") {
			importStart = i
		}

		// Count H3 headings (### ) to track position for safe rendering threshold
		if strings.HasPrefix(line, "### ") {
			h3HeadingCount++
		}

		// Check for original tfplugindocs anchor format
		if strings.HasPrefix(line, "<a id=\"nestedblock--") {
			if firstNestedAnchor < 0 {
				firstNestedAnchor = i
			}
			// Collect anchor names for later validation
			if m := anchorRegexOriginal.FindStringSubmatch(line); m != nil {
				// Store the simplified anchor name that will be generated
				anchorName := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(m[1], "_", "-"), "--", "-"))
				existingAnchors[anchorName] = true
				currentAnchorName = anchorName
				// Mark as safe if within rendering threshold
				if h3HeadingCount < maxSafeHeadings {
					safeAnchors[anchorName] = true
				}
			}
		} else if strings.HasPrefix(line, "<a id=\"") && !strings.Contains(line, "nestedblock") {
			// Check for already-transformed simplified anchor format
			if m := anchorRegexSimplified.FindStringSubmatch(line); m != nil {
				if firstNestedAnchor < 0 {
					firstNestedAnchor = i
				}
				existingAnchors[m[1]] = true
				currentAnchorName = m[1]
				// Mark as safe if within rendering threshold
				if h3HeadingCount < maxSafeHeadings {
					safeAnchors[m[1]] = true
				}
			}
		}
		// Check if current line is an attribute line (indicates the current anchor has content)
		if currentAnchorName != "" && (attrLineRegex.MatchString(line) || transformedAttrLineRegex.MatchString(line)) {
			anchorsWithContent[currentAnchorName] = true
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

	// Extract OneOf constraints from the resource Go file
	oneOfConstraintMap := extractOneOfFromGoFile(filePath)

	// Propagate OneOf constraints to all attributes mentioned in the constraint
	// This ensures that all mutually exclusive properties have the same constraint marker
	propagateOneOfConstraints(&metadataAttrs, oneOfConstraintMap)
	propagateOneOfConstraints(&specAttrs, oneOfConstraintMap)
	propagateOneOfConstraints(&readOnlyAttrs, oneOfConstraintMap)

	// Helper function to format a single attribute line
	// prefix is used for blockquote formatting (e.g., "> - " for OneOf groups)
	formatAttrLine := func(attr attrInfo, prefix string) string {
		typeStr := extractSimpleType(attr.typeInfo)

		// Clean any existing "See [X](#x) below" references from description to avoid duplication
		desc := attr.desc
		seeRefRegex := regexp.MustCompile(`See \[.+?\]\(#.+?\) below[^.]*\.?\s*`)
		desc = seeRefRegex.ReplaceAllString(desc, "")
		desc = strings.TrimSpace(desc)
		desc = strings.TrimSuffix(desc, ".")

		anchorName := toAnchorName(attr.name)
		// Only generate "See ... below" links if:
		// 1. The anchor has actual content (not empty blocks)
		// 2. The anchor is within safe rendering range (Terraform Registry truncates large pages)
		if attr.hasNested && anchorsWithContent[anchorName] && safeAnchors[anchorName] {
			if desc != "" {
				return fmt.Sprintf("%s`%s` - %s %s. See [%s](#%s) below for details.",
					prefix, attr.name, attr.reqStr, desc, toTitleCase(attr.name), anchorName)
			}
			return fmt.Sprintf("%s`%s` - %s See [%s](#%s) below for details.",
				prefix, attr.name, attr.reqStr, toTitleCase(attr.name), anchorName)
		}
		if desc != "" {
			return fmt.Sprintf("%s`%s` - %s %s (`%s`).", prefix, attr.name, attr.reqStr, desc, typeStr)
		}
		return fmt.Sprintf("%s`%s` - %s (`%s`).", prefix, attr.name, attr.reqStr, typeStr)
	}

	// Helper function to write a standalone attribute (not in OneOf group)
	writeAttr := func(attr attrInfo) {
		output.WriteString(formatAttrLine(attr, ""))
		output.WriteString("\n\n")
	}

	// Helper function to write a OneOf group with all content inside callout
	writeOneOfGroup := func(attrs []attrInfo) {
		if len(attrs) == 0 {
			return
		}
		// Write callout with bold title (no redundant "Note:" label - icon indicates it)
		// All content stays inside the callout box using <br> for line breaks
		output.WriteString("-> **Only one of the following may be set:**\n")
		for i, attr := range attrs {
			if i == 0 {
				// First attribute on next line (continuation of callout)
				output.WriteString(fmt.Sprintf("%s\n", formatAttrLine(attr, "")))
			} else {
				// Subsequent attributes with <br> prefix for line breaks inside callout
				output.WriteString(fmt.Sprintf("<br>%s\n", formatAttrLine(attr, "")))
			}
		}
		output.WriteString("\n")
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

		// Group attributes by OneOf constraint and sort for proper display
		sortedAttrs := groupAndSortAttributes(specAttrs)

		// Process attributes, grouping consecutive OneOf constraints together
		var currentGroup []attrInfo
		var currentConstraint string

		for _, attr := range sortedAttrs {
			constraintKey := normalizeOneOfKey(attr.oneOfConstraint)

			if constraintKey != currentConstraint {
				// Constraint changed - flush the previous group
				if len(currentGroup) > 0 {
					if currentConstraint != "" {
						// Previous group was a OneOf group
						writeOneOfGroup(currentGroup)
					} else {
						// Previous items were standalone
						for _, a := range currentGroup {
							writeAttr(a)
						}
					}
				}
				// Start new group
				currentGroup = []attrInfo{attr}
				currentConstraint = constraintKey
			} else {
				// Same constraint - add to current group
				currentGroup = append(currentGroup, attr)
			}
		}

		// Flush the last group
		if len(currentGroup) > 0 {
			if currentConstraint != "" {
				writeOneOfGroup(currentGroup)
			} else {
				for _, a := range currentGroup {
					writeAttr(a)
				}
			}
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

	// Process nested blocks if present - only output blocks that have content
	hasAnyContentBlocks := false
	for anchor := range anchorsWithContent {
		if existingAnchors[anchor] {
			hasAnyContentBlocks = true
			break
		}
	}

	if firstNestedAnchor > 0 && hasAnyContentBlocks {
		output.WriteString("\n---\n\n")

		nestedEnd := importStart
		if nestedEnd <= 0 {
			nestedEnd = len(lines)
		}

		// Transform each nested block
		currentBlockName := ""
		inNestedBlock := false
		skipCurrentBlock := false // NEW: track whether to skip empty blocks

		// Regex patterns for anchor detection
		anchorRegexOriginal := regexp.MustCompile(`<a id="nestedblock--([^"]+)"></a>`)
		anchorRegexSimplified := regexp.MustCompile(`<a id="([a-z0-9-]+)"></a>`)

		for i := firstNestedAnchor; i < nestedEnd; i++ {
			line := lines[i]

			// Detect anchor - handle both original tfplugindocs format and already-transformed format
			if strings.HasPrefix(line, "<a id=\"nestedblock--") {
				// Original tfplugindocs format: convert to simplified
				if m := anchorRegexOriginal.FindStringSubmatch(line); m != nil {
					attrPath := m[1] // e.g., "no_service_policies" or "active_service_policies--policies"
					currentBlockName = strings.ReplaceAll(attrPath, "--", ".")
					// Convert to simplified anchor: underscores and -- become hyphens
					anchorName := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(attrPath, "_", "-"), "--", "-"))

					// Skip empty blocks (those without content)
					if !anchorsWithContent[anchorName] {
						skipCurrentBlock = true
						inNestedBlock = false
						continue
					}

					skipCurrentBlock = false
					// Add blank line before anchor for proper markdown formatting
					output.WriteString(fmt.Sprintf("\n<a id=\"%s\"></a>\n\n", anchorName))
					inNestedBlock = true
				}
				continue
			} else if strings.HasPrefix(line, "<a id=\"") && !strings.Contains(line, "nestedblock") {
				// Already-transformed simplified format: preserve as-is
				if m := anchorRegexSimplified.FindStringSubmatch(line); m != nil {
					currentBlockName = strings.ReplaceAll(m[1], "-", ".")

					// Skip empty blocks (those without content)
					if !anchorsWithContent[m[1]] {
						skipCurrentBlock = true
						inNestedBlock = false
						continue
					}

					skipCurrentBlock = false
					// Write anchor as-is (already in simplified format)
					output.WriteString(fmt.Sprintf("\n<a id=\"%s\"></a>\n\n", m[1]))
					inNestedBlock = true
				}
				continue
			}

			// Skip everything in empty blocks
			if skipCurrentBlock {
				continue
			}

			// Transform nested schema headers to BOLD text (not H3) to prevent Registry truncation
			// Single-page rendering: H3→bold conversion allows full content to render
			if strings.HasPrefix(line, "### Nested Schema for") {
				headerRegex := regexp.MustCompile(`### Nested Schema for \x60([^\x60]+)\x60`)
				if m := headerRegex.FindStringSubmatch(line); m != nil {
					displayName := toTitleCase(strings.ReplaceAll(m[1], ".", " "))
					// Use bold text instead of H3 to prevent truncation
					output.WriteString(fmt.Sprintf("**%s**\n\n", displayName))
				}
				continue
			} else if strings.HasPrefix(line, "### ") && !strings.HasPrefix(line, "### Nested") && inNestedBlock {
				// Convert any H3 headers in nested blocks to bold text
				headerText := strings.TrimPrefix(line, "### ")
				output.WriteString(fmt.Sprintf("**%s**\n\n", headerText))
				continue
			}

			// Skip "Optional:" and "Required:" labels
			if line == "Optional:" || line == "Required:" {
				continue
			}

			// Transform h6 OneOf headings to info callouts (MD001 compliance) in nested blocks
			// Note: For nested blocks, use single-line callout with inline attribute names
			if strings.HasPrefix(line, "###### One of") {
				oneOfH6Regex := regexp.MustCompile(`###### One of the arguments from this list "([^"]+)" must be set`)
				if m := oneOfH6Regex.FindStringSubmatch(line); m != nil {
					// Parse the constraint list and format as inline backticked names
					constraintList := m[1]
					attrs := strings.Split(constraintList, ", ")
					var formattedAttrs []string
					for _, attrName := range attrs {
						formattedAttrs = append(formattedAttrs, fmt.Sprintf("`%s`", strings.TrimSpace(attrName)))
					}
					// Use Terraform Registry info callout syntax (single paragraph for light blue style)
					output.WriteString(fmt.Sprintf("-> **Note:** Only one of the following may be set: %s\n\n", strings.Join(formattedAttrs, ", ")))
				}
				continue
			}

			// Transform attribute lines in nested blocks - original tfplugindocs format
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
						// Extract the nested anchor and convert to simplified format
						nestedAnchorRegex := regexp.MustCompile(`#nestedblock--([^)]+)`)
						nestedAnchor := ""
						if am := nestedAnchorRegex.FindStringSubmatch(matches[3]); am != nil {
							// Convert to simplified anchor: underscores and -- become hyphens
							nestedAnchor = strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(am[1], "_", "-"), "--", "-"))
						}

						// Only generate "See...below" link if the nested anchor has content
						if nestedAnchor != "" && anchorsWithContent[nestedAnchor] && desc != "" {
							output.WriteString(fmt.Sprintf("`%s` - (Optional) %s. See [%s](#%s) below.\n\n",
								name, desc, toTitleCase(name), nestedAnchor))
						} else if nestedAnchor != "" && anchorsWithContent[nestedAnchor] {
							output.WriteString(fmt.Sprintf("`%s` - (Optional) See [%s](#%s) below.\n\n",
								name, toTitleCase(name), nestedAnchor))
						} else {
							// No nested content - just output the description with type
							if desc != "" {
								output.WriteString(fmt.Sprintf("`%s` - (Optional) %s (`%s`).\n\n", name, desc, typeStr))
							} else {
								output.WriteString(fmt.Sprintf("`%s` - (Optional) (`%s`).\n\n", name, typeStr))
							}
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

			// Handle already-transformed attribute lines - format: `name` - (Optional) ...
			if inNestedBlock && strings.HasPrefix(line, "`") && strings.Contains(line, "` - (") {
				// Already-transformed format: check for and remove links to empty anchors
				seeRefRegex := regexp.MustCompile(`\s*See \[([^\]]+)\]\(#([^)]+)\) below[^.]*\.?`)
				if m := seeRefRegex.FindStringSubmatch(line); m != nil {
					anchorRef := m[2]
					if !anchorsWithContent[anchorRef] {
						// Remove the "See...below" reference since anchor has no content
						line = seeRefRegex.ReplaceAllString(line, "")
						line = strings.TrimSpace(line)
						// Ensure line ends with proper format
						if !strings.HasSuffix(line, ".") && !strings.HasSuffix(line, ")") {
							line = line + "."
						}
					}
				}
				output.WriteString(line + "\n\n")
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

// NOTE: countH3Headings and shouldSplitToGuides functions removed
// Single-page rendering mode uses H3→bold conversion instead of splitting

// NOTE: extractNestedBlocks, generateNestedBlocksPage, and splitLargeDocument functions removed
// Single-page rendering mode renders all content inline with H3→bold conversion for nested blocks
