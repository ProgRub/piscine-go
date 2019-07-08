package piscine

func FindNextPrime(nb int) int {
	if nb > 1 {
		i := nb
		for i > 0 {
			if IsPrime(i) {
				return i
			}
			i++
		}
	}
	return 2
}
