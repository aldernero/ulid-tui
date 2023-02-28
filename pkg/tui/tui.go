package tui

import (
	"crypto/rand"
	"fmt"
	"github.com/aldernero/ulid-tui/pkg/util"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/oklog/ulid"
	"os"
	"strings"
	"time"
)

type Model struct {
	data    ulid.ULID
	enc     util.Enc
	isValid bool
	input   textinput.Model
}

func initialModel(text string) Model {
	var data ulid.ULID
	var valid bool
	text = strings.TrimSpace(text)
	if text == "" {
		data = ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader)
		text = data.String()
		valid = true
	} else {
		data, valid = util.ParseUlidString(text)
	}
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 26
	ti.Width = 26
	ti.Placeholder = "Enter ULID"
	ti.SetValue(text)
	return Model{
		data:    data,
		isValid: valid,
		input:   ti,
	}
}

func StartTea(text string) {
	m := initialModel(text)
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyTab:
			if m.enc == util.Dec {
				m.enc = util.Bin
			} else {
				m.enc++
			}
		}
	}
	m.data, m.isValid = util.ParseUlidString(m.input.Value())

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	var output string
	var table string
	var local, utc string
	if !m.isValid {
		output = "Invalid ULID"
		table = emptyTable()
	} else {
		local, utc = util.UlidTimes(m.data)
		table = createTable(m.data, m.enc)
	}
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s", m.input.View(), output, utc, local, table)
}
