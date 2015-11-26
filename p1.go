package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arg, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	total := 0

	for i := 1; i < arg; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	fmt.Println(total)
}
