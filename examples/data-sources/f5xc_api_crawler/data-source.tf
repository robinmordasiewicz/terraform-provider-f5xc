# Api Crawler Data Source Example
# Retrieves information about an existing Api Crawler

# Look up an existing Api Crawler by name
data "f5xc_api_crawler" "example" {
  name      = "example-api-crawler"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "api_crawler_id" {
#   value = data.f5xc_api_crawler.example.id
# }
