package main

import (
	"log"
	"tui/chat"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	chat chat.MessagesList
}

func newChat() model {

	return model{
		chat: chat.NewChat(),
	}
}

func (m model) Init() tea.Cmd {
	return m.chat.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	m.chat, cmd = m.chat.Update(msg)

	return m, cmd
}

func (m model) View() string {
	return m.chat.View()
}

func main() {
	p := tea.NewProgram(newChat())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
