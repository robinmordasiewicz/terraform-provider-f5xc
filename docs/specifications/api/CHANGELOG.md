# Changelog

## Version 2.0.13 (2026-01-05)

### Version Information
| Field | Value |
|-------|-------|
| Full Version | 2.0.13 |
| Upstream Timestamp | unknown |
| Upstream ETag | unknown |
| Enriched Version | 2.0.13 |

### Release Type
- **patch** release

### Changes
- Updated API specifications from F5 Distributed Cloud
- Applied enrichment pipeline:
  - Acronym normalization (100+ terms)
  - Grammar improvements
  - Branding updates (Volterra → F5 Distributed Cloud)
- Applied normalization pipeline:
  - Fixed orphan $ref references
  - Removed empty operations
  - Type standardization
- Validated with Spectral OpenAPI linter
- Merged specifications by domain

### Statistics
- Original specs: 270
- Domains: 38
- Total paths: 1604
- Total schemas: 9427

### API Discovery Enrichment
- Discovery timestamp: 2025-12-20T19:39:20.211392+00:00
- Endpoints explored: 300 / 1000
- Applied x-discovered-* extensions from live API exploration
- See `reports/constraint-analysis.md` for detailed comparison

### Output Structure
```text
docs/specifications/api/
├── [domain].json        # Domain-specific specs
├── openapi.json         # Master combined spec
└── index.json           # Metadata index
```

### Download
- ZIP Package: F5xc-api-(unknown-2.0.13).zip

### Source
- Source: F5 Distributed Cloud OpenAPI specifications
- Upstream: unknown (ETag: unknown)
