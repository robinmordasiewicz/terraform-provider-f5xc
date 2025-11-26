# K8s Pod Security Policy Resource Example
# Manages k8s_pod_security_policy will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

# Basic K8s Pod Security Policy configuration
resource "f5xc_k8s_pod_security_policy" "example" {
  name      = "example-k8s-pod-security-policy"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: psp_spec, yaml] Pod Security Policy Specification...
  psp_spec {
    # Configure psp_spec settings
  }
  # Capability List. List of capabilities that docker contain...
  allowed_capabilities {
    # Configure allowed_capabilities settings
  }
  # Allowed Host Paths. Restrict list of host paths, default ...
  allowed_host_paths {
    # Configure allowed_host_paths settings
  }
}
