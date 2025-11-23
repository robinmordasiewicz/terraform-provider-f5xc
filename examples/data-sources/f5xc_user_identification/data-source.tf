# User Identification Data Source Example
# Retrieves information about an existing User Identification

# Look up an existing User Identification by name
data "f5xc_user_identification" "example" {
  name      = "example-user-identification"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "user_identification_id" {
#   value = data.f5xc_user_identification.example.id
# }
