package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all tasks in the current strand project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := store.List()
		if err != nil {
			return fmt.Errorf("failed to list tasks: %w", err)
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			fmt.Println("Create one with: strand create \"Task title\"")
			return nil
		}

		if outputJSON {
			data, _ := json.MarshalIndent(tasks, "", "  ")
			fmt.Println(string(data))
			return nil
		}

		// Table output
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tTYPE\tSTATUS\tPRIORITY\tTITLE")
		fmt.Fprintln(w, "──\t────\t──────\t────────\t─────")

		for _, task := range tasks {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
				task.ID,
				task.Type,
				task.Status,
				task.Priority,
				task.Title,
			)
		}

		w.Flush()
		fmt.Printf("\nTotal: %d tasks\n", len(tasks))

		return nil
	},
}
