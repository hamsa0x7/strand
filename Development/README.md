# Strand - Task Tracking System

**A git-backed, dependency-aware task tracking system with human-readable Markdown storage**

## Project Status

🚀 **In Development** - Following Agentic AI Product Development Playbook (Beads Edition)

## Overview

Strand combines the best of Beans (human-readable Markdown) and Beads (dependency-aware graph tracking) into a unified system optimized for both humans and AI agents.

## Key Features

- ✅ Human-readable Markdown files with YAML frontmatter
- ✅ Dependency tracking and hierarchical structure
- ✅ SQLite cache for fast queries
- ✅ Beautiful TUI interface
- ✅ Multi-agent safe with hash-based IDs
- ✅ Memory decay and auto-archiving

## Documentation

All specification documents are in the `Development/Documents/` folder:

- **PID** (Product Ideation Document): `../STRAND_PID.md`
- **PDR** (Product Design Review): Coming in Phase 2
- **FSD** (Functional Specification): Coming in Phase 3
- **DSD** (Design Specification): Coming in Phase 4
- **AB** (Architecture Blueprint): Coming in Phase 5

## Development

This project follows the Agentic AI Product Development Playbook. See `AGENTS.md` for AI governance rules during development.

## Technology Stack

- **Language:** Go (single binary, cross-platform)
- **Storage:** Markdown files + SQLite
- **TUI:** Bubble Tea / Lip Gloss
- **CLI:** Cobra
- **File watching:** fsnotify

## License

TBD
