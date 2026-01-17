package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var depCmd = &cobra.Command{
	Use:   "dep",
	Short: "Manage task dependencies",
	Long:  `Add or remove dependencies between tasks.`,
}

var depAddCmd = &cobra.Command{
	Use:   "add <task-id> <depends-on-id>",
	Short: "Add a dependency",
	Long:  `Make a task depend on another task. The task won't be ready until its dependency is done.`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskID := args[0]
		dependsOnID := args[1]

		// Get the task
		task, err := store.Get(taskID)
		if err != nil {
			return err
		}

		// Verify dependency exists
		depTask, err := store.Get(dependsOnID)
		if err != nil {
			return fmt.Errorf("dependency task not found: %s", dependsOnID)
		}

		// Check if already depends on it
		for _, dep := range task.DependsOn {
			if dep == dependsOnID {
				return fmt.Errorf("task already depends on %s", dependsOnID)
			}
		}

		// Check for circular dependency (basic check)
		if dependsOnID == taskID {
			return fmt.Errorf("cannot create circular dependency: task cannot depend on itself")
		}

		// TODO: Deep cycle detection

		// Add dependency
		task.DependsOn = append(task.DependsOn, dependsOnID)
		task.Updated = time.Now()

		if err := store.Update(task); err != nil {
			return fmt.Errorf("failed to update task: %w", err)
		}

		fmt.Printf("✅ Added dependency\n")
		fmt.Printf("   Task: %s (%s)\n", task.ID, task.Title)
		fmt.Printf("   Depends on: %s (%s)\n", depTask.ID, depTask.Title)

		return nil
	},
}

var depRemoveCmd = &cobra.Command{
	Use:   "remove <task-id> <depends-on-id>",
	Short: "Remove a dependency",
	Long:  `Remove a dependency between tasks.`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskID := args[0]
		dependsOnID := args[1]

		// Get the task
		task, err := store.Get(taskID)
		if err != nil {
			return err
		}

		// Find and remove dependency
		found := false
		newDeps := []string{}
		for _, dep := range task.DependsOn {
			if dep != dependsOnID {
				newDeps = append(newDeps, dep)
			} else {
				found = true
			}
		}

		if !found {
			return fmt.Errorf("task does not depend on %s", dependsOnID)
		}

		task.DependsOn = newDeps
		task.Updated = time.Now()

		if err := store.Update(task); err != nil {
			return fmt.Errorf("failed to update task: %w", err)
		}

		fmt.Printf("✅ Removed dependency\n")
		fmt.Printf("   Task: %s (%s)\n", task.ID, task.Title)
		fmt.Printf("   No longer depends on: %s\n", dependsOnID)

		return nil
	},
}

var depListCmd = &cobra.Command{
	Use:   "list <task-id>",
	Short: "List task dependencies",
	Long:  `Show all tasks that a task depends on.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskID := args[0]

		// Get the task
		task, err := store.Get(taskID)
		if err != nil {
			return err
		}

		fmt.Printf("Task: %s (%s)\n", task.ID, task.Title)

		if len(task.DependsOn) == 0 {
			fmt.Println("No dependencies")
			return nil
		}

		fmt.Printf("\nDepends on %d task(s):\n\n", len(task.DependsOn))

		for _, depID := range task.DependsOn {
			depTask, err := store.Get(depID)
			if err != nil {
				fmt.Printf("  - %s (not found)\n", depID)
				continue
			}
			fmt.Printf("  - %s (%s) [%s]\n", depTask.ID, depTask.Title, depTask.Status)
		}

		return nil
	},
}

func init() {
	depCmd.AddCommand(depAddCmd)
	depCmd.AddCommand(depRemoveCmd)
	depCmd.AddCommand(depListCmd)
}
