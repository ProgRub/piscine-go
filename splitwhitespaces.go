package piscine

func SplitWhiteSpaces(str string) []string {
	aux := []rune(str)
	final := make([]string, 0)
	palavra := make([]rune, 0)
	for i := range aux {
		if int(aux[i]) != int(' ') && int(aux[i]) != int('\t') && int(aux[i]) != int('\n') {
			palavra = append(palavra, aux[i])
		} else {
			final = append(final, string(palavra))
			palavra = make([]rune, 0)
		}
	}
	final = append(final, string(palavra))
	return final
}
