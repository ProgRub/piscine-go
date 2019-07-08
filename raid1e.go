package piscine

import "fmt"

func Raid1e(x, y int) {
	if x > 0 {
		edgeUp := make([]rune, x)
		middle := make([]rune, x)
		edgeDown := make([]rune, x)
		for contaX := 0; contaX < x; contaX++ {
			if contaX == 0 {
				edgeUp[contaX] = 'A'
			} else if contaX == x-1 {
				edgeUp[contaX] = 'C'
			} else {
				edgeUp[contaX] = 'B'
			}
		}
		for contaX := 0; contaX < x; contaX++ {
			if contaX == 0 {
				edgeDown[contaX] = 'C'
			} else if contaX == x-1 {
				edgeDown[contaX] = 'A'
			} else {
				edgeDown[contaX] = 'B'
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
			if contaY == 0 {
				fmt.Println(string(edgeUp))
			} else if contaY == y-1 {
				fmt.Println(string(edgeDown))
			} else {
				fmt.Println(string(middle))
			}
		}
	}
}
