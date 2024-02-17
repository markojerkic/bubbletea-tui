package main

import (
	"log"
	"tui/chat"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	mChan        chan chat.Message
	messageInput chat.ChatInput
	messagesList chat.MessagesList
}

func newChat() model {

	mChan := make(chan chat.Message, 10)

	return model{
		messageInput: chat.NewChatInput(mChan),
		messagesList: chat.NewChat(mChan),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.messageInput.Init(), m.messagesList.Init())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		inputCmd tea.Cmd
		listCmd  tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	log.Println("update", msg)

	m.messageInput, inputCmd = m.messageInput.Update(msg)
	m.messagesList, listCmd = m.messagesList.Update(msg)

	return m, tea.Batch(inputCmd, listCmd)
}

func (m model) View() string {
	return m.messageInput.View() + "\n\n" + m.messagesList.View()
}

func main() {
	p := tea.NewProgram(newChat())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
