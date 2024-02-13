package main

// A simple program demonstrating the text area component from the Bubbles
// component library.

import (
	"log"
	"tui/register"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(register.InitRegister())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
