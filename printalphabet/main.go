package main

import "github.com/01-edu/z01"

func main (){
	var aRune rune = 'a'
	for  ;aRune <= 'z'; aRune++{
		z01.PrintRune(aRune)
	}
	z01.PrintRune('\n')
}