package tui

import (
	"crypto/rand"
	"fmt"
	"github.com/aldernero/ulid-tui/pkg/util"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/oklog/ulid"
	"os"
	"strings"
)

type Model struct {
	data          ulid.ULID
	enc           util.Enc
	isValid       bool
	invalidReason string
	input         textinput.Model
}

func initialModel(text string) Model {
	var data ulid.ULID
	var valid bool
	var reason string
	text = strings.TrimSpace(text)
	if text == "" {
		data = ulid.MustNew(ulid.Now(), rand.Reader)
		text = data.String()
		valid = true
	} else {
		data, valid, reason = util.ParseUlidString(text)
	}
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 26
	ti.Width = 26
	ti.Placeholder = "Enter ULID"
	ti.SetValue(text)
	return Model{
		data:          data,
		isValid:       valid,
		invalidReason: reason,
		input:         ti,
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
	m.data, m.isValid, m.invalidReason = util.ParseUlidString(m.input.Value())

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	var status string
	var table string
	if !m.isValid {
		status = invalidStyle(m.invalidReason)
		table = emptyEncodingTable()
	} else {
		status = validStyle("Valid ULID")
		table = createEncodingTable(m.data, m.enc)
	}
	return lipgloss.JoinVertical(lipgloss.Top,
		tuiStyle(lipgloss.JoinHorizontal(lipgloss.Center, m.input.View(), status)),
		tuiStyle(lipgloss.JoinHorizontal(lipgloss.Center, createUlidStringBreakdown(m.input.Value()), m.viewBaseSelector())),
		createTimeTable(m.data),
		table,
		helpMessage())
}

func (m Model) viewBaseSelector() string {
	var bin, hex, dec string
	switch m.enc {
	case util.Bin:
		bin = baseSelectedStyle("BIN")
		hex = baseUnselectedStyle("HEX")
		dec = baseUnselectedStyle("DEC")
	case util.Hex:
		bin = baseUnselectedStyle("BIN")
		hex = baseSelectedStyle("HEX")
		dec = baseUnselectedStyle("DEC")
	case util.Dec:
		bin = baseUnselectedStyle("BIN")
		hex = baseUnselectedStyle("HEX")
		dec = baseSelectedStyle("DEC")
	}
	return baseSelectorStyle(lipgloss.JoinVertical(lipgloss.Top, bin, hex, dec))
}
