package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	// a. Takes an integer from the command line.
	fmt.Println("Enter an integer: \n")
	value := os.Args[1]
	input, err := strconv.Atoi(value)

	sl := make([]int, input)
	var sum int

	// b. Generate x random numbers and store in slice
	for i := 0; i < len(sl); i++ {
		sl[i] = rand.Intn(100)
		sum += sl[i]
	}

	// c. Sort

	// d. Time
}
