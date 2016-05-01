package richtext

import (
	"fmt"
)

type ansiFormat24b struct{}

func Ansi24b() Format {
	return &ansiFormat24b{}
}

func (a *ansiFormat24b) String(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) string {
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

func (a *ansiFormat24b) Bytes(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) []byte {
	f := a.String(fg, bg, flags...)
	return func(format string, a ...interface{}) []byte {
		return []byte(f(format, a...))
	}
}
