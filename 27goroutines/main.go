package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {

	// go greeter("Hello")
	// greeter("Abhinav")

	websites := []string{
		"http://github.com",
		"http://go.dev",
		"http://google.com",
		"http://fb.com",
		"http://instagram.com",
	}

	for _, w := range websites {
		go getEndpoint(w)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond) //this will wait for go routine thread to come and print
		fmt.Println(s)
	}
}

func getEndpoint(endpoint string) {
	defer wg.Done()

	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in http call")
	}

	mutex.Lock()
	signals = append(signals, endpoint)
	mutex.Unlock()

	fmt.Printf("%v code and endpoint is: %v \n", result.StatusCode, endpoint)
}
