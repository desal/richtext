package richtext

import (
	"fmt"
	"testing"
)

type TestFormat struct{ *testing.T }

var _ Format = &TestFormat{}

func Test(t *testing.T) *TestFormat {
	f := &TestFormat{t}
	return f
}

func (f *TestFormat) MakePrintf(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) (int, error) {
	return func(format string, a ...interface{}) (int, error) {
		s := fmt.Sprintf(format, a...)
		f.T.Log(s)
		return len(s), nil
	}
}

func (f *TestFormat) PrintLine(format string, a ...interface{}) {
	f.Logf(format, a...)
}

func (f *TestFormat) WarningLine(format string, a ...interface{}) {
	f.Logf("WARNING: "+format, a...)
}

func (f *TestFormat) ErrorLine(format string, a ...interface{}) {
	f.Logf("ERROR: "+format, a...)
}

func (f *TestFormat) Print(format string, a ...interface{}) {
	f.Logf(format, a...)
}
