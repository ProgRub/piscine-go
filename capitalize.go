package piscine

func Capitalize(s string) string {
	aux := []rune(s)
	espacoAnterior := false
	maiusculaAnterior := false
	for i := range aux {
		if maiusculaAnterior {
			if int(aux[i]) >= 65 && int(aux[i]) <= 90 {
				aux[i] = aux[i] + 32
			}
		}
		if espacoAnterior || i == 0 {
			if int(aux[i]) >= 97 && int(aux[i]) <= 122 {
				aux[i] = aux[i] - 32
				maiusculaAnterior = true
			} else if int(aux[i]) >= 48 && int(aux[i]) <= 57 {
				maiusculaAnterior = true
			}
			espacoAnterior = false
		}
		if aux[i] == 32 || !((int(aux[i]) >= 97 && int(aux[i]) <= 122) || (int(aux[i]) >= 65 && int(aux[i]) <= 90) || (int(aux[i]) >= 48 && int(aux[i]) <= 57)) {
			espacoAnterior = true
			maiusculaAnterior = false
		}
	}
	return string(aux)
}
