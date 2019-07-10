package piscine

func Index(s string, toFind string) int {
	auxString := []rune(s)
	auxToFind := []rune(toFind)
	pos := 0
	beginning := pos
	for i := range auxString {
		if auxToFind[pos] == auxString[i] {
			if pos == 0 {
				beginning = i
			}
			pos++
		}
		if pos >= len(auxToFind) {
			return beginning
		}
	}
	return -1
}
