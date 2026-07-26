package tui

import (
	"strconv"
	"strings"

	"github.com/aldernero/ulid-tui/pkg/util"
	"github.com/oklog/ulid/v2"

	"charm.land/lipgloss/v2"
)

func genHeader(msg string, width int) string { //nolint:unparam // callers currently share one table width, but each derives it independently
	// +2 for the left/right border chars: lipgloss v2 Width() is border-box.
	return headerStyle.Width(width + 2).Render(util.PadWithString(msg, "═", width))
}

func genFooter(msg string, width int) string {
	return footerStyle.Width(width + 2).Render(util.PadWithString(msg, "─", width))
}

func emptyEncodingTable() string {
	return tuiStyle(lipgloss.JoinVertical(
		lipgloss.Top,
		genHeader("Encoding", tableWidth),
		t1Style(t1Placeholder),
		lipgloss.JoinHorizontal(0.0, t2Style(t2Placeholder), e1Style(e1Placeholder)),
		e2Style(e2Placeholder),
		e3Style(e3Placeholder),
		genFooter("4 bytes", tableWidth)))
}

func createEncodingTable(id ulid.ULID, enc util.Enc) string {
	d := util.NewULID(id)
	return tuiStyle(lipgloss.JoinVertical(
		lipgloss.Top,
		genHeader("Encoding", tableWidth),
		t1Style(d.ToString(util.T1, enc)),
		lipgloss.JoinHorizontal(0.0, t2Style(d.ToString(util.T2, enc)), e1Style(d.ToString(util.E1, enc))),
		e2Style(d.ToString(util.E2, enc)),
		e3Style(d.ToString(util.E3, enc)),
		genFooter("4 bytes", tableWidth)))
}

func createTimeTable(id ulid.ULID) string {
	utc, local, ms := util.UlidTimes(id)
	var epoch string
	if ms == 0 {
		utc = ""
		local = ""
	} else {
		epoch = strconv.FormatUint(ms, 10) + " milliseconds"
	}
	return tuiStyle(lipgloss.JoinVertical(
		lipgloss.Top,
		genHeader("Time", timeDescWidth+timeValWidth+1),
		lipgloss.JoinHorizontal(0.0, timeUTCDescStyle("UTC"), timeUTCValStyle(utc)),
		lipgloss.JoinHorizontal(0.0, timeLocalDescStyle("Local"), timeLocalValStyle(local)),
		lipgloss.JoinHorizontal(0.0, timeEpochDescStyle("Epoch"), timeEpochValStyle(epoch)),
	))
}

func createUlidStringBreakdown(id string) string {
	id = strings.ToUpper(id)
	var time, entropy string
	if len(id) < 26 {
		time = timeStyle.Foreground(placeholderColor).Render(strings.Repeat("░", 10))
		entropy = entropyStyle.Foreground(placeholderColor).Render(strings.Repeat("░", 16))
	} else {
		time = timeStyle.Foreground(timeColor).Render(id[:10])
		entropy = entropyStyle.Foreground(entropyColor).Render(id[10:])
	}
	return stringBreakdownStyle(lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(0.0, time, entropy),
		lipgloss.JoinHorizontal(0.0, timeLabelStyle("Time"), entropyLabelStyle("Entropy"))))
}

func helpMessage() string {
	return helpStyle(helpMsg)
}
