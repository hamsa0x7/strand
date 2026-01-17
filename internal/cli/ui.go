package cli

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hamsa0x7/strand/internal/tui"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Launch interactive TUI",
	Long:  `Launch the interactive terminal UI for managing tasks.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create TUI model
		m, err := tui.InitialModel(store)
		if err != nil {
			return fmt.Errorf("failed to initialize TUI: %w", err)
		}

		// Run TUI
		p := tea.NewProgram(m, tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			return fmt.Errorf("TUI error: %w", err)
		}

		return nil
	},
}
