package piscine

import (
	"math"
)

func BasicAtoi(s string) int {
	aux := []rune(s)
	tam := len(aux) - 1
	i := tam
	final := 0
	for index := 0; index <= tam; index++ {
		final = final + int(math.Pow(10, float64(i)))*int(aux[index]-48)
		i--
	}
	return final

}
