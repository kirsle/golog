package golog

import "github.com/kirsle/golog/ansi"

// Theme defines the color scheme for a logger. Each log level has two colors:
// a primary (for the label itself) and a secondary color. For example, if your
// log lines include a date/time this could be colored using the secondary
// color.
type Theme struct {
	Debug          ThemeColor
	DebugSecondary ThemeColor
	Info           ThemeColor
	InfoSecondary  ThemeColor
	Warn           ThemeColor
	WarnSecondary  ThemeColor
	Error          ThemeColor
	ErrorSecondary ThemeColor
}

// ThemeColor defines a color tuple for ANSI (legacy) support and modern
// 256-color support.
type ThemeColor struct {
	ANSI     string
	Extended string
}

// DarkTheme is a suitable default theme for dark terminal backgrounds.
var DarkTheme = Theme{
	Debug:          ThemeColor{ansi.BrightCyan, HexColor("#FF99FF")},
	DebugSecondary: ThemeColor{ansi.Cyan, HexColor("#996699")},
	Info:           ThemeColor{ansi.BrightGreen, HexColor("#0099FF")},
	InfoSecondary:  ThemeColor{ansi.Green, HexColor("#006699")},
	Warn:           ThemeColor{ansi.BrightYellow, HexColor("#FF9900")},
	WarnSecondary:  ThemeColor{ansi.Yellow, HexColor("#996600")},
	Error:          ThemeColor{ansi.BrightRed, HexColor("#FF0000")},
	ErrorSecondary: ThemeColor{ansi.Red, HexColor("#CC0000")},
}
