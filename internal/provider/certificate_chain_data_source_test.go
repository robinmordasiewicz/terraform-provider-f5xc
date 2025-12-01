// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccCertificateChainDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_certificate_chain.test"
	dataSourceName := "data.f5xc_certificate_chain.test"

	// Generate test certificates dynamically for CI/CD compatibility
	certs := acctest.MustGenerateTestCertificates()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateChainDataSourceConfig_basic(nsName, rName, certs),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccCertificateChainDataSourceConfig_basic(nsName, name string, certs *acctest.TestCertificates) string {
	// Use dynamically generated intermediate CA certificate for CI/CD compatibility
	// F5 XC API requires certificate_url in "string:///BASE64_ENCODED_CERT" format
	// Certificate chain requires a CA certificate (with CA:TRUE basic constraint)
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

resource "f5xc_certificate_chain" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  certificate_url = "string:///%[3]s"
}

data "f5xc_certificate_chain" "test" {
  depends_on = [f5xc_certificate_chain.test]
  name       = f5xc_certificate_chain.test.name
  namespace  = f5xc_certificate_chain.test.namespace
}
`, nsName, name, certs.IntermediateCABase64))
}
