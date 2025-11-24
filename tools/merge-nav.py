#!/usr/bin/env python3
"""Merge generated nav-api.yml into mkdocs.yml."""

import sys
import re

def main():
    mkdocs_file = "mkdocs.yml"
    nav_file = "docs/nav-api.yml"

    # Read mkdocs.yml
    with open(mkdocs_file, 'r') as f:
        mkdocs_content = f.read()

    # Read nav-api.yml (skip comment lines)
    with open(nav_file, 'r') as f:
        nav_lines = [line for line in f.readlines() if not line.startswith('#')]
    nav_content = ''.join(nav_lines)

    # Check if nav: already exists in mkdocs.yml
    if re.search(r'^nav:', mkdocs_content, re.MULTILINE):
        # Replace existing nav section - find where nav starts and where it ends
        # Nav ends when we hit a line that doesn't start with whitespace or is empty followed by non-nav config
        lines = mkdocs_content.split('\n')
        new_lines = []
        skip_nav = False
        for line in lines:
            if line.startswith('nav:'):
                skip_nav = True
                continue
            if skip_nav:
                # Stop skipping when we hit a non-indented line (except empty lines)
                if line and not line.startswith(' ') and not line.startswith('\t'):
                    skip_nav = False
                    new_lines.append(line)
                continue
            new_lines.append(line)
        mkdocs_content = '\n'.join(new_lines)

    # Append nav section at the end
    mkdocs_content = mkdocs_content.rstrip() + '\n\n' + nav_content

    # Write updated mkdocs.yml
    with open(mkdocs_file, 'w') as f:
        f.write(mkdocs_content)

    print(f"Updated {mkdocs_file} with navigation from {nav_file}")

if __name__ == "__main__":
    main()
