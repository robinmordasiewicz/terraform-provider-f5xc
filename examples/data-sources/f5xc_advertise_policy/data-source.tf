# Advertise Policy Data Source Example
# Retrieves information about an existing Advertise Policy

# Look up an existing Advertise Policy by name
data "f5xc_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "advertise_policy_id" {
#   value = data.f5xc_advertise_policy.example.id
# }
