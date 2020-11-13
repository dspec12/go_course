package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email       string
	phoneNumber string
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
		contactInfo: contactInfo{
			email:       "jim@gmail.com",
			phoneNumber: "888-888-8888",
		},
	}

	jim.updateFirstName("Jimmy")
	jim.print()
}

func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
