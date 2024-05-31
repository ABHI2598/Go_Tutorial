package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Welcome to Game world")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter Rating: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Rating given by user: ", input)
	fmt.Printf("Type of input entered %T ", input)

}
