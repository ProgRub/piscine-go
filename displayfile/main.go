package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("File name missing")
	} else if len(arguments) > 2 {
		fmt.Println("Too many arguments")
	} else {
		ficheiro, err := os.Open(arguments[1])
		check(err)
		aux, ai := ficheiro.Stat()
		check(ai)
		tam := aux.Size()
		texto := make([]byte, tam)
		ficheiro.Read(texto)
		fmt.Println(string(texto))
		ficheiro.Close()
	}
}
