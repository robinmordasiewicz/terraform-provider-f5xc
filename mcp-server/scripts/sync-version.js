#!/usr/bin/env node
// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

/**
 * Sync MCP server version with provider release version
 *
 * This script reads the version from the latest git tag and updates
 * the MCP server's package.json to match. This ensures the npm package
 * version always matches the Terraform provider release version.
 *
 * Usage:
 *   node scripts/sync-version.js
 *   node scripts/sync-version.js v2.4.3  # explicit version
 */

import { readFileSync, writeFileSync } from 'fs';
import { execSync } from 'child_process';
import { join, dirname } from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const MCP_ROOT = join(__dirname, '..');
const PACKAGE_JSON_PATH = join(MCP_ROOT, 'package.json');

/**
 * Get the version from command line argument or git tag
 */
function getVersion() {
  // Check for explicit version argument
  const explicitVersion = process.argv[2];
  if (explicitVersion) {
    return explicitVersion.replace(/^v/, '');
  }

  // Get from git tag
  try {
    const tag = execSync('git describe --tags --abbrev=0', {
      encoding: 'utf-8',
      stdio: ['pipe', 'pipe', 'pipe'],
    }).trim();
    return tag.replace(/^v/, '');
  } catch (error) {
    console.error('Error: Could not determine version from git tags');
    console.error('Usage: node scripts/sync-version.js [version]');
    process.exit(1);
  }
}

/**
 * Update package.json with the new version
 */
function updatePackageJson(version) {
  const packageJson = JSON.parse(readFileSync(PACKAGE_JSON_PATH, 'utf-8'));
  const oldVersion = packageJson.version;

  if (oldVersion === version) {
    console.log(`Version already at ${version}, no update needed`);
    return false;
  }

  packageJson.version = version;

  writeFileSync(PACKAGE_JSON_PATH, JSON.stringify(packageJson, null, 2) + '\n');
  console.log(`Updated MCP server version: ${oldVersion} → ${version}`);
  return true;
}

// Main execution
const version = getVersion();
console.log(`Syncing MCP server version to: ${version}`);
updatePackageJson(version);
