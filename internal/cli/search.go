package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/hamsa0x7/strand/internal/core"
	"github.com/spf13/cobra"
)

var (
	searchStatus   string
	searchPriority string
	searchTags     []string
)

var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search for tasks",
	Long: `Search for tasks by title, description, or metadata.
	
The query is matched against task titles and descriptions (case-insensitive).
Use flags to filter by status, priority, or tags.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := strings.Join(args, " ")
		queryLower := strings.ToLower(query)

		// Get all tasks
		allTasks, err := store.List()
		if err != nil {
			return fmt.Errorf("failed to list tasks: %w", err)
		}

		// Filter tasks
		var matches []*core.Task
		for _, task := range allTasks {
			// Text search
			titleMatch := strings.Contains(strings.ToLower(task.Title), queryLower)
			descMatch := strings.Contains(strings.ToLower(task.Description), queryLower)

			if !titleMatch && !descMatch {
				continue
			}

			// Status filter
			if searchStatus != "" && string(task.Status) != searchStatus {
				continue
			}

			// Priority filter
			if searchPriority != "" && string(task.Priority) != searchPriority {
				continue
			}

			// Tags filter
			if len(searchTags) > 0 {
				hasAllTags := true
				for _, wantTag := range searchTags {
					found := false
					for _, tag := range task.Tags {
						if tag == wantTag {
							found = true
							break
						}
					}
					if !found {
						hasAllTags = false
						break
					}
				}
				if !hasAllTags {
					continue
				}
			}

			matches = append(matches, task)
		}

		// Output results
		if len(matches) == 0 {
			fmt.Printf("No tasks found matching '%s'\n", query)
			return nil
		}

		if outputJSON {
			data, _ := json.MarshalIndent(matches, "", "  ")
			fmt.Println(string(data))
			return nil
		}

		// Table output
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tTYPE\tSTATUS\tPRIORITY\tTITLE")
		fmt.Fprintln(w, "──\t────\t──────\t────────\t─────")

		for _, task := range matches {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
				task.ID,
				task.Type,
				task.Status,
				task.Priority,
				task.Title,
			)
		}

		w.Flush()
		fmt.Printf("\nFound %d task(s) matching '%s'\n", len(matches), query)

		return nil
	},
}

func init() {
	searchCmd.Flags().StringVar(&searchStatus, "status", "", "Filter by status")
	searchCmd.Flags().StringVar(&searchPriority, "priority", "", "Filter by priority")
	searchCmd.Flags().StringSliceVar(&searchTags, "tags", []string{}, "Filter by tags (comma-separated)")
	searchCmd.Flags().BoolVar(&outputJSON, "json", false, "Output as JSON")
}
