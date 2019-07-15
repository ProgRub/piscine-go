package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	numBytes, e := strconv.Atoi(os.Args[2])
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	} else if len(arguments) > 3 {
		for i := 3; i < len(arguments); i++ {
			fileName := os.Args[i]
			ficheiro, ai := os.Open(fileName)
			if ai != nil {
				fmt.Println(ai.Error())
			} else {
				aux, erro := ficheiro.Stat()
				if erro != nil {
					fmt.Println(erro.Error())
				} else {
					if numBytes != 0 {
						if numBytes < 0 {
							numBytes = -numBytes
						}
						tam := aux.Size()
						fileBytes := make([]byte, tam)
						if numBytes > int(tam) {
							numBytes = int(tam)
						}
						ficheiro.Read(fileBytes)
						if len(arguments) != 4 {
							fmt.Print("==> ")
							fmt.Print(fileName)
							fmt.Println(" <==")
						}
						fmt.Print(string(fileBytes[len(fileBytes)-numBytes:]))
						if i != len(arguments)-1 {
							fmt.Println()
						}
					}
				}
			}
			ficheiro.Close()
		}
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
