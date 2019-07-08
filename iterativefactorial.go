package piscine

import "fmt"

func IterativeFactorial(nb int) int {
	if nb >= 0 {
		result := 1
		for nb > 0 {
			result = result * nb
			nb--
		}
		return result
	}
	return 0
