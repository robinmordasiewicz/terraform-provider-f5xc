// generate-api-docs.go generates the docs/api.md file from OpenAPI specifications
// Usage: go run tools/generate-api-docs.go [--spec-dir=path]
package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type APISpec struct {
	Category   string
	Name       string
	Filename   string
	SchemaPath string
}

func main() {
	specDir := flag.String("spec-dir", "docs/specifications/api", "Directory containing OpenAPI spec files")
	outputFile := flag.String("output", "docs/api.md", "Output markdown file")
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

	content := generateMarkdown(specs)

	if err := os.WriteFile(*outputFile, []byte(content), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s with %d APIs across %d categories\n", *outputFile, len(specs), countCategories(specs))
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
		name := formatName(parts[len(parts)-1])

		// Handle nested paths for better naming
		if len(parts) > 1 {
			// Use the last part as name, but include parent for context if it adds value
			if parts[len(parts)-1] == "subscription" && len(parts) > 1 {
				name = formatName(parts[len(parts)-2]) + " " + name
			}
		}

		specs = append(specs, APISpec{
			Category:   category,
			Name:       name,
			Filename:   entry.Name(),
			SchemaPath: schemaPath,
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

func formatName(s string) string {
	// Replace underscores with spaces and title case
	s = strings.ReplaceAll(s, "_", " ")

	// Handle common abbreviations
	abbreviations := map[string]string{
		"api":   "API",
		"cdn":   "CDN",
		"dns":   "DNS",
		"http":  "HTTP",
		"https": "HTTPS",
		"ip":    "IP",
		"k8s":   "K8s",
		"lb":    "LB",
		"nat":   "NAT",
		"nfv":   "NFV",
		"oidc":  "OIDC",
		"tcp":   "TCP",
		"tls":   "TLS",
		"udp":   "UDP",
		"usb":   "USB",
		"vpc":   "VPC",
		"vpn":   "VPN",
		"waf":   "WAF",
		"bgp":   "BGP",
		"crl":   "CRL",
		"ike":   "IKE",
		"scim":  "SCIM",
		"tpm":   "TPM",
		"acl":   "ACL",
		"asn":   "ASN",
		"aws":   "AWS",
		"gcp":   "GCP",
		"azure": "Azure",
		"rbac":  "RBAC",
		"soa":   "SOA",
		"srv6":  "SRV6",
	}

	words := strings.Fields(s)
	for i, word := range words {
		lower := strings.ToLower(word)
		if abbr, ok := abbreviations[lower]; ok {
			words[i] = abbr
		} else {
			words[i] = strings.Title(lower)
		}
	}

	return strings.Join(words, " ")
}

func generateMarkdown(specs []APISpec) string {
	var sb strings.Builder

	sb.WriteString(`---
page_title: "API Reference"
description: "Interactive F5 Distributed Cloud API documentation with Try it out functionality"
---

# API Reference

Explore the F5 Distributed Cloud API documentation with interactive "Try it out" functionality.

!!! tip "Usage"
    Expand any API section below to view the interactive Swagger UI documentation.
    Use the "Try it out" button to test API calls directly.

`)

	// Group by category
	categories := make(map[string][]APISpec)
	var categoryOrder []string

	for _, spec := range specs {
		if _, exists := categories[spec.Category]; !exists {
			categoryOrder = append(categoryOrder, spec.Category)
		}
		categories[spec.Category] = append(categories[spec.Category], spec)
	}

	// Generate content for each category
	for _, category := range categoryOrder {
		apis := categories[category]
		sb.WriteString(fmt.Sprintf("## %s\n\n", category))

		for _, api := range apis {
			sb.WriteString(fmt.Sprintf("### %s\n\n", api.Name))
			sb.WriteString(fmt.Sprintf("<swagger-ui src=\"specifications/api/%s\"/>\n\n", api.Filename))
		}
	}

	// Add footer with metadata
	sb.WriteString(fmt.Sprintf("\n<!-- Generated automatically from %d OpenAPI specifications -->\n", len(specs)))

	return sb.String()
}

func countCategories(specs []APISpec) int {
	categories := make(map[string]bool)
	for _, spec := range specs {
		categories[spec.Category] = true
	}
	return len(categories)
}
