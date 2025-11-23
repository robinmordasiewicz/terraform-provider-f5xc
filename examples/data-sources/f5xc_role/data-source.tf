# Role Data Source Example
# Retrieves information about an existing Role

# Look up an existing Role by name
data "f5xc_role" "example" {
  name      = "example-role"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "role_id" {
#   value = data.f5xc_role.example.id
# }
