# Infraprotect Internet Prefix Advertisement Data Source Example
# Retrieves information about an existing Infraprotect Internet Prefix Advertisement

# Look up an existing Infraprotect Internet Prefix Advertisement by name
data "f5xc_infraprotect_internet_prefix_advertisement" "example" {
  name      = "example-infraprotect-internet-prefix-advertisement"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "infraprotect_internet_prefix_advertisement_id" {
#   value = data.f5xc_infraprotect_internet_prefix_advertisement.example.id
# }
