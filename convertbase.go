package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return PutNbrBase(AtoiBase(nbr, baseFrom), baseTo)
}

func PutNbrBase(nbr int, base string) string {
	dividor := len([]rune(base))
	carater := 0
	valor := nbr
	tam := 0
	for nbr > 0 {
		nbr = nbr / dividor
		tam++
	}
	final := make([]rune, tam)
	pos := tam - 1
	for valor > 0 {
		carater = valor % dividor
		final[pos] = rune(carater)
		pos--
		valor = valor / dividor
	}
	numero := ""
	for i := range final {
		if int(final[i]) >= 48 && int(final[i]) <= 57 {
			numero = Concat(numero, string(final[i]+48))
		} else {
			numero = Concat(numero, string(final[i]))
		}
	}
	return numero
}
