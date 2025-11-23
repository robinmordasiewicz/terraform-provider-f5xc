# Fast Acl Data Source Example
# Retrieves information about an existing Fast Acl

# Look up an existing Fast Acl by name
data "f5xc_fast_acl" "example" {
  name      = "example-fast-acl"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "fast_acl_id" {
#   value = data.f5xc_fast_acl.example.id
# }
