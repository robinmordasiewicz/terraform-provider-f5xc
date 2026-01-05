//go:build ignore
// +build ignore

// This tool compares Terraform provider schemas between two generations
// (e.g., v1 vs v2 OpenAPI spec-based generation) to detect regressions.
//
// Usage:
//
//	go run tools/compare-schemas/main.go --baseline <dir> --candidate <dir>
//
// The tool detects:
//   - Missing resources (present in baseline but not candidate)
//   - New resources (present in candidate but not baseline)
//   - Missing attributes (attributes in baseline but not candidate)
//   - New attributes (attributes in candidate but not baseline)
//   - Type changes (attribute type differs between versions)
//
// Output:
//   - comparison-report.md: Human-readable report
//   - comparison-report.json: Machine-readable report
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// SchemaInfo holds extracted schema information for a resource
type SchemaInfo struct {
	Name       string                 `json:"name"`
	Attributes map[string]AttrInfo    `json:"attributes"`
	FilePath   string                 `json:"file_path,omitempty"`
}

// AttrInfo holds information about a single attribute
type AttrInfo struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
	Optional    bool   `json:"optional"`
	Computed    bool   `json:"computed"`
	Description string `json:"description,omitempty"`
}

// ComparisonReport holds the full comparison results
type ComparisonReport struct {
	Summary     ReportSummary        `json:"summary"`
	Missing     []ResourceDiff       `json:"missing_resources"`
	New         []ResourceDiff       `json:"new_resources"`
	Changed     []ResourceChange     `json:"changed_resources"`
	Unchanged   []string             `json:"unchanged_resources"`
}

// ReportSummary provides overview statistics
type ReportSummary struct {
	BaselineResources  int  `json:"baseline_resources"`
	CandidateResources int  `json:"candidate_resources"`
	MissingCount       int  `json:"missing_count"`
	NewCount           int  `json:"new_count"`
	ChangedCount       int  `json:"changed_count"`
	UnchangedCount     int  `json:"unchanged_count"`
	HasBreaking        bool `json:"has_breaking_changes"`
}

// ResourceDiff represents a resource that exists in one version but not the other
type ResourceDiff struct {
	Name           string `json:"name"`
	AttributeCount int    `json:"attribute_count"`
}

// ResourceChange represents changes to a resource between versions
type ResourceChange struct {
	Name               string       `json:"name"`
	MissingAttributes  []AttrDiff   `json:"missing_attributes,omitempty"`
	NewAttributes      []AttrDiff   `json:"new_attributes,omitempty"`
	TypeChanges        []TypeChange `json:"type_changes,omitempty"`
	IsBreaking         bool         `json:"is_breaking"`
}

// AttrDiff represents an attribute difference
type AttrDiff struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
}

// TypeChange represents an attribute type change
type TypeChange struct {
	Name        string `json:"name"`
	OldType     string `json:"old_type"`
	NewType     string `json:"new_type"`
	IsBreaking  bool   `json:"is_breaking"`
}

func main() {
	baselineDir := flag.String("baseline", "", "Directory containing baseline provider resources")
	candidateDir := flag.String("candidate", "", "Directory containing candidate provider resources")
	outputDir := flag.String("output", ".", "Output directory for reports")
	flag.Parse()

	if *baselineDir == "" || *candidateDir == "" {
		fmt.Fprintf(os.Stderr, "Usage: go run main.go --baseline <dir> --candidate <dir> [--output <dir>]\n")
		fmt.Fprintf(os.Stderr, "\nCompare Terraform provider schemas between two generation runs.\n\n")
		fmt.Fprintf(os.Stderr, "Arguments:\n")
		fmt.Fprintf(os.Stderr, "  --baseline   Directory with baseline *_resource.go files\n")
		fmt.Fprintf(os.Stderr, "  --candidate  Directory with candidate *_resource.go files\n")
		fmt.Fprintf(os.Stderr, "  --output     Output directory for reports (default: current)\n")
		os.Exit(1)
	}

	fmt.Printf("Comparing schemas:\n")
	fmt.Printf("  Baseline:  %s\n", *baselineDir)
	fmt.Printf("  Candidate: %s\n", *candidateDir)

	// Extract schemas from both directories
	baseline, err := extractSchemas(*baselineDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error extracting baseline schemas: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("  Baseline resources: %d\n", len(baseline))

	candidate, err := extractSchemas(*candidateDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error extracting candidate schemas: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("  Candidate resources: %d\n", len(candidate))

	// Compare schemas
	report := compareSchemas(baseline, candidate)

	// Generate reports
	mdPath := filepath.Join(*outputDir, "comparison-report.md")
	if err := writeMarkdownReport(report, mdPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing markdown report: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("  Generated: %s\n", mdPath)

	jsonPath := filepath.Join(*outputDir, "comparison-report.json")
	if err := writeJSONReport(report, jsonPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing JSON report: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("  Generated: %s\n", jsonPath)

	// Print summary
	fmt.Printf("\nSummary:\n")
	fmt.Printf("  Missing resources:  %d\n", report.Summary.MissingCount)
	fmt.Printf("  New resources:      %d\n", report.Summary.NewCount)
	fmt.Printf("  Changed resources:  %d\n", report.Summary.ChangedCount)
	fmt.Printf("  Unchanged resources: %d\n", report.Summary.UnchangedCount)

	if report.Summary.HasBreaking {
		fmt.Printf("\n  BREAKING CHANGES DETECTED\n")
		os.Exit(2)
	}
}

// extractSchemas parses Go files and extracts schema information
func extractSchemas(dir string) (map[string]SchemaInfo, error) {
	schemas := make(map[string]SchemaInfo)

	files, err := filepath.Glob(filepath.Join(dir, "*_resource.go"))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		schema, err := extractSchemaFromFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Could not extract schema from %s: %v\n", file, err)
			continue
		}
		if schema.Name != "" {
			schemas[schema.Name] = schema
		}
	}

	return schemas, nil
}

// extractSchemaFromFile parses a single Go file and extracts schema info
func extractSchemaFromFile(filePath string) (SchemaInfo, error) {
	schema := SchemaInfo{
		Attributes: make(map[string]AttrInfo),
		FilePath:   filePath,
	}

	// Extract resource name from filename
	base := filepath.Base(filePath)
	schema.Name = strings.TrimSuffix(base, "_resource.go")

	// Parse the Go file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return schema, err
	}

	// Find the model struct (usually named *Model)
	ast.Inspect(node, func(n ast.Node) bool {
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		// Look for structs ending with "Model"
		if !strings.HasSuffix(typeSpec.Name.Name, "Model") {
			return true
		}

		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		// Extract fields
		for _, field := range structType.Fields.List {
			if field.Tag == nil {
				continue
			}

			attr := extractAttrFromField(field)
			if attr.Name != "" {
				schema.Attributes[attr.Name] = attr
			}
		}

		return true
	})

	return schema, nil
}

// extractAttrFromField extracts attribute info from a struct field
func extractAttrFromField(field *ast.Field) AttrInfo {
	attr := AttrInfo{}

	// Get field name
	if len(field.Names) > 0 {
		attr.Name = field.Names[0].Name
	}

	// Get field type
	attr.Type = extractTypeString(field.Type)

	// Parse tfsdk tag for the attribute name used in Terraform
	if field.Tag != nil {
		tag := field.Tag.Value
		tfsdkRegex := regexp.MustCompile(`tfsdk:"([^"]+)"`)
		if match := tfsdkRegex.FindStringSubmatch(tag); len(match) > 1 {
			attr.Name = match[1]
		}
	}

	return attr
}

// extractTypeString converts an AST type to a string representation
func extractTypeString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		if x, ok := t.X.(*ast.Ident); ok {
			return x.Name + "." + t.Sel.Name
		}
		return t.Sel.Name
	case *ast.ArrayType:
		return "[]" + extractTypeString(t.Elt)
	case *ast.StarExpr:
		return "*" + extractTypeString(t.X)
	case *ast.MapType:
		return "map[" + extractTypeString(t.Key) + "]" + extractTypeString(t.Value)
	case *ast.IndexExpr:
		return extractTypeString(t.X) + "[" + extractTypeString(t.Index) + "]"
	default:
		return "unknown"
	}
}

// compareSchemas compares baseline and candidate schemas
func compareSchemas(baseline, candidate map[string]SchemaInfo) ComparisonReport {
	report := ComparisonReport{
		Missing:   []ResourceDiff{},
		New:       []ResourceDiff{},
		Changed:   []ResourceChange{},
		Unchanged: []string{},
	}

	// Find missing resources (in baseline but not candidate)
	for name, schema := range baseline {
		if _, ok := candidate[name]; !ok {
			report.Missing = append(report.Missing, ResourceDiff{
				Name:           name,
				AttributeCount: len(schema.Attributes),
			})
		}
	}

	// Find new resources (in candidate but not baseline)
	for name, schema := range candidate {
		if _, ok := baseline[name]; !ok {
			report.New = append(report.New, ResourceDiff{
				Name:           name,
				AttributeCount: len(schema.Attributes),
			})
		}
	}

	// Compare resources that exist in both
	for name, baseSchema := range baseline {
		candSchema, ok := candidate[name]
		if !ok {
			continue
		}

		change := compareResourceSchemas(name, baseSchema, candSchema)
		if len(change.MissingAttributes) > 0 || len(change.NewAttributes) > 0 || len(change.TypeChanges) > 0 {
			report.Changed = append(report.Changed, change)
		} else {
			report.Unchanged = append(report.Unchanged, name)
		}
	}

	// Sort results
	sort.Slice(report.Missing, func(i, j int) bool { return report.Missing[i].Name < report.Missing[j].Name })
	sort.Slice(report.New, func(i, j int) bool { return report.New[i].Name < report.New[j].Name })
	sort.Slice(report.Changed, func(i, j int) bool { return report.Changed[i].Name < report.Changed[j].Name })
	sort.Strings(report.Unchanged)

	// Calculate summary
	report.Summary = ReportSummary{
		BaselineResources:  len(baseline),
		CandidateResources: len(candidate),
		MissingCount:       len(report.Missing),
		NewCount:           len(report.New),
		ChangedCount:       len(report.Changed),
		UnchangedCount:     len(report.Unchanged),
		HasBreaking:        len(report.Missing) > 0,
	}

	// Check for breaking changes in changed resources
	for _, change := range report.Changed {
		if change.IsBreaking {
			report.Summary.HasBreaking = true
			break
		}
	}

	return report
}

// compareResourceSchemas compares two versions of a resource schema
func compareResourceSchemas(name string, baseline, candidate SchemaInfo) ResourceChange {
	change := ResourceChange{
		Name:              name,
		MissingAttributes: []AttrDiff{},
		NewAttributes:     []AttrDiff{},
		TypeChanges:       []TypeChange{},
	}

	// Find missing attributes
	for attrName, baseAttr := range baseline.Attributes {
		if _, ok := candidate.Attributes[attrName]; !ok {
			change.MissingAttributes = append(change.MissingAttributes, AttrDiff{
				Name:     attrName,
				Type:     baseAttr.Type,
				Required: baseAttr.Required,
			})
			if baseAttr.Required {
				change.IsBreaking = true
			}
		}
	}

	// Find new attributes
	for attrName, candAttr := range candidate.Attributes {
		if _, ok := baseline.Attributes[attrName]; !ok {
			change.NewAttributes = append(change.NewAttributes, AttrDiff{
				Name:     attrName,
				Type:     candAttr.Type,
				Required: candAttr.Required,
			})
			// New required attributes are breaking changes
			if candAttr.Required {
				change.IsBreaking = true
			}
		}
	}

	// Find type changes
	for attrName, baseAttr := range baseline.Attributes {
		candAttr, ok := candidate.Attributes[attrName]
		if !ok {
			continue
		}

		if baseAttr.Type != candAttr.Type {
			isBreaking := isBreakingTypeChange(baseAttr.Type, candAttr.Type)
			change.TypeChanges = append(change.TypeChanges, TypeChange{
				Name:       attrName,
				OldType:    baseAttr.Type,
				NewType:    candAttr.Type,
				IsBreaking: isBreaking,
			})
			if isBreaking {
				change.IsBreaking = true
			}
		}
	}

	// Sort results
	sort.Slice(change.MissingAttributes, func(i, j int) bool { return change.MissingAttributes[i].Name < change.MissingAttributes[j].Name })
	sort.Slice(change.NewAttributes, func(i, j int) bool { return change.NewAttributes[i].Name < change.NewAttributes[j].Name })
	sort.Slice(change.TypeChanges, func(i, j int) bool { return change.TypeChanges[i].Name < change.TypeChanges[j].Name })

	return change
}

// isBreakingTypeChange determines if a type change is breaking
func isBreakingTypeChange(oldType, newType string) bool {
	// Type changes are generally breaking, with some exceptions

	// Optional/nullable changes might not be breaking
	oldNormalized := strings.TrimPrefix(oldType, "*")
	newNormalized := strings.TrimPrefix(newType, "*")

	if oldNormalized == newNormalized {
		return false // Adding/removing pointer is usually not breaking
	}

	// Widening changes (string -> types.String) are usually not breaking
	terraformTypes := map[string]string{
		"string":    "types.String",
		"int64":     "types.Int64",
		"bool":      "types.Bool",
		"float64":   "types.Float64",
	}

	if tfType, ok := terraformTypes[oldType]; ok && newType == tfType {
		return false
	}

	return true
}

// writeMarkdownReport generates a human-readable report
func writeMarkdownReport(report ComparisonReport, path string) error {
	var sb strings.Builder

	sb.WriteString("# Schema Comparison Report\n\n")

	// Summary
	sb.WriteString("## Summary\n\n")
	sb.WriteString(fmt.Sprintf("| Metric | Count |\n"))
	sb.WriteString("| ------ | ----- |\n")
	sb.WriteString(fmt.Sprintf("| Baseline Resources | %d |\n", report.Summary.BaselineResources))
	sb.WriteString(fmt.Sprintf("| Candidate Resources | %d |\n", report.Summary.CandidateResources))
	sb.WriteString(fmt.Sprintf("| Missing Resources | %d |\n", report.Summary.MissingCount))
	sb.WriteString(fmt.Sprintf("| New Resources | %d |\n", report.Summary.NewCount))
	sb.WriteString(fmt.Sprintf("| Changed Resources | %d |\n", report.Summary.ChangedCount))
	sb.WriteString(fmt.Sprintf("| Unchanged Resources | %d |\n", report.Summary.UnchangedCount))
	sb.WriteString("\n")

	if report.Summary.HasBreaking {
		sb.WriteString("**BREAKING CHANGES DETECTED**\n\n")
	}

	// Missing resources
	if len(report.Missing) > 0 {
		sb.WriteString("## Missing Resources\n\n")
		sb.WriteString("Resources present in baseline but missing from candidate:\n\n")
		sb.WriteString("| Resource | Attributes |\n")
		sb.WriteString("| -------- | ---------- |\n")
		for _, r := range report.Missing {
			sb.WriteString(fmt.Sprintf("| `%s` | %d |\n", r.Name, r.AttributeCount))
		}
		sb.WriteString("\n")
	}

	// New resources
	if len(report.New) > 0 {
		sb.WriteString("## New Resources\n\n")
		sb.WriteString("Resources present in candidate but not in baseline:\n\n")
		sb.WriteString("| Resource | Attributes |\n")
		sb.WriteString("| -------- | ---------- |\n")
		for _, r := range report.New {
			sb.WriteString(fmt.Sprintf("| `%s` | %d |\n", r.Name, r.AttributeCount))
		}
		sb.WriteString("\n")
	}

	// Changed resources
	if len(report.Changed) > 0 {
		sb.WriteString("## Changed Resources\n\n")
		for _, change := range report.Changed {
			breakingMark := ""
			if change.IsBreaking {
				breakingMark = " **BREAKING**"
			}
			sb.WriteString(fmt.Sprintf("### `%s`%s\n\n", change.Name, breakingMark))

			if len(change.MissingAttributes) > 0 {
				sb.WriteString("**Missing Attributes:**\n\n")
				sb.WriteString("| Attribute | Type | Required |\n")
				sb.WriteString("| --------- | ---- | -------- |\n")
				for _, attr := range change.MissingAttributes {
					req := ""
					if attr.Required {
						req = "Yes"
					}
					sb.WriteString(fmt.Sprintf("| `%s` | `%s` | %s |\n", attr.Name, attr.Type, req))
				}
				sb.WriteString("\n")
			}

			if len(change.NewAttributes) > 0 {
				sb.WriteString("**New Attributes:**\n\n")
				sb.WriteString("| Attribute | Type | Required |\n")
				sb.WriteString("| --------- | ---- | -------- |\n")
				for _, attr := range change.NewAttributes {
					req := ""
					if attr.Required {
						req = "Yes"
					}
					sb.WriteString(fmt.Sprintf("| `%s` | `%s` | %s |\n", attr.Name, attr.Type, req))
				}
				sb.WriteString("\n")
			}

			if len(change.TypeChanges) > 0 {
				sb.WriteString("**Type Changes:**\n\n")
				sb.WriteString("| Attribute | Old Type | New Type | Breaking |\n")
				sb.WriteString("| --------- | -------- | -------- | -------- |\n")
				for _, tc := range change.TypeChanges {
					breaking := ""
					if tc.IsBreaking {
						breaking = "Yes"
					}
					sb.WriteString(fmt.Sprintf("| `%s` | `%s` | `%s` | %s |\n", tc.Name, tc.OldType, tc.NewType, breaking))
				}
				sb.WriteString("\n")
			}
		}
	}

	// Unchanged resources (collapsed)
	if len(report.Unchanged) > 0 {
		sb.WriteString("## Unchanged Resources\n\n")
		sb.WriteString("<details>\n")
		sb.WriteString("<summary>Click to expand list of unchanged resources</summary>\n\n")
		for _, name := range report.Unchanged {
			sb.WriteString(fmt.Sprintf("- `%s`\n", name))
		}
		sb.WriteString("\n</details>\n")
	}

	return os.WriteFile(path, []byte(sb.String()), 0644)
}

// writeJSONReport generates a machine-readable report
func writeJSONReport(report ComparisonReport, path string) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
