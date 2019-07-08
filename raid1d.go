package piscine

import "fmt"

func Raid1d(x, y int) {
	if x > 0 {
		edges := make([]rune, x)
		middle := make([]rune, x)
		for contaX := 0; contaX < x; contaX++ {
			if contaX == 0 {
				edges[contaX] = 'A'
			} else if contaX == x-1 {
				edges[contaX] = 'C'
			} else {
				edges[contaX] = 'B'
			}
		}
		for contaX := 0; contaX < x; contaX++ {
			if contaX == 0 || contaX == x-1 {
				middle[contaX] = 'B'
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
}
