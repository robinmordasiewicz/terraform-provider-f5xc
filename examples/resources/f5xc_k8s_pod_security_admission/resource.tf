# K8S Pod Security Admission Resource Example
# Manages k8s_pod_security_admission will create the object in the storage backend in F5 Distributed Cloud.

# Basic K8S Pod Security Admission configuration
resource "f5xc_k8s_pod_security_admission" "example" {
  name      = "example-k8s-pod-security-admission"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # K8s Pod Security Admission.
  pod_security_admission_specs {
    # Configure pod_security_admission_specs settings
  }
  # Enable this option
  audit {
    # Configure audit settings
  }
  # Enable this option
  baseline {
    # Configure baseline settings
  }
}
