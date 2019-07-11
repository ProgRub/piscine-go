package piscine

func ToUpper(s string) string {
	aux := []rune(s)
	for i := range aux {
		if int(aux[i]) >= 97 && int(aux[i]) <= 122 {
			aux[i] = aux[i] - 32
		}
	}
	return string(aux)
}
