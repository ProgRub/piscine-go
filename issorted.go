package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	for i := 0; i < len(tab); i++ {
		for j := i + 1; j < len(tab); j++ {
			if f(tab[i], tab[j]) > 0 {
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
