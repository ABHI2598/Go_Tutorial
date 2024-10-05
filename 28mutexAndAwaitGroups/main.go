package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to mutex and wait group")

	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}

	score := []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("ONE Race")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, m)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("TWO Race")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, m)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("THREE Race")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, m)

	wg.Wait()
	fmt.Println(score)

}
