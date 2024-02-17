package register

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type model struct {
	viewPort         viewport.Model
	textInput        textinput.Model
	selectedUserName chan string
	err              error
}

func InitRegister() model {
	ti := textinput.New()
	ti.Placeholder = "Username..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	vp := viewport.New(50, 1)
	vp.SetContent("Welcome to my chat app! Please enter your username and press enter to continue.")

	return model{
		selectedUserName: make(chan string),
		viewPort:         vp,
		textInput:        ti,
		err:              nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tCmd tea.Cmd
		vCmd tea.Cmd
	)

	m.textInput, tCmd = m.textInput.Update(msg)
	m.viewPort, vCmd = m.viewPort.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			// m.selectedUserName <- m.textInput.Value()
			m.viewPort.SetContent(fmt.Sprintf("Welcome, %s!", m.textInput.Value()))
			m.textInput.Reset()
			m.viewPort.GotoBottom()
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(tCmd, vCmd)
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n%s",
		<-m.selectedUserName,
		m.textInput.View(),
	)
}
