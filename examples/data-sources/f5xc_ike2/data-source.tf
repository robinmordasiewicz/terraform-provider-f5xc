# Ike2 Data Source Example
# Retrieves information about an existing Ike2

# Look up an existing Ike2 by name
data "f5xc_ike2" "example" {
  name      = "example-ike2"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "ike2_id" {
#   value = data.f5xc_ike2.example.id
# }
