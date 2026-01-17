package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a task",
	Long:  `Delete a task by its ID. The task file will be permanently removed.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		// Get task first to show what we're deleting
		task, err := store.Get(id)
		if err != nil {
			return err
		}

		// Confirm deletion (unless --force flag)
		if !forceDelete {
			fmt.Printf("⚠️  Delete task '%s'? (y/N): ", task.Title)
			var response string
			fmt.Scanln(&response)

			if response != "y" && response != "Y" {
				fmt.Println("Cancelled.")
				return nil
			}
		}

		// Delete
		if err := store.Delete(id); err != nil {
			return fmt.Errorf("failed to delete task: %w", err)
		}

		fmt.Printf("✅ Deleted task: %s\n", id)
		fmt.Printf("   Title: %s\n", task.Title)

		return nil
	},
}

var forceDelete bool

func init() {
	deleteCmd.Flags().BoolVarP(&forceDelete, "force", "f", false, "Skip confirmation prompt")
}
