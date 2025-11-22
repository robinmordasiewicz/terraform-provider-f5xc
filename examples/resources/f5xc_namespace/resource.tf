resource "f5xc_namespace" "example" {
  name        = "example-namespace"
  description = "Example namespace created by terraform-provider-f5xc"

  labels = {
    environment = "development"
    managed-by  = "terraform"
  }
}
