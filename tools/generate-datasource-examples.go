//go:build ignore
// +build ignore

// This tool generates example Terraform configurations for all F5 XC data sources.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// uppercaseAcronyms defines acronyms that should be fully uppercase
var uppercaseAcronyms = map[string]bool{
	"http": true, "https": true, "dns": true, "tcp": true, "udp": true,
	"tls": true, "ssl": true, "api": true, "url": true, "uri": true,
	"ip": true, "bgp": true, "jwt": true, "acl": true, "waf": true,
	"cdn": true, "aws": true, "gcp": true, "vpc": true, "tgw": true,
	"vnet": true, "ce": true, "re": true, "lb": true, "vip": true,
	"sni": true, "cors": true, "xss": true, "csrf": true, "oidc": true,
	"saml": true, "ssh": true, "nfs": true, "ntp": true, "pem": true,
	"rsa": true, "ecdsa": true, "id": true, "apm": true, "irule": true,
}

// mixedCaseAcronyms defines acronyms with specific mixed case
var mixedCaseAcronyms = map[string]string{
	"mtls": "mTLS", "oauth": "OAuth", "graphql": "GraphQL",
	"websocket": "WebSocket", "javascript": "JavaScript", "typescript": "TypeScript",
	"github": "GitHub", "gitlab": "GitLab", "devops": "DevOps",
	"fastcgi": "FastCGI", "modsecurity": "ModSecurity", "hashicorp": "HashiCorp",
	"bigip": "BigIP",
}

func main() {
	providerDir := "internal/provider"
	examplesDir := "examples/data-sources"

	files, err := filepath.Glob(filepath.Join(providerDir, "*_data_source.go"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding data source files: %v\n", err)
		os.Exit(1)
	}

	for _, file := range files {
		baseName := filepath.Base(file)
		dataSourceName := strings.TrimSuffix(baseName, "_data_source.go")

		exampleDir := filepath.Join(examplesDir, "f5xc_"+dataSourceName)
		if err := os.MkdirAll(exampleDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory %s: %v\n", exampleDir, err)
			continue
		}

		exampleFile := filepath.Join(exampleDir, "data-source.tf")
		example := generateDataSourceExample(dataSourceName)
		if err := os.WriteFile(exampleFile, []byte(example), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", exampleFile, err)
			continue
		}

		fmt.Printf("Generated: %s\n", exampleFile)
	}
}

func generateDataSourceExample(dataSourceName string) string {
	var sb strings.Builder

	humanName := toHumanName(dataSourceName)
	resourceRef := strings.ReplaceAll(dataSourceName, "_", "-")

	sb.WriteString(fmt.Sprintf("# %s Data Source Example\n", humanName))
	sb.WriteString(fmt.Sprintf("# Retrieves information about an existing %s\n\n", humanName))

	// Basic lookup example
	sb.WriteString(fmt.Sprintf("# Look up an existing %s by name\n", humanName))
	sb.WriteString(fmt.Sprintf("data \"f5xc_%s\" \"example\" {\n", dataSourceName))
	sb.WriteString(fmt.Sprintf("  name      = \"example-%s\"\n", resourceRef))
	sb.WriteString("  namespace = \"system\"\n")
	sb.WriteString("}\n\n")

	// Usage example
	sb.WriteString(fmt.Sprintf("# Example: Use the data source in another resource\n"))
	sb.WriteString(fmt.Sprintf("# output \"%s_id\" {\n", dataSourceName))
	sb.WriteString(fmt.Sprintf("#   value = data.f5xc_%s.example.id\n", dataSourceName))
	sb.WriteString("# }\n")

	// Add resource-specific examples
	addDataSourceSpecificExample(&sb, dataSourceName)

	return sb.String()
}

func addDataSourceSpecificExample(sb *strings.Builder, dataSourceName string) {
	switch dataSourceName {
	case "http_loadbalancer":
		sb.WriteString("\n# Example: Reference in another load balancer configuration\n")
		sb.WriteString("# resource \"f5xc_service_policy\" \"example\" {\n")
		sb.WriteString("#   name      = \"policy-for-lb\"\n")
		sb.WriteString("#   namespace = \"system\"\n")
		sb.WriteString("#\n")
		sb.WriteString("#   # Use the load balancer's domains\n")
		sb.WriteString("#   # domain = data.f5xc_http_loadbalancer.example.domains[0]\n")
		sb.WriteString("# }\n")

	case "origin_pool":
		sb.WriteString("\n# Example: Use origin pool data in HTTP load balancer\n")
		sb.WriteString("# resource \"f5xc_http_loadbalancer\" \"example\" {\n")
		sb.WriteString("#   name      = \"example-lb\"\n")
		sb.WriteString("#   namespace = \"system\"\n")
		sb.WriteString("#   domains   = [\"app.example.com\"]\n")
		sb.WriteString("#\n")
		sb.WriteString("#   default_route_pools {\n")
		sb.WriteString("#     pool {\n")
		sb.WriteString("#       name      = data.f5xc_origin_pool.example.name\n")
		sb.WriteString("#       namespace = data.f5xc_origin_pool.example.namespace\n")
		sb.WriteString("#     }\n")
		sb.WriteString("#   }\n")
		sb.WriteString("# }\n")

	case "namespace":
		sb.WriteString("\n# Example: Create resources in a namespace discovered via data source\n")
		sb.WriteString("# resource \"f5xc_origin_pool\" \"example\" {\n")
		sb.WriteString("#   name      = \"example-pool\"\n")
		sb.WriteString("#   namespace = data.f5xc_namespace.example.name\n")
		sb.WriteString("#   # ... other configuration\n")
		sb.WriteString("# }\n")

	case "healthcheck":
		sb.WriteString("\n# Example: Reference healthcheck in origin pool\n")
		sb.WriteString("# resource \"f5xc_origin_pool\" \"example\" {\n")
		sb.WriteString("#   name      = \"example-pool\"\n")
		sb.WriteString("#   namespace = \"system\"\n")
		sb.WriteString("#\n")
		sb.WriteString("#   healthcheck {\n")
		sb.WriteString("#     name      = data.f5xc_healthcheck.example.name\n")
		sb.WriteString("#     namespace = data.f5xc_healthcheck.example.namespace\n")
		sb.WriteString("#   }\n")
		sb.WriteString("# }\n")

	case "app_firewall":
		sb.WriteString("\n# Example: Reference WAF in HTTP load balancer\n")
		sb.WriteString("# resource \"f5xc_http_loadbalancer\" \"example\" {\n")
		sb.WriteString("#   name      = \"protected-lb\"\n")
		sb.WriteString("#   namespace = \"system\"\n")
		sb.WriteString("#\n")
		sb.WriteString("#   app_firewall {\n")
		sb.WriteString("#     name      = data.f5xc_app_firewall.example.name\n")
		sb.WriteString("#     namespace = data.f5xc_app_firewall.example.namespace\n")
		sb.WriteString("#   }\n")
		sb.WriteString("# }\n")

	case "certificate":
		sb.WriteString("\n# Example: Reference certificate in HTTPS configuration\n")
		sb.WriteString("# resource \"f5xc_http_loadbalancer\" \"example\" {\n")
		sb.WriteString("#   name      = \"https-lb\"\n")
		sb.WriteString("#   namespace = \"system\"\n")
		sb.WriteString("#\n")
		sb.WriteString("#   https {\n")
		sb.WriteString("#     tls_cert_params {\n")
		sb.WriteString("#       certificates {\n")
		sb.WriteString("#         name      = data.f5xc_certificate.example.name\n")
		sb.WriteString("#         namespace = data.f5xc_certificate.example.namespace\n")
		sb.WriteString("#       }\n")
		sb.WriteString("#     }\n")
		sb.WriteString("#   }\n")
		sb.WriteString("# }\n")

	case "aws_vpc_site", "azure_vnet_site", "gcp_vpc_site":
		sb.WriteString("\n# Example: Reference cloud site for advertising load balancer\n")
		sb.WriteString(fmt.Sprintf("# resource \"f5xc_http_loadbalancer\" \"example\" {\n"))
		sb.WriteString("#   name      = \"site-advertised-lb\"\n")
		sb.WriteString("#   namespace = \"system\"\n")
		sb.WriteString("#\n")
		sb.WriteString("#   advertise_custom {\n")
		sb.WriteString("#     advertise_where {\n")
		sb.WriteString("#       site {\n")
		sb.WriteString("#         site {\n")
		sb.WriteString(fmt.Sprintf("#           name      = data.f5xc_%s.example.name\n", dataSourceName))
		sb.WriteString(fmt.Sprintf("#           namespace = data.f5xc_%s.example.namespace\n", dataSourceName))
		sb.WriteString("#         }\n")
		sb.WriteString("#       }\n")
		sb.WriteString("#     }\n")
		sb.WriteString("#   }\n")
		sb.WriteString("# }\n")

	case "virtual_site":
		sb.WriteString("\n# Example: Reference virtual site for site selection\n")
		sb.WriteString("# resource \"f5xc_http_loadbalancer\" \"example\" {\n")
		sb.WriteString("#   name      = \"vs-advertised-lb\"\n")
		sb.WriteString("#   namespace = \"system\"\n")
		sb.WriteString("#\n")
		sb.WriteString("#   advertise_custom {\n")
		sb.WriteString("#     advertise_where {\n")
		sb.WriteString("#       virtual_site {\n")
		sb.WriteString("#         virtual_site {\n")
		sb.WriteString("#           name      = data.f5xc_virtual_site.example.name\n")
		sb.WriteString("#           namespace = data.f5xc_virtual_site.example.namespace\n")
		sb.WriteString("#         }\n")
		sb.WriteString("#       }\n")
		sb.WriteString("#     }\n")
		sb.WriteString("#   }\n")
		sb.WriteString("# }\n")

	case "service_policy":
		sb.WriteString("\n# Example: Reference service policy in HTTP load balancer\n")
		sb.WriteString("# resource \"f5xc_http_loadbalancer\" \"example\" {\n")
		sb.WriteString("#   name      = \"policy-protected-lb\"\n")
		sb.WriteString("#   namespace = \"system\"\n")
		sb.WriteString("#\n")
		sb.WriteString("#   active_service_policies {\n")
		sb.WriteString("#     policies {\n")
		sb.WriteString("#       name      = data.f5xc_service_policy.example.name\n")
		sb.WriteString("#       namespace = data.f5xc_service_policy.example.namespace\n")
		sb.WriteString("#     }\n")
		sb.WriteString("#   }\n")
		sb.WriteString("# }\n")

	case "cloud_credentials":
		sb.WriteString("\n# Example: Reference cloud credentials in site configuration\n")
		sb.WriteString("# resource \"f5xc_aws_vpc_site\" \"example\" {\n")
		sb.WriteString("#   name      = \"example-aws-site\"\n")
		sb.WriteString("#   namespace = \"system\"\n")
		sb.WriteString("#\n")
		sb.WriteString("#   aws_cred {\n")
		sb.WriteString("#     name      = data.f5xc_cloud_credentials.example.name\n")
		sb.WriteString("#     namespace = data.f5xc_cloud_credentials.example.namespace\n")
		sb.WriteString("#   }\n")
		sb.WriteString("# }\n")

	case "rate_limiter":
		sb.WriteString("\n# Example: Reference rate limiter in HTTP load balancer\n")
		sb.WriteString("# resource \"f5xc_http_loadbalancer\" \"example\" {\n")
		sb.WriteString("#   name      = \"rate-limited-lb\"\n")
		sb.WriteString("#   namespace = \"system\"\n")
		sb.WriteString("#\n")
		sb.WriteString("#   rate_limit {\n")
		sb.WriteString("#     rate_limiter {\n")
		sb.WriteString("#       name      = data.f5xc_rate_limiter.example.name\n")
		sb.WriteString("#       namespace = data.f5xc_rate_limiter.example.namespace\n")
		sb.WriteString("#     }\n")
		sb.WriteString("#   }\n")
		sb.WriteString("# }\n")

	case "dns_zone":
		sb.WriteString("\n# Example: Reference DNS zone in DNS load balancer\n")
		sb.WriteString("# resource \"f5xc_dns_load_balancer\" \"example\" {\n")
		sb.WriteString("#   name      = \"example-dns-lb\"\n")
		sb.WriteString("#   namespace = \"system\"\n")
		sb.WriteString("#\n")
		sb.WriteString("#   dns_zone {\n")
		sb.WriteString("#     name      = data.f5xc_dns_zone.example.name\n")
		sb.WriteString("#     namespace = data.f5xc_dns_zone.example.namespace\n")
		sb.WriteString("#   }\n")
		sb.WriteString("# }\n")

	case "log_receiver":
		sb.WriteString("\n# Example: Reference log receiver in site configuration\n")
		sb.WriteString("# resource \"f5xc_securemesh_site_v2\" \"example\" {\n")
		sb.WriteString("#   name      = \"example-site\"\n")
		sb.WriteString("#   namespace = \"system\"\n")
		sb.WriteString("#\n")
		sb.WriteString("#   log_receiver {\n")
		sb.WriteString("#     name      = data.f5xc_log_receiver.example.name\n")
		sb.WriteString("#     namespace = data.f5xc_log_receiver.example.namespace\n")
		sb.WriteString("#   }\n")
		sb.WriteString("# }\n")

	case "alert_receiver":
		sb.WriteString("\n# Example: Reference alert receiver in alert policy\n")
		sb.WriteString("# resource \"f5xc_alert_policy\" \"example\" {\n")
		sb.WriteString("#   name      = \"example-policy\"\n")
		sb.WriteString("#   namespace = \"system\"\n")
		sb.WriteString("#\n")
		sb.WriteString("#   receivers {\n")
		sb.WriteString("#     name      = data.f5xc_alert_receiver.example.name\n")
		sb.WriteString("#     namespace = data.f5xc_alert_receiver.example.namespace\n")
		sb.WriteString("#   }\n")
		sb.WriteString("# }\n")
	}
}

func toHumanName(name string) string {
	words := strings.Split(name, "_")
	for i, word := range words {
		lower := strings.ToLower(word)
		if uppercaseAcronyms[lower] {
			words[i] = strings.ToUpper(word)
		} else if replacement, ok := mixedCaseAcronyms[lower]; ok {
			words[i] = replacement
		} else {
			words[i] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}
