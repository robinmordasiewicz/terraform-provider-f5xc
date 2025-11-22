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
	"strings"
)

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

	// Write everything before Schema section
	for i := 0; i < schemaStart; i++ {
		output.WriteString(lines[i])
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

	// Clean up whitespace
	desc = strings.TrimSpace(desc)

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
