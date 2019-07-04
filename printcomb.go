package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb(){
i:=0
j:=0
k:=0
decI:=48
decJ:=48
decK:=48
for i<10{
if i < j && j < k{
	z01.PrintRune(rune(decI))
	z01.PrintRune(rune(decJ))
	z01.PrintRune(rune(decK))
	z01.PrintRune(',')
	z01.PrintRune(' ')
}
k++
decK++
if k==10{
	k=j
	decK=decJ
	j++
	decJ++
}
if j==10{
	j=i
	decJ=decI
	i++
	decI++
}
}
}