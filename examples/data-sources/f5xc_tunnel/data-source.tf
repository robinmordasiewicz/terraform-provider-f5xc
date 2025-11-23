# Tunnel Data Source Example
# Retrieves information about an existing Tunnel

# Look up an existing Tunnel by name
data "f5xc_tunnel" "example" {
  name      = "example-tunnel"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "tunnel_id" {
#   value = data.f5xc_tunnel.example.id
# }
