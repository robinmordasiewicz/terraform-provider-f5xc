# Cluster Data Source Example
# Retrieves information about an existing Cluster

# Look up an existing Cluster by name
data "f5xc_cluster" "example" {
  name      = "example-cluster"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "cluster_id" {
#   value = data.f5xc_cluster.example.id
# }
