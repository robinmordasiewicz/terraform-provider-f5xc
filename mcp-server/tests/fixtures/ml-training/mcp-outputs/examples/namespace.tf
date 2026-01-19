# Complete Example: Namespace (Basic Pattern)

Creates a namespace in F5 Distributed Cloud.

```hcl
resource "f5xc_namespace" "example" {
  name = "my-namespace"
}
```

## Key Syntax Notes

- Namespace is the simplest resource - just needs a name
- Most other resources require a namespace reference
