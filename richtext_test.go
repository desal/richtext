package richtext

import (
	"fmt"
	"testing"
)

func TestColors(t *testing.T) {
	f8 := Ansi8()
	f256 := Ansi256()
	f24b := Ansi24b()

	for color, name := range _Color_map {
		c8 := f8.String(color, Black)
		c256 := f256.String(color, Black)
		c24b := f24b.String(color, Black)

		fmt.Printf(" %s %s %s\n", c8("%-18s", name), c256("%-18s", name), c24b("%-18s", name))
	}
}
