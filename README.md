# Strand

**A git-backed, dependency-aware task tracking system for humans and AI agents**

[![Version](https://img.shields.io/badge/version-1.1.0-blue.svg)](https://github.com/hamsa0x7/strand/releases)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://go.dev/)
[![License](https://img.shields.io/badge/license-TBD-green.svg)](LICENSE)

---

## Why Strand?

**Markdown is your source of truth.** Tasks are stored as beautiful, editable `.md` files. AI agents and humans work together seamlessly. Zero lock-in, forever.

```markdown
---
id: strand-20260117134629
type: task
status: in-progress
priority: high
tags: [backend, auth]
---

# Implement user authentication

Build JWT-based auth system...
```

---

## Features

### 🚀 15 Powerful Commands

**Core:**
- `init` - Initialize project
- `create` - Create tasks with metadata
- `list` - View all tasks (table or JSON)
- `show` - Task details
- `update` - Change status/priority
- `delete` - Remove tasks

**Dependencies:**
- `dep add` - Create dependency
- `dep remove` - Remove dependency
- `dep list` - Show dependencies

**Advanced:**
- `ready` - Find unblocked tasks
- `search` - Full-text search with filters
- `graph` - Visualize dependency tree
- `edit` - Edit in $EDITOR
- `ui` - Interactive TUI

### ✨ What Makes Strand Special

✅ **Zero Lock-In** - Just Markdown files  
✅ **Dependency Tracking** - Never miss prerequisites  
✅ **Beautiful Graphs** - Visualize task hierarchy  
✅ **AI-Friendly** - JSON output for agents  
✅ **Human-Friendly** - Readable Markdown for you  
✅ **Git-Native** - Full version control  
✅ **Local-First** - No cloud required  
✅ **Interactive TUI** - Bubble Tea interface

---

## Quick Start

```bash
# Install
git clone https://github.com/hamsa0x7/strand.git
cd strand
go build -o strand.exe ./cmd/strand

# Use
cd my-project
strand init
strand create "My first task" --priority high
strand list
```

---

## Installation

### Option 1: Download Pre-Built Binary (Recommended)

**Latest Release:** [v1.1.0](https://github.com/hamsa0x7/strand/releases/tag/v1.1.0)

Visit the releases page above to download pre-compiled binaries for your platform (coming soon).

### Option 2: Build from Source

**Requirements:** Go 1.21+

```bash
git clone https://github.com/hamsa0x7/strand.git
cd strand
go mod download
go build -o bin/strand.exe ./cmd/strand
```

**Add to PATH** (optional):
```bash
# Windows
set PATH=%PATH%;e:\path\to\strand\bin

# Linux/macOS
export PATH=$PATH:/path/to/strand/bin
```

---

## Usage

### Basic Workflow

```bash
# Initialize
strand init

# Create tasks
strand create "Backend API" --type task --priority high
strand create "Frontend UI" --type task --priority medium

# Add dependency (Frontend depends on Backend)
strand dep add <frontend-id> <backend-id>

# See what's ready
strand ready

# Update status
strand update <task-id> --status in-progress

# Search
strand search "API" --status in-progress

# Visualize
strand graph
```

### Interactive TUI

```bash
strand ui
```

**Controls:**
- `↑/k` - Move up
- `↓/j` - Move down
- `space` - Select task
- `r` - Refresh
- `q` - Quit

---

## Examples

### Dependency Graph

```
Task Dependency Graph
=====================

└── 🔄 🟠 - Implement user authentication [in-progress]
    ├── ⚪ 🟡 - Build dashboard UI [backlog]
    └── ⚪ 🟡 - Setup database [backlog]
```

### Search with Filters

```bash
$ strand search auth --priority high --tags backend

Found 1 task(s) matching 'auth'

ID                     TYPE  STATUS        PRIORITY  TITLE
strand-20260117134629  task  in-progress   high      Implement user authentication
```

---

## Documentation

### In This Repo

- [`RELEASE_NOTES.md`](RELEASE_NOTES.md) - v1.1 release details
- [`BUILD.md`](BUILD.md) - Build instructions
- [`AGENTS.md`](AGENTS.md) - AI development governance
- [`Documents/`](Documents/) - Technical specs

### File Format

Tasks are stored as Markdown with YAML frontmatter:

```markdown
---
id: unique-task-id
type: task|epic|bug|story
status: backlog|ready|in-progress|done|blocked|cancelled
priority: critical|high|medium|low
created: "2026-01-17T13:46:29Z"
updated: "2026-01-17T14:20:00Z"
tags: [tag1, tag2]
depends_on: [other-task-id]
assignee: username
---

# Task Title

Task description in Markdown...
```

---

## Architecture

**Storage:** Markdown files with YAML frontmatter  
**Cache:** SQLite (optional, for performance)  
**CLI:** Cobra framework  
**TUI:** Bubble Tea + Lip Gloss  
**Language:** Go 1.21+

**Project Structure:**
```
.strand/
├── tasks/           # Markdown task files
│   └── *.md
├── .cache/          # SQLite cache
│   └── tasks.db
└── .gitignore       # Git ignore rules
```

---

## Roadmap

### ✅ v1.0 (MVP)
- Markdown storage
- Basic CLI commands
- Dependency tracking

### ✅ v1.1 (Full Product) `← WE ARE HERE`
- All 15 commands
- Dependency graph
- Interactive TUI
- Search & validation
- SQLite cache foundation

### 🔮 Future
- Cross-project task management
  - Global task view across multiple projects
  - Project field in task metadata
  - `strand list --all-projects` command
  - Multi-project search and filtering
- Real-time file watching
- Enhanced TUI features
- Multi-agent coordination
- Cloud sync (optional)

---

## Contributing

**Want to contribute?**
1. Fork the repo
2. Create a feature branch
3. Follow `AGENTS.md` governance
4. Submit a PR

**Technology Stack:**
- Go 1.21+
- Cobra (CLI)
- Bubble Tea (TUI)
- SQLite (Cache)
- Markdown + YAML

---

## License

TBD - Choose your license!

---

## Support

- **Issues:** https://github.com/hamsa0x7/strand/issues
- **Docs:** https://github.com/hamsa0x7/strand/tree/master/Documents

---

**Strand v1.1** - Task tracking that doesn't suck. 🚀

**Made with ❤️ using AI-assisted development**
