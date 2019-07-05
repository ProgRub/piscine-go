package main

import "fmt"

func StrLen(str string) int {
	return len([]rune(str))
}

func main() {
	str := "Hèllo!"
	nb := StrLen(str)
	fmt.Println(nb)
}
