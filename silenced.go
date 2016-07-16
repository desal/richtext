package richtext

type SilencedFormat struct{}

var _ Format = &SilencedFormat{}

var silencedFormat *SilencedFormat

func Silenced() *SilencedFormat {
	return silencedFormat
}

func (*SilencedFormat) MakePrintf(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) (int, error) {
	return func(format string, a ...interface{}) (int, error) { return 0, nil }
}

func (*SilencedFormat) PrintLine(format string, a ...interface{}) {}

func (*SilencedFormat) WarningLine(format string, a ...interface{}) {}

func (*SilencedFormat) ErrorLine(format string, a ...interface{}) {}
