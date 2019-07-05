package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	vet := make([]int, n)
	indice := 0
	cod := 48
	for indice < n {
		vet[indice] = cod
		indice++
		cod++
	}
	pos := n - 1
	aux(n, vet, pos)
	imprime(vet, n)
	z01.PrintRune('\n')
}

func aux(n int, v []int, pos int) {
	if v[0] == (58 - n) {
		return
	} else {
		imprime(v, n)
		z01.PrintRune(',')
		z01.PrintRune(' ')
		if v[pos] < 57 {
			v[pos]++
			aux(n, v, pos)
		} else {
			pos--
			if pos >= 0 {
				for pos >= 0 && v[pos] == (57-(n-pos-1)) {
					pos--
				}
				v[pos]++
			}
			for pos < n-1 && pos >= 0 {
				v[pos+1] = v[pos] + 1
				pos++
			}
			aux(n, v, pos)
		}
	}
}

func imprime(v []int, tam int) {
	indice := 0
	for indice < tam {
		z01.PrintRune(rune(v[indice]))
		indice++
	}
}
