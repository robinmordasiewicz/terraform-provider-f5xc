# API Crawler Data Source Example
# Retrieves information about an existing API Crawler

# Look up an existing API Crawler by name
data "f5xc_api_crawler" "example" {
  name      = "example-api-crawler"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "api_crawler_id" {
#   value = data.f5xc_api_crawler.example.id
# }
