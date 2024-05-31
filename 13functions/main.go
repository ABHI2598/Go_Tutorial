package main

import "fmt"

func main() {
	fmt.Println("Hello Abhinav")
	greeter()
	result := adder(3, 5)
	sumfunc := sum(1, 2, 3, 4, 5)

	fmt.Printf("Sum is: %d\n", result)
	fmt.Printf("Sum for vardiac func: %d\n", sumfunc)
}

func adder(v1 int, v2 int) int {
	return v1 + v2
}

func sum(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}
func greeter() {
	fmt.Println("Namastey Golang")
}
