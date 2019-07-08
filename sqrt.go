package piscine

func Sqrt(nb int) int {
	if nb >= 0 {
		aux := 0
		for aux*aux < nb {
			aux++
		}
		if aux*aux == nb {
			return aux
		}
		return 0
	}
	return 0
}
