package piscine

func FindNextPrime(nb int) int {
	if nb > 1 && nb < 9223372036854775807 {
		i := nb
		for i > 0 {
			if nb < 9223372036854775807 {
				if IsPrime(i) {
					return i
				}
				i++
			} else {
				return 0
			}
		}
	}
	return 2
}
