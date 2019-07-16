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
			fmt.Println(err.Error())
		} else {
			fmt.Println(in)
		}
	} else {
		for i := 1; i < len(arguments); i++ {
			ficheiro, e := os.Open(arguments[i])
			if e != nil {
				fmt.Println(e.Error())
				break
			} else {
				aux, _ := ficheiro.Stat()
				texto := make([]byte, aux.Size())
				ficheiro.Read(texto)
				fmt.Println(string(texto))
				fmt.Println("")
				ficheiro.Close()
			}
		}
	}
}
