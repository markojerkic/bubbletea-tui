package chat

import (
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
	incoming             chan Message
	messages             []Message
	messagesListViewPort viewport.Model
}

func NewChat(mChan chan Message) MessagesList {
	return MessagesList{
		incoming: mChan,
		messagesListViewPort: viewport.Model{
			Height: 10,
		},
	}
}

func (c MessagesList) Init() tea.Cmd {
	return textinput.Blink // viewport.Sync(c.messagesListViewPort)
}

func (c MessagesList) Update(msg tea.Msg) (MessagesList, tea.Cmd) {

	// newMessage := <-c.incoming
	//
	// c.messages = append(c.messages, newMessage)

	return c, nil
}

func getSender(message Message) string {
	switch message.isMine {
	case true:
		return "Me"
	default:
		return *message.from
	}
}

func (c MessagesList) View() string {
	listOfMessages := ""
	for _, message := range c.messages {
		listOfMessages += getSender(message) + ": " + message.message + "\n"
	}
	return listOfMessages
}

// A command that waits for the activity on a channel.
func waitForActivity(sub chan Message) tea.Cmd {
	return func() tea.Msg {
		return Message(<-sub)
	}
}
