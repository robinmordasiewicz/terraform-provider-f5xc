# Workload Flavor Data Source Example
# Retrieves information about an existing Workload Flavor

# Look up an existing Workload Flavor by name
data "f5xc_workload_flavor" "example" {
  name      = "example-workload-flavor"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "workload_flavor_id" {
#   value = data.f5xc_workload_flavor.example.id
# }
