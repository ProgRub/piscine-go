package piscine

import (
	"math"
	"strings"
)

func AtoiBase(s string, base string) int {
	resultado := 0
	aux := []rune(base)
	if len(aux) < 2 {
		return 0
	}
	for i := range aux {
		if aux[i] == '+' || aux[i] == '-' {
			return 0
		}
	}
	for k := range aux {
		for o := range aux {
			if aux[k] == aux[o] && k != o {
				return 0
			}
		}
	}
	auxNumero := []rune(s)
	auxBase := []rune(base)
	numBase := len(auxBase)
	potencia := len(auxNumero) - 1
	for i := range auxNumero {
		resultado = resultado + strings.Index(base, string(auxNumero[i]))*int(math.Pow(float64(numBase), float64(potencia)))
		potencia--
	}
	return resultado
}
