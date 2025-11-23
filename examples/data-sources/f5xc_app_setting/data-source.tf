# App Setting Data Source Example
# Retrieves information about an existing App Setting

# Look up an existing App Setting by name
data "f5xc_app_setting" "example" {
  name      = "example-app-setting"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "app_setting_id" {
#   value = data.f5xc_app_setting.example.id
# }
