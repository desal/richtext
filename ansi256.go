package richtext

import (
	"fmt"
	"math"
)

type AnsiFormat256 struct{ ansiFormat }

var _ Format = &AnsiFormat256{}

var ansiFormat256 *AnsiFormat256

func init() {
	ansiFormat256 := &AnsiFormat256{}
	ansiFormat256.init(ansiFormat256)
}
func Ansi256() *AnsiFormat256 {
	return ansiFormat256
}

var colour6b = [...]uint8{0, 95, 135, 175, 215, 255}
var boundary6b = [...]uint8{47, 115, 155, 195, 235, 255}

func to6b(c uint8) uint8 {
	for i, boundary := range boundary6b {
		if c <= boundary {
			return uint8(i)
		}
	}
	return 0
}

//Conversions are all completely incorrect, based on entirely false assumptions
//about orthogonality of RGB.
func rgbTo248(c Color) uint8 {
	r, g, b := rgb(c)

	ri := to6b(r)
	gi := to6b(g)
	bi := to6b(b)

	if ri == gi && ri == bi {
		rgbAvg := (float64(r) + float64(g) + float64(b)) / 3.0
		gr24 := math.Floor((rgbAvg-8)/10.0 + 0.5)
		if gr24 >= 0 && gr24 <= 23 && math.Abs(gr24*10+8-rgbAvg) < math.Abs(float64(colour6b[ri])-rgbAvg) {
			return 232 + uint8(gr24)
		}
	}
	return 16 + ri*36 + gi*6 + bi
}

func (a *AnsiFormat256) MakePrintf(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) (int, error) {
	enc := []string{}
	if fg != None {
		enc = append(enc, fmt.Sprintf("38;5;%d", rgbTo248(fg)))
	}
	if bg != None {
		enc = append(enc, fmt.Sprintf("48;5;%d", rgbTo248(bg)))
	}
	return ansiPrinter(enc, flags)
}

func (a *AnsiFormat256) MakeSprintf(fg, bg Color, flags ...Flag) func(format string, a ...interface{}) string {
	enc := []string{}
	if fg != None {
		enc = append(enc, fmt.Sprintf("38;5;%d", rgbTo248(fg)))
	}
	if bg != None {
		enc = append(enc, fmt.Sprintf("48;5;%d", rgbTo248(bg)))
	}
	return ansiSPrinter(enc, flags)
}
