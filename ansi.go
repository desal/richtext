package richtext

import (
	"fmt"
	"os"
	"strings"
)

const ansiReset = "\x1b[0m"
const ansiBold = "1"
const ansiUnderline = "4"

func ansiString(s []string) string {
	return "\x1b[" + strings.Join(s, ";") + "m"
}

func ansiFlags(flags []Flag) []string {
	result := []string{}
	for _, flag := range flags {
		switch flag {
		case Bold:
			result = append(result, ansiBold)
		case Underline:
			result = append(result, ansiUnderline)
		}
	}
	return result
}

func ansiPrinter(enc []string, flags []Flag) func(format string, a ...interface{}) string {
	ansiEnc := append(enc, ansiFlags(flags)...)
	return func(format string, a ...interface{}) string {
		return ansiString(ansiEnc) + fmt.Sprintf(format, a...) + ansiReset
	}
}

func Ansi() Format {
	term := os.Getenv("TERM")
	if strings.Contains(term, "24b") {
		return Ansi24b()
	} else if strings.Contains(term, "256") {
		return Ansi256()
	}
	return Ansi8()
}
