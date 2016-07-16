// +build !windows

package richtext

func isWindowsConsole() bool { return false }
func Console() *int          { return nil }
