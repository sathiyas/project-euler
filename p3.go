package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func isPrime(x int64) bool {

	if x == 1 || x == 2 || x == 3 {
		return true
	}
	if x%2 == 0 || x%3 == 0 {
		return false
	}

	var i int64

	for i = 3; i < int64(math.Sqrt(float64(x)))+1; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func main() {

	arg, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	var i int64 = 1

	var biggestPrime int64

	for i = 1; i < int64(math.Sqrt(float64(arg)))+1; i++ {
		if int64(arg)%i == 0 {
			//fmt.Println("Factor ", i)
			bl := isPrime(i)
			if bl {
				biggestPrime = i
				fmt.Println("Prime Factor ", i)
			}
		}
	}
	fmt.Println("Biggest Prime Factor ", biggestPrime)

}
