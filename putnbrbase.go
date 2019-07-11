package piscine

import (
	"github.com/01-edu/z01"
)

const (
	MinInt = -MaxInt - 1
)

func PrintNbrBase(nbr int, base string) {
	dividor := len([]rune(base))
	aux := []rune(base)
	if len(aux) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i := range aux {
		if aux[i] == '+' || aux[i] == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
	}
	for k := range aux {
		for o := range aux {
			if aux[k] == aux[o] && k != o {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	tricky := false
	carater := 0
	if nbr < 0 {
		z01.PrintRune('-')
		if nbr == MinInt {
			carater = nbr % dividor
			nbr = nbr / dividor
			tricky = true
		}
		nbr = -nbr
	}
	valor := nbr
	tam := 0
	for nbr > 0 {
		nbr = nbr / dividor
		tam++
	}
	if tricky {
		tam++
	}
	final := make([]rune, tam)
	pos := tam - 1
	if tricky {
		final[pos] = rune(-carater)
		pos--
	}
	for valor > 0 {
		carater = valor % dividor
		final[pos] = rune(carater)
		pos--
		valor = valor / dividor
	}
	for i := range final {
		z01.PrintRune(rune(aux[final[i]]))
	}
}
