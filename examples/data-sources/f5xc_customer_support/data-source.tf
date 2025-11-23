# Customer Support Data Source Example
# Retrieves information about an existing Customer Support

# Look up an existing Customer Support by name
data "f5xc_customer_support" "example" {
  name      = "example-customer-support"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "customer_support_id" {
#   value = data.f5xc_customer_support.example.id
# }
