package richtext

import (
	"fmt"
	"strings"
)

type (
	ansiFormat struct {
		warnText string
		errText  string
	}

	AnsiFormat interface {
		Format
		MakeSprintf(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) string
	}
)

const (
	ansiReset     = "\x1b[0m"
	ansiBold      = "1"
	ansiUnderline = "4"
)

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

func ansiSPrinter(enc []string, flags []Flag) func(format string, a ...interface{}) string {
	ansiEnc := append(enc, ansiFlags(flags)...)
	return func(format string, a ...interface{}) string {
		return ansiString(ansiEnc) + fmt.Sprintf(format, a...) + ansiReset
	}
}

func ansiPrinter(enc []string, flags []Flag) func(format string, a ...interface{}) (int, error) {
	ansiEnc := append(enc, ansiFlags(flags)...)
	return func(format string, a ...interface{}) (int, error) {
		fmt.Print(ansiString(ansiEnc))
		n, err := fmt.Printf(format, a...)
		fmt.Print(ansiReset)
		return n, err
	}
}

//Dynamic dispatch in golang eh
func (f *ansiFormat) init(ansi AnsiFormat) {
	f.warnText = ansi.MakeSprintf(Orange, None, Bold)("WARNING: ")
	f.errText = ansi.MakeSprintf(Red, None, Bold)("ERROR: ")
}

func (f *ansiFormat) PrintLine(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	fmt.Print("\n")
}

func (f *ansiFormat) WarningLine(format string, a ...interface{}) {
	fmt.Print(f.warnText)
	fmt.Printf(format, a...)
	fmt.Print("\n")
}

func (f *ansiFormat) ErrorLine(format string, a ...interface{}) {
	fmt.Print(f.errText)
	fmt.Printf(format, a...)
	fmt.Print("\n")
}

func (f *ansiFormat) Print(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
