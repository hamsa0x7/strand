# Strand - MVP Implementation

A git-backed, dependency-aware task tracking system with human-readable Markdown storage.

## 🚧 Development Status

**Phase 9.1: MVP Implementation - IN PROGRESS**

- ✅ Core Markdown storage with YAML frontmatter
- ✅ Basic CLI commands (init, create, list, show, update, ready)
- ⏸️ SQLite cache layer (deferred to post-MVP)
- ⏸️ File watcher (deferred to post-MVP)

## 🏗️ Building

**Note:** Go 1.21+ is required but not yet installed on this system.

Once Go is installed, build with:

```bash
go mod download
go build -o bin/strand ./cmd/strand
```

## 🎯 Usage

```bash
# Initialize a strand project
./strand init

# Create a task
./strand create "Implement user authentication"

# List all tasks
./strand list

# Show task details
./strand show strand-20260116123456

# Update task status
./strand update strand-20260116123456 --status in-progress

# List ready tasks (no blockers)
./strand ready
```

## 📁 Project Structure

```
.
├── cmd/strand/          # CLI entry point
├── internal/
│   ├── core/            # Domain models (Task)
│   ├── storage/         # Storage interface
│   ├── markdown/        # Markdown file storage implementation
│   └── cli/             # Cobra CLI commands
└── go.mod               # Dependencies
```

## 🧪 Testing

Once Go is installed:

```bash
go test ./...
```

## 📝 Task File Format

Tasks are stored as Markdown files with YAML frontmatter:

```markdown
---
id: strand-20260116123456
type: task
status: in-progress
priority: high
created: 2026-01-16T12:34:56Z
updated: 2026-01-16T14:22:10Z
tags: [backend, auth]
---

# Implement user authentication

## Description

Build JWT-based authentication system with refresh tokens.

## Acceptance Criteria
- [ ] Login endpoint
- [ ] Token refresh logic
- [ ] Middleware for protected routes
```

## 🌟 Features (MVP)

- ✅ Markdown-first storage (human-readable)
- ✅ YAML frontmatter for metadata
- ✅ CLI with table and JSON output
- ✅ Task statuses: backlog, ready, in-progress, done, blocked, cancelled
- ✅ Priority levels: critical, high, medium, low
- ✅ Dependency tracking (basic)
- ✅ Ready task detection (no blockers)

## 🚀 Next Steps (Post-MVP)

- SQLite cache for performance
- File watcher for auto-sync
- Beautiful TUI (Bubble Tea)
- Dependency graph visualization
- GraphQL query interface
- Auto-archiving completed tasks

## 📚 Documentation

See `Documents/` folder for:
- `01_Platform_Tech_Stack.md` - Technology decisions
- `02_Product_Design_Review.md` - Product vision and scope

## 📄 License

TBD
