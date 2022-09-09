package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// a. Takes an integer from the command line.
	fmt.Println("Enter an integer: \n")
	value := os.Args[1]
	input, err := strconv.Atoi(value)

	sl := make([]int, input)
	var sum int

	// b. Generate x random numbers and store in slice
	// c. Print sum and output time to calculate
	start := time.Now()
	for i := 0; i < len(sl); i++ {
		sl[i] = rand.Intn(100)
		sum += sl[i]
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("The sum of these %d random numbers is: %d\n", input, sum)
	fmt.Println("The time it took is: %d\n\n", elapsed)

	// d. Use go-routines
	c := make(chan int)
	startP := time.Now()
	go sumRandom(sl[:len(sl)/2], c)
	go sumRandom(sl[len(sl)/2:], c)

	// e. Print parallel
	// f. Two sums and two durations
	x, y := <-c, <-c
	var sumPar int = x + y
	tP := time.Now()
	elapsedP := t.Sub(start)
	fmt.Println("The sum using Go-Routines of these %d random numbers is: %d\n", input, sumPar)
	fmt.Println("The time it took is: %d\n", elapsedP)
}

func sumRandom(sl []int, c chan int) {
	sum := 0
	for _, i := range sl {
		sum += i
	}
	c <- sum
}
