package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		var in string
		_, err := fmt.Scanln(&in)
		if err != nil {
			panic(err)
		}
		fmt.Println(in)
	} else {
		for i := 1; i < len(arguments); i++ {
			ficheiro, err := os.Open(arguments[i])
			if err != nil {
				panic(err)
			}
			aux, ai := ficheiro.Stat()
			if ai != nil {
				panic(ai)
			}
			tam := aux.Size()
			texto := make([]byte, tam)
			ficheiro.Read(texto)
			fmt.Println(string(texto))
			fmt.Println("")
			ficheiro.Close()
		}
	}
}
