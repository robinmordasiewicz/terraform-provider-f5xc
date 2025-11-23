# Report Config Data Source Example
# Retrieves information about an existing Report Config

# Look up an existing Report Config by name
data "f5xc_report_config" "example" {
  name      = "example-report-config"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "report_config_id" {
#   value = data.f5xc_report_config.example.id
# }
