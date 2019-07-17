package piscine

func Any(f func(string) bool, arr []string) bool {
	for i := range arr {
		if f(arr[i]) {
			return true
		}
	}
	return false
}
