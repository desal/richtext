// +build !windows

package richtext

func isWindowsConsole() bool   { return false }
func Console() *SilencedFormat { return nil }
