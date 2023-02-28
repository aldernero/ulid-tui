package tui

import (
	"github.com/aldernero/ulid-tui/pkg/util"
	"github.com/charmbracelet/lipgloss"
	"github.com/oklog/ulid"
)

func genHeader() string {
	return headerStyle(util.PadWithChar("4 bytes", "<", ">", "-", tableWidth))
}

func emptyTable() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		genHeader(),
		t1Style(t1Placeholder),
		lipgloss.JoinHorizontal(0.0, t2Style(t2Placeholder), e1Style(e1Placeholder)),
		e2Style(e2Placeholder),
		e3Style(e3Placeholder))
}

func createTable(id ulid.ULID, enc util.Enc) string {
	d := util.NewULID(id)
	return lipgloss.JoinVertical(
		lipgloss.Top,
		genHeader(),
		t1Style(d.ToString(util.T1, enc)),
		lipgloss.JoinHorizontal(0.0, t2Style(d.ToString(util.T2, enc)), e1Style(d.ToString(util.E1, enc))),
		e2Style(d.ToString(util.E2, enc)),
		e3Style(d.ToString(util.E3, enc)))
}
