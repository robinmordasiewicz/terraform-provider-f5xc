# K8s Cluster Role Binding Data Source Example
# Retrieves information about an existing K8s Cluster Role Binding

# Look up an existing K8s Cluster Role Binding by name
data "f5xc_k8s_cluster_role_binding" "example" {
  name      = "example-k8s-cluster-role-binding"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "k8s_cluster_role_binding_id" {
#   value = data.f5xc_k8s_cluster_role_binding.example.id
# }
