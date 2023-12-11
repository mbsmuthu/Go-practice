package main

import "fmt"

func main() {

	var name string
	var age int
	fmt.Println("Please enter your name:")
	fmt.Scan(&name)
	fmt.Println("Please enter your age:")
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You have to be 18 or above to be eligible to vote in India")
	}
}
