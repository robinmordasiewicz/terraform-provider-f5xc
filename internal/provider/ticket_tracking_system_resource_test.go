// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

// =============================================================================
// TICKET TRACKING SYSTEM RESOURCE ACCEPTANCE TESTS
//
// These tests follow HashiCorp's acceptance testing best practices:
// https://developer.hashicorp.com/terraform/plugin/testing/testing-patterns
//
// Test categories implemented:
// 1. Basic Lifecycle Test - Create, Read, Import with API verification
// 2. All Attributes Test - Test all optional attributes
// 3. JIRA Configuration Test - Test jira_config block with adhoc_rest_api
//
// Note: ticket_tracking_system requires a custom namespace and uses time_sleep
// to avoid API race conditions.
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_API_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccTicketTrackingSystemResource -timeout 30m
// =============================================================================

// -----------------------------------------------------------------------------
// Test 1: Basic Lifecycle Test
// Verifies: Create, Read, Import operations with custom namespace
// Pattern: Basic lifecycle test with namespace creation and time_sleep
// -----------------------------------------------------------------------------

func TestAccTicketTrackingSystemResource_basic(t *testing.T) {
	t.Skip("Skipping: requires external ticketing system - ticket_tracking_system resources require pre-configured external ticketing systems (JIRA, ServiceNow, etc.) with valid credentials")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_ticket_tracking_system.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-tts")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create ticket_tracking_system with minimal configuration
			{
				Config: testAccTicketTrackingSystemResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify resource exists in Terraform state
					acctest.CheckResourceExists(resourceName),
					// Verify state attributes
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Step 2: Import state verification
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       testAccTicketTrackingSystemResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"timeouts"},
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 2: All Attributes Test
// Verifies: All optional attributes (labels, annotations, description, disable)
// Pattern: Comprehensive attribute coverage with custom namespace
// -----------------------------------------------------------------------------

func TestAccTicketTrackingSystemResource_allAttributes(t *testing.T) {
	t.Skip("Skipping: requires external ticketing system - ticket_tracking_system resources require pre-configured external ticketing systems (JIRA, ServiceNow, etc.) with valid credentials")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_ticket_tracking_system.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-tts")
	description := "Comprehensive acceptance test ticket tracking system"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccTicketTrackingSystemResourceConfig_allAttributes(nsName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					// Verify all attributes in Terraform state
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-acceptance-test"),
					resource.TestCheckResourceAttr(resourceName, "annotations.purpose", "acceptance-testing"),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "ci-cd"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Import verification for all attributes
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       testAccTicketTrackingSystemResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"timeouts", "disable"},
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 3: JIRA Configuration Test
// Verifies: jira_config block with adhoc_rest_api nested block
// Pattern: Complex nested block configuration
// Note: This test uses placeholder values for JIRA configuration
// -----------------------------------------------------------------------------

func TestAccTicketTrackingSystemResource_jiraConfig(t *testing.T) {
	t.Skip("Skipping: requires external ticketing system - ticket_tracking_system resources require pre-configured external ticketing systems (JIRA, ServiceNow, etc.) with valid credentials")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_ticket_tracking_system.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-tts")
	accountEmail := "test@example.com"
	orgDomain := "example.atlassian.net"
	apiToken := "test-api-token-placeholder"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccTicketTrackingSystemResourceConfig_jiraConfig(nsName, name, accountEmail, orgDomain, apiToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					// Verify jira_config.adhoc_rest_api attributes
					resource.TestCheckResourceAttr(resourceName, "jira_config.adhoc_rest_api.account_email", accountEmail),
					resource.TestCheckResourceAttr(resourceName, "jira_config.adhoc_rest_api.organization_domain", orgDomain),
					// Note: api_token is sensitive and returned as empty string by API
					resource.TestCheckResourceAttrSet(resourceName, "jira_config.adhoc_rest_api.api_token"),
				),
			},
			// Import verification - api_token will not match due to sensitive nature
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccTicketTrackingSystemResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{
					"timeouts",
					"jira_config.adhoc_rest_api.api_token", // API returns empty string for sensitive fields
				},
			},
		},
	})
}

// =============================================================================
// Helper Functions
// =============================================================================

// testAccTicketTrackingSystemResourceImportStateIdFunc returns the import ID in the format namespace/name
func testAccTicketTrackingSystemResourceImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes["namespace"], rs.Primary.Attributes["name"]), nil
	}
}

// =============================================================================
// Test Configuration Functions
// =============================================================================

func testAccTicketTrackingSystemResourceConfig_basic(nsName, name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_ticket_tracking_system" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name
}
`, nsName, name))
}

func testAccTicketTrackingSystemResourceConfig_allAttributes(nsName, name, description string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_ticket_tracking_system" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q
  disable     = false

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  annotations = {
    purpose = "acceptance-testing"
    owner   = "ci-cd"
  }
}
`, nsName, name, description))
}

func testAccTicketTrackingSystemResourceConfig_jiraConfig(nsName, name, accountEmail, orgDomain, apiToken string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_ticket_tracking_system" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  jira_config {
    adhoc_rest_api {
      account_email       = %[3]q
      organization_domain = %[4]q
      api_token           = %[5]q
    }
  }
}
`, nsName, name, accountEmail, orgDomain, apiToken))
}
