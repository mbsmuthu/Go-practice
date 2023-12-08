package main

import "fmt"

func main() {
	var name = "Mark"
	var email string
	fmt.Printf("Hello %v\n", name)
	fmt.Printf("Please enter your email:")
	fmt.Scan(&email)
	fmt.Printf("Notifications will be sent to %v", email)
}
