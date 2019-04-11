package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	address
}

type address struct {
	addressOne string
	addressTwo string
	city       string
	state      string
	zipcode    int
}

func main() {
	perOne := person{
		firstName: "John",
		lastName:  "Doe",
		address: address{
			addressOne: "Door No: 642",
			addressTwo: "Phase - 2",
			city:       "Mulapet",
			state:      "AP",
			zipcode:    524190,
		},
	}

	personPointer := &perOne
	fmt.Println(*personPointer)
	perOne.print()
	personPointer.updateFirstName("Jane")
	personPointer.updateLastName("Doee")
	perOne.print()
}

func (personPointer *person) updateFirstName(firstname string) {
	(*personPointer).firstName = firstname
}

func (personPointer *person) updateLastName(lastName string) {
	personPointer.lastName = lastName
}

func (personPointer *person) print() {
	fmt.Printf("%+v \n", personPointer)
}
