package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("File name missing")
	} else if len(arguments) > 2 {
		fmt.Println("Too many arguments")
	} else {
		ficheiro, err := os.Open(arguments[1])
		if err != nil {
			fmt.Println(err.Error())
		} else {
			aux, _ := ficheiro.Stat()
			texto := make([]byte, aux.Size())
			ficheiro.Read(texto)
			fmt.Println(string(texto))
			ficheiro.Close()
		}
	}
}
