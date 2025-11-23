//go:build ignore
// +build ignore

// This tool generates comprehensive example Terraform configurations for all F5 XC resources.
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
)

type SchemaInfo struct {
	ResourceName string
	TypeName     string
	Description  string
	Attributes   []AttributeInfo
	Blocks       []BlockInfo
}

type AttributeInfo struct {
	Name        string
	Type        string
	Description string
	Required    bool
	Optional    bool
}

type BlockInfo struct {
	Name        string
	Description string
	IsList      bool
	Attributes  []AttributeInfo
	Blocks      []BlockInfo
}

func main() {
	providerDir := "internal/provider"
	examplesDir := "examples/resources"

	files, err := filepath.Glob(filepath.Join(providerDir, "*_resource.go"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding resource files: %v\n", err)
		os.Exit(1)
	}

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

		// Run terraform fmt to ensure proper formatting
		cmd := exec.Command("terraform", "fmt", exampleFile)
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: terraform fmt failed for %s: %v\n", exampleFile, err)
		}

		fmt.Printf("Generated: %s\n", exampleFile)
	}
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

	// Parse attributes
	schema.Attributes = parseAttributes(text)

	// Parse blocks
	schema.Blocks = parseBlocks(text)

	return schema
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

	// Generate basic example
	sb.WriteString(fmt.Sprintf("# Basic %s configuration\n", humanName))
	sb.WriteString(fmt.Sprintf("resource \"f5xc_%s\" \"example\" {\n", resourceName))
	sb.WriteString(fmt.Sprintf("  name      = \"example-%s\"\n", strings.ReplaceAll(resourceName, "_", "-")))
	sb.WriteString("  namespace = \"system\"\n\n")

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

func addResourceSpecificConfig(sb *strings.Builder, resourceName string, schema *SchemaInfo) {
	switch resourceName {
	case "http_loadbalancer":
		sb.WriteString("\n  # HTTP Load Balancer specific configuration\n")
		sb.WriteString("  domains = [\"app.example.com\"]\n\n")
		sb.WriteString("  # Advertise on public internet\n")
		sb.WriteString("  advertise_on_internet {\n")
		sb.WriteString("    default_vip {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Default origin server\n")
		sb.WriteString("  default_route_pools {\n")
		sb.WriteString("    pool {\n")
		sb.WriteString("      name      = \"example-origin-pool\"\n")
		sb.WriteString("      namespace = \"system\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("    weight   = 1\n")
		sb.WriteString("    priority = 1\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Enable HTTP to HTTPS redirect\n")
		sb.WriteString("  http {\n")
		sb.WriteString("    dns_volterra_managed = true\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Disable rate limiting by default\n")
		sb.WriteString("  disable_rate_limit {}\n\n")
		sb.WriteString("  # No WAF by default\n")
		sb.WriteString("  disable_waf {}\n")

	case "origin_pool":
		sb.WriteString("\n  # Origin Pool specific configuration\n")
		sb.WriteString("  origin_servers {\n")
		sb.WriteString("    public_ip {\n")
		sb.WriteString("      ip = \"203.0.113.10\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  port = 443\n\n")
		sb.WriteString("  # Use TLS to origin\n")
		sb.WriteString("  use_tls {\n")
		sb.WriteString("    use_host_header_as_sni {}\n")
		sb.WriteString("    tls_config {\n")
		sb.WriteString("      default_security {}\n")
		sb.WriteString("    }\n")
		sb.WriteString("    skip_server_verification {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Endpoint selection\n")
		sb.WriteString("  endpoint_selection     = \"LOCAL_PREFERRED\"\n")
		sb.WriteString("  loadbalancer_algorithm = \"LB_OVERRIDE\"\n")

	case "healthcheck":
		sb.WriteString("\n  # Health Check specific configuration\n")
		sb.WriteString("  http_health_check {\n")
		sb.WriteString("    use_origin_server_name {}\n")
		sb.WriteString("    path                   = \"/health\"\n")
		sb.WriteString("    use_http2              = false\n")
		sb.WriteString("    expected_status_codes  = [\"200\"]\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  healthy_threshold   = 3\n")
		sb.WriteString("  unhealthy_threshold = 3\n")
		sb.WriteString("  interval            = 15\n")
		sb.WriteString("  timeout             = 5\n")

	case "app_firewall":
		sb.WriteString("\n  # Web Application Firewall configuration\n")
		sb.WriteString("  # Block malicious requests\n")
		sb.WriteString("  blocking {}\n\n")
		sb.WriteString("  # Use default detection settings\n")
		sb.WriteString("  use_default_blocking_page {}\n\n")
		sb.WriteString("  # Default bot defense configuration\n")
		sb.WriteString("  default_bot_setting {}\n\n")
		sb.WriteString("  # Default detection settings\n")
		sb.WriteString("  default_detection_settings {}\n\n")
		sb.WriteString("  # Allow all response codes\n")
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
		sb.WriteString("    namespace = \"system\"\n")
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
		sb.WriteString("    namespace = \"system\"\n")
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
		sb.WriteString("    namespace = \"system\"\n")
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
		sb.WriteString("      namespace = \"system\"\n")
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
		sb.WriteString("    namespace = \"system\"\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Rule-based load balancing\n")
		sb.WriteString("  rule_list {\n")
		sb.WriteString("    rules {\n")
		sb.WriteString("      geo_location_set {\n")
		sb.WriteString("        name      = \"us-geo\"\n")
		sb.WriteString("        namespace = \"system\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("      pool {\n")
		sb.WriteString("        name      = \"us-pool\"\n")
		sb.WriteString("        namespace = \"system\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("      score = 100\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")

	case "alert_policy":
		sb.WriteString("\n  # Alert Policy configuration\n")
		sb.WriteString("  # Alert receivers\n")
		sb.WriteString("  receivers {\n")
		sb.WriteString("    name      = \"slack-receiver\"\n")
		sb.WriteString("    namespace = \"system\"\n")
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
		sb.WriteString("      namespace = \"system\"\n")
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
		sb.WriteString("          namespace = \"system\"\n")
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
		sb.WriteString("        namespace = \"system\"\n")
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
		sb.WriteString("          namespace = \"system\"\n")
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
		sb.WriteString("    namespace = \"system\"\n")
		sb.WriteString("  }\n")

	case "voltstack_site":
		sb.WriteString("\n  # Voltstack Site configuration\n")
		sb.WriteString("  # Kubernetes configuration\n")
		sb.WriteString("  k8s_cluster {\n")
		sb.WriteString("    name      = \"example-k8s-cluster\"\n")
		sb.WriteString("    namespace = \"system\"\n")
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
		sb.WriteString("      namespace = \"system\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  cluster_wide_app_list {\n")
		sb.WriteString("    cluster_wide_apps {\n")
		sb.WriteString("      name      = \"nginx-ingress\"\n")
		sb.WriteString("      namespace = \"system\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  local_access_config {\n")
		sb.WriteString("    local_domain = \"cluster.local\"\n")
		sb.WriteString("    default_port {}\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  global_access_enable {}\n")

	case "virtual_k8s":
		sb.WriteString("\n  # Virtual Kubernetes configuration\n")
		sb.WriteString("  # Virtual site selection\n")
		sb.WriteString("  vsite_refs {\n")
		sb.WriteString("    name      = \"example-virtual-site\"\n")
		sb.WriteString("    namespace = \"system\"\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  # Disable cluster global access\n")
		sb.WriteString("  disabled {}\n")

	case "virtual_network":
		sb.WriteString("\n  # Virtual Network configuration\n")
		sb.WriteString("  site_local_network {}\n\n")
		sb.WriteString("  # DHCP range for the network\n")
		sb.WriteString("  srv6_network {\n")
		sb.WriteString("    enterprise_network {\n")
		sb.WriteString("      srv6_network_ns_params {\n")
		sb.WriteString("        namespace = \"system\"\n")
		sb.WriteString("      }\n")
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")

	case "fleet":
		sb.WriteString("\n  # Fleet configuration\n")
		sb.WriteString("  fleet_label = \"env=production\"\n\n")
		sb.WriteString("  # Network connectors\n")
		sb.WriteString("  inside_virtual_network {\n")
		sb.WriteString("    name      = \"inside-network\"\n")
		sb.WriteString("    namespace = \"system\"\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  outside_virtual_network {\n")
		sb.WriteString("    name      = \"outside-network\"\n")
		sb.WriteString("    namespace = \"system\"\n")
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
		sb.WriteString("      namespace = \"system\"\n")
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
		sb.WriteString("      namespace = \"system\"\n")
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
		sb.WriteString("    namespace = \"system\"\n")
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
		sb.WriteString("    namespace = \"system\"\n")
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
		sb.WriteString("    namespace = \"system\"\n")
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
		"http_loadbalancer":         true,
		"tcp_loadbalancer":          true,
		"origin_pool":               true,
		"aws_vpc_site":              true,
		"azure_vnet_site":           true,
		"gcp_vpc_site":              true,
		"securemesh_site_v2":        true,
		"securemesh_site":           true,
		"voltstack_site":            true,
		"service_policy":            true,
		"enhanced_firewall_policy":  true,
		"forward_proxy_policy":      true,
		"network_policy":            true,
		"cdn_loadbalancer":          true,
		"k8s_cluster":               true,
		"virtual_k8s":               true,
		"fleet":                     true,
		"workload":                  true,
		"aws_tgw_site":              true,
	}
	return complexResources[resourceName]
}

func generateAdvancedExample(resourceName string, schema *SchemaInfo) string {
	var sb strings.Builder

	humanName := toHumanName(resourceName)
	sb.WriteString(fmt.Sprintf("\n# Advanced %s with additional configuration\n", humanName))
	sb.WriteString(fmt.Sprintf("resource \"f5xc_%s\" \"advanced\" {\n", resourceName))
	sb.WriteString(fmt.Sprintf("  name      = \"advanced-%s\"\n", strings.ReplaceAll(resourceName, "_", "-")))
	sb.WriteString("  namespace = \"system\"\n\n")

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
		sb.WriteString("          namespace = \"system\"\n")
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
		sb.WriteString("      namespace = \"system\"\n")
		sb.WriteString("    }\n")
		sb.WriteString("    weight   = 80\n")
		sb.WriteString("    priority = 1\n")
		sb.WriteString("  }\n\n")
		sb.WriteString("  default_route_pools {\n")
		sb.WriteString("    pool {\n")
		sb.WriteString("      name      = \"secondary-pool\"\n")
		sb.WriteString("      namespace = \"system\"\n")
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
		sb.WriteString("    namespace = \"shared\"\n")
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
		sb.WriteString("          namespace = \"system\"\n")
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
		sb.WriteString("    namespace = \"system\"\n")
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

func toHumanName(resourceName string) string {
	words := strings.Split(resourceName, "_")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, " ")
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
