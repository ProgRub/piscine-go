package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr (n int) {	
	if n==0{
		z01.PrintRune(48)
	}else{
		conta:=0
		aux:=n
		for aux!=0{
			conta++
			aux=aux/10
		}
		tam:=conta
		vet:=make([]int,tam)
		conta--
		if n>0{
			for conta>=0{
				vet[conta]=(n%10)+48
				n=n/10
				conta--
			}
			conta=0
			for conta<tam{
				z01.PrintRune(rune(vet[conta]))
				conta++
			}
		}else{
			n=n*(-1)
			z01.PrintRune(45)
			for conta>=0{
				vet[conta]=(n%10)+48
				n=n/10
				conta--
			}
			conta=0
			for conta<tam{
				z01.PrintRune(rune(vet[conta]))
				conta++
			}
		}
	}
	z01.PrintRune('\n')
}