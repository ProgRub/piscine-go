package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 4 {
		a, err := strconv.Atoi(arguments[1])
		b, e := strconv.Atoi(arguments[3])
		if err != nil {
			fmt.Println("0")
			return
		}
		if e != nil {
			fmt.Println("0")
			return
		}
		operador := []rune(arguments[2])
		if (int(operador[0]) != 37 && int(operador[0]) != 42 && int(operador[0]) != 43 && int(operador[0]) != 45 && int(operador[0]) != 47) || len(operador) > 1 {
			fmt.Println("0")
			return
		}
		if int(operador[0]) == 37 {
			if b == 0 {
				fmt.Println("No Modulo by 0")
			} else {
				fmt.Println(a % b)
			}
		} else if int(operador[0]) == 42 {
			fmt.Println(a * b)
		} else if int(operador[0]) == 43 {
			fmt.Println(a + b)
		} else if int(operador[0]) == 45 {
			fmt.Println(a - b)
		} else if int(operador[0]) == 47 {
			if b == 0 {
				fmt.Println("No division by 0")
			} else {
				fmt.Println(a / b)
			}
		}
	}
}
