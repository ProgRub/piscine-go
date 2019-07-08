package piscine

func StrRev(s string) string {
	aux := []rune(s)
	final := make([]rune, len(aux))
	i := 0
	for index := len(aux) - 1; index >= 0; index-- {
		final[i] = aux[index]
		i++
	}
	return string(final)
}
