package piscine

func Concat(str1 string, str2 string) string {
	auxStr1 := []rune(str1)
	auxStr2 := []rune(str2)
	final := make([]rune, len(auxStr1)+len(auxStr2))
	pos := 0
	for i := range auxStr1 {
		final[pos] = auxStr1[i]
		pos++
	}
	for j := range auxStr2 {
		final[pos] = auxStr2[j]
		pos++
	}
	return string(final)
}
