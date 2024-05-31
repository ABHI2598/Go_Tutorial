package main

import "fmt"

func passByReference(n *int) {
	*n = *n * 20
}

func main() {

	var num int = 23

	var ptr *int = &num // ptr := &num

	fmt.Println("address of ptr is: ", ptr)
	fmt.Println("Values of ptr is: ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("Value of num is: ", num)

	value := 10
	passByReference(&value)
	fmt.Println("Value after pass By Reference: ", value)
}
