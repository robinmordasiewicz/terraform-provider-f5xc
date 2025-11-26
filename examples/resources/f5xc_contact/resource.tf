# Contact Resource Example
# Manages new customer's contact detail record with us, including address and phone number. in F5 Distributed Cloud.

# Basic Contact configuration
resource "f5xc_contact" "example" {
  name      = "example-contact"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }
}
