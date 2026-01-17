# Strand Development - AI Agent Governance

**Project:** Strand - Task Tracking System  
**Generated:** 2026-01-15

---

## Purpose

This file defines the governance rules for AI agents working on the Strand project. It serves as the **constitutional authority** during development (Phase 9), ensuring code quality, consistency, and alignment with the Product Ideation Document (PID).

---

## Core Laws (Never Violate)

1. **PID Authority**: The `STRAND_PID.md` is the constitutional source of truth. No code may contradict it.

2. **Markdown-First Principle**: ALL task data MUST be readable and editable as plain Markdown by humans.

3. **Dual-Format Strategy**: Store as Markdown (source of truth), cache as SQLite (auto-generated).

4. **No Silent Failures**: Every error must be surfaced to the user with context.

5. **Dependency Validation**: Dependency graph must always be consistent and acyclic.

6. **MCP Integration**: Use Context7 MCP for Go documentation, anti-bs MCP for technical claim validation.

---

## Task Tracking

Use Strand for all task management during development:

### Query Tasks
```bash
strand ready                       # Shows tasks ready to work on (no blockers)
strand list --status in-progress  # Show active tasks
strand show <id>                   # View task details
```

### Create Tasks
```bash
strand create "Task description" --priority high  # Create task
strand dep add <child> <parent>                   # Add dependency
```

### Update Tasks
```bash
strand update <id> --status in-progress  # Start working
strand update <id> --status done         # Mark complete
```

### Example Workflow
```bash
# Before starting work
strand ready

# Create tasks
strand create "Feature: Dependency Graph" --priority high
strand create "Implement graph validation" --priority medium
strand create "Add cycle detection" --priority medium

# Add dependencies
strand dep add <cycle-id> <validation-id>

# Work on ready tasks
strand update <validation-id> --status in-progress
# ... implement ...
strand update <validation-id> --status done

# Now cycle detection is unblocked
strand ready  # Returns cycle detection task
```

---

## Technology Stack Governance

### Language: Go
- **Version:** Go 1.21+
- **Style:** Follow [Effective Go](https://go.dev/doc/effective_go)
- **Linting:** Use `golangci-lint` with strict settings
- **Error Handling:** Always wrap errors with context using `fmt.Errorf("context: %w", err)`

### Key Libraries
- **CLI:** [Cobra](https://github.com/spf13/cobra)
- **TUI:** [Bubble Tea](https://github.com/charmbracelet/bubbletea) + [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- **SQLite:** [go-sqlite3](https://github.com/mattn/go-sqlite3)
- **File Watching:** [fsnotify](https://github.com/fsnotify/fsnotify)
- **Markdown:** [goldmark](https://github.com/yuin/goldmark)
- **YAML:** [go-yaml](https://github.com/go-yaml/yaml)

### Before Using ANY Library
1. Use Context7 MCP to verify the library exists and get official docs
2. Check for official examples and best practices
3. Validate API methods before using them

---

## Code Architecture Rules

### Project Structure
```
Development/
├── cmd/
│   └── strand/           # CLI entry point
├── internal/
│   ├── core/             # Core domain logic
│   │   ├── task.go       # Task struct and operations
│   │   ├── graph.go      # Dependency graph
│   │   └── store.go      # Storage interface
│   ├── markdown/         # Markdown parsing and writing
│   ├── cache/            # SQLite cache layer
│   ├── watcher/          # File watching and sync
│   ├── cli/              # CLI commands (Cobra)
│   └── tui/              # TUI interface (Bubble Tea)
├── pkg/
│   └── strand/           # Public API (if needed)
└── test/
    ├── fixtures/         # Test data
    └── integration/      # Integration tests
```

### Module Boundaries
- **core**: Pure domain logic, no external dependencies
- **markdown**: Parses/writes Markdown, no SQL, no file I/O
- **cache**: SQLite operations, implements store interface
- **watcher**: File system monitoring, triggers sync
- **cli**: User commands, orchestrates modules
- **tui**: Interactive interface, read-only state

### Dependency Flow
```
cli/tui → watcher → cache → core
      ↓              ↓       ↑
      └──────→ markdown ────┘
```

---

## Development Workflow

### Phase 9.1: MVP Implementation
**Goal:** Basic CLI with Markdown storage and SQLite cache

#### Must-Have Features
1. Markdown file storage with YAML frontmatter
2. SQLite cache layer with auto-sync
3. Basic CLI commands:
   - `strand init` - Initialize .strand/ directory
   - `strand create "title"` - Create task
   - `strand list` - List tasks
   - `strand show <id>` - Show task details
   - `strand update <id> --status=<status>` - Update task status
4. File watcher for auto-sync

#### Out of Scope for MVP
- ❌ Dependencies (Phase 2)
- ❌ TUI (Phase 3)
- ❌ GraphQL (Phase 3)
- ❌ Memory decay (Phase 3)

### Testing Requirements
- **Unit tests** for all core logic
- **Integration tests** for file operations
- **Table-driven tests** for parsing
- **Example-based tests** for CLI commands

### Git Workflow
```bash
# After completing each feature
git add .
git commit -m "feat: <feature-name>

- Detail 1
- Detail 2

Refs: bd-<task-id>"

# Ask before pushing
AI: "Feature complete. Push to GitHub? (y/n)"
```

---

## Error Handling Standards

### Pattern
```go
// Good: Contextual error wrapping
func readTaskFile(path string) (*Task, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("failed to read task file %s: %w", path, err)
    }
    // ... parse ...
}

// Bad: Silent errors
func readTaskFile(path string) *Task {
    data, _ := os.ReadFile(path)  // ❌ NEVER DO THIS
    // ...
}
```

### User-Facing Errors
```go
// Clear, actionable error messages
fmt.Errorf("task file %s is malformed: missing 'id' field in frontmatter", path)

// Not this:
fmt.Errorf("parse error")  // ❌ Too vague
```

---

## Validation Checklist

Before marking any task as complete:

- [ ] Code compiles without warnings
- [ ] All tests pass
- [ ] Error handling is comprehensive
- [ ] MCP validation (Context7 for APIs, anti-bs for claims)
- [ ] Follows Go style guidelines
- [ ] No hardcoded paths or magic numbers
- [ ] Logging is informative
- [ ] Documented in code comments
- [ ] Added to CHANGELOG

---

## Markdown File Format (Reference)

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
Detailed description here

## Acceptance Criteria
- [ ] Criterion 1
- [ ] Criterion 2

## Notes
Additional notes
```

---

## When in Doubt

1. **Check PID first** - Does this align with Strand's philosophy?
2. **Ask the user** - Don't assume, clarify!
3. **Use MCP** - Validate with Context7 or anti-bs

---

**This file is the law during development. Violations invalidate the work.**
