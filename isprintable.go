package piscine

func IsPrintable(str string) bool {
	aux := []rune(str)
	for i := range aux {
		if int(aux[i]) <= 31 {
			return false
		}
	}
	return true
}
