package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	for i := 0; i < len(tab); i++ {
		if f(tab[i], tab[i+1]) > 0 {
			if !IsSorted(f, IntRev(tab)) {
				return false
			}
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

func IntRev(s []int) []int {
	final := make([]int, len(s))
	i := 0
	for index := len(s) - 1; index >= 0; index-- {
		final[i] = s[index]
		i++
	}
	return final
}
