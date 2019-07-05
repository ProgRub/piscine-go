package piscine

import "fmt"

func PrintStr(str string) {
	for _, ch := range str {
		fmt.Printf("%v", string(ch))
	}
	fmt.Printf("\n")
}
