package core

import (
	"time"
)

// TaskType represents the type of task
type TaskType string

const (
	TaskTypeTask  TaskType = "task"
	TaskTypeEpic  TaskType = "epic"
	TaskTypeBug   TaskType = "bug"
	TaskTypeStory TaskType = "story"
)

// TaskStatus represents the current status of a task
type TaskStatus string

const (
	TaskStatusBacklog    TaskStatus = "backlog"
	TaskStatusReady      TaskStatus = "ready"
	TaskStatusInProgress TaskStatus = "in-progress"
	TaskStatusDone       TaskStatus = "done"
	TaskStatusBlocked    TaskStatus = "blocked"
	TaskStatusCancelled  TaskStatus = "cancelled"
)

// TaskPriority represents the priority level
type TaskPriority string

const (
	TaskPriorityCritical TaskPriority = "critical"
	TaskPriorityHigh     TaskPriority = "high"
	TaskPriorityMedium   TaskPriority = "medium"
	TaskPriorityLow      TaskPriority = "low"
)

// Task represents a task in the system
type Task struct {
	ID          string       `yaml:"id" json:"id"`
	Type        TaskType     `yaml:"type" json:"type"`
	Status      TaskStatus   `yaml:"status" json:"status"`
	Priority    TaskPriority `yaml:"priority,omitempty" json:"priority,omitempty"`
	Title       string       `yaml:"-" json:"title"` // Extracted from markdown
	Description string       `yaml:"-" json:"description"` // Markdown body
	DependsOn   []string     `yaml:"depends_on,omitempty" json:"depends_on,omitempty"`
	Assignee    string       `yaml:"assignee,omitempty" json:"assignee,omitempty"`
	Created     time.Time    `yaml:"created" json:"created"`
	Updated     time.Time    `yaml:"updated" json:"updated"`
	Tags        []string     `yaml:"tags,omitempty" json:"tags,omitempty"`
	FilePath    string       `yaml:"-" json:"file_path"` // Path to markdown file
}

// NewTask creates a new task with defaults
func NewTask(title string, taskType TaskType) *Task {
	now := time.Now()
	return &Task{
		ID:       generateID(),
		Type:     taskType,
		Status:   TaskStatusBacklog,
		Priority: TaskPriorityMedium,
		Title:    title,
		Created:  now,
		Updated:  now,
		Tags:     []string{},
		DependsOn: []string{},
	}
}

// generateID creates a unique hash-based ID
func generateID() string {
	// Simple implementation for MVP - timestamp-based
	// TODO: Use hash-based IDs for multi-agent safety
	return time.Now().Format("strand-20060102150405")
}

// IsReady returns true if the task has no unresolved dependencies
func (t *Task) IsReady(allTasks map[string]*Task) bool {
	if t.Status == TaskStatusDone || t.Status == TaskStatusCancelled {
		return false
	}
	
	for _, depID := range t.DependsOn {
		depTask, exists := allTasks[depID]
		if !exists {
			continue // Missing dependency, ignore for now
		}
		if depTask.Status != TaskStatusDone {
			return false // Dependency not complete
		}
	}
	
	return true
}
