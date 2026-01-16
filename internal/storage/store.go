package storage

import (
	"github.com/hamsa0x7/strand/internal/core"
)

// Store defines the interface for task storage operations
type Store interface {
	// Init initializes the storage system
	Init(strandDir string) error
	
	// Create creates a new task
	Create(task *core.Task) error
	
	// Get retrieves a task by ID
	Get(id string) (*core.Task, error)
	
	// List retrieves all tasks
	List() ([]*core.Task, error)
	
	// Update updates an existing task
	Update(task *core.Task) error
	
	// Delete deletes a task
	Delete(id string) error
	
	// Ready returns tasks that have no blocking dependencies
	Ready() ([]*core.Task, error)
	
	// Close closes the storage
	Close() error
}
