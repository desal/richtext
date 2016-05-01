package richtext

import (
	"fmt"
)

type asciiFormat struct{}

//This varies by platform, this is what windows has, the most likely source of us
//running a 16 color pallette

func Ascii() Format {
	return &asciiFormat{}
}

func (a *asciiFormat) String(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) string {
	return fmt.Sprintf
}

func (a *asciiFormat) Bytes(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) []byte {
	return func(format string, a ...interface{}) []byte {
		return []byte(fmt.Sprintf(format, a...))
	}
}
