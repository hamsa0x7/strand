package cli

import (
	"fmt"
	"time"

	"github.com/hamsa0x7/strand/internal/core"
	"github.com/spf13/cobra"
)

var (
	updateStatus   string
	updatePriority string
	updateAssignee string
)

var updateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a task",
	Long:  `Update task fields like status, priority, or assignee.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		// Validate inputs
		if updateStatus != "" {
			if err := ValidateStatus(updateStatus); err != nil {
				return err
			}
		}
		if updatePriority != "" {
			if err := ValidatePriority(updatePriority); err != nil {
				return err
			}
		}

		// Get existing task
		task, err := store.Get(id)
		if err != nil {
			return err
		}

		// Apply updates
		if updateStatus != "" {
			task.Status = core.TaskStatus(updateStatus)
		}
		if updatePriority != "" {
			task.Priority = core.TaskPriority(updatePriority)
		}
		if updateAssignee != "" {
			task.Assignee = updateAssignee
		}

		// Update timestamp
		task.Updated = time.Now()

		// Save
		if err := store.Update(task); err != nil {
			return fmt.Errorf("failed to update task: %w", err)
		}

		fmt.Printf("✅ Updated task: %s\n", task.ID)
		fmt.Printf("   Status: %s\n", task.Status)
		fmt.Printf("   Priority: %s\n", task.Priority)
		if task.Assignee != "" {
			fmt.Printf("   Assignee: %s\n", task.Assignee)
		}

		return nil
	},
}

func init() {
	updateCmd.Flags().StringVarP(&updateStatus, "status", "s", "", "Status (backlog|ready|in-progress|done|blocked|cancelled)")
	updateCmd.Flags().StringVarP(&updatePriority, "priority", "p", "", "Priority (critical|high|medium|low)")
	updateCmd.Flags().StringVarP(&updateAssignee, "assignee", "a", "", "Assignee")
}
