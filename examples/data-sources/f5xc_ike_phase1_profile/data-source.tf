# Ike Phase1 Profile Data Source Example
# Retrieves information about an existing Ike Phase1 Profile

# Look up an existing Ike Phase1 Profile by name
data "f5xc_ike_phase1_profile" "example" {
  name      = "example-ike-phase1-profile"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "ike_phase1_profile_id" {
#   value = data.f5xc_ike_phase1_profile.example.id
# }
