package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		final := make([]int, 0)
		return final
	}
	final := make([]int, max-min)
	nums := min
	for i := range final {
		final[i] = nums
		nums++
	}
	return final
}
