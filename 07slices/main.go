package main

import (
	"fmt"
	"sort"
)

func main() {

	number := []int{1, 2, 4, 3}
	fmt.Println("slice values is: ", number)

	num := make([]string, 2)

	num[0] = "apple"
	num[1] = "banana"

	num = append(num, "mango", "peach", "guava", "plum", "berry")
	fmt.Println("Fruits list: ", num)

	num = append(num, num[2:]...)
	fmt.Println("Fruits list: ", num)

	num = append(num[3:5])
	fmt.Println("Fruits list: ", num)

	highScores := make([]int, 0)

	highScores = append(highScores, 34, 21, 43, 32, 32, 98, 23)

	sort.Ints(highScores)
	fmt.Println("High Scores after Sorting: ", highScores)
	fmt.Println("High Scores are sorted : ", sort.IntsAreSorted(highScores))

	//Remove a value based on index

	var courses = []string{"java", "react", "js", "swift", "ruby", "golang"}

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Course slice after removing index 2 is : ", courses)

}
