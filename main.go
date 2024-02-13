package main

import "fmt"

type contactInfo struct {
	phoneNo int
	email   string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	person1 := person{
		firstName: "Steve",
		lastName:  "Smith",
		contactInfo: contactInfo{
			phoneNo: 123456,
			email:   "asdbfa@sdfjl.com",
		},
	}

	person1.Print()
}

func (p person) Print() {
	fmt.Printf("%+v", p)
}
