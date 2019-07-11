package piscine

import "strings"

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	/* auxString := []rune(s)
	auxToFind := []rune(toFind)
	pos := 0
	beginning := pos
	for i := range auxString {
		if auxToFind[pos] == auxString[i] && pos < len(auxToFind) {
			if pos == 0 {
				beginning = i
			}
			pos++
		}
		if pos >= len(auxToFind) {
			return beginning
		}
	}
	return -1 */
	return strings.Index(s, toFind)
}
