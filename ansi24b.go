package richtext

import (
	"fmt"
)

type AnsiFormat24b struct{ ansiFormat }

var _ Format = &AnsiFormat24b{}

var ansiFormat24b *AnsiFormat24b

func init() {
	ansiFormat24b = &AnsiFormat24b{}
	ansiFormat24b.init(ansiFormat24b)
}

func Ansi24b() *AnsiFormat24b {
	return ansiFormat24b
}

func (a *AnsiFormat24b) MakePrintf(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) (int, error) {
	enc := []string{}
	if fg != None {
		r, g, b := rgb(fg)
		enc = append(enc, fmt.Sprintf("38;2;%d;%d;%d", r, g, b))
	}
	if bg != None {
		r, g, b := rgb(bg)
		enc = append(enc, fmt.Sprintf("48;2;%d;%d;%d", r, g, b))
	}
	return ansiPrinter(enc, flags)
}

func (a *AnsiFormat24b) MakeSprintf(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) string {
	enc := []string{}
	if fg != None {
		r, g, b := rgb(fg)
		enc = append(enc, fmt.Sprintf("38;2;%d;%d;%d", r, g, b))
	}
	if bg != None {
		r, g, b := rgb(bg)
		enc = append(enc, fmt.Sprintf("48;2;%d;%d;%d", r, g, b))
	}
	return ansiSPrinter(enc, flags)
}
