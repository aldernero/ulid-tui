package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type sessionState int

const (
	hexadecimal sessionState = iota
	binary
)

type Model struct {
	state sessionState
	ulid  string
}

func StartTea() {
	m := Model{}
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
