// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore
// +build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/f5xc/terraform-provider-f5xc/tools/pkg/naming"
)

func main() {
	// Read all resource files
	providerDir := "/tmp/terraform-provider-f5xc/internal/provider"
	clientFile := "/tmp/terraform-provider-f5xc/internal/client/client.go"

	files, err := filepath.Glob(filepath.Join(providerDir, "*_resource.go"))
	if err != nil {
		fmt.Printf("Error reading provider dir: %v\n", err)
		os.Exit(1)
	}

	// Extract type names from each resource file
	typeNames := make(map[string]string) // resourceName -> TitleCase
	re := regexp.MustCompile(`type\s+(\w+)Resource\s+struct`)

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			continue
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			matches := re.FindStringSubmatch(line)
			if len(matches) > 1 {
				typeName := matches[1]
				// Skip the ones we already have
				if typeName != "Namespace" && typeName != "HTTPLoadBalancer" && typeName != "OriginPool" {
					resourceName := toSnakeCase(typeName)
					typeNames[resourceName] = typeName
					fmt.Printf("Found: %s -> %s\n", resourceName, typeName)
				}
				break
			}
		}
	}

	fmt.Printf("\n✅ Found %d new resources\n\n", len(typeNames))

	// Read current client.go and find insertion point
	clientBytes, err := os.ReadFile(clientFile)
	if err != nil {
		fmt.Printf("Error reading client: %v\n", err)
		os.Exit(1)
	}

	// Remove the malformed types if they exist
	clientContent := string(clientBytes)
	lines := strings.Split(clientContent, "\n")

	// Find the last closing brace
	lastBraceIndex := -1
	for i := len(lines) - 1; i >= 0; i-- {
		if strings.TrimSpace(lines[i]) == "}" {
			lastBraceIndex = i
			break
		}
	}

	if lastBraceIndex == -1 {
		fmt.Println("Could not find closing brace")
		os.Exit(1)
	}

	// Take everything up to (but not including) the last brace
	newContent := strings.Join(lines[:lastBraceIndex], "\n")

	// Add types for each resource
	for _, typeName := range typeNames {
		resourcePath := toSnakeCase(typeName) + "s"

		newContent += fmt.Sprintf(`

// %s represents a F5XC %s
type %s struct {
	Metadata Metadata     `+"`json:\"metadata\"`"+`
	Spec     %sSpec `+"`json:\"spec\"`"+`
}

// %sSpec defines the specification for %s
type %sSpec struct {
	Description string `+"`json:\"description,omitempty\"`"+`
}

// Create%s creates a new %s
func (c *Client) Create%s(ctx context.Context, resource *%s) (*%s, error) {
	var result %s
	path := fmt.Sprintf("/api/config/namespaces/%%s/%s", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}

// Get%s retrieves a %s
func (c *Client) Get%s(ctx context.Context, namespace, name string) (*%s, error) {
	var result %s
	path := fmt.Sprintf("/api/config/namespaces/%%s/%s/%%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}

// Update%s updates a %s
func (c *Client) Update%s(ctx context.Context, resource *%s) (*%s, error) {
	var result %s
	path := fmt.Sprintf("/api/config/namespaces/%%s/%s/%%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}

// Delete%s deletes a %s
func (c *Client) Delete%s(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%%s/%s/%%s", namespace, name)
	return c.Delete(ctx, path)
}
`, typeName, typeName,
			typeName, typeName,
			typeName, typeName,
			typeName,
			typeName, typeName,
			typeName, typeName, typeName,
			typeName, resourcePath,
			typeName, typeName,
			typeName, typeName, typeName,
			typeName, resourcePath,
			typeName, typeName,
			typeName, typeName, typeName,
			typeName, resourcePath,
			typeName, typeName,
			typeName,
			resourcePath)
	}

	// Close the file
	newContent += "\n}\n"

	// Write back
	err = os.WriteFile(clientFile, []byte(newContent), 0644)
	if err != nil {
		fmt.Printf("Error writing client: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n✅ Client types fixed successfully")
	fmt.Printf("📝 Updated: %s\n", clientFile)
}

// toSnakeCase wraps naming.ToSnakeCase for backward compatibility.
func toSnakeCase(s string) string {
	return naming.ToSnakeCase(s)
}
