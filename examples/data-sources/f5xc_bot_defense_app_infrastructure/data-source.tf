# Bot Defense App Infrastructure Data Source Example
# Retrieves information about an existing Bot Defense App Infrastructure

# Look up an existing Bot Defense App Infrastructure by name
data "f5xc_bot_defense_app_infrastructure" "example" {
  name      = "example-bot-defense-app-infrastructure"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "bot_defense_app_infrastructure_id" {
#   value = data.f5xc_bot_defense_app_infrastructure.example.id
# }
