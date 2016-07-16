//Simple package for generating rich text, the canonical example of which is
//formatted text on console.
package richtext

import (
	"os"
	"strings"
)

//go:generate stringer -type Flag
//go:generate stringer -type Color

type (
	Color  uint32
	Flag   int
	Format interface {
		MakePrintf(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) (int, error)
		WarningLine(format string, a ...interface{})
		ErrorLine(format string, a ...interface{})
		PrintLine(format string, a ...interface{})
	}
)

const (
	None Color = 1 << 24

	Bold Flag = iota
	Underline
)

func rgb(c Color) (uint8, uint8, uint8) {
	return uint8((c >> 16) & 0xFF), uint8((c >> 8) & 0xFF), uint8((c) & 0xFF)
}

func RGB(r, g, b uint8) Color {
	return Color(uint32(b) + 0x100*uint32(g) + 0x10000*uint32(r))
}

func New() Format {
	term := os.Getenv("TERM")
	if strings.Contains(term, "24b") {
		return Ansi24b()
	} else if strings.Contains(term, "256") {
		return Ansi256()
	} else if isWindowsConsole() {
		return Console()
	} else if term != "" {
		return Ansi8()
	}
	return Ascii()
}
