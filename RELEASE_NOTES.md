# Strand v1.1 - Release Notes

<p align="center">
  <img src="strand.png" alt="Strand Logo" width="200"/>
</p>

**Release Date:** January 17, 2026  
**Status:** Production Ready  
**Download:** https://github.com/hamsa0x7/strand

---

## 🎉 What is Strand?

Strand is a **git-backed, dependency-aware task tracking system** that combines human-readable Markdown storage with agent-optimized workflows. Perfect for AI-assisted developers and solo builders.

### Core Philosophy

- **Markdown is the constitution** - All tasks stored as editable `.md` files
- **Zero lock-in** - Your data is yours, forever
- **Agent-first, human-friendly** - Designed for AI collaboration
- **Local-first** - No cloud required
- **Git is the timeline** - Full version control built-in

---

## ✨ Features

### 15 CLI Commands

| Command | Purpose |
|---------|---------|
| `strand init` | Initialize a project |
| `strand create` | Create tasks with metadata |
| `strand list` | View all tasks |
| `strand show` | Task details |
| `strand update` | Change status/priority |
| `strand ready` | Find unblocked tasks |
| `strand delete` | Remove tasks |
| `strand dep add/remove/list` | Manage dependencies |
| `strand edit` | Edit in $EDITOR |
| `strand search` | Find tasks with filters |
| `strand graph` | Visualize dependency tree |
| `strand ui` | Interactive TUI |

### Advanced Capabilities

✅ **Dependency Management** - Track task relationships  
✅ **Intelligent Ready Detection** - Only show unblocked tasks  
✅ **Beautiful Graph Visualization** - See your task hierarchy  
✅ **Full-Text Search** - Find anything quickly  
✅ **Interactive TUI** - Bubble Tea interface  
✅ **SQLite Cache** - Fast queries (foundation laid)  
✅ **Input Validation** - Prevent invalid data  
✅ **JSON Output** - Perfect for AI agents

---

## 📦 Installation

### Option 1: Build from Source

```bash
git clone https://github.com/hamsa0x7/strand.git
cd strand
go build -o strand.exe ./cmd/strand
```

### Option 2: Download Binary (Coming Soon)

Check GitHub Releases for pre-built binaries.

---

## 🚀 Quick Start

```bash
# Initialize in your project
cd my-project
strand init

# Create your first task
strand create "Implement user auth" --priority high --tags backend

# Create a dependent task
strand create "Build login UI" --priority medium
strand dep add <login-ui-id> <auth-id>

# See what's ready to work on
strand ready

# Launch interactive UI
strand ui

# Visualize dependencies
strand graph
```

---

## 📊 What's New in v1.1

**Core Commands ✅**
- Delete command with confirmation
- Complete dependency management (add/remove/list)
- Edit command with $EDITOR support

**Enhanced Features ✅**
- Full-text search with filters
- Input validation for all enums
- Helpful error messages

**Advanced Features ✅**
- Dependency graph visualization (emoji + ASCII)
- Interactive TUI with Bubble Tea
- SQLite cache foundation

### Stats

- **Commands:** 15
- **Lines of Code:** ~4,000+
- **Features:** All planned features complete
- **Production Ready:** ✅ YES

---

## 🎨 Examples

### Graph Visualization

```
Task Dependency Graph
=====================

└── 🔄 🟠 - Implement user authentication [in-progress]
    └── ⚪ 🟡 - Build dashboard UI [backlog]
```

### Task Markdown File

```markdown
---
id: strand-20260117134629
type: task
status: in-progress
priority: high
created: "2026-01-17T13:46:29+04:00"
updated: "2026-01-17T13:50:04+04:00"
tags:
  - backend
  - auth
depends_on:
  - strand-20260117134500
---

# Implement user authentication

## Description
Build JWT-based auth with refresh tokens.

## Acceptance Criteria
- [ ] Login endpoint
- [ ] Token refresh
- [ ] Protected routes
```

---

## 🎯 Use Cases

### AI-Assisted Development
- **AI agents can read/write** tasks directly
- **JSON output** for programmatic access
- **Markdown format** for human review

### Solo Builders
- **Simple CLI workflow** - no complex UI
- **Git-backed** - automatic history
- **Local-first** - work offline

### Development Teams
- **Structured task tracking** - follows best practices
- **Dependency awareness** - never miss prerequisites
- **Visual feedback** - graphs and TUI

---

## 🔧 Technical Details

### Architecture

**Storage:** Markdown files with YAML frontmatter  
**Cache:** SQLite (for future performance)  
**CLI:** Cobra framework  
**TUI:** Bubble Tea + Lip Gloss  
**Language:** Go 1.21+

### File Structure

```
.strand/
├── tasks/
│   ├── strand-xxx.md  # Markdown task files
│   └── strand-yyy.md
├── .cache/
│   └── tasks.db       # SQLite cache
└── .gitignore
```

### Dependencies

- `github.com/spf13/cobra` - CLI framework
- `github.com/charmbracelet/bubbletea` - TUI
- `github.com/charmbracelet/lipgloss` - Styling
- `github.com/mattn/go-sqlite3` - Cache
- `gopkg.in/yaml.v3` - YAML parsing

---

## 📚 Documentation

See the `/Documents` folder for:
- `01_Platform_Tech_Stack.md` - Technical decisions
- `02_Product_Design_Review.md` - Product vision
- `AGENTS.md` - AI development governance

---

## 📜 License

TBD - Choose your license!

---

## 🌟 What's Next?

Strand v1.1 is **feature-complete** and **production-ready**.

**Future possibilities:**
- **Cross-project task management** - Global views, project field, multi-project search
- TUI enhancements (task editing in UI)
- SQLite cache auto-sync
- File watcher for real-time updates
- Multi-agent coordination
- Cloud sync (optional)

**But for now - Strand is DONE! 🎉**

---

**Made with ❤️ by AI + Human collaboration**
