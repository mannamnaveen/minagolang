package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// Declaring a deck type extending the slice of string
type deck []string

// print is a reference function defined for the deck type
func (d deck) print() {
	for _, c := range d {
		fmt.Println(c)
	}
}

func newDeck() deck {
	// Declaring a empty card with empty deck
	card := deck{}

	// Declaring a deck with playing card types
	cardTypes := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	// Declaring all the thirteen card values
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	// Defining a for loop to insert all the 52 cards into card deck
	for _, i := range cardTypes {
		for _, j := range cardValues {
			card = append(card, j+" of "+i)
		}
	}

	// Returning the card type
	return card
}

// Deal splits the cards into two decks as defined
func (d deck) deal(number int) (deck, deck) {
	return d[:number], d[number:]
}

// Save the deck of cards to disk file
func (d deck) writeToFile(filename string) {
	ioutil.WriteFile(filename, []byte(d.toString()), 0600)
}

// toString converts a slice into a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// readFromFile gets the input from file and generates a deck
func readFromFile(filename string) deck {
	dec, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	s := strings.Split(string(dec), ",")
	return s
}

func (d deck) shuffle() {
	for i := range d {
		newPos := rand.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
