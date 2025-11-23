# Cminstance Data Source Example
# Retrieves information about an existing Cminstance

# Look up an existing Cminstance by name
data "f5xc_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "cminstance_id" {
#   value = data.f5xc_cminstance.example.id
# }
