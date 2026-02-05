package ui

import (
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	"github.com/justynhunter/speedreader/lib"
)

type UiModel struct {
	DelayInMs     int
	HideHighlight bool
	WordProcessor lib.WordProcessor
}

type tickMsg = time.Time

func (m UiModel) Init() tea.Cmd {
	return tick(m.DelayInMs)
}

func (m UiModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}

	case tickMsg:
		eof := m.WordProcessor.Next()
		if eof {
			return m, tea.Quit
		}
		return m, tick(m.DelayInMs)
	}

	return m, nil
}

func (m UiModel) View() string {
	width, height, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		log.Fatal("unable to determin terminal size")
	}

	var text string

	if m.HideHighlight {
		text = m.WordProcessor.CurrentWord
	} else {
		text = lipgloss.StyleRunes(
			m.WordProcessor.CurrentWord,
			[]int{max(0, len(m.WordProcessor.CurrentWord)/2)},
			lipgloss.NewStyle().Foreground(lipgloss.Color("#98FF98")),
			lipgloss.NewStyle(),
		)
	}

	content := lipgloss.NewStyle().Render(text)

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, content)
}

func tick(delayInMs int) tea.Cmd {
	return tea.Tick(time.Millisecond*time.Duration(delayInMs), func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
