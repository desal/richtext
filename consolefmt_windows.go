package richtext

import (
	"fmt"
	"syscall"
	"unsafe"
)

type ConsoleFormat struct {
	hStdOut          uintptr
	isWindowsConsole bool
}

type (
	_BY_HANDLE_FILE_INFORMATION struct {
		dwFileAttributes     uint32
		ftCreationTime       _FILETIME
		ftLastAccessTime     _FILETIME
		ftLastWriteTime      _FILETIME
		dwVolumeSerialNumber uint32
		nFileSizeHigh        uint32
		nFileSizeLow         uint32
		nNumberOfLinks       uint32
		nFileIndexHigh       uint32
		nFileIndexLow        uint32
	}

	_FILETIME struct {
		dwLowDateTime  uint32
		dwHighDateTime uint32
	}

	_CONSOLE_SCREEN_BUFFER_INFO struct {
		dwSize              _COORD
		dwCursorPosition    _COORD
		wAttributes         uint16
		srWindow            _SMALL_RECT
		dwMaximumWindowSize uint16
	}

	_COORD struct {
		X int16
		Y int16
	}

	_SMALL_RECT struct {
		Left   int16
		Top    int16
		Bottom int16
		Right  int16
	}
)

const (
	_COMMON_LVB_UNDERSCORE        = 0x8000
	_FOREGROUND_INTENSITY         = 0x0008
	_BACKGROUND_INTENSITY         = 0x0080
	_STD_OUTPUT_HANDLE            = 0xFFFFFFF5
	fgMask                 uint16 = 0x000F
	bgMask                 uint16 = 0x00F0
)

var (
	kernel32                          = syscall.NewLazyDLL("kernel32.dll")
	getStdHandle                      = kernel32.NewProc("GetStdHandle")
	getFileInformationByHandle        = kernel32.NewProc("GetFileInformationByHandle")
	setConsoleTextAttribute           = kernel32.NewProc("SetConsoleTextAttribute")
	getConsoleScreenBufferInfo        = kernel32.NewProc("GetConsoleScreenBufferInfo")
	writeConsole                      = kernel32.NewProc("WriteConsoleA")
	_                          Format = &ConsoleFormat{}
	consoleFormat              *ConsoleFormat
)

func init() {
	hStdOut, _, _ := getStdHandle.Call(uintptr(_STD_OUTPUT_HANDLE), 0)
	var fi _BY_HANDLE_FILE_INFORMATION
	_, _, err := getFileInformationByHandle.Call(hStdOut, uintptr(unsafe.Pointer(&fi)), 0)
	errno := err.(syscall.Errno)
	consoleFormat = &ConsoleFormat{hStdOut: hStdOut, isWindowsConsole: errno == 1}
}

func isWindowsConsole() bool { return consoleFormat.isWindowsConsole }

func (c *ConsoleFormat) setConsoleTextAttribute(attributes uint16) {
	/*BOOL WINAPI SetConsoleTextAttribute(
	  _In_ HANDLE hConsoleOutput,
	  _In_ WORD   wAttributes
	);*/
	ok, _, err := setConsoleTextAttribute.Call(c.hStdOut, uintptr(attributes), 0)
	if ok == 0 {
		panic(err)
	}
}

func Console() *ConsoleFormat { return consoleFormat }

func (c *ConsoleFormat) rgbTo8(color Color) uint8 {
	r, g, b := rgb(color)
	r1 := r >> 7
	g1 := g >> 7
	b1 := b >> 7
	return r1<<2 + g1<<1 + b1
}

func (c *ConsoleFormat) makeAttr(currentAttributes uint16, fg, bg Color, flags ...Flag) uint16 {
	newAttributes := currentAttributes
	if fg != None {
		newAttributes = (newAttributes & ^fgMask) + uint16(c.rgbTo8(fg))

		for _, flag := range flags {
			switch flag {
			case Bold:
				newAttributes += _FOREGROUND_INTENSITY
			case Underline:
				newAttributes += _COMMON_LVB_UNDERSCORE
			}
		}
	}
	if bg != None {
		newAttributes = (newAttributes & ^bgMask) + (uint16(c.rgbTo8(bg)) << 4)
	}
	return newAttributes
}

func (c *ConsoleFormat) consoleScreenBufferInfo() _CONSOLE_SCREEN_BUFFER_INFO {
	var csbiInfo _CONSOLE_SCREEN_BUFFER_INFO
	ok, _, err := getConsoleScreenBufferInfo.Call(c.hStdOut, uintptr(unsafe.Pointer(&csbiInfo)), 0)
	if ok == 0 {
		panic(err)
	}
	return csbiInfo
}

func (c *ConsoleFormat) writeConsole(buf []byte) (int, error) {
	if len(buf) > 0xFFFF {
		return 0, fmt.Errorf("writeConsole limit of 65535 bytes")
	}
	var written uint32
	ok, _, err := writeConsole.Call(c.hStdOut, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)), uintptr(unsafe.Pointer(&written)), 0, 0)
	if ok == 0 {
		return 0, err
	}
	return int(written), nil
}

func (c *ConsoleFormat) MakePrintf(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) (int, error) {
	return func(format string, a ...interface{}) (int, error) {
		csbiInfo := c.consoleScreenBufferInfo()
		c.setConsoleTextAttribute(c.makeAttr(csbiInfo.wAttributes, fg, bg, flags...))
		n, err := c.writeConsole([]byte(fmt.Sprintf(format, a...)))
		c.setConsoleTextAttribute(csbiInfo.wAttributes)
		return n, err
	}
}

func (c *ConsoleFormat) WarningLine(format string, a ...interface{}) {
	csbiInfo := c.consoleScreenBufferInfo()
	c.setConsoleTextAttribute(c.makeAttr(csbiInfo.wAttributes, Orange, None, Bold))
	fmt.Printf(format, a...)
	c.setConsoleTextAttribute(csbiInfo.wAttributes)
}

func (c *ConsoleFormat) ErrorLine(format string, a ...interface{}) {
	csbiInfo := c.consoleScreenBufferInfo()
	c.setConsoleTextAttribute(c.makeAttr(csbiInfo.wAttributes, Red, None, Bold))
	fmt.Printf(format, a...)
	c.setConsoleTextAttribute(csbiInfo.wAttributes)
}

func (c *ConsoleFormat) PrintLine(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	fmt.Print("\n")
}

func (c *ConsoleFormat) Print(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
