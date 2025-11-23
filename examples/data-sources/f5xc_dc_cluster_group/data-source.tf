# Dc Cluster Group Data Source Example
# Retrieves information about an existing Dc Cluster Group

# Look up an existing Dc Cluster Group by name
data "f5xc_dc_cluster_group" "example" {
  name      = "example-dc-cluster-group"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "dc_cluster_group_id" {
#   value = data.f5xc_dc_cluster_group.example.id
# }
