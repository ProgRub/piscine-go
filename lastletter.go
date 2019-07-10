package piscine

func LastRune(s string) rune {
	aux := []rune(s)
	return aux[len(aux)-1]
}
