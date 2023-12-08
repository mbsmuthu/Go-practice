package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	// Declaring an array
	bookings := []string{"Tom Hanks"}
	fmt.Printf("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your last name:")
	fmt.Scan(&lastName)
	bookings = append(bookings, firstName+" "+lastName)
	// fmt.Printf("Hi %v/n", bookings)

	for _, booking := range bookings {
		fmt.Println(booking)
	}

}
