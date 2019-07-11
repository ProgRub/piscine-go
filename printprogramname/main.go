package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	if len(arguments) > 0 {
		aux := []rune(arguments[0])
		for i := range aux {
			z01.PrintRune(aux[i])
		}
		z01.PrintRune('\n')
	}
}
