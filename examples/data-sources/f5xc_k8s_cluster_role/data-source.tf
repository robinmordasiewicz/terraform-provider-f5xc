# K8s Cluster Role Data Source Example
# Retrieves information about an existing K8s Cluster Role

# Look up an existing K8s Cluster Role by name
data "f5xc_k8s_cluster_role" "example" {
  name      = "example-k8s-cluster-role"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "k8s_cluster_role_id" {
#   value = data.f5xc_k8s_cluster_role.example.id
# }
