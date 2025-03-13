package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// set seed for random number generation
	rand.Seed(time.Now().UnixNano())

	// generate a random integer between 1 and 10
	randomNumber := rand.Intn(10-1+1) + 1

	fmt.Println("Random number:", randomNumber)
}
