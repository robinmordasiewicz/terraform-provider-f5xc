# Subnet Data Source Example
# Retrieves information about an existing Subnet

# Look up an existing Subnet by name
data "f5xc_subnet" "example" {
  name      = "example-subnet"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "subnet_id" {
#   value = data.f5xc_subnet.example.id
# }
