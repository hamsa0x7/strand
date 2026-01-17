package cli

import (
	"fmt"

	"github.com/hamsa0x7/strand/internal/core"
	"github.com/spf13/cobra"
)

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Visualize task dependency graph",
	Long:  `Display a visual representation of task dependencies.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get all tasks
		tasks, err := store.List()
		if err != nil {
			return fmt.Errorf("failed to list tasks: %w", err)
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return nil
		}

		// Build task map
		taskMap := make(map[string]*core.Task)
		for _, task := range tasks {
			taskMap[task.ID] = task
		}

		// Find root tasks (no dependencies)
		var roots []*core.Task
		for _, task := range tasks {
			if len(task.DependsOn) == 0 {
				roots = append(roots, task)
			}
		}

		// Print graph
		fmt.Println("Task Dependency Graph")
		fmt.Println("=====================")
		fmt.Println()

		if len(roots) == 0 {
			fmt.Println("⚠️  All tasks have dependencies (possible circular dependency)")
			return nil
		}

		// Track visited to avoid duplicates
		visited := make(map[string]bool)

		for _, root := range roots {
			printTaskTree(root, taskMap, visited, "", true)
		}

		// Show orphaned tasks (those that are depended on but we haven't shown)
		orphans := []*core.Task{}
		for _, task := range tasks {
			if !visited[task.ID] {
				orphans = append(orphans, task)
			}
		}

		if len(orphans) > 0 {
			fmt.Println()
			fmt.Println("Tasks with dependencies (shown above):")
			for _, task := range orphans {
				fmt.Printf("  %s - %s [%s]\n", task.ID, task.Title, task.Status)
			}
		}

		return nil
	},
}

func printTaskTree(task *core.Task, taskMap map[string]*core.Task, visited map[string]bool, prefix string, isLast bool) {
	if visited[task.ID] {
		return
	}
	visited[task.ID] = true

	// Print current task
	connector := "├── "
	if isLast {
		connector = "└── "
	}

	statusIcon := getStatusIcon(task.Status)
	priorityColor := getPrioritySymbol(task.Priority)

	fmt.Printf("%s%s%s %s - %s [%s]\n",
		prefix,
		connector,
		statusIcon,
		priorityColor,
		task.Title,
		task.Status,
	)

	// Prepare prefix for children
	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	// Find tasks that depend on this one
	var dependents []*core.Task
	for _, t := range taskMap {
		for _, depID := range t.DependsOn {
			if depID == task.ID {
				dependents = append(dependents, t)
				break
			}
		}
	}

	// Print dependents
	for i, dep := range dependents {
		isLastChild := i == len(dependents)-1
		printTaskTree(dep, taskMap, visited, childPrefix, isLastChild)
	}
}

func getStatusIcon(status core.TaskStatus) string {
	switch status {
	case core.TaskStatusDone:
		return "✅"
	case core.TaskStatusInProgress:
		return "🔄"
	case core.TaskStatusBlocked:
		return "🚫"
	case core.TaskStatusCancelled:
		return "❌"
	case core.TaskStatusReady:
		return "🟢"
	case core.TaskStatusBacklog:
		return "⚪"
	default:
		return "?"
	}
}

func getPrioritySymbol(priority core.TaskPriority) string {
	switch priority {
	case core.TaskPriorityCritical:
		return "🔴"
	case core.TaskPriorityHigh:
		return "🟠"
	case core.TaskPriorityMedium:
		return "🟡"
	case core.TaskPriorityLow:
		return "🟢"
	default:
		return "⚪"
	}
}

// Alternative ASCII-only version (for terminals without emoji support)
var graphAsciiCmd = &cobra.Command{
	Use:   "ascii",
	Short: "Show ASCII-only dependency graph",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get all tasks
		tasks, err := store.List()
		if err != nil {
			return fmt.Errorf("failed to list tasks: %w", err)
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return nil
		}

		// Build task map
		taskMap := make(map[string]*core.Task)
		for _, task := range tasks {
			taskMap[task.ID] = task
		}

		// Find root tasks
		var roots []*core.Task
		for _, task := range tasks {
			if len(task.DependsOn) == 0 {
				roots = append(roots, task)
			}
		}

		fmt.Println("Task Dependency Graph (ASCII)")
		fmt.Println("==============================")
		fmt.Println()

		if len(roots) == 0 {
			fmt.Println("WARNING: All tasks have dependencies")
			return nil
		}

		visited := make(map[string]bool)

		for _, root := range roots {
			printTaskTreeAscii(root, taskMap, visited, "", true)
		}

		return nil
	},
}

func printTaskTreeAscii(task *core.Task, taskMap map[string]*core.Task, visited map[string]bool, prefix string, isLast bool) {
	if visited[task.ID] {
		return
	}
	visited[task.ID] = true

	connector := "+-- "
	if isLast {
		connector = "`-- "
	}

	statusChar := getStatusChar(task.Status)
	priorityChar := getPriorityChar(task.Priority)

	fmt.Printf("%s%s[%s][%s] %s\n",
		prefix,
		connector,
		statusChar,
		priorityChar,
		task.Title,
	)

	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "|   "
	}

	var dependents []*core.Task
	for _, t := range taskMap {
		for _, depID := range t.DependsOn {
			if depID == task.ID {
				dependents = append(dependents, t)
				break
			}
		}
	}

	for i, dep := range dependents {
		isLastChild := i == len(dependents)-1
		printTaskTreeAscii(dep, taskMap, visited, childPrefix, isLastChild)
	}
}

func getStatusChar(status core.TaskStatus) string {
	switch status {
	case core.TaskStatusDone:
		return "X"
	case core.TaskStatusInProgress:
		return ">"
	case core.TaskStatusBlocked:
		return "!"
	case core.TaskStatusCancelled:
		return "-"
	case core.TaskStatusReady:
		return "R"
	case core.TaskStatusBacklog:
		return "."
	default:
		return "?"
	}
}

func getPriorityChar(priority core.TaskPriority) string {
	switch priority {
	case core.TaskPriorityCritical:
		return "!!"
	case core.TaskPriorityHigh:
		return "H"
	case core.TaskPriorityMedium:
		return "M"
	case core.TaskPriorityLow:
		return "L"
	default:
		return "?"
	}
}

func init() {
	graphCmd.AddCommand(graphAsciiCmd)
}
