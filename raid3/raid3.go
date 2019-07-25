package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"sync"
)

var wg sync.WaitGroup // 1

func main() {

	esquerda, e := ioutil.ReadAll(os.Stdin)
	if e != nil {
		fmt.Println(e.Error())
		return
	} else {
		x := 0
		i := 0
		for esquerda[i] != 10 {
			x++
			i++
		}
		y := 0
		for j := i; j < len(esquerda); j += x + 1 {
			y++
		}
		cmdA := exec.Command("./raid1a", strconv.Itoa(x), strconv.Itoa(y))
		cmdB := exec.Command("./raid1b", strconv.Itoa(x), strconv.Itoa(y))
		cmdC := exec.Command("./raid1c", strconv.Itoa(x), strconv.Itoa(y))
		cmdD := exec.Command("./raid1d", strconv.Itoa(x), strconv.Itoa(y))
		cmdE := exec.Command("./raid1e", strconv.Itoa(x), strconv.Itoa(y))
		outA, _ := cmdA.Output()
		outB, _ := cmdB.Output()
		outC, _ := cmdC.Output()
		outD, _ := cmdD.Output()
		outE, _ := cmdE.Output()
		raids := make([]string, 0)
		if bytes.Equal(outA, esquerda) {
			raids = append(raids, "raid1a")
		} else if bytes.Equal(outB, esquerda) {
			raids = append(raids, "raid1b")
		} else {
			if bytes.Equal(outC, esquerda) {
				raids = append(raids, "raid1c")
			}
			if bytes.Equal(outD, esquerda) {
				raids = append(raids, "raid1d")
			}
			if bytes.Equal(outE, esquerda) {
				raids = append(raids, "raid1e")
			}
		}
		if len(raids) == 0 {
			fmt.Println("Not a Raid function")
			return
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
}
