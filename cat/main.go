package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		reader := bufio.NewReader(os.Stdin)
		util, err := reader.ReadString('\n')
		if err != nil && util != "" {
			fmt.Println(err.Error())
		} else {
			if util != "\n" {
				fmt.Println(util)
			}
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
