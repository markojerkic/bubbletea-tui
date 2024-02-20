package chat

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Message struct {
	isMine  bool
	from    *string
	date    time.Time
	message string
}

type MessagesList struct {
	messages     []Message
	messagesList viewport.Model
	messageInput textinput.Model
}

func NewChat() MessagesList {
	textInput := textinput.New()
	textInput.Placeholder = "Type a message and press enter to send"
	textInput.Focus()

	messagesList := viewport.Model{
		Height: 10,
	}

	messagesList.Style = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FF00FF")).Margin(0, 1).Padding(1, 1)

	return MessagesList{
		messagesList: messagesList,
		messageInput: textInput,
	}
}

func (c MessagesList) Init() tea.Cmd {
	c.messageInput.Focus()
	return textinput.Blink
}

func (c MessagesList) Update(msg tea.Msg) (MessagesList, tea.Cmd) {

	cmds := make([]tea.Cmd, 0)
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.messagesList.Width = msg.Width - 1
		c.messagesList.Height = msg.Height - 10
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			message := Message{from: nil, message: c.messageInput.Value(),
				isMine: true, date: time.Now()}
			c.messages = append(c.messages, message)

			c.messagesList.SetContent(c.getMessagesView())
			c.messagesList.GotoBottom()

			c.messageInput.Reset()
		}
	}

	c.messageInput, cmd = c.messageInput.Update(msg)
	cmds = append(cmds, cmd)

	return c, tea.Batch(cmds...)
}

func getSender(message Message) string {
	switch message.isMine {
	case true:
		return "Me"
	default:
		return *message.from
	}
}

func getRandomLeftOrRightPosition() lipgloss.Position {
	random := rand.Intn(2)
	if random == 1 {
		return lipgloss.Right
	}
	return lipgloss.Right

}

func (c MessagesList) getMessagesView() string {
	messages := make([]string, len(c.messages))
	for _, message := range c.messages {

		line := lipgloss.NewStyle().Width(c.messagesList.Width) //.Align(getRandomLeftOrRightPosition())

		messageBubble := lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("2")).MaxWidth(60)

		mess := line.Render(messageBubble.Render(lipgloss.NewStyle().Foreground(lipgloss.Color("5")).Render(getSender(message)) + "\n" + message.message))

		messages = append(messages, mess)
	}

	return strings.Join(messages, "\n")
}

func (c MessagesList) View() string {
	return fmt.Sprintf("%s\n%s", c.messagesList.View(), c.messageInput.View())
}
