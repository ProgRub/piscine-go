package piscine

func IsPrime(nb int) bool {
	if nb > 1 {
		if nb > 3 {
			if nb%2 == 0 {
				return false
			}

			for i := 5; i*i < nb; i = i + 6 {
				if nb%i == 0 || nb%(i+2) == 0 {
					return true
				}
			}
			return false
		}
		return true
	}
	return false
}
