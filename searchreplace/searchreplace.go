package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) != 4 {
		fmt.Println()
		return
	}
	aux := []rune(arguments[1])
	procura := []rune(arguments[2])
	substitui := []rune(arguments[3])
	for i := range aux {
		if aux[i] == procura[0] {
			aux[i] = substitui[0]
		}
	}
	fmt.Println(string(aux))
}
