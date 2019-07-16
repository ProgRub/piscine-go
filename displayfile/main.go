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
			return
		}
		aux, ai := ficheiro.Stat()
		if ai != nil {
			fmt.Println(ai.Error())
			return
		}
		tam := aux.Size()
		texto := make([]byte, tam)
		ficheiro.Read(texto)
		fmt.Println(string(texto))
		ficheiro.Close()
	}
}
