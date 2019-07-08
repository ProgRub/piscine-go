package piscine

import "fmt"

func Raid1b(x, y int) {
	if x > 0 {
		edgeUp := make([]rune, x)
		middle := make([]rune, x)
		edgeDown := make([]rune, x)
		for contaX := 0; contaX < x; contaX++ {
			if contaX == 0 {
				edgeUp[contaX] = '/'
			} else if contaX == x-1 {
				edgeUp[contaX] = '\\'
			} else {
				edgeUp[contaX] = '*'
			}
		}
		for contaX := 0; contaX < x; contaX++ {
			if contaX == 0 {
				edgeDown[contaX] = '\\'
			} else if contaX == x-1 {
				edgeDown[contaX] = '/'
			} else {
				edgeDown[contaX] = '*'
			}
		}
		for contaX := 0; contaX < x; contaX++ {
			if contaX == 0 || contaX == x-1 {
				middle[contaX] = '*'
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
