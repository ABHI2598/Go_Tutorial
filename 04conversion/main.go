package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num int = 432
	var num1 float64 = float64(num)
	fmt.Println("Number converted int float: ", num1)

	name := "1234"

	n, _ := strconv.Atoi(name)
	fmt.Println("String after conversion into Integer: ", n)
	fmt.Printf("Type of converted Integer: %T\n", n)

	nu := 123
	na := strconv.Itoa(nu)
	fmt.Println("Integer after conversion into String: ", na)
	fmt.Printf("Type of na: %T\n", na)

	nm := "1323"
	l, _ := strconv.ParseFloat(strings.TrimSpace(nm), 64)
	fmt.Println("String after conversion into float: ", l)
	fmt.Printf("Type of converted Float: %T\n", l)

	fl := "true"
	h, _ := strconv.ParseBool(fl)
	fmt.Println("String after conversion into Boolean: ", h)
	fmt.Printf("Type of converted Boolean: %T", h)

}
