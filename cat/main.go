package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		/* Reader := bufio.NewReader(os.Stdin)
		in, err := Reader.ReadString('\n')
		in = strings.Replace(in, "\n", "", -1)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if in != "" {
				fmt.Println(in)
			}
		} */
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fmt.Println(scanner.Text())
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
