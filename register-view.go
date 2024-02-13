package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type register struct {
	selectedId chan string
	textinput  textinput.Model
	viewport   viewport.Model
	err        error
}

func initialRegistrationView() register {
	t := textinput.New()
	t.Placeholder = "Username"
	t.Focus()
	t.Width = 20

	vp := viewport.New(30, 5)
	vp.SetContent("Welcome to the chat app! Please enter your username")

	return register{
		selectedId: make(chan string),
		textinput:  t,
		viewport:   vp,
	}
}

func (r register) View() string {
	return r.View()
}

func (r register) Init() tea.Cmd {
	return textinput.Blink
}

func (r register) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	r.textinput, tiCmd = r.textinput.Update(msg)
	r.viewport, vpCmd = r.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		{
			switch msg.Type {
			case tea.KeyCtrlC, tea.KeyEsc:
				return r, tea.Quit
			case tea.KeyEnter:
				r.selectedId <- r.textinput.Value()
				r.textinput.Reset()
			}

		}

	case error:
		r.err = msg
		return r, nil

	}

	return r, tea.Batch(tiCmd, vpCmd)
}
