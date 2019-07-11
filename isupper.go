package piscine

func IsUpper(str string) bool {
	aux := []rune(str)
	for i := range aux {
		if !(int(aux[i]) >= 65 && int(aux[i]) <= 90) {
			return false
		}
	}
	return true
}
