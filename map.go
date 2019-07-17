package piscine

func Map(f func(int) bool, arr []int) []bool {
	final := make([]bool, len(arr))
	for i := range arr {
		final[i] = f(arr[i])
	}
	return final
}
