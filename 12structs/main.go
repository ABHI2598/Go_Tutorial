package main

import "fmt"

func main() {

	abhinav := User{"Abhinav", "abhi@go.dev", true, 25}

	fmt.Printf("User details are: %+v\n", abhinav)
	fmt.Printf("User name is: %v and Age is: %v\n", abhinav.Name, abhinav.Age)

	user2 := new(User)
	user2.Name = "Abhi"
	user2.Age = 23
	user2.Status = false

	fmt.Printf("User2 details are: %+v\n", *user2)
	fmt.Printf("User2 name is: %v and Age is: %v\n", user2.Name, user2.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
