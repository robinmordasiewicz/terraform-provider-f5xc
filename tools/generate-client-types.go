// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

// ResourceInfo for client type generation
type ResourceInfo struct {
	Name      string // snake_case: http_loadbalancer
	TitleCase string // TitleCase: HTTPLoadBalancer
	APIPath   string // API path: http_loadbalancers
}

var clientTypeTemplate = `
// ===== {{.TitleCase}} =====
type {{.TitleCase}} struct {
	Metadata Metadata         ` + "`json:\"metadata\"`" + `
	Spec     {{.TitleCase}}Spec ` + "`json:\"spec\"`" + `
}
type {{.TitleCase}}Spec struct {
	Description string ` + "`json:\"description,omitempty\"`" + `
}
func (c *Client) Create{{.TitleCase}}(ctx context.Context, resource *{{.TitleCase}}) (*{{.TitleCase}}, error) {
	var result {{.TitleCase}}
	path := fmt.Sprintf("/api/config/namespaces/%s/{{.APIPath}}", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) Get{{.TitleCase}}(ctx context.Context, namespace, name string) (*{{.TitleCase}}, error) {
	var result {{.TitleCase}}
	path := fmt.Sprintf("/api/config/namespaces/%s/{{.APIPath}}/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}
func (c *Client) Update{{.TitleCase}}(ctx context.Context, resource *{{.TitleCase}}) (*{{.TitleCase}}, error) {
	var result {{.TitleCase}}
	path := fmt.Sprintf("/api/config/namespaces/%s/{{.APIPath}}/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}
func (c *Client) Delete{{.TitleCase}}(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/{{.APIPath}}/%s", namespace, name)
	return c.Delete(ctx, path)
}
`

func main() {
	fmt.Println("🔨 F5XC Client Type Generator")
	fmt.Println(strings.Repeat("=", 50))

	// Find all resource files
	resourceFiles, err := filepath.Glob("/tmp/terraform-provider-f5xc/internal/provider/*_resource.go")
	if err != nil {
		fmt.Printf("Error finding resource files: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("📄 Found %d resource files\n\n", len(resourceFiles))

	// Extract resource info from each file
	resources := []ResourceInfo{}
	tmpl := template.Must(template.New("client").Parse(clientTypeTemplate))

	// Types that already exist in client.go (manually implemented)
	existingTypes := map[string]bool{
		"namespace":         true,
		"http_loadbalancer": true,
		"origin_pool":       true,
	}

	for _, file := range resourceFiles {
		// Extract resource name from filename
		base := filepath.Base(file)
		name := strings.TrimSuffix(base, "_resource.go")

		// Skip types that already exist
		if existingTypes[name] {
			fmt.Printf("  Skipping %s (already exists)...\n", name)
			continue
		}

		// Read file to extract TitleCase name
		content, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		// Extract TitleCase from "type XxxResource struct"
		re := regexp.MustCompile(`type\s+([A-Z][a-zA-Z0-9]+)Resource\s+struct`)
		matches := re.FindSubmatch(content)
		if len(matches) < 2 {
			continue
		}

		titleCase := string(matches[1])

		// API path is plural form of name
		apiPath := toAPIPath(name)

		resources = append(resources, ResourceInfo{
			Name:      name,
			TitleCase: titleCase,
			APIPath:   apiPath,
		})
	}

	fmt.Printf("✅ Identified %d resources to add client types\n\n", len(resources))

	// Generate client types
	var buf bytes.Buffer
	for _, res := range resources {
		fmt.Printf("  Generating client types for %s (%s)...\n", res.Name, res.TitleCase)
		err := tmpl.Execute(&buf, res)
		if err != nil {
			fmt.Printf("  ❌ Error generating %s: %v\n", res.Name, err)
			continue
		}
	}

	// Read existing client.go
	clientPath := "/tmp/terraform-provider-f5xc/internal/client/client.go"
	existing, err := os.ReadFile(clientPath)
	if err != nil {
		fmt.Printf("Error reading client.go: %v\n", err)
		os.Exit(1)
	}

	// Simply append new types to existing content
	newContent := string(existing) + buf.String()

	// Write updated client.go
	err = os.WriteFile(clientPath, []byte(newContent), 0644)
	if err != nil {
		fmt.Printf("Error writing client.go: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n🎉 Client type generation complete!")
	fmt.Printf("📊 Added %d resource types to client.go\n", len(resources))
}

func toAPIPath(name string) string {
	// Convert snake_case name to API path (usually plural)
	if strings.HasSuffix(name, "y") && !strings.HasSuffix(name, "ey") {
		return strings.TrimSuffix(name, "y") + "ies"
	}
	if strings.HasSuffix(name, "s") || strings.HasSuffix(name, "x") {
		return name + "es"
	}
	return name + "s"
}
