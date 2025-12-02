// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package mocks

import (
	"fmt"
	"time"
)

// Fixture generators for common F5 XC resource types

// NamespaceResponse creates a mock namespace response
func NamespaceResponse(name string, labels, annotations map[string]string, description string) map[string]interface{} {
	return map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":        name,
			"namespace":   "system",
			"uid":         generateUID(),
			"labels":      labels,
			"annotations": annotations,
			"description": description,
		},
		"spec": map[string]interface{}{},
		"system_metadata": map[string]interface{}{
			"uid":                    generateUID(),
			"creation_timestamp":     time.Now().UTC().Format(time.RFC3339),
			"modification_timestamp": time.Now().UTC().Format(time.RFC3339),
			"creator_class":          "API",
			"creator_id":             "mock-server",
			"tenant":                 "mock-tenant",
		},
	}
}

// AWSVPCSiteResponse creates a mock AWS VPC site response
func AWSVPCSiteResponse(namespace, name string, opts ...AWSVPCSiteOption) map[string]interface{} {
	cfg := &awsVPCSiteConfig{
		awsRegion:     "us-east-1",
		vpcID:         "vpc-12345678",
		instanceType:  "t3.xlarge",
		sshKey:        "mock-ssh-key",
		diskSize:      80,
		nodesPerAZ:    1,
		siteType:      "ingress_egress_gw",
		description:   "Mock AWS VPC Site",
	}
	for _, opt := range opts {
		opt(cfg)
	}

	response := map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":        name,
			"namespace":   namespace,
			"uid":         generateUID(),
			"description": cfg.description,
		},
		"spec": map[string]interface{}{
			"aws_region":       cfg.awsRegion,
			"vpc_id":           cfg.vpcID,
			"instance_type":    cfg.instanceType,
			"ssh_key":          cfg.sshKey,
			"disk_size":        cfg.diskSize,
			"nodes_per_az":     cfg.nodesPerAZ,
		},
		"system_metadata": map[string]interface{}{
			"uid":                    generateUID(),
			"creation_timestamp":     time.Now().UTC().Format(time.RFC3339),
			"modification_timestamp": time.Now().UTC().Format(time.RFC3339),
			"creator_class":          "API",
			"creator_id":             "mock-server",
			"tenant":                 "mock-tenant",
		},
		"status": map[string]interface{}{
			"site_state": "ONLINE",
			"vpc_id":     cfg.vpcID,
		},
	}

	// Add site type specific configuration
	if cfg.siteType == "ingress_egress_gw" {
		response["spec"].(map[string]interface{})["ingress_egress_gw"] = map[string]interface{}{
			"az_nodes": []map[string]interface{}{
				{
					"aws_az_name": cfg.awsRegion + "a",
					"disk_size":   cfg.diskSize,
				},
			},
		}
	}

	return response
}

type awsVPCSiteConfig struct {
	awsRegion     string
	vpcID         string
	instanceType  string
	sshKey        string
	diskSize      int
	nodesPerAZ    int
	siteType      string
	description   string
}

// AWSVPCSiteOption configures AWS VPC site mock responses
type AWSVPCSiteOption func(*awsVPCSiteConfig)

// WithAWSRegion sets the AWS region
func WithAWSRegion(region string) AWSVPCSiteOption {
	return func(c *awsVPCSiteConfig) { c.awsRegion = region }
}

// WithVPCID sets the VPC ID
func WithVPCID(vpcID string) AWSVPCSiteOption {
	return func(c *awsVPCSiteConfig) { c.vpcID = vpcID }
}

// WithInstanceType sets the EC2 instance type
func WithInstanceType(instanceType string) AWSVPCSiteOption {
	return func(c *awsVPCSiteConfig) { c.instanceType = instanceType }
}

// WithSiteDescription sets the site description
func WithSiteDescription(desc string) AWSVPCSiteOption {
	return func(c *awsVPCSiteConfig) { c.description = desc }
}

// AzureVNETSiteResponse creates a mock Azure VNET site response
func AzureVNETSiteResponse(namespace, name string, opts ...AzureVNETSiteOption) map[string]interface{} {
	cfg := &azureVNETSiteConfig{
		azureRegion:   "eastus",
		vnetID:        "/subscriptions/mock-sub/resourceGroups/mock-rg/providers/Microsoft.Network/virtualNetworks/mock-vnet",
		machineType:   "Standard_D3_v2",
		diskSize:      80,
		siteType:      "ingress_egress_gw",
		description:   "Mock Azure VNET Site",
	}
	for _, opt := range opts {
		opt(cfg)
	}

	return map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":        name,
			"namespace":   namespace,
			"uid":         generateUID(),
			"description": cfg.description,
		},
		"spec": map[string]interface{}{
			"azure_region":  cfg.azureRegion,
			"vnet":          map[string]interface{}{"existing_vnet": map[string]interface{}{"resource_group": "mock-rg", "vnet_name": "mock-vnet"}},
			"machine_type":  cfg.machineType,
			"disk_size":     cfg.diskSize,
		},
		"system_metadata": map[string]interface{}{
			"uid":                    generateUID(),
			"creation_timestamp":     time.Now().UTC().Format(time.RFC3339),
			"modification_timestamp": time.Now().UTC().Format(time.RFC3339),
			"creator_class":          "API",
			"creator_id":             "mock-server",
			"tenant":                 "mock-tenant",
		},
		"status": map[string]interface{}{
			"site_state": "ONLINE",
		},
	}
}

type azureVNETSiteConfig struct {
	azureRegion   string
	vnetID        string
	machineType   string
	diskSize      int
	siteType      string
	description   string
}

// AzureVNETSiteOption configures Azure VNET site mock responses
type AzureVNETSiteOption func(*azureVNETSiteConfig)

// WithAzureRegion sets the Azure region
func WithAzureRegion(region string) AzureVNETSiteOption {
	return func(c *azureVNETSiteConfig) { c.azureRegion = region }
}

// GCPVPCSiteResponse creates a mock GCP VPC site response
func GCPVPCSiteResponse(namespace, name string, opts ...GCPVPCSiteOption) map[string]interface{} {
	cfg := &gcpVPCSiteConfig{
		gcpRegion:    "us-central1",
		projectID:    "mock-project",
		networkName:  "mock-network",
		machineType:  "n1-standard-4",
		diskSize:     80,
		siteType:     "ingress_egress_gw",
		description:  "Mock GCP VPC Site",
	}
	for _, opt := range opts {
		opt(cfg)
	}

	return map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":        name,
			"namespace":   namespace,
			"uid":         generateUID(),
			"description": cfg.description,
		},
		"spec": map[string]interface{}{
			"gcp_region":    cfg.gcpRegion,
			"gcp_project":   cfg.projectID,
			"network_name":  cfg.networkName,
			"machine_type":  cfg.machineType,
			"disk_size":     cfg.diskSize,
		},
		"system_metadata": map[string]interface{}{
			"uid":                    generateUID(),
			"creation_timestamp":     time.Now().UTC().Format(time.RFC3339),
			"modification_timestamp": time.Now().UTC().Format(time.RFC3339),
			"creator_class":          "API",
			"creator_id":             "mock-server",
			"tenant":                 "mock-tenant",
		},
		"status": map[string]interface{}{
			"site_state": "ONLINE",
		},
	}
}

type gcpVPCSiteConfig struct {
	gcpRegion    string
	projectID    string
	networkName  string
	machineType  string
	diskSize     int
	siteType     string
	description  string
}

// GCPVPCSiteOption configures GCP VPC site mock responses
type GCPVPCSiteOption func(*gcpVPCSiteConfig)

// WithGCPRegion sets the GCP region
func WithGCPRegion(region string) GCPVPCSiteOption {
	return func(c *gcpVPCSiteConfig) { c.gcpRegion = region }
}

// WithGCPProject sets the GCP project ID
func WithGCPProject(projectID string) GCPVPCSiteOption {
	return func(c *gcpVPCSiteConfig) { c.projectID = projectID }
}

// CloudCredentialsResponse creates a mock cloud credentials response
func CloudCredentialsResponse(namespace, name, cloudType string) map[string]interface{} {
	spec := map[string]interface{}{}

	switch cloudType {
	case "aws":
		spec["aws_secret_credentials"] = map[string]interface{}{
			"access_key_id":     "AKIAIOSFODNN7EXAMPLE",
			"secret_access_key": map[string]interface{}{"blindfold_secret_info": map[string]interface{}{"location": "string:///mock"}},
		}
	case "azure":
		spec["azure_client_secret"] = map[string]interface{}{
			"subscription_id": "mock-subscription-id",
			"tenant_id":       "mock-tenant-id",
			"client_id":       "mock-client-id",
			"client_secret":   map[string]interface{}{"blindfold_secret_info": map[string]interface{}{"location": "string:///mock"}},
		}
	case "gcp":
		spec["gcp_credentials"] = map[string]interface{}{
			"credential_file": map[string]interface{}{"blindfold_secret_info": map[string]interface{}{"location": "string:///mock"}},
		}
	}

	return map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":        name,
			"namespace":   namespace,
			"uid":         generateUID(),
			"description": fmt.Sprintf("Mock %s cloud credentials", cloudType),
		},
		"spec": spec,
		"system_metadata": map[string]interface{}{
			"uid":                    generateUID(),
			"creation_timestamp":     time.Now().UTC().Format(time.RFC3339),
			"modification_timestamp": time.Now().UTC().Format(time.RFC3339),
			"creator_class":          "API",
			"creator_id":             "mock-server",
			"tenant":                 "mock-tenant",
		},
	}
}

// OriginPoolResponse creates a mock origin pool response
func OriginPoolResponse(namespace, name string, opts ...OriginPoolOption) map[string]interface{} {
	cfg := &originPoolConfig{
		port:           443,
		healthcheckOn: "PORT",
		description:   "Mock origin pool",
	}
	for _, opt := range opts {
		opt(cfg)
	}

	return map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":        name,
			"namespace":   namespace,
			"uid":         generateUID(),
			"description": cfg.description,
		},
		"spec": map[string]interface{}{
			"origin_servers": []map[string]interface{}{
				{
					"public_ip": map[string]interface{}{"ip": "93.184.216.34"},
				},
			},
			"port":              cfg.port,
			"no_tls":            true,
			"endpoint_selection": "LOCAL_PREFERRED",
			"loadbalancer_algorithm": "LB_OVERRIDE",
		},
		"system_metadata": map[string]interface{}{
			"uid":                    generateUID(),
			"creation_timestamp":     time.Now().UTC().Format(time.RFC3339),
			"modification_timestamp": time.Now().UTC().Format(time.RFC3339),
			"creator_class":          "API",
			"creator_id":             "mock-server",
			"tenant":                 "mock-tenant",
		},
	}
}

type originPoolConfig struct {
	port           int
	healthcheckOn  string
	description    string
}

// OriginPoolOption configures origin pool mock responses
type OriginPoolOption func(*originPoolConfig)

// WithOriginPoolPort sets the port
func WithOriginPoolPort(port int) OriginPoolOption {
	return func(c *originPoolConfig) { c.port = port }
}

// HTTPLoadBalancerResponse creates a mock HTTP load balancer response
func HTTPLoadBalancerResponse(namespace, name string, domains []string) map[string]interface{} {
	return map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":        name,
			"namespace":   namespace,
			"uid":         generateUID(),
			"description": "Mock HTTP load balancer",
		},
		"spec": map[string]interface{}{
			"domains":     domains,
			"http": map[string]interface{}{
				"port": 80,
			},
			"advertise_on_public_default_vip": map[string]interface{}{},
			"default_route_pools": []map[string]interface{}{},
		},
		"system_metadata": map[string]interface{}{
			"uid":                    generateUID(),
			"creation_timestamp":     time.Now().UTC().Format(time.RFC3339),
			"modification_timestamp": time.Now().UTC().Format(time.RFC3339),
			"creator_class":          "API",
			"creator_id":             "mock-server",
			"tenant":                 "mock-tenant",
		},
	}
}

// AppFirewallResponse creates a mock app firewall response
func AppFirewallResponse(namespace, name string) map[string]interface{} {
	return map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":        name,
			"namespace":   namespace,
			"uid":         generateUID(),
			"description": "Mock app firewall",
		},
		"spec": map[string]interface{}{
			"blocking": map[string]interface{}{},
			"default_detection_settings": map[string]interface{}{},
			"default_bot_setting": map[string]interface{}{},
			"allow_all_response_codes": map[string]interface{}{},
		},
		"system_metadata": map[string]interface{}{
			"uid":                    generateUID(),
			"creation_timestamp":     time.Now().UTC().Format(time.RFC3339),
			"modification_timestamp": time.Now().UTC().Format(time.RFC3339),
			"creator_class":          "API",
			"creator_id":             "mock-server",
			"tenant":                 "mock-tenant",
		},
	}
}

// HealthcheckResponse creates a mock healthcheck response
func HealthcheckResponse(namespace, name string, healthcheckType string) map[string]interface{} {
	spec := map[string]interface{}{
		"timeout":  3,
		"interval": 15,
	}

	switch healthcheckType {
	case "http":
		spec["http_health_check"] = map[string]interface{}{
			"path": "/health",
		}
	case "tcp":
		spec["tcp_health_check"] = map[string]interface{}{}
	}

	return map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":        name,
			"namespace":   namespace,
			"uid":         generateUID(),
			"description": "Mock healthcheck",
		},
		"spec": spec,
		"system_metadata": map[string]interface{}{
			"uid":                    generateUID(),
			"creation_timestamp":     time.Now().UTC().Format(time.RFC3339),
			"modification_timestamp": time.Now().UTC().Format(time.RFC3339),
			"creator_class":          "API",
			"creator_id":             "mock-server",
			"tenant":                 "mock-tenant",
		},
	}
}

// TenantConfigurationResponse creates a mock tenant configuration response
func TenantConfigurationResponse(name string) map[string]interface{} {
	return map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":      name,
			"namespace": "system",
			"uid":       generateUID(),
		},
		"spec": map[string]interface{}{
			"company_name":     "Mock Company",
			"company_domain":   "mock.example.com",
			"contact_email":    "admin@mock.example.com",
		},
		"system_metadata": map[string]interface{}{
			"uid":                    generateUID(),
			"creation_timestamp":     time.Now().UTC().Format(time.RFC3339),
			"modification_timestamp": time.Now().UTC().Format(time.RFC3339),
			"creator_class":          "SYSTEM",
			"creator_id":             "system",
			"tenant":                 "mock-tenant",
		},
	}
}

// VK8sClusterResponse creates a mock virtual K8s cluster response
func VK8sClusterResponse(namespace, name string) map[string]interface{} {
	return map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":        name,
			"namespace":   namespace,
			"uid":         generateUID(),
			"description": "Mock vK8s cluster",
		},
		"spec": map[string]interface{}{
			"site_selector": map[string]interface{}{
				"expressions": []string{"site in (mock-site)"},
			},
		},
		"system_metadata": map[string]interface{}{
			"uid":                    generateUID(),
			"creation_timestamp":     time.Now().UTC().Format(time.RFC3339),
			"modification_timestamp": time.Now().UTC().Format(time.RFC3339),
			"creator_class":          "API",
			"creator_id":             "mock-server",
			"tenant":                 "mock-tenant",
		},
		"status": map[string]interface{}{
			"cluster_state": "ONLINE",
		},
	}
}

// DNSZoneResponse creates a mock DNS zone response
func DNSZoneResponse(namespace, name, domain string) map[string]interface{} {
	return map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":        name,
			"namespace":   namespace,
			"uid":         generateUID(),
			"description": "Mock DNS zone",
		},
		"spec": map[string]interface{}{
			"primary": map[string]interface{}{
				"default_rr_set_group": []interface{}{},
				"default_soa_parameters": map[string]interface{}{},
			},
		},
		"system_metadata": map[string]interface{}{
			"uid":                    generateUID(),
			"creation_timestamp":     time.Now().UTC().Format(time.RFC3339),
			"modification_timestamp": time.Now().UTC().Format(time.RFC3339),
			"creator_class":          "API",
			"creator_id":             "mock-server",
			"tenant":                 "mock-tenant",
		},
	}
}

// NFVServiceResponse creates a mock NFV service response
func NFVServiceResponse(namespace, name string) map[string]interface{} {
	return map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":        name,
			"namespace":   namespace,
			"uid":         generateUID(),
			"description": "Mock NFV service",
		},
		"spec": map[string]interface{}{
			"nfv_type": "BIG_IP",
		},
		"system_metadata": map[string]interface{}{
			"uid":                    generateUID(),
			"creation_timestamp":     time.Now().UTC().Format(time.RFC3339),
			"modification_timestamp": time.Now().UTC().Format(time.RFC3339),
			"creator_class":          "API",
			"creator_id":             "mock-server",
			"tenant":                 "mock-tenant",
		},
	}
}

// GenericResourceResponse creates a generic mock resource response
func GenericResourceResponse(namespace, name, resourceType string, spec map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":        name,
			"namespace":   namespace,
			"uid":         generateUID(),
			"description": fmt.Sprintf("Mock %s resource", resourceType),
		},
		"spec": spec,
		"system_metadata": map[string]interface{}{
			"uid":                    generateUID(),
			"creation_timestamp":     time.Now().UTC().Format(time.RFC3339),
			"modification_timestamp": time.Now().UTC().Format(time.RFC3339),
			"creator_class":          "API",
			"creator_id":             "mock-server",
			"tenant":                 "mock-tenant",
		},
	}
}
