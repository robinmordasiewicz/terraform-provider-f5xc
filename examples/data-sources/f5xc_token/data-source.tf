# Token Data Source Example
# Retrieves information about an existing Token

# Look up an existing Token by name
data "f5xc_token" "example" {
  name      = "example-token"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "token_id" {
#   value = data.f5xc_token.example.id
# }
