package main

import (
	"fmt"
	"time"
)

func main() {

	go greeter("Hello")
	greeter("Abhinav")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond) //this will wait for go routine thread to come and print
		fmt.Println(s)
	}
}
