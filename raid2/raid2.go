package piscine

import (
	"fmt"
	"os"
)

func Sudoku() {
	arguments := os.Args
	carateresCorretos := true
	if len(arguments) > 10 {
		carateresCorretos = false
	} else if len(arguments) < 10 {
		carateresCorretos = false
	} else {
		board := [9][9]int{}
		for i := 1; i < 10; i++ {
			for j := 0; j < 9; j++ {
				if arguments[i][j] >= '1' && arguments[i][j] <= '9' {
					board[i-1][j] = int(arguments[i][j] - 48)
				} else if arguments[i][j] != '.' {
					carateresCorretos = false
				}
			}
		}
		if carateresCorretos {
			if !Solucao(board) {
				fmt.Println("Error")
			}
		}
	}
	if !carateresCorretos {
		fmt.Println("Error")
	}
}

func print(board [9][9]int) {
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			if i < 8 {
				fmt.Print(board[j][i], " ")
			} else {
				fmt.Println(board[j][i])
			}
		}
	}
}

func podeColocar(coluna, linha, numero int, board [9][9]int) bool {
	for i := 0; i < 9; i++ {
		if board[linha][i] == numero {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if board[i][coluna] == numero {
			return false
		}
	}
	for i := linha - linha%3; i <= linha-linha%3+2; i++ {
		for j := coluna - coluna%3; j <= coluna-coluna%3+2; j++ {
			if board[i][j] == numero {
				return false
			}
		}
	}
	return true
}

func Solucao(board [9][9]int) bool {
	coluna, linha := 0, 0
	if !posicaoVazia(board, &linha, &coluna) {
		print(board)
		return true
	}
	for i := 1; i <= 9; i++ {
		if podeColocar(coluna, linha, i, board) {
			board[linha][coluna] = i
			if Solucao(board) {
				return true
			}
			board[linha][coluna] = 0
		}
	}
	return false
}

func matrizCheia(board [9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func posicaoVazia(board [9][9]int, row, column *int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				*row = i
				*column = j
				return true
			}
		}
	}
	return false
}

//
