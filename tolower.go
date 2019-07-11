package piscine

func ToLower(s string) string {
	aux := []rune(s)
	for i := range aux {
		if int(aux[i]) >= 65 && int(aux[i]) <= 90 {
			aux[i] = aux[i] + 32
		}
	}
	return string(aux)
}
