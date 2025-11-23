# Aws Tgw Site Data Source Example
# Retrieves information about an existing Aws Tgw Site

# Look up an existing Aws Tgw Site by name
data "f5xc_aws_tgw_site" "example" {
  name      = "example-aws-tgw-site"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "aws_tgw_site_id" {
#   value = data.f5xc_aws_tgw_site.example.id
# }
