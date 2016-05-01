package richtext

import (
	"fmt"
)

type ansiFormat8 struct{}

//This varies by platform, this is what windows has, the most likely source of us
//running a 16 color pallette

func Ansi8() Format {
	return &ansiFormat8{}
}

func rgbTo8(c Color) uint8 {
	r, g, b := rgb(c)
	r1 := r >> 7
	g1 := g >> 7
	b1 := b >> 7
	return b1<<2 + g1<<1 + r1
}

func (a *ansiFormat8) String(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) string {
	enc := []string{}
	if fg != None {
		enc = append(enc, fmt.Sprintf("%d", 30+rgbTo8(fg)))
	}
	if bg != None {
		enc = append(enc, fmt.Sprintf("%d", 40+rgbTo8(bg)))
	}

	return ansiPrinter(enc, flags)
}

func (a *ansiFormat8) Bytes(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) []byte {
	f := a.String(fg, bg, flags...)
	return func(format string, a ...interface{}) []byte {
		return []byte(f(format, a...))
	}
}
