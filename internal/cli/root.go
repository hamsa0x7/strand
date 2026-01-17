package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hamsa0x7/strand/internal/markdown"
	"github.com/hamsa0x7/strand/internal/storage"
	"github.com/spf13/cobra"
)

var (
	store     storage.Store
	strandDir string
)

var rootCmd = &cobra.Command{
	Use:   "strand",
	Short: "Strand - Task tracking for humans and AI agents",
	Long: `Strand is a git-backed, dependency-aware task tracking system 
that stores tasks as human-readable Markdown with embedded metadata.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Initialize storage for all commands except init
		if cmd.Name() == "init" {
			return nil
		}

		// Find .strand directory
		dir, err := findStrandDir()
		if err != nil {
			return fmt.Errorf("not in a strand project (run 'strand init'): %w", err)
		}
		strandDir = dir

		// Initialize store
		store = markdown.NewStore()
		if err := store.Init(strandDir); err != nil {
			return fmt.Errorf("failed to initialize storage: %w", err)
		}

		return nil
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(readyCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(depCmd)
	rootCmd.AddCommand(editCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(graphCmd)
}

// findStrandDir searches for .strand directory in current or parent directories
func findStrandDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		strandPath := filepath.Join(dir, ".strand")
		if info, err := os.Stat(strandPath); err == nil && info.IsDir() {
			return strandPath, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf(".strand directory not found")
		}
		dir = parent
	}
}
