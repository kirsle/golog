package golog

import (
	"fmt"

	"github.com/tomnomnom/xtermcolor"
)

type colorLevel int

// Options for color support in your logger.
const (
	// NoColor doesn't use any color codes at all (plain text).
	NoColor colorLevel = iota

	// ANSIColor uses the standard 16 colors supported by most terminals. This
	// option is the most portable across platforms.
	ANSIColor

	// ExtendedColor allows the use of 256 colors supported by most modern
	// terminals (24-bit color codes).
	ExtendedColor
)

// HexColor is a convenient wrapper around `xtermcolor.FromHexStr` to define colors
// for themes for xterm-256color codes.
func HexColor(hex string) string {
	code, err := xtermcolor.FromHexStr(hex)
	if err != nil {
		code = 201 // bright magenta seems like a good default
	}

	return fmt.Sprintf("38;5;%d", code)
}
