package main

import "fmt"

const Token string = "kjndkndaskjfdsfkjdbhkfbkdj" //Public var since first letter is capital

func main() {

	var userName string = "abhinav"
	fmt.Println(userName)
	fmt.Printf("Variable is of Type: %T \n", userName)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of Type: %T \n", isLoggedIn)

	var data uint8 = 255
	fmt.Println(data)
	fmt.Printf("Variable is of Type: %T \n", data)

	var users int = 37484
	fmt.Println(users)
	fmt.Printf("Variable is of Type: %T \n", users)

	var count float32 = 453.9887821323
	fmt.Println(count)
	fmt.Printf("Variable is of Type: %T \n", count)

	var count1 float64 = 453.9887821323
	fmt.Println(count1)
	fmt.Printf("Variable is of Type: %T \n", count1)

	//Implicit type

	var name = "abhinav.n"
	fmt.Println(name)

	// no-var type

	userfname := "adarsh"
	fmt.Println(userfname)

	fmt.Printf("Const Variable: %s \n", Token)

}
