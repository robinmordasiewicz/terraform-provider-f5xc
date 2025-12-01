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

func TestAccCertificateResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-cert")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_certificate.test"

	// Generate test certificates dynamically for CI/CD compatibility
	certs := acctest.MustGenerateTestCertificates()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("f5xc_certificate"),
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateConfig_basic(nsName, rName, certs),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "private_key"},
				ImportStateIdFunc:       testAccCertificateImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccCertificateImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		namespace := rs.Primary.Attributes["namespace"]
		name := rs.Primary.Attributes["name"]
		return fmt.Sprintf("%s/%s", namespace, name), nil
	}
}

func testAccCertificateConfig_basic(nsName, name string, certs *acctest.TestCertificates) string {
	// Use dynamically generated certificates for CI/CD compatibility
	// F5 XC API requires certificate_url in "string:///BASE64_ENCODED_CERT" format
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

resource "f5xc_certificate" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  certificate_url = "string:///%[3]s"

  private_key {
    clear_secret_info {
      url = "string:///%[4]s"
    }
  }

  disable_ocsp_stapling {}
}
`, nsName, name, certs.ServerCertBase64, certs.ServerKeyBase64))
}
