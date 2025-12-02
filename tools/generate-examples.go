//go:build ignore
// +build ignore

// This tool generates comprehensive example Terraform configurations for all F5 XC resources.
// Examples follow the Volterra provider style with inline OneOf comments showing mutually exclusive options.
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/f5xc/terraform-provider-f5xc/tools/pkg/naming"
	"github.com/f5xc/terraform-provider-f5xc/tools/pkg/namespace"
)

type SchemaInfo struct {
	ResourceName string
	TypeName     string
	Description  string
	Attributes   []AttributeInfo
	Blocks       []BlockInfo
	OneOfGroups  []OneOfGroup
}

type AttributeInfo struct {
	Name        string
	Type        string
	Description string
	Required    bool
	Optional    bool
	OneOfGroup  string // Which OneOf group this attribute belongs to
}

type BlockInfo struct {
	Name        string
	Description string
	IsList      bool
	Attributes  []AttributeInfo
	Blocks      []BlockInfo
	OneOfGroup  string // Which OneOf group this block belongs to
}

// OneOfGroup represents a group of mutually exclusive options
type OneOfGroup struct {
	Options     []string // List of mutually exclusive option names
	Description string   // Description of the group
}

// getNamespaceForResource wraps namespace.ForResource for backward compatibility
func getNamespaceForResource(resourceName string) (namespace.Type, string) {
	return namespace.ForResource(resourceName)
}

// getNamespaceForReference wraps namespace.ForReference for backward compatibility
func getNamespaceForReference(referencedResourceType string) string {
	return namespace.ForReference(referencedResourceType)
}

func main() {
	providerDir := "internal/provider"
	examplesDir := "examples/resources"

	files, err := filepath.Glob(filepath.Join(providerDir, "*_resource.go"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding resource files: %v\n", err)
		os.Exit(1)
	}

	var generatedCount int
	var formatErrors []string

	for _, file := range files {
		baseName := filepath.Base(file)
		resourceName := strings.TrimSuffix(baseName, "_resource.go")

		schema := parseResourceFile(file)
		if schema == nil {
			continue
		}

		exampleDir := filepath.Join(examplesDir, "f5xc_"+resourceName)
		if err := os.MkdirAll(exampleDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory %s: %v\n", exampleDir, err)
			continue
		}

		exampleFile := filepath.Join(exampleDir, "resource.tf")
		example := generateExample(resourceName, schema)
		if err := os.WriteFile(exampleFile, []byte(example), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", exampleFile, err)
			continue
		}

		// Run terraform fmt to ensure proper formatting and validate HCL syntax
		// terraform fmt will fail if the file has syntax errors
		cmd := exec.Command("terraform", "fmt", exampleFile)
		output, err := cmd.CombinedOutput()
		if err != nil {
			errMsg := fmt.Sprintf("%s: %v", exampleFile, err)
			if len(output) > 0 {
				errMsg = fmt.Sprintf("%s\n%s", errMsg, string(output))
			}
			formatErrors = append(formatErrors, errMsg)
			fmt.Fprintf(os.Stderr, "❌ Format/syntax error in %s: %v\n", exampleFile, err)
			if len(output) > 0 {
				fmt.Fprintf(os.Stderr, "   %s\n", string(output))
			}
		} else {
			generatedCount++
			fmt.Printf("✅ Generated: %s\n", exampleFile)
		}
	}

	// Print summary
	fmt.Printf("\n=== Generation Summary ===\n")
	fmt.Printf("Successfully generated: %d examples\n", generatedCount)

	if len(formatErrors) > 0 {
		fmt.Printf("Format/syntax errors: %d\n", len(formatErrors))
		fmt.Fprintf(os.Stderr, "\n❌ FAILED: %d example(s) have formatting or syntax errors:\n", len(formatErrors))
		for _, err := range formatErrors {
			fmt.Fprintf(os.Stderr, "  - %s\n", err)
		}
		os.Exit(1)
	}

	fmt.Printf("✅ All examples generated and validated successfully\n")
}

func parseResourceFile(filename string) *SchemaInfo {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	schema := &SchemaInfo{}
	scanner := bufio.NewScanner(file)

	// Read the entire file content
	var content strings.Builder
	for scanner.Scan() {
		content.WriteString(scanner.Text() + "\n")
	}

	text := content.String()

	// Extract resource type name
	typeNameRe := regexp.MustCompile(`TypeName = req\.ProviderTypeName \+ "_([^"]+)"`)
	if match := typeNameRe.FindStringSubmatch(text); match != nil {
		schema.TypeName = match[1]
	}

	// Extract description
	descRe := regexp.MustCompile(`MarkdownDescription:\s*"([^"]+)"`)
	if match := descRe.FindStringSubmatch(text); match != nil {
		schema.Description = match[1]
	}

	// Parse OneOf groups from MarkdownDescription fields
	schema.OneOfGroups = parseOneOfGroups(text)

	// Parse attributes
	schema.Attributes = parseAttributes(text)

	// Parse blocks
	schema.Blocks = parseBlocks(text)

	return schema
}

// parseOneOfGroups extracts OneOf groups from [OneOf: option1, option2] markers in MarkdownDescription
func parseOneOfGroups(content string) []OneOfGroup {
	var groups []OneOfGroup
	seen := make(map[string]bool)

	// Pattern to match [OneOf: option1, option2, option3] in MarkdownDescription
	oneOfRe := regexp.MustCompile(`\[OneOf:\s*([^\]]+)\]`)
	matches := oneOfRe.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		optionsStr := match[1]
		// Split by comma and clean up
		options := strings.Split(optionsStr, ",")
		var cleanOptions []string
		for _, opt := range options {
			opt = strings.TrimSpace(opt)
			if opt != "" {
				cleanOptions = append(cleanOptions, opt)
			}
		}

		if len(cleanOptions) > 1 {
			// Create a key to deduplicate
			key := strings.Join(cleanOptions, "|")
			if !seen[key] {
				seen[key] = true
				groups = append(groups, OneOfGroup{
					Options: cleanOptions,
				})
			}
		}
	}

	return groups
}

// formatOneOfComment formats options into Volterra-style comment
func formatOneOfComment(options []string) string {
	return fmt.Sprintf("// One of the arguments from this list \"%s\" must be set", strings.Join(options, " "))
}

// getOneOfGroupForBlock finds the OneOf group that contains this block name
func getOneOfGroupForBlock(blockName string, groups []OneOfGroup) *OneOfGroup {
	for i := range groups {
		for _, opt := range groups[i].Options {
			if opt == blockName {
				return &groups[i]
			}
		}
	}
	return nil
}

func parseAttributes(content string) []AttributeInfo {
	var attrs []AttributeInfo

	// Common patterns for attribute extraction
	attrRe := regexp.MustCompile(`"(\w+)":\s*schema\.(String|Bool|Int64|Float64|List|Map)Attribute\{[^}]*MarkdownDescription:\s*"([^"]*)"[^}]*(Required:\s*true)?(Optional:\s*true)?`)
	matches := attrRe.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		attr := AttributeInfo{
			Name:        match[1],
			Type:        strings.ToLower(match[2]),
			Description: match[3],
			Required:    strings.Contains(match[0], "Required: true"),
			Optional:    strings.Contains(match[0], "Optional: true"),
		}
		// Skip computed-only attributes
		if attr.Name != "id" {
			attrs = append(attrs, attr)
		}
	}

	return attrs
}

func parseBlocks(content string) []BlockInfo {
	var blocks []BlockInfo

	// Find block definitions
	blockRe := regexp.MustCompile(`"(\w+)":\s*schema\.(SingleNestedBlock|ListNestedBlock)\{[^{]*MarkdownDescription:\s*"([^"]*)"`)
	matches := blockRe.FindAllStringSubmatch(content, -1)

	seen := make(map[string]bool)
	for _, match := range matches {
		if seen[match[1]] {
			continue
		}
		seen[match[1]] = true

		block := BlockInfo{
			Name:        match[1],
			IsList:      match[2] == "ListNestedBlock",
			Description: match[3],
		}
		blocks = append(blocks, block)
	}

	return blocks
}

func generateExample(resourceName string, schema *SchemaInfo) string {
	var sb strings.Builder

	humanName := toHumanName(resourceName)
	sb.WriteString(fmt.Sprintf("# %s Resource Example\n", humanName))
	sb.WriteString(fmt.Sprintf("# %s\n\n", schema.Description))

	// Get the appropriate namespace for this resource type
	_, namespace := getNamespaceForResource(resourceName)

	// Generate basic example
	sb.WriteString(fmt.Sprintf("# Basic %s configuration\n", humanName))
	sb.WriteString(fmt.Sprintf("resource \"f5xc_%s\" \"example\" {\n", resourceName))
	sb.WriteString(fmt.Sprintf("  name      = \"example-%s\"\n", strings.ReplaceAll(resourceName, "_", "-")))
	sb.WriteString(fmt.Sprintf("  namespace = \"%s\"\n\n", namespace))

	// Add labels
	sb.WriteString("  labels = {\n")
	sb.WriteString("    environment = \"production\"\n")
	sb.WriteString("    managed_by  = \"terraform\"\n")
	sb.WriteString("  }\n\n")

	// Add annotations
	sb.WriteString("  annotations = {\n")
	sb.WriteString("    \"owner\" = \"platform-team\"\n")
	sb.WriteString("  }\n")

	// Add resource-specific configurations based on type
	addResourceSpecificConfig(&sb, resourceName, schema)

	sb.WriteString("}\n")

	// Per HashiCorp standards: "Avoid multiple examples unless configuration is particularly complex"
	// We generate only the basic example - advanced configurations should be in user guides, not reference docs

	return sb.String()
}

// addResourceSpecificConfig adds resource-specific example configurations.
// OneOf comments are manually curated to exclude deprecated options that are
// no longer valid in the current API version.
func addResourceSpecificConfig(sb *strings.Builder, resourceName string, schema *SchemaInfo) {
	switch resourceName {
	case "http_loadbalancer":
		sb.WriteString("\n  // One of the arguments from this list \"advertise_custom advertise_on_public advertise_on_public_default_vip do_not_advertise\" must be set\n\n")
		sb.WriteString("  advertise_on_public_default_vip = true\n\n")
		sb.WriteString("  // One of the arguments from this list \"api_specification disable_api_definition\" must be set\n\n")
		sb.WriteString("  disable_api_definition = true\n\n")
		sb.WriteString("  // One of the arguments from this list \"disable_api_discovery enable_api_discovery\" must be set\n\n")
		sb.WriteString("  enable_api_discovery {\n")
		sb.WriteString("    // One of the arguments from this list \"api_discovery_from_code_scan api_discovery_from_discovered_schema api_discovery_from_live_traffic\" must be set\n\n")
		sb.WriteString("    api_discovery_from_live_traffic {}\n\n")
		sb.WriteString("    discovered_api_settings {\n")
		sb.WriteString("      purge_duration_for_inactive_discovered_apis = \"30\"\n")
		sb.WriteString("    }\n\n")
		sb.WriteString("    // One of the arguments from this list \"disable_learn_from_redirect_traffic enable_learn_from_redirect_traffic\" must be set\n\n")
		sb.WriteString("    disable_learn_from_redirect_traffic = true\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  // One of the arguments from this list \"api_testing disable_api_testing\" must be set\n\n")
		sb.WriteString("  disable_api_testing = true\n\n")
		sb.WriteString("  // One of the arguments from this list \"captcha_challenge enable_challenge js_challenge no_challenge policy_based_challenge\" must be set\n\n")
		sb.WriteString("  js_challenge {\n")
		sb.WriteString("    cookie_expiry     = 3600\n")
		sb.WriteString("    custom_page       = \"\"\n")
		sb.WriteString("    js_script_delay   = 5000\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  domains = [\"app.example.com\", \"www.example.com\"]\n\n")
		sb.WriteString("  // One of the arguments from this list \"cookie_stickiness least_active random ring_hash round_robin source_ip_stickiness\" must be set\n\n")
		sb.WriteString("  round_robin = true\n\n")
		sb.WriteString("  // One of the arguments from this list \"http https https_auto_cert\" must be set\n\n")
		sb.WriteString("  https_auto_cert {\n")
		sb.WriteString("    http_redirect = true\n")
		sb.WriteString("    add_hsts      = true\n\n")
		sb.WriteString("    // One of the arguments from this list \"default_header no_headers server_name\" must be set\n\n")
		sb.WriteString("    default_header {}\n\n")
		sb.WriteString("    tls_config {\n")
		sb.WriteString("      // One of the arguments from this list \"custom_security default_security low_security medium_security\" must be set\n\n")
		sb.WriteString("      default_security {}\n")
		sb.WriteString("    }\n\n")
		sb.WriteString("    // One of the arguments from this list \"no_mtls use_mtls\" must be set\n\n")
		sb.WriteString("    no_mtls {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  // One of the arguments from this list \"disable_malicious_user_detection enable_malicious_user_detection\" must be set\n\n")
		sb.WriteString("  enable_malicious_user_detection = true\n\n")
		sb.WriteString("  // One of the arguments from this list \"disable_malware_protection malware_protection_settings\" must be set\n\n")
		sb.WriteString("  disable_malware_protection = true\n\n")
		sb.WriteString("  // One of the arguments from this list \"api_rate_limit disable_rate_limit rate_limit\" must be set\n\n")
		sb.WriteString("  rate_limit {\n")
		sb.WriteString("    rate_limiter {\n")
		sb.WriteString("      name      = \"example-rate-limiter\"\n")
		sb.WriteString(fmt.Sprintf("      namespace = \"%s\"\n", getNamespaceForReference("rate_limiter")))
		sb.WriteString("    }\n")
		sb.WriteString("    no_ip_allowed_list {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  // One of the arguments from this list \"default_sensitive_data_policy sensitive_data_policy\" must be set\n\n")
		sb.WriteString("  default_sensitive_data_policy = true\n\n")
		sb.WriteString("  // One of the arguments from this list \"active_service_policies no_service_policies service_policies_from_namespace\" must be set\n\n")
		sb.WriteString("  active_service_policies {\n")
		sb.WriteString("    policies {\n")
		sb.WriteString("      name      = \"example-service-policy\"\n")
		sb.WriteString(fmt.Sprintf("      namespace = \"%s\"\n", getNamespaceForReference("service_policy")))
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  // One of the arguments from this list \"disable_threat_mesh enable_threat_mesh\" must be set\n\n")
		sb.WriteString("  enable_threat_mesh = true\n\n")
		sb.WriteString("  // One of the arguments from this list \"disable_trust_client_ip_headers enable_trust_client_ip_headers\" must be set\n\n")
		sb.WriteString("  disable_trust_client_ip_headers = true\n\n")
		sb.WriteString("  // One of the arguments from this list \"user_id_client_ip user_identification\" must be set\n\n")
		sb.WriteString("  user_identification {\n")
		sb.WriteString("    name      = \"example-user-identification\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("user_identification")))
		sb.WriteString("  }\n\n")
		sb.WriteString("  // One of the arguments from this list \"app_firewall disable_waf\" must be set\n\n")
		sb.WriteString("  app_firewall {\n")
		sb.WriteString("    name      = \"example-app-firewall\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("app_firewall")))
		sb.WriteString("  }\n\n")
		sb.WriteString("  // One of the arguments from this list \"bot_defense bot_defense_advanced disable_bot_defense\" must be set\n\n")
		sb.WriteString("  bot_defense {\n")
		sb.WriteString("    policy {\n")
		sb.WriteString("      // One of the arguments from this list \"js_download_path js_insert_all_pages js_insert_all_pages_except\" must be set\n\n")
		sb.WriteString("      js_insert_all_pages {\n")
		sb.WriteString("        javascript_location = \"AFTER_HEAD\"\n")
		sb.WriteString("      }\n\n")
		sb.WriteString("      // One of the arguments from this list \"disable_mobile_sdk enable_mobile_sdk\" must be set\n\n")
		sb.WriteString("      disable_mobile_sdk {}\n")
		sb.WriteString("    }\n")
		sb.WriteString("    regional_endpoint = \"US\"\n")
		sb.WriteString("    timeout           = 1000\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  // Default route pools configuration\n")
		sb.WriteString("  default_route_pools {\n")
		sb.WriteString("    pool {\n")
		sb.WriteString("      name      = \"example-origin-pool\"\n")
		sb.WriteString(fmt.Sprintf("      namespace = \"%s\"\n", getNamespaceForReference("origin_pool")))
		sb.WriteString("    }\n")
		sb.WriteString("    weight   = 1\n")
		sb.WriteString("    priority = 1\n")
		sb.WriteString("  }\n")

	case "origin_pool":
		sb.WriteString("\n  // Origin servers configuration\n")
		sb.WriteString("  origin_servers {\n")
		sb.WriteString("    // One of the arguments from this list \"consul_service custom_endpoint_object k8s_service private_ip private_name public_ip public_name vn_private_ip vn_private_name\" must be set\n\n")
		sb.WriteString("    public_name {\n")
		sb.WriteString("      dns_name         = \"origin.example.com\"\n")
		sb.WriteString("      refresh_interval = 60\n")
		sb.WriteString("    }\n\n")
		sb.WriteString("    labels = {\n")
		sb.WriteString("      \"app\" = \"backend\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  origin_servers {\n")
		sb.WriteString("    // One of the arguments from this list \"consul_service custom_endpoint_object k8s_service private_ip private_name public_ip public_name vn_private_ip vn_private_name\" must be set\n\n")
		sb.WriteString("    k8s_service {\n")
		sb.WriteString("      service_name = \"backend-svc\"\n\n")
		sb.WriteString("      // One of the arguments from this list \"inside_network outside_network vk8s_networks\" must be set\n\n")
		sb.WriteString("      vk8s_networks {}\n\n")
		sb.WriteString("      site_locator {\n")
		sb.WriteString("        // One of the arguments from this list \"site virtual_site\" must be set\n\n")
		sb.WriteString("        site {\n")
		sb.WriteString("          name      = \"example-site\"\n")
		sb.WriteString(fmt.Sprintf("          namespace = \"%s\"\n", getNamespaceForReference("securemesh_site")))
		sb.WriteString("        }\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  port = 443\n\n")
		sb.WriteString("  // One of the arguments from this list \"no_tls use_tls\" must be set\n\n")
		sb.WriteString("  use_tls {\n")
		sb.WriteString("    // One of the arguments from this list \"disable_sni sni use_host_header_as_sni\" must be set\n\n")
		sb.WriteString("    sni = \"backend.example.com\"\n\n")
		sb.WriteString("    tls_config {\n")
		sb.WriteString("      // One of the arguments from this list \"custom_security default_security low_security medium_security\" must be set\n\n")
		sb.WriteString("      default_security {}\n")
		sb.WriteString("    }\n\n")
		sb.WriteString("    // One of the arguments from this list \"no_mtls use_mtls use_mtls_obj\" must be set\n\n")
		sb.WriteString("    no_mtls {}\n\n")
		sb.WriteString("    // One of the arguments from this list \"skip_server_verification use_server_verification volterra_trusted_ca\" must be set\n\n")
		sb.WriteString("    volterra_trusted_ca {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  // Health check configuration\n")
		sb.WriteString("  healthcheck {\n")
		sb.WriteString("    name      = \"example-healthcheck\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("healthcheck")))
		sb.WriteString("  }\n\n")
		sb.WriteString("  // Load balancing configuration\n")
		sb.WriteString("  endpoint_selection     = \"LOCAL_PREFERRED\"\n")
		sb.WriteString("  loadbalancer_algorithm = \"ROUND_ROBIN\"\n")

	case "healthcheck":
		sb.WriteString("\n  // One of the arguments from this list \"http_health_check tcp_health_check udp_icmp_health_check\" must be set\n\n")
		sb.WriteString("  http_health_check {\n")
		sb.WriteString("    // One of the arguments from this list \"host_header use_origin_server_name\" must be set\n\n")
		sb.WriteString("    use_origin_server_name {}\n\n")
		sb.WriteString("    path                  = \"/health\"\n")
		sb.WriteString("    use_http2             = false\n")
		sb.WriteString("    expected_status_codes = [\"200\"]\n\n")
		sb.WriteString("    // One of the arguments from this list \"headers request_headers_to_remove\" must be set\n\n")
		sb.WriteString("    headers = {\n")
		sb.WriteString("      \"x-health-check\" = \"true\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  healthy_threshold   = 3\n")
		sb.WriteString("  unhealthy_threshold = 3\n")
		sb.WriteString("  interval            = 15\n")
		sb.WriteString("  timeout             = 5\n")

	case "app_firewall":
		sb.WriteString("\n  // One of the arguments from this list \"blocking monitoring\" must be set\n\n")
		sb.WriteString("  blocking {}\n\n")
		sb.WriteString("  // One of the arguments from this list \"blocking_page use_default_blocking_page\" must be set\n\n")
		sb.WriteString("  use_default_blocking_page {}\n\n")
		sb.WriteString("  // One of the arguments from this list \"bot_protection_setting default_bot_setting\" must be set\n\n")
		sb.WriteString("  bot_protection_setting {\n")
		sb.WriteString("    malicious_bot_action  = \"BLOCK\"\n")
		sb.WriteString("    suspicious_bot_action = \"REPORT\"\n")
		sb.WriteString("    good_bot_action       = \"REPORT\"\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  // One of the arguments from this list \"ai_risk_based_blocking default_detection_settings detection_settings\" must be set\n\n")
		sb.WriteString("  default_detection_settings {}\n\n")
		sb.WriteString("  // One of the arguments from this list \"allow_all_response_codes allowed_response_codes\" must be set\n\n")
		sb.WriteString("  allow_all_response_codes {}\n")

	case "service_policy":
		sb.WriteString("\n  # Service Policy configuration\n")
		sb.WriteString("  algo = \"FIRST_MATCH\"\n\n")
		sb.WriteString("  # Allow specific paths\n")
		sb.WriteString("  rules {\n")
		sb.WriteString("    metadata {\n")
		sb.WriteString("      name = \"allow-api\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("    spec {\n")
		sb.WriteString("      action = \"ALLOW\"\n")
		sb.WriteString("      path {\n")
		sb.WriteString("        prefix = \"/api/\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")

	case "aws_vpc_site":
		sb.WriteString("\n  # AWS VPC Site configuration\n")
		sb.WriteString("  aws_region = \"us-west-2\"\n\n")
		sb.WriteString("  # AWS credentials reference\n")
		sb.WriteString("  aws_cred {\n")
		sb.WriteString("    name      = \"aws-credentials\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("cloud_credentials")))
		sb.WriteString("  }\n\n")
		sb.WriteString("  # VPC configuration\n")
		sb.WriteString("  vpc {\n")
		sb.WriteString("    new_vpc {\n")
		sb.WriteString("      name_tag     = \"f5xc-vpc\"\n")
		sb.WriteString("      primary_ipv4 = \"10.0.0.0/16\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Instance type\n")
		sb.WriteString("  instance_type = \"t3.xlarge\"\n\n")
		sb.WriteString("  # Ingress/Egress gateway\n")
		sb.WriteString("  ingress_egress_gw {\n")
		sb.WriteString("    aws_certified_hw = \"aws-byol-multi-nic-voltmesh\"\n")
		sb.WriteString("    az_nodes {\n")
		sb.WriteString("      aws_az_name = \"us-west-2a\"\n")
		sb.WriteString("      inside_subnet {\n")
		sb.WriteString("        subnet_param {\n")
		sb.WriteString("          ipv4 = \"10.0.1.0/24\"\n")
		sb.WriteString("        }\n")
		sb.WriteString("      }\n")
		sb.WriteString("      outside_subnet {\n")
		sb.WriteString("        subnet_param {\n")
		sb.WriteString("          ipv4 = \"10.0.2.0/24\"\n")
		sb.WriteString("        }\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # No worker nodes by default\n")
		sb.WriteString("  no_worker_nodes {}\n")

	case "azure_vnet_site":
		sb.WriteString("\n  # Azure VNET Site configuration\n")
		sb.WriteString("  azure_region = \"westus2\"\n\n")
		sb.WriteString("  # Azure credentials reference\n")
		sb.WriteString("  azure_cred {\n")
		sb.WriteString("    name      = \"azure-credentials\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("cloud_credentials")))
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Resource group\n")
		sb.WriteString("  resource_group = \"f5xc-rg\"\n\n")
		sb.WriteString("  # VNET configuration\n")
		sb.WriteString("  vnet {\n")
		sb.WriteString("    new_vnet {\n")
		sb.WriteString("      name       = \"f5xc-vnet\"\n")
		sb.WriteString("      primary_ipv4 = \"10.0.0.0/16\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Machine type\n")
		sb.WriteString("  machine_type = \"Standard_D3_v2\"\n\n")
		sb.WriteString("  # Ingress/Egress gateway\n")
		sb.WriteString("  ingress_egress_gw {\n")
		sb.WriteString("    azure_certified_hw = \"azure-byol-multi-nic-voltmesh\"\n")
		sb.WriteString("    az_nodes {\n")
		sb.WriteString("      azure_az = \"1\"\n")
		sb.WriteString("      inside_subnet {\n")
		sb.WriteString("        subnet_param {\n")
		sb.WriteString("          ipv4 = \"10.0.1.0/24\"\n")
		sb.WriteString("        }\n")
		sb.WriteString("      }\n")
		sb.WriteString("      outside_subnet {\n")
		sb.WriteString("        subnet_param {\n")
		sb.WriteString("          ipv4 = \"10.0.2.0/24\"\n")
		sb.WriteString("        }\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # No worker nodes by default\n")
		sb.WriteString("  no_worker_nodes {}\n")

	case "gcp_vpc_site":
		sb.WriteString("\n  # GCP VPC Site configuration\n")
		sb.WriteString("  gcp_region = \"us-west1\"\n\n")
		sb.WriteString("  # GCP credentials reference\n")
		sb.WriteString("  cloud_credentials {\n")
		sb.WriteString("    name      = \"gcp-credentials\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("cloud_credentials")))
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Instance type\n")
		sb.WriteString("  instance_type = \"n1-standard-4\"\n\n")
		sb.WriteString("  # Ingress/Egress gateway\n")
		sb.WriteString("  ingress_egress_gw {\n")
		sb.WriteString("    gcp_certified_hw = \"gcp-byol-multi-nic-voltmesh\"\n")
		sb.WriteString("    node_number      = 1\n")
		sb.WriteString("    inside_network {\n")
		sb.WriteString("      new_network {\n")
		sb.WriteString("        name = \"inside-network\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("    outside_network {\n")
		sb.WriteString("      new_network {\n")
		sb.WriteString("        name = \"outside-network\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("    inside_subnet {\n")
		sb.WriteString("      new_subnet {\n")
		sb.WriteString("        subnet_name  = \"inside-subnet\"\n")
		sb.WriteString("        primary_ipv4 = \"10.0.1.0/24\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("    outside_subnet {\n")
		sb.WriteString("      new_subnet {\n")
		sb.WriteString("        subnet_name  = \"outside-subnet\"\n")
		sb.WriteString("        primary_ipv4 = \"10.0.2.0/24\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # No worker nodes by default\n")
		sb.WriteString("  no_worker_nodes {}\n")

	case "tcp_loadbalancer":
		sb.WriteString("\n  # TCP Load Balancer specific configuration\n")
		sb.WriteString("  listen_port = 8443\n\n")
		sb.WriteString("  # Advertise on public internet\n")
		sb.WriteString("  advertise_on_internet {\n")
		sb.WriteString("    default_vip {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Origin pools\n")
		sb.WriteString("  origin_pools_weights {\n")
		sb.WriteString("    pool {\n")
		sb.WriteString("      name      = \"example-tcp-pool\"\n")
		sb.WriteString(fmt.Sprintf("      namespace = \"%s\"\n", getNamespaceForReference("origin_pool")))
		sb.WriteString("    }\n")
		sb.WriteString("    weight = 1\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # DNS for TCP load balancer\n")
		sb.WriteString("  dns_volterra_managed = true\n\n")
		sb.WriteString("  # No retract cluster by default\n")
		sb.WriteString("  retract_cluster {}\n")

	case "certificate":
		sb.WriteString("\n  # Certificate configuration\n")
		sb.WriteString("  certificate_url = \"string:///LS0tLS1CRUdJTi...\"\n")
		sb.WriteString("  private_key {\n")
		sb.WriteString("    clear_secret_info {\n")
		sb.WriteString("      url = \"string:///LS0tLS1CRUdJTi...\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")

	case "cloud_credentials":
		sb.WriteString("\n  # Cloud Credentials configuration\n")
		sb.WriteString("  # AWS credentials example\n")
		sb.WriteString("  aws_secret_key {\n")
		sb.WriteString("    access_key = \"AKIAIOSFODNN7EXAMPLE\"\n")
		sb.WriteString("    secret_key {\n")
		sb.WriteString("      clear_secret_info {\n")
		sb.WriteString("        url = \"string:///d0phbmVzc2VjcmV0a2V5\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")

	case "virtual_site":
		sb.WriteString("\n  # Virtual Site configuration\n")
		sb.WriteString("  site_type = \"CUSTOMER_EDGE\"\n\n")
		sb.WriteString("  # Site selector expression\n")
		sb.WriteString("  site_selector {\n")
		sb.WriteString("    expressions = [\"region in (us-west-2, us-east-1)\"]\n")
		sb.WriteString("  }\n")

	case "dns_load_balancer":
		sb.WriteString("\n  # DNS Load Balancer configuration\n")
		sb.WriteString("  record_type = \"A\"\n\n")
		sb.WriteString("  # DNS zone reference\n")
		sb.WriteString("  dns_zone {\n")
		sb.WriteString("    name      = \"example-dns-zone\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("dns_zone")))
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Rule-based load balancing\n")
		sb.WriteString("  rule_list {\n")
		sb.WriteString("    rules {\n")
		sb.WriteString("      geo_location_set {\n")
		sb.WriteString("        name      = \"us-geo\"\n")
		sb.WriteString(fmt.Sprintf("        namespace = \"%s\"\n", getNamespaceForReference("geo_location_set")))
		sb.WriteString("      }\n")
		sb.WriteString("      pool {\n")
		sb.WriteString("        name      = \"us-pool\"\n")
		sb.WriteString(fmt.Sprintf("        namespace = \"%s\"\n", getNamespaceForReference("dns_lb_pool")))
		sb.WriteString("      }\n")
		sb.WriteString("      score = 100\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")

	case "alert_policy":
		sb.WriteString("\n  # Alert Policy configuration\n")
		sb.WriteString("  # Alert receivers\n")
		sb.WriteString("  receivers {\n")
		sb.WriteString("    name      = \"slack-receiver\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("alert_receiver")))
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Alert routes\n")
		sb.WriteString("  routes {\n")
		sb.WriteString("    any {}\n")
		sb.WriteString("    send {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Notification parameters\n")
		sb.WriteString("  notification_parameters {\n")
		sb.WriteString("    default {}\n")
		sb.WriteString("    group_wait     = \"30s\"\n")
		sb.WriteString("    group_interval = \"1m\"\n")
		sb.WriteString("  }\n")

	case "alert_receiver":
		sb.WriteString("\n  # Alert Receiver configuration\n")
		sb.WriteString("  # Slack configuration\n")
		sb.WriteString("  slack {\n")
		sb.WriteString("    url = \"https://your-slack-webhook-url\"\n")
		sb.WriteString("  }\n")

	case "namespace":
		sb.WriteString("\n  # Namespace configuration\n")
		sb.WriteString("  description = \"Example namespace for application workloads\"\n")

	case "role":
		sb.WriteString("\n  # Role configuration\n")
		sb.WriteString("  role_type = \"CUSTOM\"\n\n")
		sb.WriteString("  # API groups for this role\n")
		sb.WriteString("  api_groups {\n")
		sb.WriteString("    name = \"read-http-lb\"\n")
		sb.WriteString("    api_group_elements {\n")
		sb.WriteString("      api_group     = \"ves.io.schema.views.http_loadbalancer\"\n")
		sb.WriteString("      resource_type = \"http_loadbalancer\"\n")
		sb.WriteString("      verbs {\n")
		sb.WriteString("        get  = true\n")
		sb.WriteString("        list = true\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")

	case "network_connector":
		sb.WriteString("\n  # Network Connector configuration\n")
		sb.WriteString("  # Direct connection\n")
		sb.WriteString("  sli_to_global_dr {\n")
		sb.WriteString("    global_vn {\n")
		sb.WriteString("      name      = \"global-network\"\n")
		sb.WriteString(fmt.Sprintf("      namespace = \"%s\"\n", getNamespaceForReference("virtual_network")))
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Disable forward proxy\n")
		sb.WriteString("  disable_forward_proxy {}\n")

	case "route":
		sb.WriteString("\n  # Route configuration\n")
		sb.WriteString("  routes {\n")
		sb.WriteString("    match {\n")
		sb.WriteString("      path {\n")
		sb.WriteString("        prefix = \"/api/\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("    route_destination {\n")
		sb.WriteString("      destinations {\n")
		sb.WriteString("        cluster {\n")
		sb.WriteString("          name      = \"api-cluster\"\n")
		sb.WriteString(fmt.Sprintf("          namespace = \"%s\"\n", getNamespaceForReference("cluster")))
		sb.WriteString("        }\n")
		sb.WriteString("        weight = 100\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")

	case "discovery":
		sb.WriteString("\n  # Discovery configuration\n")
		sb.WriteString("  discovery_k8s {\n")
		sb.WriteString("    access_info {\n")
		sb.WriteString("      kubeconfig_url {\n")
		sb.WriteString("        clear_secret_info {\n")
		sb.WriteString("          url = \"string:///base64-kubeconfig\"\n")
		sb.WriteString("        }\n")
		sb.WriteString("      }\n")
		sb.WriteString("      isolated {}\n")
		sb.WriteString("    }\n")
		sb.WriteString("    publish_info {\n")
		sb.WriteString("      disable {}\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Site selection\n")
		sb.WriteString("  where {\n")
		sb.WriteString("    site {\n")
		sb.WriteString("      ref {\n")
		sb.WriteString("        name      = \"example-site\"\n")
		sb.WriteString(fmt.Sprintf("        namespace = \"%s\"\n", getNamespaceForReference("securemesh_site")))
		sb.WriteString("      }\n")
		sb.WriteString("      network_type = \"VIRTUAL_NETWORK_SITE_LOCAL_INSIDE\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")

	case "rate_limiter":
		sb.WriteString("\n  # Rate Limiter configuration\n")
		sb.WriteString("  total_number  = 100\n")
		sb.WriteString("  unit         = \"MINUTE\"\n")
		sb.WriteString("  burst_multiplier = 10\n")

	case "cdn_loadbalancer":
		sb.WriteString("\n  # CDN Load Balancer configuration\n")
		sb.WriteString("  domains = [\"cdn.example.com\"]\n\n")
		sb.WriteString("  # Origin pool\n")
		sb.WriteString("  origin_pool {\n")
		sb.WriteString("    public_name {\n")
		sb.WriteString("      dns_name = \"origin.example.com\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("    follow_origin_redirect = true\n")
		sb.WriteString("    no_tls {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Cache TTL settings\n")
		sb.WriteString("  cache_ttl_options {\n")
		sb.WriteString("    cache_ttl_default = \"1h\"\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # HTTP protocol\n")
		sb.WriteString("  https_auto_cert {\n")
		sb.WriteString("    http_redirect = true\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Add location header\n")
		sb.WriteString("  add_location = true\n")

	case "log_receiver":
		sb.WriteString("\n  # Log Receiver configuration\n")
		sb.WriteString("  # HTTP receiver example\n")
		sb.WriteString("  http_receiver {\n")
		sb.WriteString("    uri = \"https://logs.example.com/ingest\"\n")
		sb.WriteString("    batch {\n")
		sb.WriteString("      max_bytes   = 1048576\n")
		sb.WriteString("      max_events  = 100\n")
		sb.WriteString("      timeout_seconds = 5\n")
		sb.WriteString("    }\n")
		sb.WriteString("    no_tls_verify_hostname {}\n")
		sb.WriteString("    no_compression {}\n")
		sb.WriteString("  }\n")

	case "forward_proxy_policy":
		sb.WriteString("\n  # Forward Proxy Policy configuration\n")
		sb.WriteString("  proxy_label_selector {\n")
		sb.WriteString("    expressions = [\"app in (web, api)\"]\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  drp_http_connect {\n")
		sb.WriteString("    any_proxy {}\n")
		sb.WriteString("    rule_list {\n")
		sb.WriteString("      rules {\n")
		sb.WriteString("        metadata {\n")
		sb.WriteString("          name = \"allow-external\"\n")
		sb.WriteString("        }\n")
		sb.WriteString("        spec {\n")
		sb.WriteString("          action = \"ALLOW\"\n")
		sb.WriteString("          dst_list {\n")
		sb.WriteString("            any_dst {}\n")
		sb.WriteString("          }\n")
		sb.WriteString("        }\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")

	case "network_policy":
		sb.WriteString("\n  # Network Policy configuration\n")
		sb.WriteString("  endpoint {\n")
		sb.WriteString("    any {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  ingress_rules {\n")
		sb.WriteString("    metadata {\n")
		sb.WriteString("      name = \"allow-http\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("    spec {\n")
		sb.WriteString("      action = \"ALLOW\"\n")
		sb.WriteString("      any   = {}\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  egress_rules {\n")
		sb.WriteString("    metadata {\n")
		sb.WriteString("      name = \"allow-all-egress\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("    spec {\n")
		sb.WriteString("      action = \"ALLOW\"\n")
		sb.WriteString("      any   = {}\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")

	case "enhanced_firewall_policy":
		sb.WriteString("\n  # Enhanced Firewall Policy configuration\n")
		sb.WriteString("  rule_list {\n")
		sb.WriteString("    rules {\n")
		sb.WriteString("      metadata {\n")
		sb.WriteString("        name = \"allow-web-traffic\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("      allow {}\n")
		sb.WriteString("      advanced_action {\n")
		sb.WriteString("        action = \"LOG\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("      source_prefix_list {\n")
		sb.WriteString("        ip_prefix_set {\n")
		sb.WriteString("          name      = \"trusted-ips\"\n")
		sb.WriteString(fmt.Sprintf("          namespace = \"%s\"\n", getNamespaceForReference("ip_prefix_set")))
		sb.WriteString("        }\n")
		sb.WriteString("      }\n")
		sb.WriteString("      all_traffic {}\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")

	case "secret_policy":
		sb.WriteString("\n  # Secret Policy configuration\n")
		sb.WriteString("  algo = \"DENY_OVERRIDES\"\n\n")
		sb.WriteString("  rules {\n")
		sb.WriteString("    metadata {\n")
		sb.WriteString("      name = \"allow-team-secrets\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("    spec {\n")
		sb.WriteString("      action = \"ALLOW\"\n")
		sb.WriteString("      secret_match {\n")
		sb.WriteString("        regex {\n")
		sb.WriteString("          regex = \"team-.*\"\n")
		sb.WriteString("        }\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")

	case "site_mesh_group":
		sb.WriteString("\n  # Site Mesh Group configuration\n")
		sb.WriteString("  type = \"SITE_MESH_GROUP_TYPE_FULL_MESH\"\n\n")
		sb.WriteString("  # Control and data plane settings\n")
		sb.WriteString("  full_mesh {\n")
		sb.WriteString("    control_and_data_plane_mesh {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Hub status\n")
		sb.WriteString("  hub {}\n\n")
		sb.WriteString("  # Virtual site reference\n")
		sb.WriteString("  virtual_site {\n")
		sb.WriteString("    name      = \"example-virtual-site\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("virtual_site")))
		sb.WriteString("  }\n")

	case "voltstack_site":
		sb.WriteString("\n  # Voltstack Site configuration\n")
		sb.WriteString("  # Kubernetes configuration\n")
		sb.WriteString("  k8s_cluster {\n")
		sb.WriteString("    name      = \"example-k8s-cluster\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("k8s_cluster")))
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Master nodes configuration\n")
		sb.WriteString("  master_nodes = [\"master1.example.com\"]\n\n")
		sb.WriteString("  # Default fleet configuration\n")
		sb.WriteString("  default_fleet_config {\n")
		sb.WriteString("    no_bond_devices {}\n")
		sb.WriteString("    no_dc_cluster_group {}\n")
		sb.WriteString("    default_storage_config {}\n")
		sb.WriteString("    no_gpu {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Disable HA by default\n")
		sb.WriteString("  disable_ha {}\n\n")
		sb.WriteString("  # No worker nodes\n")
		sb.WriteString("  no_worker_nodes {}\n")

	case "k8s_cluster":
		sb.WriteString("\n  # Kubernetes Cluster configuration\n")
		sb.WriteString("  # Use custom local domain\n")
		sb.WriteString("  use_custom_cluster_role_bindings {\n")
		sb.WriteString("    cluster_role_bindings {\n")
		sb.WriteString("      name      = \"admin-binding\"\n")
		sb.WriteString(fmt.Sprintf("      namespace = \"%s\"\n", getNamespaceForReference("k8s_cluster_role_binding")))
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  cluster_wide_app_list {\n")
		sb.WriteString("    cluster_wide_apps {\n")
		sb.WriteString("      name      = \"nginx-ingress\"\n")
		sb.WriteString(fmt.Sprintf("      namespace = \"%s\"\n", getNamespaceForReference("k8s_cluster")))
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  local_access_config {\n")
		sb.WriteString("    local_domain = \"cluster.local\"\n")
		sb.WriteString("    default_port {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  global_access_enable {}\n")

	case "virtual_k8s":
		sb.WriteString("\n  // Virtual site selection for workload deployment\n")
		sb.WriteString("  vsite_refs {\n")
		sb.WriteString("    name      = \"example-virtual-site\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("virtual_site")))
		sb.WriteString("  }\n\n")
		sb.WriteString("  // One of the arguments from this list \"disabled isolated\" must be set\n\n")
		sb.WriteString("  isolated {}\n\n")
		sb.WriteString("  // Default workload flavor reference\n")
		sb.WriteString("  default_flavor_ref {\n")
		sb.WriteString("    name      = \"example-workload-flavor\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("workload_flavor")))
		sb.WriteString("  }\n")

	case "virtual_network":
		sb.WriteString("\n  // One of the arguments from this list \"global_network legacy_type site_local_inside_network site_local_network srv6_network\" must be set\n\n")
		sb.WriteString("  site_local_network {}\n\n")
		sb.WriteString("  // Static routes configuration (optional)\n")
		sb.WriteString("  static_routes {\n")
		sb.WriteString("    ip_prefixes = [\"10.0.0.0/8\"]\n\n")
		sb.WriteString("    // One of the arguments from this list \"default_gateway ip_address node_interface\" must be set\n\n")
		sb.WriteString("    default_gateway {}\n\n")
		sb.WriteString("    attrs = [\"ROUTE_ATTR_INSTALL_FORWARDING\"]\n")
		sb.WriteString("  }\n")

	case "fleet":
		sb.WriteString("\n  # Fleet configuration\n")
		sb.WriteString("  fleet_label = \"env=production\"\n\n")
		sb.WriteString("  # Network connectors\n")
		sb.WriteString("  inside_virtual_network {\n")
		sb.WriteString("    name      = \"inside-network\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("virtual_network")))
		sb.WriteString("  }\n\n")
		sb.WriteString("  outside_virtual_network {\n")
		sb.WriteString("    name      = \"outside-network\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("virtual_network")))
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Default config\n")
		sb.WriteString("  default_config {}\n")

	case "workload":
		sb.WriteString("\n  # Workload configuration\n")
		sb.WriteString("  # Container configuration\n")
		sb.WriteString("  containers {\n")
		sb.WriteString("    name = \"web\"\n")
		sb.WriteString("    image {\n")
		sb.WriteString("      name       = \"nginx\"\n")
		sb.WriteString("      public     = {}\n")
		sb.WriteString("      pull_policy = \"IMAGE_PULL_POLICY_ALWAYS\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Deploy on regional edge\n")
		sb.WriteString("  deploy_on_re {\n")
		sb.WriteString("    virtual_site {\n")
		sb.WriteString("      name      = \"example-virtual-site\"\n")
		sb.WriteString(fmt.Sprintf("      namespace = \"%s\"\n", getNamespaceForReference("virtual_site")))
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")

	case "udp_loadbalancer":
		sb.WriteString("\n  # UDP Load Balancer configuration\n")
		sb.WriteString("  listen_port = 53\n\n")
		sb.WriteString("  # Advertise on public internet\n")
		sb.WriteString("  advertise_on_internet {\n")
		sb.WriteString("    default_vip {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Origin pools\n")
		sb.WriteString("  origin_pools_weights {\n")
		sb.WriteString("    pool {\n")
		sb.WriteString("      name      = \"dns-pool\"\n")
		sb.WriteString(fmt.Sprintf("      namespace = \"%s\"\n", getNamespaceForReference("origin_pool")))
		sb.WriteString("    }\n")
		sb.WriteString("    weight = 1\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # DNS for UDP load balancer\n")
		sb.WriteString("  dns_volterra_managed = true\n")

	case "dns_zone":
		sb.WriteString("\n  # DNS Zone configuration\n")
		sb.WriteString("  # Primary DNS zone\n")
		sb.WriteString("  primary {\n")
		sb.WriteString("    soa_record_parameters {\n")
		sb.WriteString("      refresh    = 86400\n")
		sb.WriteString("      retry      = 7200\n")
		sb.WriteString("      expire     = 3600000\n")
		sb.WriteString("      ttl        = 86400\n")
		sb.WriteString("      neg_ttl    = 1800\n")
		sb.WriteString("    }\n")
		sb.WriteString("    default_rr_set_group {}\n")
		sb.WriteString("    default_soa_parameters {}\n")
		sb.WriteString("    dnssec_mode {\n")
		sb.WriteString("      disable {}\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")

	case "api_definition":
		sb.WriteString("\n  # API Definition configuration\n")
		sb.WriteString("  # OpenAPI spec\n")
		sb.WriteString("  swagger_specs = [\"string:///base64-openapi-spec\"]\n\n")
		sb.WriteString("  # Non-validation mode\n")
		sb.WriteString("  non_validation_mode {}\n")

	case "api_credential":
		sb.WriteString("\n  # API Credential configuration\n")
		sb.WriteString("  api_credential_type = \"API_CERTIFICATE\"\n\n")
		sb.WriteString("  # Expiration settings\n")
		sb.WriteString("  expiration_timestamp = \"2025-12-31T23:59:59Z\"\n\n")
		sb.WriteString("  # Active state\n")
		sb.WriteString("  active = true\n")

	case "malicious_user_mitigation":
		sb.WriteString("\n  # Malicious User Mitigation configuration\n")
		sb.WriteString("  # Detection rules\n")
		sb.WriteString("  rules {\n")
		sb.WriteString("    threat_level = \"HIGH\"\n")
		sb.WriteString("    mitigation_action {\n")
		sb.WriteString("      block {\n")
		sb.WriteString("        body = \"Access denied\"\n")
		sb.WriteString("        status = \"403\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")

	case "user_identification":
		sb.WriteString("\n  # User Identification configuration\n")
		sb.WriteString("  rules {\n")
		sb.WriteString("    identifier_type = \"CLIENT_IP\"\n")
		sb.WriteString("    any_client {}\n")
		sb.WriteString("  }\n")

	case "ip_prefix_set":
		sb.WriteString("\n  # IP Prefix Set configuration\n")
		sb.WriteString("  prefix = [\"192.168.1.0/24\", \"10.0.0.0/8\"]\n")

	case "geo_location_set":
		sb.WriteString("\n  # Geo Location Set configuration\n")
		sb.WriteString("  country_codes = [\"US\", \"CA\", \"GB\"]\n")

	case "bgp":
		sb.WriteString("\n  # BGP configuration\n")
		sb.WriteString("  bgp_router_id = \"192.168.1.1\"\n\n")
		sb.WriteString("  bgp_peers {\n")
		sb.WriteString("    metadata {\n")
		sb.WriteString("      name = \"upstream-peer\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("    spec {\n")
		sb.WriteString("      peer_asn = 65000\n")
		sb.WriteString("      peer_address = \"192.168.1.2\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  local_asn = 65001\n\n")
		sb.WriteString("  # Site reference\n")
		sb.WriteString("  site {\n")
		sb.WriteString("    name      = \"example-site\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("securemesh_site")))
		sb.WriteString("  }\n")

	case "tunnel":
		sb.WriteString("\n  # Tunnel configuration\n")
		sb.WriteString("  remote_ip_address = \"203.0.113.1\"\n")
		sb.WriteString("  local_ip_address  = \"203.0.113.2\"\n\n")
		sb.WriteString("  # IPsec tunnel\n")
		sb.WriteString("  ipsec {\n")
		sb.WriteString("    psk = \"pre-shared-key-here\"\n")
		sb.WriteString("    ike_params {\n")
		sb.WriteString("      ike_version = \"IKE_V2\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Site reference\n")
		sb.WriteString("  site {\n")
		sb.WriteString("    name      = \"example-site\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("securemesh_site")))
		sb.WriteString("  }\n")

	case "securemesh_site":
		sb.WriteString("\n  # Secure Mesh Site configuration\n")
		sb.WriteString("  # Generic provider\n")
		sb.WriteString("  generic {\n")
		sb.WriteString("    not_managed {\n")
		sb.WriteString("      node_list {\n")
		sb.WriteString("        hostname  = \"node1.example.com\"\n")
		sb.WriteString("        public_ip = \"203.0.113.10\"\n")
		sb.WriteString("        type      = \"Control\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Master nodes\n")
		sb.WriteString("  master_nodes_count = 1\n\n")
		sb.WriteString("  # Default fleet config\n")
		sb.WriteString("  default_fleet_config {}\n\n")
		sb.WriteString("  # Disable HA\n")
		sb.WriteString("  disable_ha {}\n")

	case "aws_tgw_site":
		sb.WriteString("\n  # AWS TGW Site configuration\n")
		sb.WriteString("  aws_region = \"us-west-2\"\n\n")
		sb.WriteString("  # AWS credentials\n")
		sb.WriteString("  aws_cred {\n")
		sb.WriteString("    name      = \"aws-credentials\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("cloud_credentials")))
		sb.WriteString("  }\n\n")
		sb.WriteString("  # VPC configuration\n")
		sb.WriteString("  vpc {\n")
		sb.WriteString("    new_vpc {\n")
		sb.WriteString("      name_tag     = \"f5xc-tgw-vpc\"\n")
		sb.WriteString("      primary_ipv4 = \"10.0.0.0/16\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # TGW configuration\n")
		sb.WriteString("  tgw {\n")
		sb.WriteString("    new_tgw {\n")
		sb.WriteString("      name = \"f5xc-tgw\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Instance type\n")
		sb.WriteString("  instance_type = \"t3.xlarge\"\n\n")
		sb.WriteString("  # Service VPC\n")
		sb.WriteString("  services_vpc {\n")
		sb.WriteString("    aws_certified_hw = \"aws-byol-voltmesh\"\n")
		sb.WriteString("    az_nodes {\n")
		sb.WriteString("      aws_az_name = \"us-west-2a\"\n")
		sb.WriteString("      inside_subnet {\n")
		sb.WriteString("        subnet_param {\n")
		sb.WriteString("          ipv4 = \"10.0.1.0/24\"\n")
		sb.WriteString("        }\n")
		sb.WriteString("      }\n")
		sb.WriteString("      outside_subnet {\n")
		sb.WriteString("        subnet_param {\n")
		sb.WriteString("          ipv4 = \"10.0.2.0/24\"\n")
		sb.WriteString("        }\n")
		sb.WriteString("      }\n")
		sb.WriteString("      workload_subnet {\n")
		sb.WriteString("        subnet_param {\n")
		sb.WriteString("          ipv4 = \"10.0.3.0/24\"\n")
		sb.WriteString("        }\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # No worker nodes\n")
		sb.WriteString("  no_worker_nodes {}\n")

	case "proxy":
		sb.WriteString("\n  # Proxy configuration\n")
		sb.WriteString("  proxy_url = \"http://proxy.example.com:8080\"\n")

	case "token":
		sb.WriteString("\n  # Token configuration\n")
		sb.WriteString("  token_type = \"REGISTRATION_TOKEN\"\n")

	case "trusted_ca_list":
		sb.WriteString("\n  # Trusted CA List configuration\n")
		sb.WriteString("  trusted_ca_url = \"string:///LS0tLS1CRUdJTi...\"\n")

	default:
		// Generic blocks based on schema
		if len(schema.Blocks) > 0 {
			sb.WriteString("\n  # Resource-specific configuration\n")
			addGenericBlocks(sb, schema.Blocks, 1, 1)
		}
	}
}

func addGenericBlocks(sb *strings.Builder, blocks []BlockInfo, maxDepth, currentIndent int) {
	if maxDepth <= 0 {
		return
	}

	indent := strings.Repeat("  ", currentIndent)

	for i, block := range blocks {
		if i >= 3 { // Limit to first 3 blocks for readability
			break
		}

		// Skip very common or unnecessary blocks
		if block.Name == "metadata" || block.Name == "spec" || block.Name == "status" {
			continue
		}

		sb.WriteString(fmt.Sprintf("%s# %s\n", indent, truncateDescription(block.Description, 60)))
		sb.WriteString(fmt.Sprintf("%s%s {\n", indent, block.Name))

		// Add placeholder content
		innerIndent := strings.Repeat("  ", currentIndent+1)
		sb.WriteString(fmt.Sprintf("%s# Configure %s settings\n", innerIndent, block.Name))

		sb.WriteString(fmt.Sprintf("%s}\n", indent))
	}
}

func isComplexResource(resourceName string) bool {
	complexResources := map[string]bool{
		"http_loadbalancer":        true,
		"tcp_loadbalancer":         true,
		"origin_pool":              true,
		"aws_vpc_site":             true,
		"azure_vnet_site":          true,
		"gcp_vpc_site":             true,
		"securemesh_site_v2":       true,
		"securemesh_site":          true,
		"voltstack_site":           true,
		"service_policy":           true,
		"enhanced_firewall_policy": true,
		"forward_proxy_policy":     true,
		"network_policy":           true,
		"cdn_loadbalancer":         true,
		"k8s_cluster":              true,
		"virtual_k8s":              true,
		"fleet":                    true,
		"workload":                 true,
		"aws_tgw_site":             true,
	}
	return complexResources[resourceName]
}

func generateAdvancedExample(resourceName string, schema *SchemaInfo) string {
	var sb strings.Builder

	// Get the appropriate namespace for this resource type
	_, namespace := getNamespaceForResource(resourceName)

	humanName := toHumanName(resourceName)
	sb.WriteString(fmt.Sprintf("\n# Advanced %s with additional configuration\n", humanName))
	sb.WriteString(fmt.Sprintf("resource \"f5xc_%s\" \"advanced\" {\n", resourceName))
	sb.WriteString(fmt.Sprintf("  name      = \"advanced-%s\"\n", strings.ReplaceAll(resourceName, "_", "-")))
	sb.WriteString(fmt.Sprintf("  namespace = \"%s\"\n\n", namespace))

	sb.WriteString("  labels = {\n")
	sb.WriteString("    environment = \"staging\"\n")
	sb.WriteString("    team        = \"platform\"\n")
	sb.WriteString("    cost_center = \"engineering\"\n")
	sb.WriteString("  }\n\n")

	sb.WriteString("  annotations = {\n")
	sb.WriteString("    \"created_by\" = \"terraform\"\n")
	sb.WriteString("    \"version\"    = \"2.0\"\n")
	sb.WriteString("  }\n")

	// Add advanced resource-specific configs
	addAdvancedConfig(&sb, resourceName)

	sb.WriteString("}\n")

	return sb.String()
}

func addAdvancedConfig(sb *strings.Builder, resourceName string) {
	switch resourceName {
	case "http_loadbalancer":
		sb.WriteString("\n  # Multiple domains\n")
		sb.WriteString("  domains = [\"app.example.com\", \"www.example.com\", \"api.example.com\"]\n\n")
		sb.WriteString("  # Custom advertise configuration\n")
		sb.WriteString("  advertise_custom {\n")
		sb.WriteString("    advertise_where {\n")
		sb.WriteString("      site {\n")
		sb.WriteString("        site {\n")
		sb.WriteString("          name      = \"example-ce-site\"\n")
		sb.WriteString(fmt.Sprintf("          namespace = \"%s\"\n", getNamespaceForReference("securemesh_site")))
		sb.WriteString("        }\n")
		sb.WriteString("        network = \"SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("      use_default_port {}\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Multiple origin pools with weights\n")
		sb.WriteString("  default_route_pools {\n")
		sb.WriteString("    pool {\n")
		sb.WriteString("      name      = \"primary-pool\"\n")
		sb.WriteString(fmt.Sprintf("      namespace = \"%s\"\n", getNamespaceForReference("origin_pool")))
		sb.WriteString("    }\n")
		sb.WriteString("    weight   = 80\n")
		sb.WriteString("    priority = 1\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  default_route_pools {\n")
		sb.WriteString("    pool {\n")
		sb.WriteString("      name      = \"secondary-pool\"\n")
		sb.WriteString(fmt.Sprintf("      namespace = \"%s\"\n", getNamespaceForReference("origin_pool")))
		sb.WriteString("    }\n")
		sb.WriteString("    weight   = 20\n")
		sb.WriteString("    priority = 2\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # HTTPS with auto certificate\n")
		sb.WriteString("  https_auto_cert {\n")
		sb.WriteString("    http_redirect           = true\n")
		sb.WriteString("    add_hsts               = true\n")
		sb.WriteString("    tls_config {\n")
		sb.WriteString("      default_security {}\n")
		sb.WriteString("    }\n")
		sb.WriteString("    no_mtls {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # WAF configuration\n")
		sb.WriteString("  app_firewall {\n")
		sb.WriteString("    name      = \"example-waf\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("app_firewall")))
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Service policies\n")
		sb.WriteString("  active_service_policies {\n")
		sb.WriteString("    policies {\n")
		sb.WriteString("      name      = \"api-policy\"\n")
		sb.WriteString("      namespace = \"shared\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Rate limiting\n")
		sb.WriteString("  rate_limit {\n")
		sb.WriteString("    rate_limiter {\n")
		sb.WriteString("      name      = \"api-rate-limit\"\n")
		sb.WriteString("      namespace = \"shared\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("    no_ip_allowed_list {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # User identification\n")
		sb.WriteString("  user_identification {\n")
		sb.WriteString("    name      = \"user-id\"\n")
		sb.WriteString("    namespace = \"shared\"\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Bot defense\n")
		sb.WriteString("  bot_defense {\n")
		sb.WriteString("    policy {\n")
		sb.WriteString("      js_insert_all_pages {\n")
		sb.WriteString("        javascript_location = \"AFTER_HEAD\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("    regional_endpoint = \"US\"\n")
		sb.WriteString("    timeout = 1000\n")
		sb.WriteString("  }\n")

	case "origin_pool":
		sb.WriteString("\n  # Multiple origin servers\n")
		sb.WriteString("  origin_servers {\n")
		sb.WriteString("    public_name {\n")
		sb.WriteString("      dns_name         = \"origin1.example.com\"\n")
		sb.WriteString("      refresh_interval = 30\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  origin_servers {\n")
		sb.WriteString("    public_name {\n")
		sb.WriteString("      dns_name         = \"origin2.example.com\"\n")
		sb.WriteString("      refresh_interval = 30\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  origin_servers {\n")
		sb.WriteString("    k8s_service {\n")
		sb.WriteString("      service_name = \"backend-svc\"\n")
		sb.WriteString("      site_locator {\n")
		sb.WriteString("        site {\n")
		sb.WriteString("          name      = \"example-site\"\n")
		sb.WriteString(fmt.Sprintf("          namespace = \"%s\"\n", getNamespaceForReference("securemesh_site")))
		sb.WriteString("        }\n")
		sb.WriteString("      }\n")
		sb.WriteString("      vk8s_networks {}\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  port = 443\n\n")
		sb.WriteString("  # TLS with custom certificate\n")
		sb.WriteString("  use_tls {\n")
		sb.WriteString("    sni = \"backend.example.com\"\n")
		sb.WriteString("    tls_config {\n")
		sb.WriteString("      custom_security {\n")
		sb.WriteString("        cipher_suites = [\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\"]\n")
		sb.WriteString("        min_version   = \"TLS12\"\n")
		sb.WriteString("        max_version   = \"TLS13\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("    volterra_trusted_ca {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Health check\n")
		sb.WriteString("  healthcheck {\n")
		sb.WriteString("    name      = \"example-healthcheck\"\n")
		sb.WriteString(fmt.Sprintf("    namespace = \"%s\"\n", getNamespaceForReference("healthcheck")))
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Advanced load balancing\n")
		sb.WriteString("  endpoint_selection     = \"LOCAL_PREFERRED\"\n")
		sb.WriteString("  loadbalancer_algorithm = \"ROUND_ROBIN\"\n\n")
		sb.WriteString("  # Automatic origin server subsets\n")
		sb.WriteString("  origin_servers_subset_rule_list {\n")
		sb.WriteString("    origin_server_subset_rules {\n")
		sb.WriteString("      keys = [\"version\"]\n")
		sb.WriteString("      any_asn {}\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")
	}
}

func truncateDescription(desc string, maxLen int) string {
	if len(desc) <= maxLen {
		return desc
	}
	return desc[:maxLen-3] + "..."
}

// toHumanName wraps naming.ToHumanName for backward compatibility
func toHumanName(resourceName string) string {
	return naming.ToHumanName(resourceName)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func getKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
