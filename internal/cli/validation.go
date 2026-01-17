package cli

import (
	"fmt"

	"github.com/hamsa0x7/strand/internal/core"
)

// ValidateStatus validates a status string
func ValidateStatus(status string) error {
	validStatuses := []core.TaskStatus{
		core.TaskStatusBacklog,
		core.TaskStatusReady,
		core.TaskStatusInProgress,
		core.TaskStatusDone,
		core.TaskStatusBlocked,
		core.TaskStatusCancelled,
	}

	for _, valid := range validStatuses {
		if status == string(valid) {
			return nil
		}
	}

	return fmt.Errorf("invalid status '%s'. Valid: backlog, ready, in-progress, done, blocked, cancelled", status)
}

// ValidatePriority validates a priority string
func ValidatePriority(priority string) error {
	validPriorities := []core.TaskPriority{
		core.TaskPriorityCritical,
		core.TaskPriorityHigh,
		core.TaskPriorityMedium,
		core.TaskPriorityLow,
	}

	for _, valid := range validPriorities {
		if priority == string(valid) {
			return nil
		}
	}

	return fmt.Errorf("invalid priority '%s'. Valid: critical, high, medium, low", priority)
}

// ValidateType validates a task type string
func ValidateType(taskType string) error {
	validTypes := []core.TaskType{
		core.TaskTypeTask,
		core.TaskTypeEpic,
		core.TaskTypeBug,
		core.TaskTypeStory,
	}

	for _, valid := range validTypes {
		if taskType == string(valid) {
			return nil
		}
	}

	return fmt.Errorf("invalid type '%s'. Valid: task, epic, bug, story", taskType)
}
