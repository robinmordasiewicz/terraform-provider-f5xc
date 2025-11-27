---
page_title: "Test: Unescaped HTML Tags"
subcategory: "Testing"
description: |-
  Test guide with unescaped HTML tags to verify Registry markdown parsing behavior.
---

# Test: Unescaped HTML Tags

This guide tests whether unescaped HTML tags break Terraform Registry markdown parsing.

## Section Before HTML Tags

This section appears before any HTML tags and should render correctly.

### Argument Reference

- `name` - Required String. Name of the resource
- `description` - Optional String. Description of the resource

## Section With Unescaped HTML Tags

This section contains unescaped HTML tags that may break the markdown parser.

### JavaScript Location Settings

The JavaScript can be inserted at different locations:

- Insert JavaScript after <head> tag
- Insert JavaScript after </title> tag
- Insert JavaScript before first <script> tag

These HTML-like tags above are unescaped and may cause parsing issues.

## Section After HTML Tags

If this section renders on the Terraform Registry, then unescaped HTML tags do NOT break parsing.

### Additional Arguments

- `enabled` - Optional Bool. Enable the feature
- `timeout` - Optional Number. Timeout in seconds

## Final Section

This is the final section of the test guide. If you can see this on the Registry, the entire document rendered successfully.
