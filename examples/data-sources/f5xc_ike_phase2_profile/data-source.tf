# Ike Phase2 Profile Data Source Example
# Retrieves information about an existing Ike Phase2 Profile

# Look up an existing Ike Phase2 Profile by name
data "f5xc_ike_phase2_profile" "example" {
  name      = "example-ike-phase2-profile"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "ike_phase2_profile_id" {
#   value = data.f5xc_ike_phase2_profile.example.id
# }
