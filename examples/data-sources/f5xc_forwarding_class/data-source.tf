# Forwarding Class Data Source Example
# Retrieves information about an existing Forwarding Class

# Look up an existing Forwarding Class by name
data "f5xc_forwarding_class" "example" {
  name      = "example-forwarding-class"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "forwarding_class_id" {
#   value = data.f5xc_forwarding_class.example.id
# }
