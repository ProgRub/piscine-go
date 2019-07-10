package piscine

func NRune(s string, n int) rune {
	aux := []rune(s)
	if n > 0 && n <= len(aux) {
		return aux[n-1]
	}
	return 95
}
