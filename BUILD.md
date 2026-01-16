# Strand MVP - Build Instructions

## Prerequisites

### Install Go  1.21+

**Method 1: Manual Download (Recommended)**
1. Download from: https://go.dev/dl/
2. Run the installer as Administrator
3. Verify installation: `go version`

**Method 2: Chocolatey (Admin Terminal Required)**
```powershell
# Open PowerShell as Administrator
choco install golang -y
refreshenv
go version
```

## Building Strand

Once Go is installed:

```bash
# Navigate to Strand directory
cd e:\Data\Users\JoJo\Desktop\Neura\Strand

# Download dependencies
go mod download

# Build the binary
go build -o bin/strand.exe ./cmd/strand

# Or install globally
go install ./cmd/strand
```

## Quick Start

```bash
# Navigate to a project directory
cd my-project

# Initialize strand
.\bin\strand.exe init

# Create your first task
.\bin\strand.exe create "Implement user authentication" --type task --priority high

# List tasks
.\bin\strand.exe list

# Show task details
.\bin\strand.exe show strand-20260116123456

# Update task status
.\bin\strand.exe update strand-20260116123456 --status in-progress

# See what's ready to work on
.\bin\strand.exe ready
```

## Development

```bash
# Run tests
go test ./...

# Run locally without building
go run ./cmd/strand <command>

# Example:
go run ./cmd/strand init
go run ./cmd/strand create "Test task"
go run ./cmd/strand list
```

## Project Structure

```
Strand/
├── cmd/strand/          # Main CLI entry point
├── internal/
│   ├── core/            # Task domain model
│   ├── storage/         # Storage interface
│   ├── markdown/        # Markdown file implementation
│   └── cli/             # CLI commands
├── AGENTS.md            # AI development governance
├── Documents/           # Specifications
├── go.mod               # Go dependencies
└── README.md
```

## Task File Example

After running `strand create`, you'll find Markdown files in `.strand/tasks/`:

```markdown
---
id: strand-20260116123456
type: task
status: backlog
priority: medium
created: 2026-01-16T12:34:56Z
updated: 2026-01-16T12:34:56Z
tags: []
depends_on: []
---

# Implement user authentication

```

You can edit these files directly in VS Code or any text editor!

## Next Steps (Post-MVP)

After testing the MVP:
1. Add SQLite cache for performance
2. Implement file watcher for auto-sync
3. Build TUI with Bubble Tea
4. Add dependency graph visualization
5. Implement GraphQL queries

## Troubleshooting

**`go: command not found`**
- Ensure Go is installed and in your PATH
- Try closing and reopening your terminal
- Run `refreshenv` (if using Chocolatey)

**Build errors**
- Run `go mod tidy` to clean up dependencies
- Ensure you're in the project root directory

**Permission errors**
- Run terminal/PowerShell as Administrator for first-time setup
