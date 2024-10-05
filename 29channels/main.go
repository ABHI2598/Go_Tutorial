package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to channels")

	var mych = make(chan int, 2)
	var wg = &sync.WaitGroup{}

	wg.Add(2)
	//RECEIVE ONLY CHANNEL
	go func(ch <-chan int, wg *sync.WaitGroup) {
		//close(ch)  //this will lead to infinite loop so to avoid this we will add <-chan to make this RECEIVE Only channel
		value, isChannelOpen := <-ch

		fmt.Println(isChannelOpen)
		fmt.Println(value)
		fmt.Println(<-ch)
		wg.Done()
	}(mych, wg)

	//SEND ONLY CHANNEL
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 7
		close(ch)
		// ch <- 7
		// ch <- 8
		wg.Done()
	}(mych, wg)

	wg.Wait()
}
