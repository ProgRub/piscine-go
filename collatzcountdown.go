package piscine

func CollatzCountdown(start int) int {
	conta := 1
	for start != 1 {
		if start%2 == 0 {
			start /= 2
		} else {
			start = start*3 + 1
		}
		conta++
	}
	return conta
}
