package tui

import (
	"strings"
	"testing"

	"github.com/aldernero/ulid-tui/pkg/util"
	"github.com/oklog/ulid/v2"

	tea "charm.land/bubbletea/v2"
)

func TestInitialModelGeneratesULIDWhenEmpty(t *testing.T) {
	m := initialModel("")
	if !m.isValid {
		t.Fatalf("expected generated ULID to be valid, reason=%q", m.invalidReason)
	}
	if m.input.Value() == "" {
		t.Fatal("expected input to be populated with the generated ULID string")
	}
}

func TestInitialModelWithValidText(t *testing.T) {
	text := "01ARZ3NDEKTSV4RRFFQ69G5FAV"
	m := initialModel(text)
	if !m.isValid {
		t.Fatalf("expected %q to be valid, reason=%q", text, m.invalidReason)
	}
	want := ulid.MustParseStrict(text)
	if m.data != want {
		t.Errorf("m.data = %v, want %v", m.data, want)
	}
}

func TestInitialModelWithInvalidText(t *testing.T) {
	m := initialModel("not-a-ulid")
	if m.isValid {
		t.Fatal("expected invalid input to be marked invalid")
	}
	if m.invalidReason == "" {
		t.Fatal("expected a non-empty invalid reason")
	}
}

func TestUpdateTabCyclesEncoding(t *testing.T) {
	m := initialModel("01ARZ3NDEKTSV4RRFFQ69G5FAV")
	if m.enc != util.Bin {
		t.Fatalf("expected initial encoding to be Bin, got %v", m.enc)
	}

	next, _ := m.Update(keyPress("tab"))
	m = next.(Model)
	if m.enc != util.Hex {
		t.Fatalf("after one tab, enc = %v, want Hex", m.enc)
	}

	next, _ = m.Update(keyPress("tab"))
	m = next.(Model)
	if m.enc != util.Dec {
		t.Fatalf("after two tabs, enc = %v, want Dec", m.enc)
	}

	next, _ = m.Update(keyPress("tab"))
	m = next.(Model)
	if m.enc != util.Bin {
		t.Fatalf("after three tabs, enc = %v, want Bin (wrapped around)", m.enc)
	}
}

func TestUpdateQuitKeys(t *testing.T) {
	for _, key := range []string{"enter", "ctrl+c", "esc"} {
		t.Run(key, func(t *testing.T) {
			m := initialModel("01ARZ3NDEKTSV4RRFFQ69G5FAV")
			_, cmd := m.Update(keyPress(key))
			if cmd == nil {
				t.Fatalf("expected a Cmd for key %q, got nil", key)
			}
			msg := cmd()
			if _, ok := msg.(tea.QuitMsg); !ok {
				t.Errorf("expected tea.QuitMsg for key %q, got %T", key, msg)
			}
		})
	}
}

func TestUpdateRevalidatesOnInputChange(t *testing.T) {
	m := initialModel("01ARZ3NDEKTSV4RRFFQ69G5FAV")
	m.input.SetValue("garbage")

	next, _ := m.Update(tea.WindowSizeMsg{})
	m = next.(Model)
	if m.isValid {
		t.Fatal("expected model to become invalid after input changed to garbage")
	}
}

func TestViewProducesAltScreenContent(t *testing.T) {
	m := initialModel("01ARZ3NDEKTSV4RRFFQ69G5FAV")
	v := m.View()
	if !v.AltScreen {
		t.Error("expected View() to request the alt screen")
	}
	if !strings.Contains(v.Content, "Valid ULID") {
		t.Error("expected view content to report a valid ULID")
	}
}

func TestViewInvalidShowsReason(t *testing.T) {
	m := initialModel("not-a-ulid")
	v := m.View()
	if !strings.Contains(v.Content, m.invalidReason) {
		t.Errorf("expected view content to contain invalid reason %q", m.invalidReason)
	}
}

// keyPress builds a KeyPressMsg whose String() matches the given keystroke,
// for keys where Code alone (as used in tui.go's Update switch) is sufficient.
func keyPress(s string) tea.KeyPressMsg {
	switch s {
	case "tab":
		return tea.KeyPressMsg{Code: tea.KeyTab}
	case "enter":
		return tea.KeyPressMsg{Code: tea.KeyEnter}
	case "esc":
		return tea.KeyPressMsg{Code: tea.KeyEscape}
	case "ctrl+c":
		return tea.KeyPressMsg{Code: 'c', Mod: tea.ModCtrl}
	}
	panic("unsupported key in test helper: " + s)
}
