package main

import "fmt"

func main() {

	var fruits [4]string
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[3] = "kiwi"

	fmt.Println("fruits list is: ", fruits)
	fmt.Println("len of fruits is: ", len(fruits))

	var vegs = [3]string{"batata", "onion"}
	fmt.Println("Vegs list is: ", vegs)
	fmt.Println("Length of vegs list is: ", len(vegs))

}
