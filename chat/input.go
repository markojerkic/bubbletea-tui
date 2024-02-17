package chat

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ChatInput struct {
	messageInput textinput.Model
	outgoing     chan Message
}

func NewChatInput(mChan chan Message) ChatInput {
	ti := textinput.New()
	ti.Placeholder = "Type a message and press enter to send"
	return ChatInput{
		messageInput: ti,
		outgoing:     mChan,
	}
}

func (ci ChatInput) Init() tea.Cmd {
	return textinput.Blink
}

func (ci ChatInput) Update(msg tea.Msg) (ChatInput, tea.Cmd) {
	var cmd tea.Cmd

	ci.messageInput, cmd = ci.messageInput.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			ci.messageInput.Reset()
			message := Message{from: nil, message: ci.messageInput.Value(),
				isMine: true, date: time.Now()}

			ci.outgoing <- message
		}
	}

	return ci, cmd
}

func (ci ChatInput) View() string {
	return fmt.Sprintf(">: %s\n", ci.messageInput.View())
}
