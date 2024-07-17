package main

import (
	"fmt"
	"net/url"
)

var myurl string = "https://abhi.dev:8080/learn?name=abhinav&course=golang"

func main() {

	u, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	// fmt.Println(u.Scheme)
	// fmt.Println(u.Host)
	// fmt.Println(u.Port())
	// fmt.Println(u.Path)
	// fmt.Println(u.RawQuery)

	q := u.Query()

	fmt.Println(q.Get("name"))
	fmt.Println(q.Get("course"))
	fmt.Println(q["name"])

}
