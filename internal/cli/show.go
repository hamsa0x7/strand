package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show <id>",
	Short: "Show details of a task",
	Long:  `Display full details of a task including its description.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		task, err := store.Get(id)
		if err != nil {
			return err
		}

		if outputJSON {
			data, _ := json.MarshalIndent(task, "", "  ")
			fmt.Println(string(data))
			return nil
		}

		// Pretty output
		fmt.Printf("ID:          %s\n", task.ID)
		fmt.Printf("Type:        %s\n", task.Type)
		fmt.Printf("Status:      %s\n", task.Status)
		fmt.Printf("Priority:    %s\n", task.Priority)
		fmt.Printf("Title:       %s\n", task.Title)

		if task.Assignee != "" {
			fmt.Printf("Assignee:    %s\n", task.Assignee)
		}

		if len(task.Tags) > 0 {
			fmt.Printf("Tags:        %v\n", task.Tags)
		}

		if len(task.DependsOn) > 0 {
			fmt.Printf("Depends On:  %v\n", task.DependsOn)
		}

		fmt.Printf("Created:     %s\n", task.Created.Format("2006-01-02 15:04:05"))
		fmt.Printf("Updated:     %s\n", task.Updated.Format("2006-01-02 15:04:05"))
		fmt.Printf("File:        %s\n", task.FilePath)

		if task.Description != "" {
			fmt.Println("\nDescription:")
			fmt.Println(task.Description)
		}

		return nil
	},
}
