package markdown

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hamsa0x7/strand/internal/core"
	"gopkg.in/yaml.v3"
)

// Store implements storage.Store using Markdown files
type Store struct {
	tasksDir string
}

// NewStore creates a new markdown-based store
func NewStore() *Store {
	return &Store{}
}

// Init initializes the markdown store
func (s *Store) Init(strandDir string) error {
	s.tasksDir = filepath.Join(strandDir, "tasks")

	// Create tasks directory if it doesn't exist
	if err := os.MkdirAll(s.tasksDir, 0755); err != nil {
		return fmt.Errorf("failed to create tasks directory: %w", err)
	}

	return nil
}

// Create creates a new task as a markdown file
func (s *Store) Create(task *core.Task) error {
	filename := s.taskFilename(task.ID)
	task.FilePath = filename

	content, err := s.taskToMarkdown(task)
	if err != nil {
		return fmt.Errorf("failed to convert task to markdown: %w", err)
	}

	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write task file: %w", err)
	}

	return nil
}

// Get retrieves a task by ID
func (s *Store) Get(id string) (*core.Task, error) {
	filename := s.taskFilename(id)

	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("task not found: %s", id)
		}
		return nil, fmt.Errorf("failed to read task file: %w", err)
	}

	task, err := s.markdownToTask(data, filename)
	if err != nil {
		return nil, fmt.Errorf("failed to parse task: %w", err)
	}

	return task, nil
}

// List retrieves all tasks
func (s *Store) List() ([]*core.Task, error) {
	entries, err := os.ReadDir(s.tasksDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []*core.Task{}, nil
		}
		return nil, fmt.Errorf("failed to read tasks directory: %w", err)
	}

	var tasks []*core.Task
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		filename := filepath.Join(s.tasksDir, entry.Name())
		data, err := os.ReadFile(filename)
		if err != nil {
			continue // Skip unreadable files
		}

		task, err := s.markdownToTask(data, filename)
		if err != nil {
			continue // Skip unparseable files
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// Update updates an existing task
func (s *Store) Update(task *core.Task) error {
	// Same as Create - overwrite the file
	return s.Create(task)
}

// Delete deletes a task
func (s *Store) Delete(id string) error {
	filename := s.taskFilename(id)

	if err := os.Remove(filename); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("task not found: %s", id)
		}
		return fmt.Errorf("failed to delete task: %w", err)
	}

	return nil
}

// Ready returns tasks with no blocking dependencies
func (s *Store) Ready() ([]*core.Task, error) {
	allTasks, err := s.List()
	if err != nil {
		return nil, err
	}

	// Build a map for quick lookup
	taskMap := make(map[string]*core.Task)
	for _, t := range allTasks {
		taskMap[t.ID] = t
	}

	var readyTasks []*core.Task
	for _, task := range allTasks {
		if task.IsReady(taskMap) {
			readyTasks = append(readyTasks, task)
		}
	}

	return readyTasks, nil
}

// Close closes the store (no-op for markdown)
func (s *Store) Close() error {
	return nil
}

// taskFilename returns the full path to a task's markdown file
func (s *Store) taskFilename(id string) string {
	return filepath.Join(s.tasksDir, id+".md")
}

// taskToMarkdown converts a task to markdown format
func (s *Store) taskToMarkdown(task *core.Task) (string, error) {
	var buf bytes.Buffer

	// Write YAML frontmatter
	buf.WriteString("---\n")

	// Create frontmatter struct (only metadata,not title/description)
	type Frontmatter struct {
		ID        string            `yaml:"id"`
		Type      core.TaskType     `yaml:"type"`
		Status    core.TaskStatus   `yaml:"status"`
		Priority  core.TaskPriority `yaml:"priority,omitempty"`
		DependsOn []string          `yaml:"depends_on,omitempty"`
		Assignee  string            `yaml:"assignee,omitempty"`
		Created   string            `yaml:"created"`
		Updated   string            `yaml:"updated"`
		Tags      []string          `yaml:"tags,omitempty"`
	}

	fm := Frontmatter{
		ID:        task.ID,
		Type:      task.Type,
		Status:    task.Status,
		Priority:  task.Priority,
		DependsOn: task.DependsOn,
		Assignee:  task.Assignee,
		Created:   task.Created.Format("2006-01-02T15:04:05Z07:00"),
		Updated:   task.Updated.Format("2006-01-02T15:04:05Z07:00"),
		Tags:      task.Tags,
	}

	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)
	if err := encoder.Encode(&fm); err != nil {
		return "", err
	}
	encoder.Close()

	buf.WriteString("---\n\n")

	// Write title
	buf.WriteString("# ")
	buf.WriteString(task.Title)
	buf.WriteString("\n\n")

	// Write description
	if task.Description != "" {
		buf.WriteString(task.Description)
		buf.WriteString("\n")
	}

	return buf.String(), nil
}

// markdownToTask parses a markdown file into a Task
func (s *Store) markdownToTask(data []byte, filename string) (*core.Task, error) {
	content := string(data)

	// Split frontmatter and body
	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid markdown format: missing frontmatter")
	}

	// Parse YAML frontmatter
	type Frontmatter struct {
		ID        string            `yaml:"id"`
		Type      core.TaskType     `yaml:"type"`
		Status    core.TaskStatus   `yaml:"status"`
		Priority  core.TaskPriority `yaml:"priority"`
		DependsOn []string          `yaml:"depends_on"`
		Assignee  string            `yaml:"assignee"`
		Created   string            `yaml:"created"`
		Updated   string            `yaml:"updated"`
		Tags      []string          `yaml:"tags"`
	}

	var fm Frontmatter
	if err := yaml.Unmarshal([]byte(parts[1]), &fm); err != nil {
		return nil, fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	// Parse body (title + description)
	body := strings.TrimSpace(parts[2])
	lines := strings.Split(body, "\n")

	var title string
	var description string

	if len(lines) > 0 && strings.HasPrefix(lines[0], "# ") {
		title = strings.TrimPrefix(lines[0], "# ")
		if len(lines) > 1 {
			description = strings.Join(lines[1:], "\n")
			description = strings.TrimSpace(description)
		}
	}

	// Build task
	task := &core.Task{
		ID:          fm.ID,
		Type:        fm.Type,
		Status:      fm.Status,
		Priority:    fm.Priority,
		Title:       title,
		Description: description,
		DependsOn:   fm.DependsOn,
		Assignee:    fm.Assignee,
		Tags:        fm.Tags,
		FilePath:    filename,
	}

	// Parse timestamps
	if created, err := parseTime(fm.Created); err == nil {
		task.Created = created
	}
	if updated, err := parseTime(fm.Updated); err == nil {
		task.Updated = updated
	}

	return task, nil
}

// parseTime parses an ISO 8601 timestamp
func parseTime(s string) (time.Time, error) {
	// Try RFC3339 format first
	t, err := time.Parse(time.RFC3339, s)
	if err == nil {
		return t, nil
	}

	// Try without timezone
	t, err = time.Parse("2006-01-02T15:04:05", s)
	return t, err
}
