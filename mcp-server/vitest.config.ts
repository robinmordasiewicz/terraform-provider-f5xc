// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

import { defineConfig } from 'vitest/config';

export default defineConfig({
  test: {
    globals: true,
    environment: 'node',
    include: ['tests/**/*.test.ts'],
    coverage: {
      provider: 'v8',
      reporter: ['text', 'json', 'html'],
      include: ['src/**/*.ts'],
      exclude: ['src/**/*.d.ts', 'src/services/**'],
    },
    // Default timeout for unit/integration tests
    testTimeout: 30000,
    hookTimeout: 30000,
    // E2E tests may need longer timeouts - individual tests can override
    // Resource creation tests: 5 minutes per test
    // Determinism tests: 1 minute per test
    pool: 'forks',
    isolate: true,
    // Retry failed tests once (for network flakiness, not config errors)
    retry: 0, // IMPORTANT: No retries - determinism requires first-try success
    // Sequence tests to avoid resource conflicts
    sequence: {
      shuffle: false,
    },
  },
});
