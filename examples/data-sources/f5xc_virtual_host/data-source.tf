# Virtual Host Data Source Example
# Retrieves information about an existing Virtual Host

# Look up an existing Virtual Host by name
data "f5xc_virtual_host" "example" {
  name      = "example-virtual-host"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "virtual_host_id" {
#   value = data.f5xc_virtual_host.example.id
# }
