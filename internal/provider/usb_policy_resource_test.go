// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

// =============================================================================
// USB_POLICY RESOURCE ACCEPTANCE TESTS
//
// These tests verify the f5xc_usb_policy resource implementation following
// HashiCorp's acceptance testing best practices.
//
// Test categories implemented:
// 1. Basic Lifecycle Test - Create, Read, Update, Delete with API verification
// 2. Import Test - Verify import produces identical state
// 3. Update Tests - Test all mutable attributes
// 4. Custom Namespace Pattern - Create namespace with time_sleep delay
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccUsbPolicyResource -timeout 30m
// =============================================================================

// -----------------------------------------------------------------------------
// Test 1: Basic Lifecycle Test
// Verifies: Create, Read, Update, Delete operations with API verification
// Pattern: Basic lifecycle test with custom namespace and time_sleep
// -----------------------------------------------------------------------------

func TestAccUsbPolicyResource_basic(t *testing.T) {
	t.Skip("Skipping: requires USB policy infrastructure - usb_policy resources require F5 XC Edge sites with USB device management capabilities that are not available in standard test environments")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_usb_policy.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-usb-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		Steps: []resource.TestStep{
			// Step 1: Create usb_policy with minimal configuration
			{
				Config: testAccUsbPolicyResourceConfig_basic(nsName, name),
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
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccUsbPolicyResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 2: All Attributes Test
// Verifies: All optional attributes including allowed_devices block
// Pattern: Comprehensive attribute coverage
// -----------------------------------------------------------------------------

func TestAccUsbPolicyResource_allAttributes(t *testing.T) {
	t.Skip("Skipping: requires USB policy infrastructure - usb_policy resources require F5 XC Edge sites with USB device management capabilities that are not available in standard test environments")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_usb_policy.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-usb-policy")
	description := "Comprehensive acceptance test USB policy"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccUsbPolicyResourceConfig_allAttributes(nsName, name, description),
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
					// Verify allowed_devices block
					resource.TestCheckResourceAttr(resourceName, "allowed_devices.0.id_vendor", "0x046d"),
					resource.TestCheckResourceAttr(resourceName, "allowed_devices.0.id_product", "0xc52b"),
					resource.TestCheckResourceAttr(resourceName, "allowed_devices.0.b_device_class", "0x00"),
					resource.TestCheckResourceAttr(resourceName, "allowed_devices.1.id_vendor", "0x413c"),
					resource.TestCheckResourceAttr(resourceName, "allowed_devices.1.id_product", "0x2003"),
				),
			},
			// Import verification for all attributes
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "disable"},
				ImportStateIdFunc:       testAccUsbPolicyResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 3: Update Test - Labels
// Verifies: In-place update of labels attribute
// Pattern: Update test with multiple steps
// -----------------------------------------------------------------------------

func TestAccUsbPolicyResource_updateLabels(t *testing.T) {
	t.Skip("Skipping: requires USB policy infrastructure - usb_policy resources require F5 XC Edge sites with USB device management capabilities that are not available in standard test environments")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_usb_policy.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-usb-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with initial labels
			{
				Config: testAccUsbPolicyResourceConfig_withLabels(nsName, name, "test", "terraform"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform"),
				),
			},
			// Step 2: Update labels
			{
				Config: testAccUsbPolicyResourceConfig_withLabels(nsName, name, "staging", "terraform-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "staging"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-updated"),
				),
			},
			// Step 3: Remove all labels
			{
				Config: testAccUsbPolicyResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckNoResourceAttr(resourceName, "labels.environment"),
					resource.TestCheckNoResourceAttr(resourceName, "labels.managed_by"),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 4: Update Test - Description
// Verifies: Update and removal of description attribute
// Pattern: Update test for optional string attribute
// -----------------------------------------------------------------------------

func TestAccUsbPolicyResource_updateDescription(t *testing.T) {
	t.Skip("Skipping: requires USB policy infrastructure - usb_policy resources require F5 XC Edge sites with USB device management capabilities that are not available in standard test environments")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_usb_policy.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-usb-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		Steps: []resource.TestStep{
			// Step 1: Create without description
			{
				Config: testAccUsbPolicyResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckNoResourceAttr(resourceName, "description"),
				),
			},
			// Step 2: Add description
			{
				Config: testAccUsbPolicyResourceConfig_withDescription(nsName, name, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			// Step 3: Update description
			{
				Config: testAccUsbPolicyResourceConfig_withDescription(nsName, name, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
			// Step 4: Remove description
			{
				Config: testAccUsbPolicyResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 5: Update Test - Allowed Devices
// Verifies: Update and modification of allowed_devices block
// Pattern: Update test for list of nested blocks
// -----------------------------------------------------------------------------

func TestAccUsbPolicyResource_updateAllowedDevices(t *testing.T) {
	t.Skip("Skipping: requires USB policy infrastructure - usb_policy resources require F5 XC Edge sites with USB device management capabilities that are not available in standard test environments")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_usb_policy.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-usb-policy")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with single allowed device
			{
				Config: testAccUsbPolicyResourceConfig_withSingleDevice(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "allowed_devices.0.id_vendor", "0x046d"),
					resource.TestCheckResourceAttr(resourceName, "allowed_devices.0.id_product", "0xc52b"),
				),
			},
			// Step 2: Add another device
			{
				Config: testAccUsbPolicyResourceConfig_withMultipleDevices(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "allowed_devices.0.id_vendor", "0x046d"),
					resource.TestCheckResourceAttr(resourceName, "allowed_devices.0.id_product", "0xc52b"),
					resource.TestCheckResourceAttr(resourceName, "allowed_devices.1.id_vendor", "0x413c"),
					resource.TestCheckResourceAttr(resourceName, "allowed_devices.1.id_product", "0x2003"),
				),
			},
			// Step 3: Update device attributes
			{
				Config: testAccUsbPolicyResourceConfig_withDeviceAttributes(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "allowed_devices.0.id_vendor", "0x046d"),
					resource.TestCheckResourceAttr(resourceName, "allowed_devices.0.id_product", "0xc52b"),
					resource.TestCheckResourceAttr(resourceName, "allowed_devices.0.b_device_class", "0x00"),
					resource.TestCheckResourceAttr(resourceName, "allowed_devices.0.b_device_sub_class", "0x00"),
					resource.TestCheckResourceAttr(resourceName, "allowed_devices.0.b_device_protocol", "0x00"),
				),
			},
			// Step 4: Remove all allowed devices
			{
				Config: testAccUsbPolicyResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
		},
	})
}

// =============================================================================
// Test Configuration Functions
// =============================================================================

func testAccUsbPolicyResourceImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes["namespace"], rs.Primary.Attributes["name"]), nil
	}
}

func testAccUsbPolicyResourceConfig_basic(nsName, name string) string {
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

resource "f5xc_usb_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name
}
`, nsName, name))
}

func testAccUsbPolicyResourceConfig_allAttributes(nsName, name, description string) string {
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

resource "f5xc_usb_policy" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  annotations = {
    purpose = "acceptance-testing"
    owner   = "ci-cd"
  }

  allowed_devices {
    id_vendor       = "0x046d"
    id_product      = "0xc52b"
    b_device_class  = "0x00"
  }

  allowed_devices {
    id_vendor  = "0x413c"
    id_product = "0x2003"
  }
}
`, nsName, name, description))
}

func testAccUsbPolicyResourceConfig_withLabels(nsName, name, environment, managedBy string) string {
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

resource "f5xc_usb_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  labels = {
    environment = %[3]q
    managed_by  = %[4]q
  }
}
`, nsName, name, environment, managedBy))
}

func testAccUsbPolicyResourceConfig_withDescription(nsName, name, description string) string {
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

resource "f5xc_usb_policy" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q
}
`, nsName, name, description))
}

func testAccUsbPolicyResourceConfig_withSingleDevice(nsName, name string) string {
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

resource "f5xc_usb_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  allowed_devices {
    id_vendor  = "0x046d"
    id_product = "0xc52b"
  }
}
`, nsName, name))
}

func testAccUsbPolicyResourceConfig_withMultipleDevices(nsName, name string) string {
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

resource "f5xc_usb_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  allowed_devices {
    id_vendor  = "0x046d"
    id_product = "0xc52b"
  }

  allowed_devices {
    id_vendor  = "0x413c"
    id_product = "0x2003"
  }
}
`, nsName, name))
}

func testAccUsbPolicyResourceConfig_withDeviceAttributes(nsName, name string) string {
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

resource "f5xc_usb_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  allowed_devices {
    id_vendor          = "0x046d"
    id_product         = "0xc52b"
    b_device_class     = "0x00"
    b_device_sub_class = "0x00"
    b_device_protocol  = "0x00"
  }
}
`, nsName, name))
}
