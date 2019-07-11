package main

import (
	"os"
	"sort"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	if len(arguments) > 0 {
		toOrder := make([]string, len(arguments)-1)
		for i := 1; i < len(arguments); i++ {
			toOrder[i-1] = arguments[i]
		}
		sort.Strings(toOrder)
		for a := range toOrder {
			aux := []rune(toOrder[a])
			for j := range aux {
				z01.PrintRune(aux[j])
			}
			z01.PrintRune('\n')
		}
	}
}
