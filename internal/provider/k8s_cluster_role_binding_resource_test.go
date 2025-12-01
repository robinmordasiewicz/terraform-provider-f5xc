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

func TestAccK8SClusterRoleBindingResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-k8s-binding")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_k8s_cluster_role_binding.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_k8s_cluster_role_binding"),
		Steps: []resource.TestStep{
			{
				Config: testAccK8SClusterRoleBindingConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "k8s_cluster_role", "subjects"},
				ImportStateIdFunc:       testAccK8SClusterRoleBindingImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccK8SClusterRoleBindingImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccK8SClusterRoleBindingConfig_basic(nsName, name string) string {
	// K8s cluster role binding must be in system namespace
	// Requires k8s_cluster_role reference and subjects
	// Using shared ves-io-admin-cluster-role from ves-io tenant
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_k8s_cluster_role_binding" "test" {
  name      = %[1]q
  namespace = "system"

  k8s_cluster_role {
    name      = "ves-io-admin-cluster-role"
    namespace = "shared"
    tenant    = "ves-io"
  }

  subjects {
    user = "test-user"
  }
}
`, name))
}
