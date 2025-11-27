---
page_title: "Test: Escaped HTML Tags"
description: |-
  Test guide with escaped HTML tags (using backticks) to verify Registry markdown parsing behavior.
---

# Test: Escaped HTML Tags

This guide tests whether escaped HTML tags (using backticks) render correctly on Terraform Registry.

## Section Before HTML Tags

This section appears before any HTML tags and should render correctly.

### Argument Reference

- `name` - Required String. Name of the resource
- `description` - Optional String. Description of the resource

## Section With Escaped HTML Tags

This section contains HTML tags escaped with backticks, which should render safely.

### JavaScript Location Settings

The JavaScript can be inserted at different locations:

- Insert JavaScript after `<head>` tag
- Insert JavaScript after `</title>` tag
- Insert JavaScript before first `<script>` tag

These HTML-like tags above are escaped with backticks and should render as code.

## Section After HTML Tags

If this section renders on the Terraform Registry, then escaped HTML tags are handled correctly.

### Additional Arguments

- `enabled` - Optional Bool. Enable the feature
- `timeout` - Optional Number. Timeout in seconds

## Final Section

This is the final section of the test guide. If you can see this on the Registry, the entire document rendered successfully.
