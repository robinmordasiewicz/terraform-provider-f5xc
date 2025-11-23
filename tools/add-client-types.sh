#!/bin/bash
#
# Add Client Types for Generated Resources
#

set -e

CLIENT_FILE="/tmp/terraform-provider-f5xc/internal/client/client.go"
PROVIDER_DIR="/tmp/terraform-provider-f5xc/internal/provider"

# Extract resource names from generated files
# Using find instead of ls | grep to avoid SC2010
RESOURCES=$(find "$PROVIDER_DIR" -name '*_resource.go' -type f ! -name 'namespace_resource.go' ! -name 'http_loadbalancer_resource.go' ! -name 'origin_pool_resource.go' -exec basename {} \; | sed 's/_resource.go//')

echo "🔨 Adding Client Types for Generated Resources"
echo "=============================================="
echo ""

# Read the existing client.go to find insertion point
TEMP_FILE=$(mktemp)

# Find the line number before the final closing brace
LINE_NUM=$(grep -n "^}$" $CLIENT_FILE | tail -1 | cut -d: -f1)

# Insert before the last line
head -n $((LINE_NUM - 1)) $CLIENT_FILE > $TEMP_FILE

# Add each resource's client types
for resource in $RESOURCES; do
    # Convert to TitleCase
    title_case=$(echo $resource | sed -r 's/(^|_)([a-z])/\U\2/g' | sed 's/_//g')

    echo "Adding types for: $resource ($title_case)"

    cat >> $TEMP_FILE << EOF

// $title_case represents a F5XC $title_case
type $title_case struct {
	Metadata Metadata        \`json:"metadata"\`
	Spec     ${title_case}Spec \`json:"spec"\`
}

// ${title_case}Spec defines the specification for $title_case
type ${title_case}Spec struct {
	Description string \`json:"description,omitempty"\`
}

// Create$title_case creates a new $title_case
func (c *Client) Create$title_case(ctx context.Context, resource *$title_case) (*$title_case, error) {
	var result $title_case
	path := fmt.Sprintf("/api/config/namespaces/%s/${resource}s", resource.Metadata.Namespace)
	err := c.Post(ctx, path, resource, &result)
	return &result, err
}

// Get$title_case retrieves a $title_case
func (c *Client) Get$title_case(ctx context.Context, namespace, name string) (*$title_case, error) {
	var result $title_case
	path := fmt.Sprintf("/api/config/namespaces/%s/${resource}s/%s", namespace, name)
	err := c.Get(ctx, path, &result)
	return &result, err
}

// Update$title_case updates a $title_case
func (c *Client) Update$title_case(ctx context.Context, resource *$title_case) (*$title_case, error) {
	var result $title_case
	path := fmt.Sprintf("/api/config/namespaces/%s/${resource}s/%s", resource.Metadata.Namespace, resource.Metadata.Name)
	err := c.Put(ctx, path, resource, &result)
	return &result, err
}

// Delete$title_case deletes a $title_case
func (c *Client) Delete$title_case(ctx context.Context, namespace, name string) error {
	path := fmt.Sprintf("/api/config/namespaces/%s/${resource}s/%s", namespace, name)
	return c.Delete(ctx, path)
}
EOF
done

# Close the file
echo "}" >> $TEMP_FILE

# Replace original file
mv $TEMP_FILE $CLIENT_FILE

echo ""
echo "✅ Client types added successfully"
echo "📝 Updated: $CLIENT_FILE"
