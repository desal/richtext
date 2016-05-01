package richtext_test

import (
	"fmt"
	"testing"

	"github.com/desal/richtext"
)

func TestSyntax(t *testing.T) {
	format := richtext.Ansi256()
	warning := format.String(richtext.DarkRed, richtext.Black, richtext.Bold)
	fmt.Printf(warning("WARNING"))
	fmt.Printf(" Yo\n")
}
