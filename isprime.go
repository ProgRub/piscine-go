package piscine

func IsPrime(nb int) bool {
	if nb > 1 {
		multiplier := nb / 6
		if nb > 3 {
			if nb%2 == 0 {
				return false
			}
			for 6*multiplier-1 <= nb || 6*multiplier+1 <= nb {
				if 6*multiplier-1 == nb || 6*multiplier+1 == nb {
					return true
				}
				multiplier++
			}
			return false
		}
		return true
	}
	return false
}
