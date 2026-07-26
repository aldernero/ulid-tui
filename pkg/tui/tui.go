package tui

import (
	"crypto/rand"
	"fmt"
	"os"
	"strings"

	"github.com/aldernero/ulid-tui/pkg/util"
	"github.com/oklog/ulid/v2"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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
	ti.SetWidth(26)
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
	p := tea.NewProgram(m)
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

	if msg, ok := msg.(tea.KeyPressMsg); ok {
		switch msg.String() {
		case "enter", "ctrl+c", "esc":
			return m, tea.Quit
		case "tab":
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

func (m Model) View() tea.View {
	var status string
	var table string
	if !m.isValid {
		status = invalidStyle(m.invalidReason)
		table = emptyEncodingTable()
	} else {
		status = validStyle("Valid ULID")
		table = createEncodingTable(m.data, m.enc)
	}
	content := lipgloss.JoinVertical(lipgloss.Top,
		tuiStyle(lipgloss.JoinHorizontal(lipgloss.Center, m.input.View(), status)),
		tuiStyle(lipgloss.JoinHorizontal(lipgloss.Center, createUlidStringBreakdown(m.input.Value()), m.viewBaseSelector())),
		createTimeTable(m.data),
		table,
		helpMessage())
	v := tea.NewView(content)
	v.AltScreen = true
	return v
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
