package piscine

func IterativePower(nb, power int) int {
	if power > 0 {
		result := 1
		for i := 1; i <= power; i++ {
			result = result * nb
		}
		return result
	}
	return 0
}
