package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b

}
func main() {
	fmt.Println("Welcome to my calculator CLI ")
	fmt.Println("select your choice")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	var input int
	fmt.Println("Enter your choice")
	fmt.Scan(&input)
	if input == 1 {
		fmt.Println("Enter the first number")
		var a int
		fmt.Scan(&a)
		fmt.Println("Enter the second number")
		var b int
		fmt.Scan(&b)
		sum := add(a, b)
		fmt.Println("The sum is:", sum)

	}
	if input == 2 {
		fmt.Println("Enter the first number")
		var a int
		fmt.Scan(&a)
		fmt.Println("Enter the second number")
		var b int
		fmt.Scan(&b)
		diff := sub(a, b)
		fmt.Println("The diff is:", diff)

	}
}
