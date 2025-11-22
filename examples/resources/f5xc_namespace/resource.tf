resource "f5xc_namespace" "example" {
  name        = "my-namespace"
  description = "Example namespace for application workloads"
  
  labels = {
    environment = "production"
    team        = "platform"
  }
}
