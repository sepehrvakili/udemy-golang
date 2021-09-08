package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Jones",
		contactInfo: contactInfo{
			email:   "jim@mail.com",
			zipCode: 90000,
		},
	}

	jim.updateName("jimmy")

	jim.print()
	name := "bill"
	fmt.Println(*&name)
}

func (pPointer *person) updateName(newFirstName string) {
	(*pPointer).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
