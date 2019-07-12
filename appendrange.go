package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		var final []int
		return final
	}
	var auxilia []int
	final := auxilia[:]
	nums := min
	for i := 0; i < max-min; i++ {
		final = append(final, nums)
		nums++
	}
	return final
}
