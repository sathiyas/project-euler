package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func findfib(x int) int {

	if x == 1 || x == 2 {
		return x
	} else {
		m := findfib(x - 1)
		n := findfib(x - 2)
		return (m + n)
	}
}

func main() {

	arg, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	total := 0

	idx := 1
	for {
		j := findfib(idx)
		if j%2 == 0 {
			total += j
		}
		if j > arg {
			break
		}
		idx++
	}

	fmt.Println(total)

}
