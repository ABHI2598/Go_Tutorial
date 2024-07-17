package main

import "fmt"

func main() {

	abhinav := User{"Abhinav", "Pandey", true, 23}

	fmt.Printf("User Name is: %s and Age is: %d and LastName is: %s\n", abhinav.FirstName, abhinav.Age, abhinav.LastName)

	abhinav.getStatus()
	abhinav.getLastName()
	abhinav.getFirstName()
	fmt.Printf("User Name is: %s and Age is: %d and LastName is: %s\n", abhinav.FirstName, abhinav.Age, abhinav.LastName)

}

type User struct {
	FirstName string
	LastName  string
	Status    bool
	Age       int
}

func (u User) getStatus() {
	fmt.Printf("User is Online or not: %v\n", u.Status)
}

func (u User) getLastName() {
	u.LastName = "Pandey1"
	fmt.Printf("This User2 last name is: %s\n", u.LastName)
}

func (u User) getFirstName() {
	u.FirstName = "Adarsh"
	fmt.Printf("User FirstName is: %s\n", u.FirstName)
}
