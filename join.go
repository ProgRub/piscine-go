package piscine

func Join(strs []string, sep string) string {
	var final string = strs[0]
	for i := range strs {
		if i == 0 {
			final = Concat(final, sep)
		} else if i != len(strs)-1 {
			final = Concat(final, strs[i])
			final = Concat(final, sep)
		} else if i == len(strs)-1 {
			final = Concat(final, strs[i])
		}
	}
	return final
}
