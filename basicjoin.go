package piscine

func BasicJoin(strs []string) string {
	comprimento := 0
	for i := range strs {
		aux := []rune(strs[i])
		comprimento = comprimento + len(aux)
	}
	final := make([]rune, comprimento)
	pos := 0
	for i := range strs {
		aux := []rune(strs[i])
		for j := range aux {
			final[pos] = aux[j]
			pos++
		}
	}
	return string(final)
}
