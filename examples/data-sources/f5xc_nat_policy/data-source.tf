# Nat Policy Data Source Example
# Retrieves information about an existing Nat Policy

# Look up an existing Nat Policy by name
data "f5xc_nat_policy" "example" {
  name      = "example-nat-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "nat_policy_id" {
#   value = data.f5xc_nat_policy.example.id
# }
