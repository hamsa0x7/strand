package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit <id>",
	Short: "Edit a task in your editor",
	Long: `Open a task's markdown file in your default editor.
	
The editor is determined by the $EDITOR environment variable.
Falls back to 'notepad' on Windows, 'nano' on Unix.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		// Get task to verify it exists and get file path
		task, err := store.Get(id)
		if err != nil {
			return err
		}

		// Determine editor
		editor := os.Getenv("EDITOR")
		if editor == "" {
			// Default editors by platform
			if _, err := exec.LookPath("code"); err == nil {
				editor = "code"
			} else if _, err := exec.LookPath("notepad"); err == nil {
				editor = "notepad"
			} else if _, err := exec.LookPath("nano"); err == nil {
				editor = "nano"
			} else {
				editor = "vi" // Last resort
			}
		}

		// Open editor
		fmt.Printf("Opening %s in %s...\n", task.FilePath, editor)

		editorCmd := exec.Command(editor, task.FilePath)
		editorCmd.Stdin = os.Stdin
		editorCmd.Stdout = os.Stdout
		editorCmd.Stderr = os.Stderr

		if err := editorCmd.Run(); err != nil {
			return fmt.Errorf("failed to open editor: %w", err)
		}

		fmt.Println("✅ File saved. Task will be updated on next read.")

		return nil
	},
}
