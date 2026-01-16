# Phase 1: Platform & Technology Stack Selection

**Document Type:** Planning  
**Phase:** 1 of 9 (Planning)  
**Status:** Complete  
**Date:** 2026-01-16  
**Authority:** Derived from STRAND_PID.md

---

## 1. Platform Requirements

### 1.1 Target Platforms

**Primary Platform:** Cross-platform CLI tool  
**Supported Operating Systems:**
- ✅ Windows (primary development environment)
- ✅ macOS
- ✅ Linux

**Deployment Model:** Single binary executable (no dependencies)

### 1.2 Platform Rationale

From the PID (Section 3, Markdown-First Principle):
> "ALL task data MUST be readable and editable as plain Markdown by humans."

This requires:
- Cross-platform file system access
- Git integration (platform-agnostic)
- Terminal/CLI interface (universally available)
- No cloud dependencies (local-first)

---

## 2. Technology Stack Selection

### 2.1 Programming Language: **Go**

**Selection Rationale:**
- ✅ Single binary compilation (no runtime dependencies)
- ✅ Excellent cross-platform support (Windows, macOS, Linux)
- ✅ Strong standard library for file I/O and CLI
- ✅ Fast compilation and execution
- ✅ Built-in concurrency for file watching
- ✅ Mature ecosystem for CLI tools

**Version:** Go 1.21+ (with modern generics support)

**Alternatives Considered:**
- ❌ **Rust**: Great performance but steeper learning curve, longer compile times
- ❌ **Python**: Requires runtime, harder to distribute as single binary
- ❌ **Node.js**: Requires runtime, larger binary size

### 2.2 Core Libraries

#### CLI Framework: **Cobra**
- **Purpose:** Command-line interface structure
- **Why:** Industry standard for Go CLI tools (used by kubectl, gh, hugo)
- **Repository:** https://github.com/spf13/cobra
- **License:** Apache 2.0

#### TUI Framework: **Bubble Tea + Lip Gloss**
- **Purpose:** Terminal user interface (interactive mode)
- **Why:** Modern, composable TUI framework with excellent aesthetics
- **Repositories:**
  - Bubble Tea: https://github.com/charmbracelet/bubbletea
  - Lip Gloss: https://github.com/charmbracelet/lipgloss
- **License:** MIT

#### Database: **SQLite** (via go-sqlite3)
- **Purpose:** Cache layer for fast queries
- **Why:** Embedded, zero-configuration, ACID-compliant
- **Repository:** https://github.com/mattn/go-sqlite3
- **License:** MIT
- **Note:** Uses CGo (requires C compiler during build)

#### File Watching: **fsnotify**
- **Purpose:** Monitor file system changes for auto-sync
- **Why:** Cross-platform, battle-tested, part of Go ecosystem
- **Repository:** https://github.com/fsnotify/fsnotify
- **License:** BSD-3-Clause

#### Markdown Processing: **goldmark**
- **Purpose:** Parse and generate Markdown
- **Why:** CommonMark compliant, extensible, pure Go
- **Repository:** https://github.com/yuin/goldmark
- **License:** MIT

#### YAML Processing: **go-yaml**
- **Purpose:** Parse YAML frontmatter in Markdown files
- **Why:** De facto standard for Go, full YAML 1.2 support
- **Repository:** https://github.com/go-yaml/yaml
- **License:** Apache 2.0 / MIT

---

## 3. Architecture Decisions

### 3.1 Data Storage Strategy

**Dual-Format Approach** (from PID Section 4):
```
.strand/
├── tasks/              ← Markdown files (source of truth)
│   ├── epic-001.md
│   ├── task-001-1.md
│   └── task-001-1-1.md
└── .cache/
    └── graph.db        ← SQLite cache (auto-generated)
```

**Key Principles:**
1. Markdown is the **source of truth**
2. SQLite is **derived** and **regeneratable**
3. File watcher keeps them in sync
4. Humans edit Markdown, agents query SQLite

### 3.2 File Format

**Markdown with YAML Frontmatter:**
```markdown
---
id: strand-a3f8.1
type: task
status: in-progress
priority: high
depends_on: [strand-a3f8]
assignee: agent
created: 2026-01-15T07:00:00Z
updated: 2026-01-15T07:05:00Z
tags: [frontend, ui]
---

# Task Title

## Description
...
```

### 3.3 Dependency Management

**Go Modules** (go.mod)
- No vendoring (rely on Go module cache)
- Semantic versioning for dependencies
- Minimal dependency tree

---

## 4. Development Tools

### 4.1 Build System

**Make** (optional, for convenience)
- Cross-platform build targets
- Test runners
- Linting integration

### 4.2 Code Quality

**golangci-lint**
- Comprehensive linter aggregator
- Enforce Go best practices
- Pre-commit hooks

### 4.3 Testing Framework

**Go standard library `testing`**
- Table-driven tests
- Subtests for organization
- Benchmark support

**testify** (assertions library)
- Cleaner assertions
- Mock generation

---

## 5. Distribution Strategy

### 5.1 Binary Distribution

**GitHub Releases** with **GoReleaser**
- Automated multi-platform builds
- Checksums and signatures
- Homebrew tap generation

**Platforms:**
- Windows: `strand.exe` (amd64, arm64)
- macOS: `strand` (amd64, arm64)
- Linux: `strand` (amd64, arm64, arm)

### 5.2 Package Managers

**Phase 1 (MVP):**
- Direct binary download from GitHub Releases

**Phase 2 (Post-MVP):**
- Homebrew tap (macOS/Linux)
- Scoop bucket (Windows)
- Go install (`go install github.com/hamsa0x7/strand@latest`)

---

## 6. Technology Validation

### 6.1 MCP Validation Required

Before finalizing, the AI must:
1. ✅ Use **Context7 MCP** to validate:
   - Cobra API compatibility
   - Bubble Tea best practices
   - go-sqlite3 usage patterns
   - goldmark extensions

2. ✅ Use **anti-bs MCP** to validate:
   - Performance claims (single binary, fast queries)
   - Cross-platform compatibility claims
   - SQLite suitability for this use case

---

## 7. Risk Assessment

### 7.1 Technical Risks

| Risk | Likelihood | Impact | Mitigation |
|------|-----------|--------|------------|
| SQLite sync performance issues | Low | Medium | Batch updates, debounce file watcher |
| Cross-platform file system edge cases | Medium | Medium | Extensive testing on all platforms |
| go-sqlite3 CGo build complexity | Medium | Low | Document build requirements, consider pure Go alternatives |
| YAML frontmatter parsing conflicts | Low | High | Strict validation, schema enforcement |

### 7.2 Dependency Risks

All chosen libraries are:
- ✅ Actively maintained
- ✅ Widely used in production
- ✅ Permissive licenses (MIT, Apache 2.0, BSD)
- ✅ No known security vulnerabilities

---

## 8. Compliance with PID

### 8.1 Core Principles Alignment

✅ **Markdown-First Principle** (PID Section 3)
- Storage format is Markdown with YAML frontmatter

✅ **Dual-Format Strategy** (PID Section 4)
- Markdown source of truth + SQLite cache

✅ **Technology Stack** (PID Section 10)
- Language: Go ✅
- Storage: Markdown + SQLite ✅
- TUI: Bubble Tea / Lip Gloss ✅
- CLI: Cobra ✅
- File watching: fsnotify ✅

---

## 9. Phase 1 Deliverables

✅ **Technology Stack Defined:**
- Language: Go 1.21+
- Core libraries selected and justified
- Architecture patterns established

✅ **Platform Support Confirmed:**
- Windows, macOS, Linux
- Single binary distribution

✅ **Risk Assessment Complete:**
- Technical risks identified and mitigated
- Dependency risks evaluated

---

## 10. Next Steps

**Phase 2: Product Design Review (PDR)**
- Define product vision and philosophy
- Establish core feature list (MVP vs Full)
- Document non-goals and boundaries

**Commit Message:**
```
Phase 1: Platform & Tech Stack complete

- Selected Go 1.21+ as primary language
- Defined core libraries (Cobra, Bubble Tea, SQLite, goldmark)
- Established dual-format storage strategy
- Validated cross-platform compatibility
```

---

**This document is constitutionally aligned with STRAND_PID.md and ready for Phase 2.**
