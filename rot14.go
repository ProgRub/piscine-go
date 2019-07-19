package piscine

func Rot14(str string) string {
	aux := []rune(str)
	for i := range aux {
		if (int(aux[i]) >= 65 && int(aux[i]) <= 90) || (int(aux[i]) >= 97 && int(aux[i]) <= 122) {
			if int(aux[i]) > 76 && int(aux[i]) <= 90 {
				aux[i] = 65 + (14 - (90 - aux[i])) - 1
			} else if int(aux[i]) > 108 && int(aux[i]) <= 122 {
				aux[i] = 97 + (14 - (122 - aux[i])) - 1
			} else {
				aux[i] += 14
			}
		}
	}
	return string(aux)
}
