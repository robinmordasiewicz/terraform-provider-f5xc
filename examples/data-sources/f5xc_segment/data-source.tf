# Segment Data Source Example
# Retrieves information about an existing Segment

# Look up an existing Segment by name
data "f5xc_segment" "example" {
  name      = "example-segment"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "segment_id" {
#   value = data.f5xc_segment.example.id
# }
