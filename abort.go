package piscine

import "sort"

func Abort(a, b, c, d, e int) int {
	aux := [5]int{a, b, c, d, e}
	sort.Ints(aux[:])
	return aux[2]
}
