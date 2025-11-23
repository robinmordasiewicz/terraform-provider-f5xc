# Infraprotect Tunnel Data Source Example
# Retrieves information about an existing Infraprotect Tunnel

# Look up an existing Infraprotect Tunnel by name
data "f5xc_infraprotect_tunnel" "example" {
  name      = "example-infraprotect-tunnel"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "infraprotect_tunnel_id" {
#   value = data.f5xc_infraprotect_tunnel.example.id
# }
