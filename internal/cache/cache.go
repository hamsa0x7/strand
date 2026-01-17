package cache

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"time"

	"github.com/hamsa0x7/strand/internal/core"
	_ "github.com/mattn/go-sqlite3"
)

// Cache provides a SQLite-backed cache for fast task queries
type Cache struct {
	db *sql.DB
}

// NewCache creates a new cache instance
func NewCache(strandDir string) (*Cache, error) {
	dbPath := filepath.Join(strandDir, ".cache", "tasks.db")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open cache database: %w", err)
	}

	cache := &Cache{db: db}

	if err := cache.initSchema(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to initialize schema: %w", err)
	}

	return cache, nil
}

// initSchema creates the database schema
func (c *Cache) initSchema() error {
	schema := `
	CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		type TEXT NOT NULL,
		status TEXT NOT NULL,
		priority TEXT,
		title TEXT NOT NULL,
		description TEXT,
		assignee TEXT,
		created TEXT NOT NULL,
		updated TEXT NOT NULL,
		file_path TEXT NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS task_dependencies (
		task_id TEXT NOT NULL,
		depends_on_id TEXT NOT NULL,
		PRIMARY KEY (task_id, depends_on_id),
		FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE,
		FOREIGN KEY (depends_on_id) REFERENCES tasks(id) ON DELETE CASCADE
	);
	
	CREATE TABLE IF NOT EXISTS task_tags (
		task_id TEXT NOT NULL,
		tag TEXT NOT NULL,
		PRIMARY KEY (task_id, tag),
		FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
	);
	
	CREATE INDEX IF NOT EXISTS idx_tasks_status ON tasks(status);
	CREATE INDEX IF NOT EXISTS idx_tasks_priority ON tasks(priority);
	CREATE INDEX IF NOT EXISTS idx_tasks_type ON tasks(type);
	CREATE INDEX IF NOT EXISTS idx_task_tags_tag ON task_tags(tag);
	`

	_, err := c.db.Exec(schema)
	return err
}

// Sync synchronizes a task into the cache
func (c *Cache) Sync(task *core.Task) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Upsert task
	_, err = tx.Exec(`
		INSERT OR REPLACE INTO tasks 
		(id, type, status, priority, title, description, assignee, created, updated, file_path)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		task.ID,
		task.Type,
		task.Status,
		task.Priority,
		task.Title,
		task.Description,
		task.Assignee,
		task.Created.Format(time.RFC3339),
		task.Updated.Format(time.RFC3339),
		task.FilePath,
	)
	if err != nil {
		return err
	}

	// Clear and re-insert dependencies
	_, err = tx.Exec("DELETE FROM task_dependencies WHERE task_id = ?", task.ID)
	if err != nil {
		return err
	}

	for _, depID := range task.DependsOn {
		_, err = tx.Exec("INSERT INTO task_dependencies (task_id, depends_on_id) VALUES (?, ?)", task.ID, depID)
		if err != nil {
			return err
		}
	}

	// Clear and re-insert tags
	_, err = tx.Exec("DELETE FROM task_tags WHERE task_id = ?", task.ID)
	if err != nil {
		return err
	}

	for _, tag := range task.Tags {
		_, err = tx.Exec("INSERT INTO task_tags (task_id, tag) VALUES (?, ?)", task.ID, tag)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// Query runs a SQL query and returns matching tasks
func (c *Cache) Query(query string, args ...interface{}) ([]*core.Task, error) {
	rows, err := c.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*core.Task
	for rows.Next() {
		task, err := c.scanTask(rows)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, rows.Err()
}

// Get retrieves a task by ID from cache
func (c *Cache) Get(id string) (*core.Task, error) {
	row := c.db.QueryRow(`
		SELECT id, type, status, priority, title, description, assignee, created, updated, file_path
		FROM tasks WHERE id = ?
	`, id)

	return c.scanTask(row)
}

// scanTask scans a database row into a Task
func (c *Cache) scanTask(scanner interface {
	Scan(dest ...interface{}) error
}) (*core.Task, error) {
	var task core.Task
	var created, updated string

	err := scanner.Scan(
		&task.ID,
		&task.Type,
		&task.Status,
		&task.Priority,
		&task.Title,
		&task.Description,
		&task.Assignee,
		&created,
		&updated,
		&task.FilePath,
	)
	if err != nil {
		return nil, err
	}

	// Parse timestamps
	task.Created, _ = time.Parse(time.RFC3339, created)
	task.Updated, _ = time.Parse(time.RFC3339, updated)

	// Load dependencies
	rows, err := c.db.Query("SELECT depends_on_id FROM task_dependencies WHERE task_id = ?", task.ID)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var depID string
			if rows.Scan(&depID) == nil {
				task.DependsOn = append(task.DependsOn, depID)
			}
		}
	}

	// Load tags
	rows, err = c.db.Query("SELECT tag FROM task_tags WHERE task_id = ?", task.ID)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var tag string
			if rows.Scan(&tag) == nil {
				task.Tags = append(task.Tags, tag)
			}
		}
	}

	return &task, nil
}

// Clear removes all cached data
func (c *Cache) Clear() error {
	_, err := c.db.Exec("DELETE FROM tasks")
	return err
}

// Close closes the cache database
func (c *Cache) Close() error {
	return c.db.Close()
}
