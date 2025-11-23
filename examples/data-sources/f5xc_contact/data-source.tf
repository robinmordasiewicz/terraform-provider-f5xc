# Contact Data Source Example
# Retrieves information about an existing Contact

# Look up an existing Contact by name
data "f5xc_contact" "example" {
  name      = "example-contact"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "contact_id" {
#   value = data.f5xc_contact.example.id
# }
