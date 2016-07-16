package richtext

import (
	"fmt"
)

type AsciiFormat struct{ ansiFormat }

var asciiFormat *AsciiFormat

var _ Format = &AsciiFormat{}

func init() {
	asciiFormat := &AsciiFormat{}
	asciiFormat.init(asciiFormat)
}

func Ascii() *AsciiFormat {
	return asciiFormat
}

func (a *AsciiFormat) MakePrintf(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) (int, error) {
	return fmt.Printf
}

func (a *AsciiFormat) MakeSprintf(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) string {
	return fmt.Sprintf
}
