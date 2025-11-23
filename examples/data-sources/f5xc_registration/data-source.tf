# Registration Data Source Example
# Retrieves information about an existing Registration

# Look up an existing Registration by name
data "f5xc_registration" "example" {
  name      = "example-registration"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "registration_id" {
#   value = data.f5xc_registration.example.id
# }
