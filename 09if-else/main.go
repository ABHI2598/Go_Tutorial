package main

import "fmt"

func main() {

	x := 10

	if x > 5 {
		fmt.Println("X values is greater than 5")
	}

	y := 3

	if y < 2 {
		fmt.Println("Y is smaller")
	} else if y >= 3 {
		fmt.Println("Y is greater or equal")
	} else {
		fmt.Println("No cases match")
	}
}
