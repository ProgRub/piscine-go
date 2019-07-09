package piscine

import (
	"math/bits"
)

const (
	IntSize = bits.UintSize - 1
	MaxInt  = 1<<IntSize - 1
)

func FindNextPrime(nb int) int {
	if nb > 1 && nb < MaxInt {
		if nb > 3 {
			i := 0
			if nb%2 == 0 {
				i = nb + 1
			} else {
				i = nb
			}
			for true {
				if i < MaxInt {
					if IsPrime(i) {
						return i
					}
					i = i + 2
				} else {
					return 0
				}
			}
		}
		return nb
	}
	return 2
}
