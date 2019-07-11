package piscine

func IsLower(str string) bool {
	aux := []rune(str)
	for i := range aux {
		if !(int(aux[i]) >= 97 && int(aux[i]) <= 122) {
			return false
		}
	}
	return true
}
