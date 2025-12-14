#!/usr/bin/env node
/**
 * Copy documentation files to dist/ for npm distribution
 *
 * This script copies the auto-generated documentation from the parent
 * terraform-provider-f5xc project into the MCP server's dist directory
 * so they are included in the npm package.
 */

import { cpSync, existsSync, mkdirSync, rmSync, readdirSync, statSync } from 'fs';
import { join, dirname } from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

// Paths
const MCP_ROOT = join(__dirname, '..');
const PROJECT_ROOT = join(MCP_ROOT, '..');
const DIST_DOCS = join(MCP_ROOT, 'dist', 'docs');

// Source paths (from parent project)
const SOURCES = {
  resources: join(PROJECT_ROOT, 'docs', 'resources'),
  dataSources: join(PROJECT_ROOT, 'docs', 'data-sources'),
  functions: join(PROJECT_ROOT, 'docs', 'functions'),
  guides: join(PROJECT_ROOT, 'docs', 'guides'),
  apiSpecs: join(PROJECT_ROOT, 'docs', 'specifications', 'api'),
};

// Destination paths (in dist/)
const DESTINATIONS = {
  resources: join(DIST_DOCS, 'resources'),
  dataSources: join(DIST_DOCS, 'data-sources'),
  functions: join(DIST_DOCS, 'functions'),
  guides: join(DIST_DOCS, 'guides'),
  apiSpecs: join(DIST_DOCS, 'specifications', 'api'),
};

console.log('Bundling documentation for npm distribution...\n');

// Clean existing docs in dist
if (existsSync(DIST_DOCS)) {
  console.log('Cleaning existing dist/docs...');
  rmSync(DIST_DOCS, { recursive: true });
}

// Create dist/docs directory
mkdirSync(DIST_DOCS, { recursive: true });

// Copy each documentation type
let totalFiles = 0;

for (const [key, source] of Object.entries(SOURCES)) {
  const dest = DESTINATIONS[key];

  if (!existsSync(source)) {
    console.log(`  [SKIP] ${key}: Source not found (${source})`);
    continue;
  }

  try {
    cpSync(source, dest, { recursive: true });
    const count = countFiles(dest);
    totalFiles += count;
    console.log(`  [OK] ${key}: Copied ${count} files`);
  } catch (error) {
    console.error(`  [ERROR] ${key}: ${error.message}`);
  }
}

console.log(`\nBundled ${totalFiles} documentation files to dist/docs/`);

/**
 * Count files in a directory recursively
 */
function countFiles(dir) {
  let count = 0;

  try {
    const items = readdirSync(dir);
    for (const item of items) {
      const itemPath = join(dir, item);
      if (statSync(itemPath).isDirectory()) {
        count += countFiles(itemPath);
      } else {
        count++;
      }
    }
  } catch {
    // Ignore errors
  }

  return count;
}
