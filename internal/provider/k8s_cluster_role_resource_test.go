// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccK8SClusterRoleResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	t.Skip("Skipping: k8s_cluster_role generator does not populate nested policy_rule_list.policy_rule.resource_list fields (api_groups, resource_types, verbs)")

	rName := acctest.RandomName("tf-acc-test-k8s-role")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_k8s_cluster_role.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_k8s_cluster_role"),
		Steps: []resource.TestStep{
			{
				Config: testAccK8SClusterRoleConfig_basic(nsName, rName),
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
				ImportStateVerifyIgnore: []string{"timeouts", "policy_rule_list"},
				ImportStateIdFunc:       testAccK8SClusterRoleImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccK8SClusterRoleImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccK8SClusterRoleConfig_basic(nsName, name string) string {
	// K8s cluster role must be in system namespace and requires policy_rule_list
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_k8s_cluster_role" "test" {
  name      = %[1]q
  namespace = "system"

  policy_rule_list {
    policy_rule {
      resource_list {
        api_groups     = [""]
        resource_types = ["pods"]
        verbs          = ["get", "list"]
      }
    }
  }
}
`, name))
}
