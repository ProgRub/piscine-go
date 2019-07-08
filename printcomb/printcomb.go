package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb(){
for i:=48;i<58;i++{
	for j:=48;j<58;j++{
		for k:=48;k<58;k++{
			if i==55 && j==56 && k==57 {
				z01.PrintRune(rune(i))
				z01.PrintRune(rune(j))
				z01.PrintRune(rune(k))	
				z01.PrintRune('\n')			
			} else if i < j && j < k{
				z01.PrintRune(rune(i))
				z01.PrintRune(rune(j))
				z01.PrintRune(rune(k))
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
}