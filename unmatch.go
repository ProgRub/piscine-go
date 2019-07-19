package piscine

import (
	"sort"
)

func Unmatch(arr []int) int {
	sort.Ints(arr)
	for i := 0; i < len(arr); i += 2 {
		if i == len(arr)-1 {
			return arr[i]
		}
		if arr[i] != arr[i+1] {
			return arr[i]
		}
	}
	return -1
}
