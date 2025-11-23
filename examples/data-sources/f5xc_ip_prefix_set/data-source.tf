# Ip Prefix Set Data Source Example
# Retrieves information about an existing Ip Prefix Set

# Look up an existing Ip Prefix Set by name
data "f5xc_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "ip_prefix_set_id" {
#   value = data.f5xc_ip_prefix_set.example.id
# }
