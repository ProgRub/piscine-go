package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	if len(arguments) > 0 {
		for i := 1; i < len(arguments); i++ {
			aux := []rune(arguments[i])
			for j := range aux {
				z01.PrintRune(aux[j])
			}
			z01.PrintRune('\n')
		}
	}
}
