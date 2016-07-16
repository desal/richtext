package richtext

import (
	"fmt"
)

type AnsiFormat8 struct{ ansiFormat }

var _ Format = &AnsiFormat8{}

var ansiFormat8 *AnsiFormat8

func init() {
	ansiFormat8 := &AnsiFormat8{}
	ansiFormat8.init(ansiFormat8)
}

func Ansi8() *AnsiFormat8 { return ansiFormat8 }

func rgbTo8(c Color) uint8 {
	r, g, b := rgb(c)
	r1 := r >> 7
	g1 := g >> 7
	b1 := b >> 7
	return b1<<2 + g1<<1 + r1
}

func (a *AnsiFormat8) MakePrintf(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) (int, error) {
	enc := []string{}
	if fg != None {
		enc = append(enc, fmt.Sprintf("%d", 30+rgbTo8(fg)))
	}
	if bg != None {
		enc = append(enc, fmt.Sprintf("%d", 40+rgbTo8(bg)))
	}

	return ansiPrinter(enc, flags)
}

func (a *AnsiFormat8) MakeSprintf(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) string {
	enc := []string{}
	if fg != None {
		enc = append(enc, fmt.Sprintf("%d", 30+rgbTo8(fg)))
	}
	if bg != None {
		enc = append(enc, fmt.Sprintf("%d", 40+rgbTo8(bg)))
	}

	return ansiSPrinter(enc, flags)
}
