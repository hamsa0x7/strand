# Product Ideation Document: Strand

**Status:** Conceptual  
**Authority Level:** Constitutional  
**Version:** 0.1

---

## 1. Product Identity

### 1.1 Product Name
**Strand** - A unified task tracking system for coding agents and humans

### 1.2 Product Definition (What It IS)
Strand is a git-backed, dependency-aware task tracking system that stores tasks as **human-readable Markdown with embedded metadata**, combining the simplicity of Beans with the power of Beads.

### 1.3 What It is NOT
- Not a separate application or web service
- Not a replacement for project management tools for large teams
- Not a cloud-based SaaS solution

---

## 2. Problem Statement

**Problem:** Existing agent task trackers force a choice between:
- **Beans:** Simple, readable, but lacks dependencies and hierarchy
- **Beads:** Powerful graphs, but JSONL is not human-friendly

**Users need:** A system that is BOTH human-readable AND feature-complete for agents.

---

## 3. Core Principles

### The Markdown-First Principle
**ALL task data MUST be readable and editable as plain Markdown by humans.**

### The Dual-Format Strategy
```
.strand/
├── tasks/
│   ├── epic-001.md          ← Human-readable source of truth
│   ├── task-001-1.md        ← Hierarchical markdown files
│   └── task-001-1-1.md
└── .cache/
    └── graph.db             ← SQLite cache (auto-generated)
```

**Key insight:** Store as Markdown, cache as SQLite.  
- Humans read/edit Markdown
- Agents query SQLite for speed
- Markdown is source of truth, SQLite regenerated automatically

---

## 4. Feature Set

### Best of Beans ✅
- ✅ Human-readable Markdown files
- ✅ Built-in TUI for browsing
- ✅ GraphQL query engine (via SQLite)
- ✅ Project memory (archived tasks)
- ✅ Auto-generates roadmap

### Best of Beads ✅
- ✅ Dependency tracking (stored in YAML frontmatter)
- ✅ Hierarchical structure (filename-based)
- ✅ Multi-agent safe (hash-based IDs)
- ✅ Performance (SQLite cache)
- ✅ Memory decay (auto-archive)
- ✅ Stealth mode

### Enhanced Features (Neither has) ✨
- ✨ **Bi-directional sync:** Edit Markdown → auto-update cache
- ✨ **Visual dependency graph:** TUI shows task graph
- ✨ **Markdown templates:** Customizable task templates
- ✨ **Smart auto-complete:** Agent-friendly CLI
- ✨ **Integration hooks:** Git hooks, CI/CD integration

---

## 5. Technical Architecture

### File Format: Markdown with YAML Frontmatter
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

# Implement Dashboard UI

## Description
Create the main dashboard interface with widget system.

## Acceptance Criteria
- [ ] Grid layout responsive
- [ ] Widgets draggable
- [ ] State persists

## Notes
- Use shadcn/ui components
- Follow DSD mockup design
```

### Cache Structure: SQLite
```sql
CREATE TABLE tasks (
    id TEXT PRIMARY KEY,
    parent_id TEXT,
    title TEXT,
    status TEXT,
    priority INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    markdown_path TEXT,
    content_hash TEXT
);

CREATE TABLE dependencies (
    child_id TEXT,
    parent_id TEXT,
    PRIMARY KEY (child_id, parent_id)
);
```

### Sync Mechanism
- **File watcher** detects Markdown changes
- **Background daemon** updates SQLite cache
- **Hash-based** change detection (only update if changed)

---

## 6. User Experience

### For Humans
```bash
# Browse with beautiful TUI
strand tui

# Create task (opens editor with template)
strand new "Fix login bug" --type=bug --priority=high

# View task
strand show strand-a3f8.1

# Edit directly in vim/vscode
vim .strand/tasks/strand-a3f8.1.md
# Changes auto-sync to cache
```

### For Agents
```bash
# Query ready tasks (uses SQLite for speed)
strand ready --json

# Create with dependencies
strand create "Build API" --depends-on=strand-a3f8 --json

# Update status
strand update strand-a3f8.1 --status=done

# Query dependency graph
strand deps strand-a3f8 --graph
```

---

## 7. Integration with AGENTS.md

```markdown
# AGENTS.md

## Task Tracking

Use Strand for all task management:

```bash
# Before starting work
strand ready

# Create tasks as you discover work
strand create "Fix performance issue"

# Update progress
strand update <id> --status=in-progress

# Mark complete
strand done <id>
```

Strand files are in `.strand/tasks/` - you can read them as Markdown.
```

---

## 8. Advantages Over Beans/Beads

| Feature | Beans | Beads | **Strand** |
|---------|-------|-------|------------|
| Human-readable | ✅ Markdown | ❌ JSONL | ✅ **Markdown** |
| Dependency graph | ❌ | ✅ | ✅ **In frontmatter** |
| Hierarchical | ❌ | ✅ | ✅ **File-based** |
| Fast queries | ❌ | ✅ | ✅ **SQLite cache** |
| TUI | ✅ | ❌ | ✅ **Enhanced** |
| Multi-agent safe | ⚠️ | ✅ | ✅ **Hash IDs** |
| Memory decay | ❌ | ✅ | ✅ **Auto-archive** |
| GraphQL | ✅ | ❌ | ✅ **Via SQLite** |
| Edit in editor | ✅ | ❌ | ✅ **Full support** |

---

## 9. Implementation Phases

### Phase 1: Core (MVP)
- Markdown storage with YAML frontmatter
- SQLite cache layer
- Basic CLI (create, list, show, update)
- File watcher + auto-sync

### Phase 2: Dependencies
- Dependency graph in frontmatter
- Dependency validation
- `strand ready` (returns tasks with no blockers)
- Visual graph in TUI

### Phase 3: Polish
- Beautiful TUI
- GraphQL query interface
- Templates system
- Memory decay/auto-archive

### Phase 4: Integration
- AGENTS.md templates
- Git hooks
- CI/CD examples
- IDE extensions

---

## 10. Technology Stack

**Language:** Go (single binary, cross-platform)  
**Storage:** Markdown files + SQLite  
**TUI:** Bubble Tea / Lip Gloss  
**CLI:** Cobra  
**File watching:** fsnotify  
**Sync:** Background daemon (optional)

---

## 11. Next Steps

If approved:
1. Create detailed FSD (Functional Specification)
2. Design DSD (TUI mockups)
3. Define architecture (Module boundaries)
4. Implement MVP
5. Test with playbook-generated projects

---

**This combines the best of both worlds with NO cons!**
- ✅ Human-readable (Markdown source)
- ✅ Agent-optimized (SQLite cache)
- ✅ Dependency tracking (YAML frontmatter)
- ✅ Performance (Background sync)
- ✅ Simple setup (Single binary)
- ✅ Beautiful TUI (Enhanced from Beans)

**Shall I proceed with full specification using the playbook?**
