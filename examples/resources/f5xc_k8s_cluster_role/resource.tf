# K8s Cluster Role Resource Example
# Manages k8s_cluster_role will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

# Basic K8s Cluster Role configuration
resource "f5xc_k8s_cluster_role" "example" {
  name      = "example-k8s-cluster-role"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: k8s_cluster_role_selector, policy_rule_list, yaml...
  k8s_cluster_role_selector {
    # Configure k8s_cluster_role_selector settings
  }
  # Policy Rule List. List of rules for role permissions
  policy_rule_list {
    # Configure policy_rule_list settings
  }
  # Policy Rules. List of rules for role permissions
  policy_rule {
    # Configure policy_rule settings
  }
}
