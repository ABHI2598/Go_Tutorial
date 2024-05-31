package main

import "fmt"

func main() {

	studGrades := make(map[string]int)

	studGrades["alice"] = 90
	studGrades["bob"] = 89
	studGrades["abhi"] = 97
	studGrades["ayu"] = 87

	fmt.Println("Map is:", studGrades)

	fmt.Println(studGrades["alice"])

	delete(studGrades, "ayu")
	fmt.Println(studGrades)

	grade, exists := studGrades["abhi"]
	fmt.Println("Grade exists: ", exists)
	fmt.Println("Grade is: ", grade)

	for name, grade := range studGrades {
		fmt.Printf("%s %d\n", name, grade)
	}

}
