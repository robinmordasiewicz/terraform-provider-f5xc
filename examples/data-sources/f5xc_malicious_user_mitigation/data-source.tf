# Malicious User Mitigation Data Source Example
# Retrieves information about an existing Malicious User Mitigation

# Look up an existing Malicious User Mitigation by name
data "f5xc_malicious_user_mitigation" "example" {
  name      = "example-malicious-user-mitigation"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "malicious_user_mitigation_id" {
#   value = data.f5xc_malicious_user_mitigation.example.id
# }
