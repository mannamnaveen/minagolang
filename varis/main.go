package main

import "fmt"

func main() {
	fmt.Println("Declaring variables in go lang.")
	var firstName string
	var lastName string
	var age int
	var height float32
	var weight float32

	firstName = "John"
	lastName = "Doe"
	age = 23
	height = 1.8
	weight = 196.5
	fmt.Println("Full Name :", firstName, lastName)
	fmt.Println("Age : ", age)
	fmt.Println("Height : ", height)
	fmt.Println("Weight : ", weight)
}
