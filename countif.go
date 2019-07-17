package piscine

func CountIf(f func(string) bool, tab []string) int {
	conta := 0
	for i := range tab {
		if f(tab[i]) {
			conta++
		}
	}
	return conta
}
