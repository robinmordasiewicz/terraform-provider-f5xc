package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func main() {
	providerDir := "/tmp/terraform-provider-f5xc/internal/provider"
	providerFile := filepath.Join(providerDir, "provider.go")

	// Read all resource files
	files, err := filepath.Glob(filepath.Join(providerDir, "*_resource.go"))
	if err != nil {
		fmt.Printf("Error reading provider dir: %v\n", err)
		os.Exit(1)
	}

	// Extract function names
	funcRe := regexp.MustCompile(`^func\s+New(\w+)Resource\(\)`)
	var resourceFuncs []string

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			continue
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			matches := funcRe.FindStringSubmatch(line)
			if len(matches) > 1 {
				funcName := "New" + matches[1] + "Resource"
				resourceFuncs = append(resourceFuncs, funcName)
				break
			}
		}
	}

	sort.Strings(resourceFuncs)
	fmt.Printf("Found %d resources\n\n", len(resourceFuncs))

	// Read provider.go
	content, err := os.ReadFile(providerFile)
	if err != nil {
		fmt.Printf("Error reading provider.go: %v\n", err)
		os.Exit(1)
	}

	// Replace Resources function
	lines := strings.Split(string(content), "\n")
	newLines := []string{}
	inResourcesFunc := false
	funcDepth := 0

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.Contains(line, "func (p *F5XCProvider) Resources(") {
			inResourcesFunc = true
			newLines = append(newLines, line)
			continue
		}

		if inResourcesFunc {
			if strings.Contains(trimmed, "{") {
				funcDepth++
			}
			if strings.Contains(trimmed, "}") {
				funcDepth--
				if funcDepth == 0 {
					// End of Resources function - insert our resources
					newLines = append(newLines, "\treturn []func() resource.Resource{")
					for _, fn := range resourceFuncs {
						newLines = append(newLines, fmt.Sprintf("\t\t%s,", fn))
					}
					newLines = append(newLines, "\t}")
					newLines = append(newLines, "}")
					inResourcesFunc = false
					continue
				}
			}
			continue // Skip old content
		}

		newLines = append(newLines, line)
	}

	// Write back
	err = os.WriteFile(providerFile, []byte(strings.Join(newLines, "\n")), 0644)
	if err != nil {
		fmt.Printf("Error writing provider.go: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Resources registered successfully")
	fmt.Printf("📝 Updated: %s\n", providerFile)
	fmt.Printf("📊 Total resources: %d\n", len(resourceFuncs))
}
