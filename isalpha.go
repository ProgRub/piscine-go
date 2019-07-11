package piscine

func IsAlpha(str string) bool {
	aux := []rune(str)
	for i := range aux {
		if !((int(aux[i]) >= 48 && int(aux[i]) <= 57) || (int(aux[i]) >= 65 && int(aux[i]) <= 90) || (int(aux[i]) >= 97 && int(aux[i]) <= 122)) {
			return false
		}
	}
	return true
}
