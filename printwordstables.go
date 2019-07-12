package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(table []string) {
	toPrint := []rune(Join(table, "\n"))
	for i := range toPrint {
		z01.PrintRune(toPrint[i])
	}
	z01.PrintRune('\n')
}
