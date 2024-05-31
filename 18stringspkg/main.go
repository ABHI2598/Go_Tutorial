package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "Hello,abhinav,Pandey"

	newstr := strings.Split(str, ",")
	fmt.Println(newstr)

	s := "one,two,three,two,one,two"

	count := strings.Count(s, "two")
	fmt.Printf("Count of Two substr is: %d\n", count)

	stri := "   hello!g   "
	fmt.Println(strings.TrimSpace(stri))

	str1 := "Hello"
	str2 := "World"

	str3 := strings.Join([]string{str1, str2}, " ")
	fmt.Println(str3)
}
