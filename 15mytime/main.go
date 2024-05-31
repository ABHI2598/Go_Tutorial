package main

import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	createdDate := time.Date(2020, time.August, 25, 23, 12, 07, 0, time.UTC)
	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 Mon 15:04:05"))
}
