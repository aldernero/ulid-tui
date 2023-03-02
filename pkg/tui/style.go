package tui

import "github.com/charmbracelet/lipgloss"

const (
	tuiWidth            = 50
	tableWidth          = 42
	ulidWidth           = 26
	helpMsg             = "[Tab] switch base • [Esc] quit"
	t1Placeholder       = "32_bit_uint_time_high"
	t2Placeholder       = "16_bit_uint_time_low"
	e1Placeholder       = "16_bit_uint_random"
	e2Placeholder       = "32_bit_uint_random"
	e3Placeholder       = "32_bit_uint_random"
	timeDescWidth       = 9
	timeValWidth        = 32
	timeWidth           = 10
	entropyWidth        = 16
	helpColor           = lipgloss.Color("#888888")
	borderColor         = lipgloss.Color("#de3e93")
	placeholderColor    = lipgloss.Color("#888888")
	baseSelectedColor   = lipgloss.Color("#47a4ac")
	baseUnselectedColor = lipgloss.Color("#baebda")
	invalidColor        = lipgloss.Color("#ff0000")
	validColor          = lipgloss.Color("#08e225")
	timeColor           = lipgloss.Color("#ffdf80")
	entropyColor        = lipgloss.Color("#80a0ff")
	headerColor         = lipgloss.Color("#3ede89")
	footerColor         = lipgloss.Color("#3ede89")
	crtLightGray        = lipgloss.Color("#c9c9c9")
)

var headerBorder = lipgloss.Border{
	Left:  "╔",
	Right: "╗",
}

var footerBorder = lipgloss.Border{
	Left:  "└",
	Right: "┘",
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

var timeUTCDescBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "",
	Left:        "│",
	Right:       "│",
	TopLeft:     "┌",
	TopRight:    "┬",
	BottomRight: "┴",
	BottomLeft:  "╷",
}

var timeLocalDescBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "├",
	TopRight:    "┼",
	BottomRight: "┴",
	BottomLeft:  "└",
}

var timeEpochDescBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "├",
	TopRight:    "┼",
	BottomRight: "┴",
	BottomLeft:  "└",
}

var timeUTCValBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "",
	Left:        "│",
	Right:       "│",
	TopLeft:     "┌",
	TopRight:    "┐",
	BottomRight: "╷",
	BottomLeft:  "╷",
}

var timeLocalValBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "┌",
	TopRight:    "┤",
	BottomRight: "┘",
	BottomLeft:  "╷",
}
var timeEpochValBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "┌",
	TopRight:    "┤",
	BottomRight: "┘",
	BottomLeft:  "╷",
}

var tuiStyle = lipgloss.NewStyle().
	Width(tuiWidth).
	Align(lipgloss.Center).
	Render

var baseSelectorStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Padding(0, 1, 0, 1).
	Margin(0, 0, 0, 7).
	Border(lipgloss.NormalBorder(), true).
	BorderForeground(baseSelectedColor).
	Render

var baseSelectedStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Background(baseSelectedColor).
	Foreground(baseUnselectedColor).
	Render

var baseUnselectedStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Foreground(baseSelectedColor).
	Render

var helpStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Width(tuiWidth).
	Margin(1, 0, 0, 0).
	Foreground(helpColor).
	Render

var invalidStyle = lipgloss.NewStyle().
	Align(lipgloss.Left).
	Foreground(invalidColor).
	Italic(true).
	Render

var validStyle = lipgloss.NewStyle().
	Align(lipgloss.Left).
	Foreground(validColor).
	Italic(true).
	Render

var stringBreakdownStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Margin(0, 0, 0, 2).
	Width(ulidWidth).
	Render

var timeStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	BorderForeground(timeColor).
	Margin(1, 0, 0, 0).
	Width(timeWidth).
	Border(lipgloss.NormalBorder(), false, false, true, false)

var timeLabelStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Foreground(timeColor).
	Margin(0, 0, 1, 0).
	Width(timeWidth).
	Render

var entropyStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	BorderForeground(entropyColor).
	Margin(1, 0, 0, 0).
	Width(entropyWidth).
	Border(lipgloss.NormalBorder(), false, false, true, false)

var entropyLabelStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Foreground(entropyColor).
	Margin(0, 0, 1, 0).
	Width(entropyWidth).
	Render

var headerStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Foreground(headerColor).
	BorderForeground(headerColor).
	Border(headerBorder, false, true)

var footerStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Foreground(footerColor).
	BorderForeground(footerColor).
	Border(footerBorder, false, true)

var baseStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Padding(0).
	Margin(0).
	BorderForeground(borderColor)

var t1Style = lipgloss.NewStyle().
	Inherit(baseStyle).
	Foreground(timeColor).
	Width(tableWidth).
	Border(t1Border, true, true, false, true).
	Render

var t2Style = lipgloss.NewStyle().
	Inherit(baseStyle).
	Foreground(timeColor).
	Width(tableWidth/2).
	Border(t2Border, true, true, true, true).
	Render

var e1Style = lipgloss.NewStyle().
	Inherit(baseStyle).
	Foreground(entropyColor).
	Width(tableWidth/2-1).
	Border(e1Border, true, true, true, false).
	Render

var e2Style = lipgloss.NewStyle().
	Inherit(baseStyle).
	Foreground(entropyColor).
	Width(tableWidth).
	Border(e2Border, false, true, true, true).
	Render

var e3Style = lipgloss.NewStyle().
	Inherit(baseStyle).
	Foreground(entropyColor).
	Width(tableWidth).
	Border(e3Border, false, true, true, true).
	Render

var timeUTCDescStyle = lipgloss.NewStyle().
	Align(lipgloss.Right).
	Width(timeDescWidth).
	Padding(0).
	Margin(0).
	BorderForeground(borderColor).
	//Foreground(crtLightGray).
	Border(timeUTCDescBorder, true, true, false, true).
	Render

var timeLocalDescStyle = lipgloss.NewStyle().
	Align(lipgloss.Right).
	Width(timeDescWidth).
	Padding(0).
	Margin(0).
	BorderForeground(borderColor).
	//Foreground(crtLightGray).
	Border(timeLocalDescBorder, true, true, false, true).
	Render

var timeUTCValStyle = lipgloss.NewStyle().
	Align(lipgloss.Left).
	Width(timeValWidth).
	Padding(0).
	Margin(0).
	BorderForeground(borderColor).
	//Foreground(crtLightGray).
	Border(timeUTCValBorder, true, true, false, false).
	Render

var timeLocalValStyle = lipgloss.NewStyle().
	Align(lipgloss.Left).
	Width(timeValWidth).
	Padding(0).
	Margin(0).
	BorderForeground(borderColor).
	//Foreground(crtLightGray).
	Border(timeLocalValBorder, true, true, false, false).
	Render

var timeEpochDescStyle = lipgloss.NewStyle().
	Align(lipgloss.Right).
	Width(timeDescWidth).
	Padding(0).
	Margin(0).
	BorderForeground(borderColor).
	//Foreground(crtLightGray).
	Border(timeEpochDescBorder, true, true, true, true).
	Render

var timeEpochValStyle = lipgloss.NewStyle().
	Align(lipgloss.Left).
	Width(timeValWidth).
	Padding(0).
	Margin(0).
	BorderForeground(borderColor).
	//Foreground(crtLightGray).
	Border(timeEpochValBorder, true, true, true, false).
	Render
