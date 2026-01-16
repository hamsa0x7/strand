# Phase 2: Product Design Review (PDR)

**Document Type:** Planning  
**Phase:** 2 of 9 (Planning)  
**Status:** Complete  
**Date:** 2026-01-16  
**Authority:** Derived from STRAND_PID.md

---

## 1. Product Vision

### 1.1 The Problem

**Current State of Task Tracking for AI Agents:**

Developers and AI agents face a **false choice** between two task tracking paradigms:

**Option A: Beans** (Human-Friendly)
- ✅ Human-readable Markdown files
- ✅ Beautiful TUI interface
- ❌ No dependency tracking
- ❌ No hierarchical structure
- ❌ Performance issues with large task sets

**Option B: Beads** (Agent-Optimized)
- ✅ Dependency-aware graph database
- ✅ Hierarchical task structure
- ✅ Fast queries via JSONL storage
- ❌ JSONL is **not human-readable**
- ❌ No TUI interface
- ❌ Requires specialized knowledge to edit manually

**The Gap:** No tool combines **human readability** with **agent optimization**.

### 1.2 The Strand Solution

**Core Insight:** You don't have to choose. Store as Markdown, cache as SQLite.

Strand is a **dual-format task tracker** that:
- Stores tasks as **beautiful, editable Markdown files** (source of truth)
- Maintains a **SQLite cache** for lightning-fast queries (auto-generated)
- Syncs bi-directionally with **zero manual intervention**
- Provides **both TUI and CLI** interfaces
- Tracks **dependencies and hierarchy** via YAML frontmatter

**Result:** Humans edit Markdown in their favorite editor. Agents query a high-performance graph database. Everyone wins.

---

## 2. Product Philosophy

### 2.1 Core Principles

**1. Markdown is the Constitution**
- Markdown files are the **only** source of truth
- SQLite is **derived** and **regeneratable**
- If cache and Markdown conflict, Markdown wins

**2. Zero Lock-In**
- Tasks are portable plain text files
- No proprietary formats
- No cloud dependencies
- Git-friendly by design

**3. Agent-First, Human-Friendly**
- Optimize for AI agent workflows (JSON output, dependency queries)
- But never sacrifice human readability
- Both personas are first-class citizens

**4. Local-First Philosophy**
- All data stays on disk
- No authentication required
- No network requests (except Git)
- Works offline, forever

**5. Git is the Timeline**
- Tasks are tracked in Git
- Branching enables parallel work
- History is preserved
- Collaboration via Git workflows

---

## 3. Target Users

### 3.1 Primary Personas

**Persona 1: The AI-Assisted Developer**
- Uses agentic coding assistants (Claude, GPT, Gemini)
- Needs persistent task memory across sessions
- Wants agents to autonomously track work
- Values human oversight and control

**Example:** Sarah uses Claude to build a React app. Claude creates tasks in Strand as it discovers work. Sarah reviews tasks in TUI, edits them in VS Code, and approves agent execution.

**Persona 2: The Solo Builder**
- Works on side projects or startups
- Juggles multiple concurrent tasks
- Needs dependency tracking to avoid blockers
- Wants beautiful, inspectable task lists

**Example:** Alex builds a SaaS product across frontend, backend, and infrastructure. Strand's dependency graph ensures auth is built before dashboard, API before frontend.

**Persona 3: The Playbook-Driven Team**
- Uses structured development playbooks (like Agentic AI Playbook)
- Needs hierarchical task breakdown (epics → features → tasks)
- Wants progress tracking across phases
- Values Markdown for documentation integration

**Example:** A team builds products using the Agentic AI Playbook. Strand tracks all 9 phases, auto-generates progress reports, and integrates with AGENTS.md governance.

### 3.2 Explicit Non-Users

**Who Strand is NOT for:**
- ❌ Large enterprise teams (use Jira, Linear)
- ❌ Non-technical users (requires CLI comfort)
- ❌ Cloud-first workflows (Strand is local-first)
- ❌ Real-time collaboration (use Notion, ClickUp)

---

## 4. Core Features

### 4.1 MVP Features (Phase 9.1)

**Must-Have for Initial Release:**

1. **Markdown Task Storage**
   - YAML frontmatter for metadata
   - Markdown body for description, notes, acceptance criteria
   - Hierarchical file structure (`epic-001.md`, `task-001-1.md`)

2. **SQLite Cache Layer**
   - Auto-generated from Markdown files
   - Stores task graph, dependencies, metadata
   - Rebuilt on file changes

3. **File Watcher & Auto-Sync**
   - Monitors `.strand/tasks/` directory
   - Updates SQLite on Markdown changes
   - Debounced for performance

4. **Basic CLI Commands**
   - `strand init` - Initialize .strand/ directory
   - `strand create "title"` - Create task with template
   - `strand list` - List all tasks
   - `strand show <id>` - Show task details
   - `strand update <id> --status=<status>` - Update task
   - `strand ready` - Show tasks with no blockers

5. **Task States**
   - `backlog`, `ready`, `in-progress`, `done`, `blocked`, `cancelled`

6. **Priority Levels**
   - `critical`, `high`, `medium`, `low`

### 4.2 Post-MVP Features (Phase 9.3)

**Deferred to Full Development:**

7. **Dependency Graph**
   - Add dependencies via `depends_on: [strand-xxx]` in frontmatter
   - Cycle detection and validation
   - Visual graph in TUI

8. **Beautiful TUI**
   - Interactive task browser (Bubble Tea)
   - Keyboard navigation
   - Syntax-highlighted Markdown preview
   - Dependency graph visualization

9. **GraphQL Query Interface**
   - Query tasks via GraphQL (optional)
   - Enables advanced agent workflows
   - Integrates with other tools

10. **Memory Decay / Auto-Archive**
    - Archive completed tasks after N days
    - Keep recent history, compress old tasks
    - Configurable retention policy

11. **Templates System**
    - Custom task templates
    - Project-specific templates
    - Template inheritance

12. **Multi-Agent Support**
    - Hash-based IDs for conflict-free merges
    - Lock-free concurrent writes
    - Git-based synchronization

### 4.3 Explicit Non-Features

**Will NEVER be supported:**

- ❌ Cloud sync / SaaS offering
- ❌ Real-time collaboration UI
- ❌ Mobile apps
- ❌ Calendar/time tracking integration
- ❌ Built-in chat/comments
- ❌ User authentication/permissions
- ❌ Kanban boards with drag-and-drop
- ❌ Issue importing from Jira/GitHub

**Rationale:** These conflict with Strand's philosophy of simplicity, local-first design, and Markdown ownership.

---

## 5. User Experience Principles

### 5.1 CLI Design

**Conventions:**
- Follow POSIX standards (e.g., `--help`, `-h`)
- Predictable command structure: `strand <verb> <noun> [flags]`
- JSON output for agent consumption: `--json` flag on all commands
- Human-friendly default output with colors and tables

**Examples:**
```bash
# Human-friendly
strand list
# → Table with colors

# Agent-friendly
strand list --json
# → JSON array for parsing
```

### 5.2 TUI Design

**Aesthetics:**
- Clean, minimal interface (inspired by lazygit, k9s)
- Vim-style keybindings (`j`/`k` for navigation, `/` for search)
- Syntax-highlighted Markdown preview
- Responsive to terminal resizing

**Navigation:**
- Tree view of task hierarchy
- Quick filters (by status, priority, tags)
- Dependency graph visualization

### 5.3 File Format Design

**Readability First:**
- YAML frontmatter is minimal (only metadata)
- Markdown body is free-form (no rigid structure)
- Acceptance criteria use checkboxes (`- [ ]`)
- Notes and descriptions are plain text

**Example:**
```markdown
---
id: strand-a3f8
type: epic
status: in-progress
priority: high
created: 2026-01-16T20:00:00Z
tags: [mvp, core]
---

# Implement Core Storage Layer

## Description
Build the Markdown file storage system with YAML frontmatter parsing.

## Acceptance Criteria
- [ ] Parse YAML frontmatter correctly
- [ ] Generate hierarchical file names
- [ ] Validate task IDs are unique

## Notes
- Use goldmark for Markdown parsing
- Use go-yaml for frontmatter
```

---

## 6. Success Metrics

### 6.1 MVP Success Criteria

**Functional:**
- ✅ Can create, list, show, update tasks via CLI
- ✅ Markdown files are git-committable and readable
- ✅ SQLite cache updates automatically on file changes
- ✅ `strand ready` returns unblocked tasks correctly

**Performance:**
- ✅ File watcher responds within 100ms
- ✅ CLI commands complete in < 200ms for 1000 tasks
- ✅ Cache rebuild completes in < 1s for 100 tasks

**Usability:**
- ✅ Non-technical user can read task files in VS Code
- ✅ Agent can parse JSON output without errors
- ✅ Tasks survive Git merges without conflicts (hash IDs)

### 6.2 Post-MVP Success Criteria

**Adoption:**
- 100+ GitHub stars within 3 months
- 10+ community contributions
- Used in 5+ real projects

**Quality:**
- TUI rated "beautiful" in user feedback
- No data loss bugs reported
- Dependency graph handles 1000+ tasks

---

## 7. Product Boundaries

### 7.1 What Strand IS

✅ A task tracker for developers and AI agents  
✅ A Markdown-first, Git-friendly system  
✅ A CLI/TUI tool for local workflows  
✅ A dependency-aware task graph database  

### 7.2 What Strand is NOT

❌ A project management SaaS  
❌ A Jira replacement for large teams  
❌ A time tracking or calendar tool  
❌ A real-time collaboration platform  

---

## 8. Design Tradeoffs

### 8.1 Simplicity vs Power

**Decision:** Favor simplicity in MVP, add power in phases

**Examples:**
- MVP: Basic dependency tracking (parent/child only)
- Post-MVP: Complex dependency graphs (blocks/blocked-by)

**Rationale:** Ship fast, validate user need, iterate

### 8.2 Human Readability vs Agent Performance

**Decision:** Optimize for both via dual-format strategy

**Implementation:**
- Markdown for humans (slow but readable)
- SQLite for agents (fast but opaque)
- File watcher bridges the gap

**Rationale:** This is Strand's core value proposition—no compromise needed

### 8.3 Local-First vs Collaboration

**Decision:** Local-first wins, collaboration via Git

**Implementation:**
- No built-in sync
- Git is the collaboration layer
- Hash-based IDs prevent merge conflicts

**Rationale:** Aligns with developer workflows, avoids complexity

---

## 9. Risks and Mitigation

### 9.1 Adoption Risks

**Risk:** Developers prefer existing tools (Beans, Beads, Taskwarrior)

**Mitigation:**
- Clear differentiation: "Best of both worlds"
- Migration guides from Beans/Beads
- Beautiful TUI to wow users
- Agent integration examples

### 9.2 Technical Risks

**Risk:** SQLite sync performance degrades with large task sets

**Mitigation:**
- Batch file watcher updates
- Debounce rapid changes
- Benchmark with 10,000 tasks
- Optimize indexes

**Risk:** File format conflicts during Git merges

**Mitigation:**
- Hash-based IDs (collision-resistant)
- Conflict resolution guide
- Automatic ID regeneration on collisions

### 9.3 Scope Creep Risks

**Risk:** Feature requests push toward bloated "second Jira"

**Mitigation:**
- Strict adherence to PID philosophy
- Public "non-features" list
- Governance via AGENTS.md
- Regular PID review

---

## 10. Competitive Landscape

### 10.1 Existing Tools

| Tool | Human-Readable | Dependencies | TUI | Agent-Friendly | Local-First |
|------|----------------|--------------|-----|----------------|-------------|
| **Beans** | ✅ Markdown | ❌ | ✅ | ⚠️ Limited | ✅ |
| **Beads** | ❌ JSONL | ✅ | ❌ | ✅ | ✅ |
| **Taskwarrior** | ❌ Custom | ✅ | ⚠️ Basic | ⚠️ Limited | ✅ |
| **Org-mode** | ⚠️ Org syntax | ⚠️ Manual | ✅ (Emacs) | ❌ | ✅ |
| **Linear/Jira** | ❌ Web UI | ✅ | ❌ | ✅ API | ❌ Cloud |
| **Strand** | ✅ Markdown | ✅ | ✅ | ✅ JSON | ✅ |

**Strand's Advantage:** Only tool with ALL checkmarks.

### 10.2 Differentiation Strategy

**Key Messages:**
1. "Markdown you can read, SQLite agents can query—no compromise"
2. "Git-native task tracking for developer workflows"
3. "Beautiful TUI meets powerful CLI"
4. "Zero setup, zero cloud, zero lock-in"

---

## 11. MVP Scope Definition

### 11.1 In-Scope for MVP

**Core Loop:**
1. Developer/agent creates task → Markdown file written
2. File watcher detects change → SQLite cache updated
3. Developer queries via CLI → SQLite returns results
4. Developer edits Markdown directly → Cache auto-syncs

**Commands:**
- `init`, `create`, `list`, `show`, `update`, `ready`

**File Format:**
- Markdown with YAML frontmatter
- Basic fields: `id`, `type`, `status`, `priority`, `created`, `updated`

**Storage:**
- `.strand/tasks/` directory
- `.strand/.cache/graph.db` SQLite file

### 11.2 Out-of-Scope for MVP

- ❌ TUI interface (CLI only)
- ❌ Dependency graph (basic hierarchy only)
- ❌ GraphQL queries
- ❌ Templates
- ❌ Auto-archiving
- ❌ Multi-agent locking

**Rationale:** Validate core value proposition first, add features based on feedback.

---

## 12. Development Roadmap

### Phase 9.1: MVP (2-3 weeks)
- Core Markdown storage
- SQLite cache layer
- File watcher
- Basic CLI commands
- Unit + integration tests

### Phase 9.5: MVP Review & Refinement (1 week)
- User testing
- Bug fixes
- Performance tuning
- Documentation

### Phase 9.3: Full Product (4-6 weeks)
- Dependency graph with cycle detection
- Beautiful TUI (Bubble Tea)
- GraphQL query interface
- Templates system
- Memory decay/archiving
- Polish and optimization

---

## 13. Compliance with PID

### 13.1 Alignment Check

✅ **Product Identity** (PID Section 1)
- Matches definition: "git-backed, dependency-aware task tracking system"

✅ **Problem Statement** (PID Section 2)
- Addresses the Beans vs Beads dilemma

✅ **Core Principles** (PID Section 3)
- Markdown-First Principle: ✅
- Dual-Format Strategy: ✅

✅ **Feature Set** (PID Section 4)
- Best of Beans: ✅
- Best of Beads: ✅
- Enhanced Features: ✅

✅ **Technology Stack** (PID Section 10)
- Go, SQLite, Markdown, Bubble Tea: ✅

---

## 14. Next Steps

**Phase 3: Functional Specification Document (FSD)**
- Define exact user workflows
- Specify CLI command syntax
- Document file format schema
- Define error handling

**Commit Message:**
```
Phase 2: Product Design Review complete

- Defined product vision and philosophy
- Established target user personas
- Documented MVP vs full feature set
- Defined success metrics and risks
- Scoped MVP development clearly
```

---

**This PDR is approved and ready for Phase 3: Functional Specification.**
