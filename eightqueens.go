package piscine

import (
	"github.com/01-edu/z01"
)

func EightQueens() {
	var NaoPodeTabuleiro [8]int
	var solucao [8]int
	solucao = [...]int{0, 0, 0, 0, 0, 0, 0, 0}
	for i := 1; i <= 8; i++ {
		solucao[0] = i
		NaoPodeTabuleiro[0] = i
		verificaSolucoes(NaoPodeTabuleiro, solucao, 1)
	}
}

func verificaSolucoes(vetorNaoPode [8]int, VetorNumeroAtual [8]int, numeroAtual int) {
	if numeroAtual == 8 { //Ultima posicao do array
		imprime := true
		for i := 0; i <= 7; i++ {
			if VetorNumeroAtual[i] == 0 {
				imprime = false
			}
		}
		if imprime {
			final(VetorNumeroAtual)
		}
	} else {
		var auxiliar [8]int
		auxiliar = criaVetorNaoPode(vetorNaoPode, numeroAtual)
		for i := 0; i <= 8; i++ {
			if verificaSePodeInserir(auxiliar, i) {
				VetorNumeroAtual[numeroAtual] = i
				vetorNaoPode[numeroAtual] = i
				verificaSolucoes(vetorNaoPode, VetorNumeroAtual, numeroAtual+1)
			}
		}
	}

}

func final(resposta [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(resposta[i] + 48))
	}
	z01.PrintRune('\n')
}

func verificaSePodeInserir(vetorNaoPode [8]int, numero int) bool {
	for i := 0; i < 8; i++ {
		if vetorNaoPode[i] == numero {
			return false
		}
	}
	return true
}

func criaVetorNaoPode(vetorNaoPode [8]int, posicao int) [8]int {
	var horizontais [8]int //Diagonal Descendente
	horizontais = [...]int{0, 0, 0, 0, 0, 0, 0, 0}
	var diagonais [8]int //Diagonal Ascendente
	diagonais = [...]int{0, 0, 0, 0, 0, 0, 0, 0}
	var final [8]int
	final = [...]int{0, 0, 0, 0, 0, 0, 0, 0}
	for e := 0; e <= 7; e++ {
		if vetorNaoPode[e] != 0 {
			horizontais[e] = vetorNaoPode[e]
			diagonais[e] = vetorNaoPode[e]
			final[e] = vetorNaoPode[e]
		}
	}
	for e := 0; e <= posicao; e++ {
		if horizontais[e] != 0 {
			if ((posicao - e) + vetorNaoPode[e]) <= 8 {
				horizontais[e] = (posicao - e) + vetorNaoPode[e]
			} else {
				horizontais[e] = 0
			}
		}

	}
	for e := 0; e <= posicao; e++ {
		if diagonais[e] != 0 {
			if (-(posicao - e) + vetorNaoPode[e]) >= 0 {
				diagonais[e] = -(posicao - e) + vetorNaoPode[e]
			} else {
				diagonais[e] = 0
			}
		}

	}
	for e := 0; e <= 7; e++ {
		if final[e] == 0 {
			for i := 0; i <= 7; i++ {
				if !numeroEstaVetor(final, horizontais[i]) && e < 8 {
					final[e] = horizontais[i]
					e++
				}
				if !numeroEstaVetor(final, diagonais[i]) && e < 8 {
					final[e] = diagonais[i]
					e++
				}
			}
		}
	}
	return final
}

func numeroEstaVetor(vetor [8]int, numero int) bool {
	for i := 0; i < 8; i++ {
		if vetor[i] == numero {
			return true
		}
	}
	return false
}
