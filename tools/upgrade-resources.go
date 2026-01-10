// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

//go:build ignore
// +build ignore

// This script upgrades all resource files to include HashiCorp best practices:
// - Configurable timeouts using terraform-plugin-framework-timeouts
// - ModifyPlan for unknown value warnings
// - StateUpgraders for schema versioning
// - Private state for API metadata storage
// - ValidateConfig for configuration validation
//
// Usage: go run tools/upgrade-resources.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ResourceInfo contains information about a resource for transformation
type ResourceInfo struct {
	FileName      string
	ResourceName  string
	StructName    string
	ModelName     string
	IsLongRunning bool
}

// LongRunningResources that require extended timeouts
var longRunningResources = map[string]bool{
	"aws_vpc_site":       true,
	"azure_vnet_site":    true,
	"gcp_vpc_site":       true,
	"aws_tgw_site":       true,
	"voltstack_site":     true,
	"securemesh_site":    true,
	"securemesh_site_v2": true,
	"k8s_cluster":        true,
	"virtual_k8s":        true,
}

func main() {
	providerDir := "internal/provider"

	// Find all resource files
	files, err := filepath.Glob(filepath.Join(providerDir, "*_resource.go"))
	if err != nil {
		fmt.Printf("Error finding resource files: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d resource files to upgrade\n", len(files))

	// Skip namespace_resource.go as it's already the reference implementation
	upgraded := 0
	skipped := 0
	errors := 0

	for _, file := range files {
		if strings.HasSuffix(file, "namespace_resource.go") {
			fmt.Printf("Skipping %s (reference implementation)\n", file)
			skipped++
			continue
		}

		if strings.HasSuffix(file, "resource_base.go") {
			fmt.Printf("Skipping %s (helper file)\n", file)
			skipped++
			continue
		}

		err := upgradeResourceFile(file)
		if err != nil {
			fmt.Printf("Error upgrading %s: %v\n", file, err)
			errors++
		} else {
			fmt.Printf("Upgraded %s\n", file)
			upgraded++
		}
	}

	fmt.Printf("\nUpgrade complete: %d upgraded, %d skipped, %d errors\n", upgraded, skipped, errors)
}

func upgradeResourceFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	fileContent := string(content)

	// Skip if already upgraded (has timeouts import)
	if strings.Contains(fileContent, "terraform-plugin-framework-timeouts") {
		return nil // Already upgraded
	}

	// Extract resource information
	info := extractResourceInfo(filePath, fileContent)

	// Apply transformations
	newContent := applyTransformations(fileContent, info)

	// Write back
	return os.WriteFile(filePath, []byte(newContent), 0644)
}

func extractResourceInfo(filePath string, content string) ResourceInfo {
	info := ResourceInfo{
		FileName: filepath.Base(filePath),
	}

	// Extract resource name from filename
	baseName := strings.TrimSuffix(info.FileName, "_resource.go")
	info.ResourceName = baseName

	// Check if long-running
	info.IsLongRunning = longRunningResources[baseName]

	// Extract struct name (e.g., HTTPLoadBalancerResource)
	structRe := regexp.MustCompile(`type\s+(\w+Resource)\s+struct`)
	if matches := structRe.FindStringSubmatch(content); len(matches) > 1 {
		info.StructName = matches[1]
	}

	// Extract model name (e.g., HTTPLoadBalancerResourceModel)
	modelRe := regexp.MustCompile(`type\s+(\w+ResourceModel)\s+struct`)
	if matches := modelRe.FindStringSubmatch(content); len(matches) > 1 {
		info.ModelName = matches[1]
	}

	return info
}

func applyTransformations(content string, info ResourceInfo) string {
	// 1. Add imports
	content = addImports(content)

	// 2. Add interface assertions
	content = addInterfaceAssertions(content, info)

	// 3. Add schema version constant
	content = addSchemaVersion(content, info)

	// 4. Add Timeouts field to model
	content = addTimeoutsToModel(content, info)

	// 5. Update Schema function to add timeouts block and version
	content = updateSchemaFunction(content, info)

	// 6. Add ValidateConfig method
	content = addValidateConfigMethod(content, info)

	// 7. Add ModifyPlan method
	content = addModifyPlanMethod(content, info)

	// 8. Add UpgradeState method
	content = addUpgradeStateMethod(content, info)

	// 9. Update CRUD methods with timeouts
	content = updateCRUDMethods(content, info)

	return content
}

func addImports(content string) string {
	// Find import block
	importRe := regexp.MustCompile(`(?s)import\s*\(\s*(.+?)\s*\)`)

	// Add new imports if not present
	newImports := []string{
		`"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"`,
		`inttimeouts "github.com/f5xc/terraform-provider-f5xc/internal/timeouts"`,
		`"github.com/f5xc/terraform-provider-f5xc/internal/privatestate"`,
	}

	return importRe.ReplaceAllStringFunc(content, func(match string) string {
		for _, imp := range newImports {
			if !strings.Contains(match, imp) {
				// Add before closing paren
				match = strings.TrimSuffix(match, ")") + "\n\t" + imp + "\n)"
			}
		}
		return match
	})
}

func addInterfaceAssertions(content string, info ResourceInfo) string {
	// Find existing var block with interface assertions
	varBlockRe := regexp.MustCompile(`(?s)var\s*\(\s*_\s+resource\.Resource\s*.+?\)`)

	newAssertions := fmt.Sprintf(`var (
	_ resource.Resource                   = &%s{}
	_ resource.ResourceWithConfigure      = &%s{}
	_ resource.ResourceWithImportState    = &%s{}
	_ resource.ResourceWithModifyPlan     = &%s{}
	_ resource.ResourceWithUpgradeState   = &%s{}
	_ resource.ResourceWithValidateConfig = &%s{}
)`, info.StructName, info.StructName, info.StructName, info.StructName, info.StructName, info.StructName)

	return varBlockRe.ReplaceAllString(content, newAssertions)
}

func addSchemaVersion(content string, info ResourceInfo) string {
	// Add schema version constant after var block
	constName := strings.ToLower(strings.TrimSuffix(info.StructName, "Resource")) + "SchemaVersion"

	if strings.Contains(content, constName) {
		return content
	}

	// Find var block and add constant after it
	varBlockRe := regexp.MustCompile(`(?s)(var\s*\(.+?\))`)

	return varBlockRe.ReplaceAllString(content, fmt.Sprintf("$1\n\n// %s schema version for state upgrades\nconst %s int64 = 1", info.StructName, constName))
}

func addTimeoutsToModel(content string, info ResourceInfo) string {
	// Add Timeouts field to model struct
	modelRe := regexp.MustCompile(fmt.Sprintf(`(type\s+%s\s+struct\s*\{[^}]+)(ID\s+types\.String\s+`+"`"+`tfsdk:"id"`+"`"+`)`, info.ModelName))

	if !modelRe.MatchString(content) {
		return content
	}

	return modelRe.ReplaceAllString(content, "$1$2\n\tTimeouts    timeouts.Value `tfsdk:\"timeouts\"`")
}

func updateSchemaFunction(content string, info ResourceInfo) string {
	// Add Version to schema
	schemaRe := regexp.MustCompile(`(resp\.Schema\s*=\s*schema\.Schema\s*\{)`)
	constName := strings.ToLower(strings.TrimSuffix(info.StructName, "Resource")) + "SchemaVersion"
	content = schemaRe.ReplaceAllString(content, fmt.Sprintf("$1\n\t\tVersion: %s,", constName))

	// Add timeouts block if not present
	if !strings.Contains(content, `"timeouts":`) {
		// Find Blocks or add one
		blocksRe := regexp.MustCompile(`(Blocks:\s*map\[string\]schema\.Block\s*\{)(\s*\},?)`)
		if blocksRe.MatchString(content) {
			content = blocksRe.ReplaceAllString(content, "$1\n\t\t\t\"timeouts\": timeouts.Block(ctx, timeouts.Opts{\n\t\t\t\tCreate: true,\n\t\t\t\tRead:   true,\n\t\t\t\tUpdate: true,\n\t\t\t\tDelete: true,\n\t\t\t}),$2")
		} else {
			// Add Blocks section
			attrsEndRe := regexp.MustCompile(`(\},\s*)(}\s*}\s*$)`)
			content = attrsEndRe.ReplaceAllString(content, "$1Blocks: map[string]schema.Block{\n\t\t\t\"timeouts\": timeouts.Block(ctx, timeouts.Opts{\n\t\t\t\tCreate: true,\n\t\t\t\tRead:   true,\n\t\t\t\tUpdate: true,\n\t\t\t\tDelete: true,\n\t\t\t}),\n\t\t},$2")
		}
	}

	return content
}

func addValidateConfigMethod(content string, info ResourceInfo) string {
	if strings.Contains(content, "func (r *"+info.StructName+") ValidateConfig") {
		return content
	}

	method := fmt.Sprintf(`

// ValidateConfig implements resource.ResourceWithValidateConfig
func (r *%s) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data %s
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Additional validation can be added here
}
`, info.StructName, info.ModelName)

	// Add before Create method
	createRe := regexp.MustCompile(`(\n)(func \(r \*` + info.StructName + `\) Create)`)
	return createRe.ReplaceAllString(content, method+"$1$2")
}

func addModifyPlanMethod(content string, info ResourceInfo) string {
	if strings.Contains(content, "func (r *"+info.StructName+") ModifyPlan") {
		return content
	}

	resourceDesc := strings.ReplaceAll(strings.TrimSuffix(info.StructName, "Resource"), "_", " ")

	method := fmt.Sprintf(`

// ModifyPlan implements resource.ResourceWithModifyPlan
func (r *%s) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Handle resource destruction warnings (Terraform 1.3+)
	if req.Plan.Raw.IsNull() {
		resp.Diagnostics.AddWarning(
			"Resource Destruction",
			"This will permanently delete the %s from F5 Distributed Cloud.",
		)
		return
	}
}
`, info.StructName, resourceDesc)

	// Add before ValidateConfig or Create method
	insertRe := regexp.MustCompile(`(\n)(func \(r \*` + info.StructName + `\) (ValidateConfig|Create))`)
	return insertRe.ReplaceAllString(content, method+"$1$2")
}

func addUpgradeStateMethod(content string, info ResourceInfo) string {
	if strings.Contains(content, "func (r *"+info.StructName+") UpgradeState") {
		return content
	}

	method := fmt.Sprintf(`

// UpgradeState implements resource.ResourceWithUpgradeState
func (r *%s) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// Version 0 to 1 upgrade: add timeouts block
		0: {
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				// Timeouts block is optional with defaults, no data migration needed
			},
		},
	}
}
`, info.StructName)

	// Add before Create method
	createRe := regexp.MustCompile(`(\n)(func \(r \*` + info.StructName + `\) Create)`)
	return createRe.ReplaceAllString(content, method+"$1$2")
}

func updateCRUDMethods(content string, info ResourceInfo) string {
	// Update Create method
	content = addTimeoutToMethod(content, info, "Create", "inttimeouts.DefaultCreate")

	// Update Read method
	content = addTimeoutToMethod(content, info, "Read", "inttimeouts.DefaultRead")

	// Update Update method
	content = addTimeoutToMethod(content, info, "Update", "inttimeouts.DefaultUpdate")

	// Update Delete method
	content = addTimeoutToMethod(content, info, "Delete", "inttimeouts.DefaultDelete")

	return content
}

func addTimeoutToMethod(content string, info ResourceInfo, method, defaultTimeout string) string {
	methodLower := strings.ToLower(method)

	// Check if timeout already added
	if strings.Contains(content, methodLower+"Timeout, diags :=") {
		return content
	}

	// Find the method and add timeout after data extraction
	methodRe := regexp.MustCompile(fmt.Sprintf(`(func \(r \*%s\) %s\(ctx context\.Context, req resource\.%sRequest, resp \*resource\.%sResponse\) \{\s*var data %s\s*resp\.Diagnostics\.Append\(req\.(Plan|State)\.Get\(ctx, &data\)\.\.\.\)\s*if resp\.Diagnostics\.HasError\(\) \{\s*return\s*\})`, info.StructName, method, method, method, info.ModelName))

	timeoutCode := fmt.Sprintf(`$1

	// Apply configurable timeout
	%sTimeout, diags := data.Timeouts.%s(ctx, %s)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx, cancel := context.WithTimeout(ctx, %sTimeout)
	defer cancel()`, methodLower, method, defaultTimeout, methodLower)

	return methodRe.ReplaceAllString(content, timeoutCode)
}

func readFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sb.WriteString(scanner.Text() + "\n")
	}
	return sb.String(), scanner.Err()
}
