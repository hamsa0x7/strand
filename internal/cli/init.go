package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new strand project",
	Long: `Initialize a new strand project by creating the .strand directory structure.
	
This creates:
  .strand/tasks/     - Markdown task files
  .strand/.cache/    - SQLite cache (future)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get current directory
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current directory: %w", err)
		}

		strandDir := filepath.Join(cwd, ".strand")

		// Check if already initialized
		if _, err := os.Stat(strandDir); err == nil {
			return fmt.Errorf("strand project already initialized")
		}

		// Create directory structure
		tasksDir := filepath.Join(strandDir, "tasks")
		cacheDir := filepath.Join(strandDir, ".cache")

		if err := os.MkdirAll(tasksDir, 0755); err != nil {
			return fmt.Errorf("failed to create tasks directory: %w", err)
		}

		if err := os.MkdirAll(cacheDir, 0755); err != nil {
			return fmt.Errorf("failed to create cache directory: %w", err)
		}

		// Create .gitignore for cache
		gitignore := filepath.Join(strandDir, ".gitignore")
		gitignoreContent := `.cache/
*.tmp
`
		if err := os.WriteFile(gitignore, []byte(gitignoreContent), 0644); err != nil {
			return fmt.Errorf("failed to create .gitignore: %w", err)
		}

		fmt.Println("✅ Initialized strand project in", strandDir)
		fmt.Println()
		fmt.Println("Next steps:")
		fmt.Println("  strand create \"My first task\"")
		fmt.Println("  strand list")

		return nil
	},
}
