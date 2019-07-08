package piscine

import (
	"math"
)

func Atoi(s string) int {
	aux := []rune(s)
	tam := len(aux) - 1
	i := tam
	final := 0
	negativo := false
	for index := 0; index <= tam; index++ {
		if index > 0 && !(int(aux[index]) > 47 && int(aux[index]) < 58) {
			return 0
		}
	}
	for index := 0; index <= tam; index++ {
		if int(aux[index]) > 47 && int(aux[index]) < 58 {
			final = final + int(math.Pow(10, float64(i)))*int(aux[index]-48)
		} else if int(aux[index]) == 45 {
			negativo = true
		}
		i--
	}
	if negativo {
		return -final
	}
	return final

}
