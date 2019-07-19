package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	for i := range arguments {
		if arguments[i] == "01" || arguments[i] == "galaxy 01" || arguments[i] == "galaxy" {
			fmt.Println("Alert!!!")
			return
		}
	}
	fmt.Println()
}
