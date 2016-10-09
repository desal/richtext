package richtext

import (
	"fmt"
	"io"
)

type DebugFormat struct {
	w io.Writer
}

var _ Format = &DebugFormat{}

func Debug(w io.Writer) *DebugFormat {
	f := &DebugFormat{w}
	return f
}

func (f *DebugFormat) MakePrintf(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) (int, error) {
	return func(format string, a ...interface{}) (int, error) {
		s := fmt.Sprintf(format, a...)
		f.w.Write([]byte(fmt.Sprintf("[%v,%v,%v]", fg, bg, flags)))
		f.w.Write([]byte(fmt.Sprintf(format, a...)))
		f.w.Write([]byte("[]"))
		return len(s), nil
	}
}

func (f *DebugFormat) PrintLine(format string, a ...interface{}) {
	f.w.Write([]byte(fmt.Sprintf(format, a...)))
	f.w.Write([]byte("\n"))
}

func (f *DebugFormat) WarningLine(format string, a ...interface{}) {
	f.w.Write([]byte("[WARN]"))
	f.w.Write([]byte(fmt.Sprintf(format, a...)))
	f.w.Write([]byte("[]\n"))
}

func (f *DebugFormat) ErrorLine(format string, a ...interface{}) {
	f.w.Write([]byte("[ERR]"))
	f.w.Write([]byte(fmt.Sprintf(format, a...)))
	f.w.Write([]byte("[]\n"))
}
