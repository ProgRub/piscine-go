package piscine

func IsNumeric(str string) bool {
	aux := []rune(str)
	for i := range aux {
		if !(int(aux[i]) >= 48 && int(aux[i]) <= 57) {
			return false
		}
	}
	return true
}
