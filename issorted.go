package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	return menorMaior(f, tab) || maiorMenor(f, tab)
}

func menorMaior(f func(a, b int) int, tab []int) bool {
	for i := 0; i < len(tab)-1; i++ {
		if f(tab[i], tab[i+1]) > 0 {
			return false
		}
	}
	return true
}

func maiorMenor(f func(a, b int) int, tab []int) bool {
	for i := len(tab) - 1; i > 0; i-- {
		if f(tab[i], tab[i-1]) > 0 {
			return false
		}
	}
	return true
}
func comp(a, b int) int {
	if a > b {
		return 15
	} else if a < b {
		return -15
	} else {
		return 0
	}
}
