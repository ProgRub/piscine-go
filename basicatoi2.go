package piscine

import (
	"math"
)

func BasicAtoi2(s string) int {
	aux := []rune(s)
	tam := len(aux) - 1
	i := tam
	final := 0
	for index := 0; index <= tam; index++ {
		if !(int(aux[index]) > 47 && int(aux[index]) < 58) {
			return 0
		}
	}
	for index := 0; index <= tam; index++ {
		if int(aux[index]) > 47 && int(aux[index]) < 58 {
			final = final + int(math.Pow(10, float64(i)))*int(aux[index]-48)
		}
		i--
	}
	return final

}
