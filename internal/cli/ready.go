package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var readyCmd = &cobra.Command{
	Use:   "ready",
	Short: "List tasks ready to work on",
	Long:  `List tasks that have no blocking dependencies and are ready to be worked on.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := store.Ready()
		if err != nil {
			return fmt.Errorf("failed to get ready tasks: %w", err)
		}

		if len(tasks) == 0 {
			fmt.Println("No ready tasks found.")
			fmt.Println("All tasks either:")
			fmt.Println("  - Are already done/cancelled")
			fmt.Println("  - Have blocking dependencies")
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
		fmt.Printf("\n✅ %d tasks ready to work on\n", len(tasks))

		return nil
	},
}

func init() {
	readyCmd.Flags().BoolVar(&outputJSON, "json", false, "Output as JSON")
}
