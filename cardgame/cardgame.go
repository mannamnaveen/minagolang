package main

import "fmt"

func main() {
	fmt.Println("----------------------")
	fmt.Println("The King of Card Games")
	fmt.Println("----------------------")
	// cards := newDeck()
	// cards.writeToFile("newdeckcards")
	// cards := readFromFile("newdeckcards")
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	dealCards, leftCards := cards.deal(10)
	fmt.Println("----------------------")
	dealCards.print()
	fmt.Println("----------------------")
	leftCards.print()
}
