package main

import "fmt"

// Declaring an interface

type bot interface {
	getGreeting() string
}

// Defining similar structs
type englishBot struct{}
type spanishBot struct{}

func main() {
	fmt.Println("Interfaces in Go")
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// printGreeting

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// If we are not using the receiver reference inside the function
// it can be removed like eb englishBot to englishBot
func (englishBot) getGreeting() string {
	return "Hello!!!"
}

func (spanishBot) getGreeting() string {
	return "Hola!!!"
}
