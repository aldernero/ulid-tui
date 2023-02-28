package tui

import "github.com/charmbracelet/lipgloss"

const (
	tableWidth    = 50
	t1Placeholder = "32_bit_uint_time_high"
	t2Placeholder = "16_bit_uint_time_low"
	e1Placeholder = "16_bit_uint_random"
	e2Placeholder = "32_bit_uint_random"
	e3Placeholder = "32_bit_uint_random"
	timeSymbol    = "🕙"
	entropySymbol = "🎲"
)

var headerBorder = lipgloss.Border{
	Left:  "│",
	Right: "│",
}

var t1Border = lipgloss.Border{
	Top:         "─",
	Bottom:      "",
	Left:        "│",
	Right:       "│",
	TopLeft:     "┌",
	TopRight:    "┐",
	BottomRight: "╷",
	BottomLeft:  "╷",
}

var t2Border = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "├",
	TopRight:    "┬",
	BottomRight: "┴",
	BottomLeft:  "├",
}

var e1Border = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "├",
	TopRight:    "┤",
	BottomRight: "┤",
	BottomLeft:  "╷",
}

var e2Border = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "├",
	TopRight:    "┤",
	BottomRight: "┤",
	BottomLeft:  "├",
}

var e3Border = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "├",
	TopRight:    "┤",
	BottomRight: "┘",
	BottomLeft:  "└",
}

var borderColor = lipgloss.Color("#ff00ff")
var timeColor = lipgloss.Color("#00ffff")
var entropyColor = lipgloss.Color("#ffff00")

var headerStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Width(tableWidth).
	BorderForeground(borderColor).
	Border(headerBorder, false, true, false, true).
	Render

var baseStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Padding(0).
	Margin(0).
	BorderForeground(borderColor)

var t1Style = lipgloss.NewStyle().
	Inherit(baseStyle).
	Width(tableWidth).
	Border(t1Border, true, true, false, true).
	Render

var t2Style = lipgloss.NewStyle().
	Inherit(baseStyle).
	Width(tableWidth/2).
	Border(t2Border, true, true, true, true).
	Render

var e1Style = lipgloss.NewStyle().
	Inherit(baseStyle).
	Width(tableWidth/2-1).
	Border(e1Border, true, true, true, false).
	Render

var e2Style = lipgloss.NewStyle().
	Inherit(baseStyle).
	Width(tableWidth).
	Border(e2Border, false, true, true, true).
	Render

var e3Style = lipgloss.NewStyle().
	Inherit(baseStyle).
	Width(tableWidth).
	Border(e3Border, false, true, true, true).
	Render
