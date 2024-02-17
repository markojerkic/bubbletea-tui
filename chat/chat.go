package chat

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Message struct {
	isMine  bool
	from    *string
	date    time.Time
	message string
}

type MessagesList struct {
	messages             []Message
	messagesList viewport.Model
	messageInput         textinput.Model
}

func NewChat() MessagesList {
	textInput := textinput.New()
	textInput.Focus()
	return MessagesList{
		messagesList: viewport.Model{
			Height: 10,
		},
		messageInput: textInput,
	}
}

func (c MessagesList) Init() tea.Cmd {
	c.messageInput.Focus()
	return textinput.Blink
}

func (c MessagesList) Update(msg tea.Msg) (MessagesList, tea.Cmd) {

	var (
		vCmd tea.Cmd
		iCmd tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			message := Message{from: nil, message: c.messageInput.Value(),
				isMine: true, date: time.Now()}
			c.messageInput.Reset()

			c.messages = append(c.messages, message)

			c.messagesList.SetContent(c.getMessagesView())
			c.messagesList.GotoBottom()
		}
	}

	c.messagesList, vCmd = c.messagesList.Update(msg)
	c.messageInput, iCmd = c.messageInput.Update(msg)

	return c, tea.Batch(vCmd, iCmd)
}

func getSender(message Message) string {
	switch message.isMine {
	case true:
		return "Me"
	default:
		return *message.from
	}
}

func (c MessagesList) getMessagesView() string {
	listOfMessages := ""
	for _, message := range c.messages {
		listOfMessages = fmt.Sprintf("%s%s: %s\n", listOfMessages, getSender(message), message.message)
	}

	return listOfMessages
}

func (c MessagesList) View() string {
	c.getMessagesView()
	return fmt.Sprintf("%s\n%s", c.messagesList.View(), c.messageInput.View())
}
