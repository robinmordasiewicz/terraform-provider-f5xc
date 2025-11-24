// generate-api-docs.go generates API documentation pages from OpenAPI specifications
// Creates separate pages per API for better performance and navigation
// Usage: go run tools/generate-api-docs.go [--spec-dir=path] [--output-dir=path]
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type APISpec struct {
	Category    string
	Name        string
	Filename    string
	SchemaPath  string
	SlugName    string
	Description string
}

func main() {
	specDir := flag.String("spec-dir", "docs/specifications/api", "Directory containing OpenAPI spec files")
	outputDir := flag.String("output-dir", "docs/api", "Output directory for API documentation")
	flag.Parse()

	specs, err := scanSpecs(*specDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning specs: %v\n", err)
		os.Exit(1)
	}

	if len(specs) == 0 {
		fmt.Fprintf(os.Stderr, "No OpenAPI specs found in %s\n", *specDir)
		os.Exit(1)
	}

	// Create output directory
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Clean existing .md files in output directory (except index.md initially)
	cleanOutputDir(*outputDir)

	// Generate index page
	if err := generateIndexPage(*outputDir, specs); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating index: %v\n", err)
		os.Exit(1)
	}

	// Generate individual API pages
	for _, spec := range specs {
		if err := generateAPIPage(*outputDir, spec); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating page for %s: %v\n", spec.Name, err)
			os.Exit(1)
		}
	}

	// Generate navigation YAML for mkdocs.yml
	if err := generateNavYAML(*outputDir, specs); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating nav YAML: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %d API pages in %s\n", len(specs), *outputDir)
}

func cleanOutputDir(dir string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return // Directory might not exist yet
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
			os.Remove(filepath.Join(dir, entry.Name()))
		}
	}
}

func scanSpecs(dir string) ([]APISpec, error) {
	var specs []APISpec

	// Pattern: docs-cloud-f5-com.XXXX.public.ves.io.schema.CATEGORY[.SUBCATEGORY].ves-swagger.json
	pattern := regexp.MustCompile(`docs-cloud-f5-com\.(\d+)\.public\.ves\.io\.schema\.(.+)\.ves-swagger\.json`)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		match := pattern.FindStringSubmatch(entry.Name())
		if match == nil {
			continue
		}

		schemaPath := match[2]
		parts := strings.Split(schemaPath, ".")

		// Determine category and name
		category := formatName(parts[0])
		name := buildDisplayName(parts)

		// Create slug for filename
		slug := strings.ToLower(strings.ReplaceAll(schemaPath, ".", "-"))

		specs = append(specs, APISpec{
			Category:    category,
			Name:        name,
			Filename:    entry.Name(),
			SchemaPath:  schemaPath,
			SlugName:    slug,
			Description: fmt.Sprintf("API documentation for %s", name),
		})
	}

	// Sort by category, then by name
	sort.Slice(specs, func(i, j int) bool {
		if specs[i].Category != specs[j].Category {
			return specs[i].Category < specs[j].Category
		}
		return specs[i].Name < specs[j].Name
	})

	return specs, nil
}

func buildDisplayName(parts []string) string {
	if len(parts) == 1 {
		return formatName(parts[0])
	}

	// For nested paths, create a readable name
	var nameParts []string
	for _, p := range parts {
		nameParts = append(nameParts, formatName(p))
	}

	// If last part is generic like "subscription", include parent
	lastPart := parts[len(parts)-1]
	if lastPart == "subscription" || lastPart == "ves-swagger" {
		return strings.Join(nameParts, " - ")
	}

	// For views.*, show "View Name"
	if parts[0] == "views" && len(parts) > 1 {
		return formatName(parts[len(parts)-1])
	}

	return strings.Join(nameParts, " - ")
}

func formatName(s string) string {
	s = strings.ReplaceAll(s, "_", " ")

	abbreviations := map[string]string{
		"api": "API", "cdn": "CDN", "dns": "DNS", "http": "HTTP",
		"https": "HTTPS", "ip": "IP", "k8s": "K8s", "lb": "LB",
		"nat": "NAT", "nfv": "NFV", "oidc": "OIDC", "tcp": "TCP",
		"tls": "TLS", "udp": "UDP", "usb": "USB", "vpc": "VPC",
		"vpn": "VPN", "waf": "WAF", "bgp": "BGP", "crl": "CRL",
		"ike": "IKE", "scim": "SCIM", "tpm": "TPM", "acl": "ACL",
		"asn": "ASN", "aws": "AWS", "gcp": "GCP", "azure": "Azure",
		"rbac": "RBAC", "soa": "SOA", "srv6": "SRV6", "apm": "APM",
		"cds": "CDS", "dhcp": "DHCP", "lte": "LTE", "nginx": "NGINX",
		"pbac": "PBAC", "sdk": "SDK", "ssh": "SSH", "ui": "UI",
		"vnet": "VNet", "tgw": "TGW", "csg": "CSG",
	}

	words := strings.Fields(s)
	for i, word := range words {
		lower := strings.ToLower(word)
		if abbr, ok := abbreviations[lower]; ok {
			words[i] = abbr
		} else if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}

	return strings.Join(words, " ")
}

func generateIndexPage(outputDir string, specs []APISpec) error {
	var sb strings.Builder

	sb.WriteString(`---
page_title: "API Reference"
description: "Interactive F5 Distributed Cloud API documentation"
---

# API Reference

Explore the F5 Distributed Cloud API documentation with interactive "Try it out" functionality.

!!! tip "Navigation"
    Select an API from the left navigation menu to view its interactive documentation.
    Each API page includes a Swagger UI with "Try it out" capability.

## Quick Links

`)

	// Group by category for quick links
	categories := make(map[string][]APISpec)
	var categoryOrder []string

	for _, spec := range specs {
		if _, exists := categories[spec.Category]; !exists {
			categoryOrder = append(categoryOrder, spec.Category)
		}
		categories[spec.Category] = append(categories[spec.Category], spec)
	}

	for _, category := range categoryOrder {
		apis := categories[category]
		sb.WriteString(fmt.Sprintf("### %s\n\n", category))
		for _, api := range apis {
			sb.WriteString(fmt.Sprintf("- [%s](%s.md)\n", api.Name, api.SlugName))
		}
		sb.WriteString("\n")
	}

	sb.WriteString(fmt.Sprintf("\n---\n\n*%d APIs available*\n", len(specs)))

	return os.WriteFile(filepath.Join(outputDir, "index.md"), []byte(sb.String()), 0644)
}

func generateAPIPage(outputDir string, spec APISpec) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(`---
page_title: "%s API - F5XC Provider"
description: "%s"
---

# %s

%s

<swagger-ui src="../specifications/api/%s"/>
`, spec.Name, spec.Description, spec.Name, spec.Description, spec.Filename))

	filename := filepath.Join(outputDir, spec.SlugName+".md")
	return os.WriteFile(filename, []byte(sb.String()), 0644)
}

func generateNavYAML(outputDir string, specs []APISpec) error {
	var sb strings.Builder

	docsDir := filepath.Dir(outputDir)

	// Generate YAML navigation structure
	sb.WriteString("# Auto-generated navigation - DO NOT EDIT\n")
	sb.WriteString("# Generated by tools/generate-api-docs.go\n")
	sb.WriteString("nav:\n")
	sb.WriteString("  - Home: index.md\n")

	// Scan and add Resources
	resources, err := scanDocsDir(filepath.Join(docsDir, "resources"))
	if err == nil && len(resources) > 0 {
		sb.WriteString("  - Resources:\n")
		for _, r := range resources {
			sb.WriteString(fmt.Sprintf("    - %s: resources/%s\n", r.DisplayName, r.Filename))
		}
	}

	// Scan and add Data Sources
	dataSources, err := scanDocsDir(filepath.Join(docsDir, "data-sources"))
	if err == nil && len(dataSources) > 0 {
		sb.WriteString("  - Data Sources:\n")
		for _, ds := range dataSources {
			sb.WriteString(fmt.Sprintf("    - %s: data-sources/%s\n", ds.DisplayName, ds.Filename))
		}
	}

	// Add API Reference section
	sb.WriteString("  - API Reference:\n")
	sb.WriteString("    - Overview: api/index.md\n")

	// Add all APIs alphabetically by name
	for _, spec := range specs {
		sb.WriteString(fmt.Sprintf("    - %s: api/%s.md\n", spec.Name, spec.SlugName))
	}

	// Write to a separate file that can be used to update mkdocs.yml
	navFile := filepath.Join(docsDir, "nav-api.yml")
	return os.WriteFile(navFile, []byte(sb.String()), 0644)
}

type DocFile struct {
	Filename    string
	DisplayName string
}

func scanDocsDir(dir string) ([]DocFile, error) {
	var docs []DocFile

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		// Convert filename to display name
		// e.g., "http_loadbalancer.md" -> "HTTP Loadbalancer"
		name := strings.TrimSuffix(entry.Name(), ".md")
		displayName := formatName(name)

		docs = append(docs, DocFile{
			Filename:    entry.Name(),
			DisplayName: displayName,
		})
	}

	// Sort alphabetically by display name
	sort.Slice(docs, func(i, j int) bool {
		return docs[i].DisplayName < docs[j].DisplayName
	})

	return docs, nil
}
