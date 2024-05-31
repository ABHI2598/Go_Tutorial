package main

import "fmt"

func main() {

	//days := []string{"sun", "mon", "tue", "wed", "thurs", "fri", "sat"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for d := range days {
	// 	fmt.Println(days[d])
	// }

	// for i, d := range days {
	// 	fmt.Printf("Index is: %v and value is: %v\n", i, d)
	// }

	rv := 1

	for rv < 10 {

		// if rv == 5 {
		// 	break
		// }
		// if rv == 5 {
		// 	rv++
		// 	continue
		// }

		if rv == 5 {
			goto abhi
		}
		fmt.Println(rv)
		rv++
	}

abhi:
	fmt.Println("Hey hello bhinav this side")
}
