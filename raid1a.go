package piscine

import "fmt"

func Raid1a(x, y int) {
	edges := make([]rune, x)
	middle := make([]rune, x)
	for contaX := 0; contaX < x; contaX++ {
		if contaX == 0 || contaX == x-1 {
			edges[contaX] = 'o'
		} else {
			edges[contaX] = '-'
		}
	}
	for contaX := 0; contaX < x; contaX++ {
		if contaX == 0 || contaX == x-1 {
			middle[contaX] = '|'
		} else {
			middle[contaX] = ' '
		}
	}
	for contaY := 0; contaY < y; contaY++ {
		if contaY == 0 || contaY == y-1 {
			fmt.Println(string(edges))
		} else {
			fmt.Println(string(middle))
		}
	}

}
