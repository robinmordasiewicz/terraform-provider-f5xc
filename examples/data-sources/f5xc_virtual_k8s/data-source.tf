# Virtual K8s Data Source Example
# Retrieves information about an existing Virtual K8s

# Look up an existing Virtual K8s by name
data "f5xc_virtual_k8s" "example" {
  name      = "example-virtual-k8s"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "virtual_k8s_id" {
#   value = data.f5xc_virtual_k8s.example.id
# }
