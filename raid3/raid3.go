package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	esquerda, _ := ioutil.ReadAll(os.Stdin)
	if len(esquerda) == 0 {
		fmt.Println("Not a Raid function")
		return
	}
	x := 0
	i := 0
	for esquerda[i] != 10 && i < len(esquerda) {
		x++
		i++
	}
	y := 0
	for j := i; j < len(esquerda); j += x + 1 {
		y++
	}
	raids := make([]string, 0)
	raids = append(raids, "raid1a")
	raids = append(raids, "raid1b")
	raids = append(raids, "raid1c")
	raids = append(raids, "raid1d")
	raids = append(raids, "raid1e")
	if len(raids) == 0 {
		fmt.Println("Not a Raid function")
	} else if len(raids) == 1 {
		fmt.Printf("[%v] [%v] [%v]", raids[0], x, y)
		fmt.Println()
	} else {
		for i, z := range raids {
			if i != len(raids)-1 {
				fmt.Printf("[%v] [%v] [%v] || ", z, x, y)
			} else {
				fmt.Printf("[%v] [%v] [%v]", z, x, y)
				fmt.Println()
			}
		}
	}
}
