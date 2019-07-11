package piscine

func BasicJoin(strs []string) string {
	var final string = strs[0]
	for i := range strs {
		if i != 0 {
			final = Concat(final, strs[i])
		}
	}
	return final
}
