package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hamsa0x7/strand/internal/core"
	"github.com/hamsa0x7/strand/internal/storage"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			MarginBottom(1)

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("170")).
			Bold(true)

	statusColors = map[core.TaskStatus]lipgloss.Color{
		core.TaskStatusDone:       lipgloss.Color("10"),
		core.TaskStatusInProgress: lipgloss.Color("12"),
		core.TaskStatusBlocked:    lipgloss.Color("9"),
		core.TaskStatusBacklog:    lipgloss.Color("245"),
		core.TaskStatusReady:      lipgloss.Color("11"),
		core.TaskStatusCancelled:  lipgloss.Color("8"),
	}
)

type model struct {
	tasks    []*core.Task
	cursor   int
	selected map[string]bool
	store    storage.Store
	width    int
	height   int
}

func InitialModel(store storage.Store) (model, error) {
	tasks, err := store.List()
	if err != nil {
		return model{}, err
	}

	return model{
		tasks:    tasks,
		cursor:   0,
		selected: make(map[string]bool),
		store:    store,
		width:    80,
		height:   24,
	}, nil
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}

		case "enter", " ":
			if len(m.tasks) > 0 {
				id := m.tasks[m.cursor].ID
				m.selected[id] = !m.selected[id]
			}

		case "r":
			// Refresh task list
			tasks, err := m.store.List()
			if err == nil {
				m.tasks = tasks
				if m.cursor >= len(m.tasks) {
					m.cursor = len(m.tasks) - 1
				}
				if m.cursor < 0 {
					m.cursor = 0
				}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	var b strings.Builder

	// Title
	b.WriteString(titleStyle.Render("Strand Task Manager"))
	b.WriteString("\n\n")

	if len(m.tasks) == 0 {
		b.WriteString("No tasks found. Press 'q' to quit.\n")
		return b.String()
	}

	// Task list
	for i, task := range m.tasks {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}

		checked := " "
		if m.selected[task.ID] {
			checked = "x"
		}

		// Status color
		statusColor := statusColors[task.Status]
		statusStyle := lipgloss.NewStyle().Foreground(statusColor)

		line := fmt.Sprintf("%s [%s] %s - %s [%s]",
			cursor,
			checked,
			statusStyle.Render(string(task.Status)),
			task.Title,
			task.Priority,
		)

		if i == m.cursor {
			line = selectedStyle.Render(line)
		}

		b.WriteString(line)
		b.WriteString("\n")
	}

	// Help text
	b.WriteString("\n")
	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	help := helpStyle.Render("↑/k: up • ↓/j: down • space: select • r: refresh • q: quit")
	b.WriteString(help)

	return b.String()
}
