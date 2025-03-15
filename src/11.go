package main

import "fmt"

func main() {
	const (
		a = 10
		b = 20
	)

	if a > b {
		fmt.Println("a is greater than b")
	} else if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a and b are equal")
	}
}
