/**
 * Documentation Service
 * Loads and manages Terraform provider documentation
 */

import { readFileSync, existsSync, readdirSync } from 'fs';
import { join, dirname, basename, extname } from 'path';
import { fileURLToPath } from 'url';
import type { ResourceDoc, SearchResult, SubscriptionMetadata, SubscriptionTier } from '../types.js';

// Get the package root directory
const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

// When installed via npm, docs are in dist/docs/ relative to package root
// When running in development, docs are in parent project's docs/
const PACKAGE_ROOT = join(__dirname, '..', '..'); // mcp-server/
const BUNDLED_DOCS = join(PACKAGE_ROOT, 'dist', 'docs');
const PROJECT_ROOT = join(PACKAGE_ROOT, '..'); // terraform-provider-f5xc/
const PROJECT_DOCS = join(PROJECT_ROOT, 'docs');

// Use bundled docs if available (npm install), otherwise fall back to project docs (development)
function getDocsRoot(): string {
  if (existsSync(BUNDLED_DOCS)) {
    return BUNDLED_DOCS;
  }
  return PROJECT_DOCS;
}

// Subscription metadata paths
const BUNDLED_SUBSCRIPTION_METADATA = join(PACKAGE_ROOT, 'dist', 'subscription-tiers.json');
const PROJECT_SUBSCRIPTION_METADATA = join(PROJECT_ROOT, 'tools', 'subscription-tiers.json');

// Cache for subscription metadata
let subscriptionMetadata: SubscriptionMetadata | null = null;

/**
 * Load subscription tier metadata
 */
function loadSubscriptionMetadata(): SubscriptionMetadata | null {
  if (subscriptionMetadata !== null) {
    return subscriptionMetadata;
  }

  // Try bundled first, then project path
  const paths = [BUNDLED_SUBSCRIPTION_METADATA, PROJECT_SUBSCRIPTION_METADATA];

  for (const metadataPath of paths) {
    if (existsSync(metadataPath)) {
      try {
        const content = readFileSync(metadataPath, 'utf-8');
        subscriptionMetadata = JSON.parse(content) as SubscriptionMetadata;
        return subscriptionMetadata;
      } catch (e) {
        console.error(`Failed to load subscription metadata from ${metadataPath}:`, e);
      }
    }
  }

  return null;
}

/**
 * Find resource metadata with name transformations
 */
function findResourceMetadata(resourceName: string): { tier: SubscriptionTier; service: string; advancedFeatures?: string[] } | null {
  const metadata = loadSubscriptionMetadata();
  if (!metadata) {
    return null;
  }

  // Try exact match first
  if (metadata.resources[resourceName]) {
    const res = metadata.resources[resourceName];
    return {
      tier: res.minimum_tier as SubscriptionTier,
      service: res.service,
      advancedFeatures: res.advanced_features,
    };
  }

  // Try common name transformations
  const nameVariants = [
    resourceName.replace(/^f5xc_/, ''),
    resourceName.replace(/_resource$/, ''),
    resourceName.replace(/bot_defense_app_/, 'bot_'),
    resourceName.replace(/bot_defense_/, ''),
    `shape_${resourceName}`,
    resourceName.replace(/^shape_/, ''),
    resourceName.replace(/client_side_defense_/, ''),
    resourceName.replace(/__/g, '_'),
  ];

  for (const variant of nameVariants) {
    if (metadata.resources[variant]) {
      const res = metadata.resources[variant];
      return {
        tier: res.minimum_tier as SubscriptionTier,
        service: res.service,
        advancedFeatures: res.advanced_features,
      };
    }
  }

  return null;
}

/**
 * Enrich a ResourceDoc with subscription tier information
 */
function enrichWithSubscriptionInfo(doc: ResourceDoc): ResourceDoc {
  const resourceMeta = findResourceMetadata(doc.name);
  if (resourceMeta) {
    return {
      ...doc,
      subscriptionTier: resourceMeta.tier,
      addonService: resourceMeta.service,
      advancedFeatures: resourceMeta.advancedFeatures,
    };
  }
  return doc;
}

// Documentation paths relative to docs root
function getDocsPaths() {
  const docsRoot = getDocsRoot();
  return {
    resources: join(docsRoot, 'resources'),
    dataSources: join(docsRoot, 'data-sources'),
    functions: join(docsRoot, 'functions'),
    guides: join(docsRoot, 'guides'),
  };
}

const DOCS_PATHS = getDocsPaths();

// Cache for loaded documentation
const docsCache = new Map<string, ResourceDoc>();
let allDocs: ResourceDoc[] | null = null;

/**
 * Load all documentation files
 */
export function loadAllDocumentation(): ResourceDoc[] {
  if (allDocs) {
    return allDocs;
  }

  allDocs = [];

  // Load resources (enriched with subscription tier info)
  if (existsSync(DOCS_PATHS.resources)) {
    const files = readdirSync(DOCS_PATHS.resources).filter(f => f.endsWith('.md'));
    for (const file of files) {
      const doc: ResourceDoc = {
        name: basename(file, '.md'),
        path: join(DOCS_PATHS.resources, file),
        type: 'resource',
      };
      allDocs.push(enrichWithSubscriptionInfo(doc));
    }
  }

  // Load data sources (enriched with subscription tier info)
  if (existsSync(DOCS_PATHS.dataSources)) {
    const files = readdirSync(DOCS_PATHS.dataSources).filter(f => f.endsWith('.md'));
    for (const file of files) {
      const doc: ResourceDoc = {
        name: basename(file, '.md'),
        path: join(DOCS_PATHS.dataSources, file),
        type: 'data-source',
      };
      allDocs.push(enrichWithSubscriptionInfo(doc));
    }
  }

  // Load functions
  if (existsSync(DOCS_PATHS.functions)) {
    const files = readdirSync(DOCS_PATHS.functions).filter(f => f.endsWith('.md'));
    for (const file of files) {
      allDocs.push({
        name: basename(file, '.md'),
        path: join(DOCS_PATHS.functions, file),
        type: 'function',
      });
    }
  }

  // Load guides
  if (existsSync(DOCS_PATHS.guides)) {
    const files = readdirSync(DOCS_PATHS.guides).filter(f => f.endsWith('.md'));
    for (const file of files) {
      allDocs.push({
        name: basename(file, '.md'),
        path: join(DOCS_PATHS.guides, file),
        type: 'guide',
      });
    }
  }

  return allDocs;
}

/**
 * Get documentation content for a specific resource
 */
export function getDocumentation(name: string, type?: string): ResourceDoc | null {
  const cacheKey = `${type || 'any'}:${name}`;

  if (docsCache.has(cacheKey)) {
    return docsCache.get(cacheKey)!;
  }

  const docs = loadAllDocumentation();
  let doc = docs.find(d => {
    const nameMatch = d.name === name || d.name === name.replace(/_/g, '-') || d.name === name.replace(/-/g, '_');
    const typeMatch = !type || d.type === type;
    return nameMatch && typeMatch;
  });

  if (doc && existsSync(doc.path)) {
    const content = readFileSync(doc.path, 'utf-8');
    doc = { ...doc, content };
    docsCache.set(cacheKey, doc);
    return doc;
  }

  return null;
}

/**
 * List all available documentation
 */
export function listDocumentation(type?: string): ResourceDoc[] {
  const docs = loadAllDocumentation();
  if (type) {
    return docs.filter(d => d.type === type);
  }
  return docs;
}

/**
 * Search documentation content
 */
export function searchDocumentation(query: string, type?: string, limit: number = 20): SearchResult[] {
  const docs = loadAllDocumentation();
  const results: SearchResult[] = [];
  const queryLower = query.toLowerCase();
  const queryTerms = queryLower.split(/\s+/).filter(t => t.length > 0);

  for (const doc of docs) {
    if (type && doc.type !== type) {
      continue;
    }

    // Check name match
    const nameLower = doc.name.toLowerCase();
    let score = 0;

    // Exact name match
    if (nameLower === queryLower) {
      score += 100;
    }
    // Name contains query
    else if (nameLower.includes(queryLower)) {
      score += 50;
    }
    // Name contains any query term
    else {
      for (const term of queryTerms) {
        if (nameLower.includes(term)) {
          score += 20;
        }
      }
    }

    // Search content if file exists
    if (existsSync(doc.path)) {
      const content = readFileSync(doc.path, 'utf-8').toLowerCase();

      // Count term occurrences in content
      for (const term of queryTerms) {
        const occurrences = (content.match(new RegExp(term, 'gi')) || []).length;
        score += Math.min(occurrences * 2, 30); // Cap content score per term
      }

      if (score > 0) {
        // Extract relevant snippet
        const contentOriginal = readFileSync(doc.path, 'utf-8');
        const snippet = extractSnippet(contentOriginal, queryTerms);

        results.push({
          name: doc.name,
          type: doc.type,
          path: doc.path,
          snippet,
          score,
        });
      }
    } else if (score > 0) {
      results.push({
        name: doc.name,
        type: doc.type,
        path: doc.path,
        snippet: `${doc.type}: ${doc.name}`,
        score,
      });
    }
  }

  // Sort by score descending
  results.sort((a, b) => b.score - a.score);

  return results.slice(0, limit);
}

/**
 * Extract a relevant snippet from content
 */
function extractSnippet(content: string, terms: string[], maxLength: number = 200): string {
  const contentLower = content.toLowerCase();

  // Find the first occurrence of any term
  let firstIndex = content.length;
  for (const term of terms) {
    const index = contentLower.indexOf(term);
    if (index !== -1 && index < firstIndex) {
      firstIndex = index;
    }
  }

  if (firstIndex === content.length) {
    // No term found, return beginning of content
    const lines = content.split('\n').filter(l => l.trim().length > 0);
    return lines.slice(0, 3).join(' ').slice(0, maxLength) + '...';
  }

  // Extract snippet around the term
  const start = Math.max(0, firstIndex - 50);
  const end = Math.min(content.length, firstIndex + maxLength - 50);
  let snippet = content.slice(start, end);

  // Clean up snippet
  snippet = snippet.replace(/\n+/g, ' ').replace(/\s+/g, ' ').trim();

  if (start > 0) {
    snippet = '...' + snippet;
  }
  if (end < content.length) {
    snippet = snippet + '...';
  }

  return snippet;
}

/**
 * Get resource categories summary
 */
export function getDocumentationSummary(): Record<string, number> {
  const docs = loadAllDocumentation();
  const summary: Record<string, number> = {};

  for (const doc of docs) {
    summary[doc.type] = (summary[doc.type] || 0) + 1;
  }

  return summary;
}

/**
 * Get subscription tier information for a specific resource
 */
export function getResourceSubscriptionInfo(resourceName: string): {
  tier: SubscriptionTier;
  service: string;
  advancedFeatures?: string[];
  requiresAdvanced: boolean;
} | null {
  const meta = findResourceMetadata(resourceName);
  if (!meta) {
    return null;
  }

  return {
    tier: meta.tier,
    service: meta.service,
    advancedFeatures: meta.advancedFeatures,
    requiresAdvanced: meta.tier === 'ADVANCED' || Boolean(meta.advancedFeatures && meta.advancedFeatures.length > 0),
  };
}

/**
 * Get all resources that require Advanced subscription tier
 */
export function getAdvancedTierResources(): ResourceDoc[] {
  const docs = loadAllDocumentation();
  return docs.filter(doc => {
    if (doc.type !== 'resource' && doc.type !== 'data-source') {
      return false;
    }
    // Check if resource itself requires ADVANCED tier
    if (doc.subscriptionTier === 'ADVANCED') {
      return true;
    }
    // Check if resource has any features requiring ADVANCED tier
    if (doc.advancedFeatures && doc.advancedFeatures.length > 0) {
      return true;
    }
    return false;
  });
}

/**
 * Get subscription metadata summary
 */
export function getSubscriptionSummary(): {
  totalResources: number;
  advancedOnlyResources: number;
  resourcesWithAdvancedFeatures: number;
  advancedFeaturesList: string[];
} {
  const docs = loadAllDocumentation();
  const resourceDocs = docs.filter(d => d.type === 'resource' || d.type === 'data-source');

  const advancedOnly = resourceDocs.filter(d => d.subscriptionTier === 'ADVANCED');
  const withAdvancedFeatures = resourceDocs.filter(d => d.advancedFeatures && d.advancedFeatures.length > 0);

  // Collect all unique advanced features
  const allFeatures = new Set<string>();
  for (const doc of resourceDocs) {
    if (doc.advancedFeatures) {
      doc.advancedFeatures.forEach(f => allFeatures.add(f));
    }
  }

  return {
    totalResources: resourceDocs.length,
    advancedOnlyResources: advancedOnly.length,
    resourcesWithAdvancedFeatures: withAdvancedFeatures.length,
    advancedFeaturesList: Array.from(allFeatures).sort(),
  };
}
