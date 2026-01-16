package cli

import (
	"encoding/json"
	"fmt"

	"github.com/hamsa0x7/strand/internal/core"
	"github.com/spf13/cobra"
)

var (
	createType     string
	createPriority string
	createTags     []string
	createAssignee string
	outputJSON     bool
)

var createCmd = &cobra.Command{
	Use:   "create <title>",
	Short: "Create a new task",
	Long:  `Create a new task with the given title.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		title := args[0]

		// Create task
		taskType := core.TaskType(createType)
		task := core.NewTask(title, taskType)

		// Apply optional fields
		if createPriority != "" {
			task.Priority = core.TaskPriority(createPriority)
		}
		if createAssignee != "" {
			task.Assignee = createAssignee
		}
		if len(createTags) > 0 {
			task.Tags = createTags
		}

		// Save to storage
		if err := store.Create(task); err != nil {
			return fmt.Errorf("failed to create task: %w", err)
		}

		// Output
		if outputJSON {
			data, _ := json.MarshalIndent(task, "", "  ")
			fmt.Println(string(data))
		} else {
			fmt.Printf("✅ Created task: %s\n", task.ID)
			fmt.Printf("   Title: %s\n", task.Title)
			fmt.Printf("   Type: %s\n", task.Type)
			fmt.Printf("   Status: %s\n", task.Status)
			fmt.Printf("   File: .strand/tasks/%s.md\n", task.ID)
		}

		return nil
	},
}

func init() {
	createCmd.Flags().StringVarP(&createType, "type", "t", "task", "Task type (task|epic|bug|story)")
	createCmd.Flags().StringVarP(&createPriority, "priority", "p", "medium", "Priority (critical|high|medium|low)")
	createCmd.Flags().StringSliceVar(&createTags, "tags", []string{}, "Comma-separated tags")
	createCmd.Flags().StringVarP(&createAssignee, "assignee", "a", "", "Assignee")
	createCmd.Flags().BoolVar(&outputJSON, "json", false, "Output as JSON")
}
