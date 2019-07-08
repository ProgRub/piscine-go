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
		i := nb
		for i > 0 {
			if i < MaxInt {
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
