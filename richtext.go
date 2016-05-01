//Simple package for generating rich text, the canonical example of which is
//formatted text on console.
package richtext

//	"math"

type Format interface {
	String(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) string
	Bytes(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) []byte
}

type Color uint32

const (
	None Color = 1 << 24
)

func rgb(c Color) (uint8, uint8, uint8) {
	return uint8((c >> 16) & 0xFF), uint8((c >> 8) & 0xFF), uint8((c) & 0xFF)
}

//func RGB(r, g, b uint8) color { return color(uint32(b) + 0x100*uint32(g) + 0x10000*uint32(r)) }

/*
Some usage examples
format := richtext.Ansi256(...)
warning := format.Bytes(r, g, b, richtext.Bold)
warning := format.Bytes(richtext.RGB{r, g, b}, richtext.RGB{r, g, b}, richtext.Bold)

*/

type Flag int

const (
	Bold Flag = iota
	Underline
)

//func rgbTo16(r, g, b uint8) uint8 {

//}

//type ansiFormat16 struct{}
