package main

import (
	"fmt"
	"io/ioutil"
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
			texto, erro := ioutil.ReadFile(arguments[i])
			if erro != nil {
				panic(erro)
			}
			fmt.Println(string(texto))
			fmt.Println("")
		}
	}
}
