# K8s Pod Security Admission Resource Example
# Manages k8s_pod_security_admission will create the object in the storage backend in F5 Distributed Cloud.

# Basic K8s Pod Security Admission configuration
resource "f5xc_k8s_pod_security_admission" "example" {
  name      = "example-k8s-pod-security-admission"
  namespace = "system"

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
  # Empty. This can be used for messages where no values are ...
  audit {
    # Configure audit settings
  }
  # Empty. This can be used for messages where no values are ...
  baseline {
    # Configure baseline settings
  }
}
