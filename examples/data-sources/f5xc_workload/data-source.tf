# Workload Data Source Example
# Retrieves information about an existing Workload

# Look up an existing Workload by name
data "f5xc_workload" "example" {
  name      = "example-workload"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "workload_id" {
#   value = data.f5xc_workload.example.id
# }
