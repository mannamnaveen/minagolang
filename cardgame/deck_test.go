package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "A of Spades" {
		t.Errorf("Expected first card as A of Spades but got %v ", d[0])
	}

	if d[len(d)-1] != "K of Clubs" {
		t.Errorf("Expected last card as K of Clubs but got %v ", d[len(d)-1])
	}
}

func TestWriteToFileReadFromFile(t *testing.T) {
	os.Remove("_decktest")

	card := newDeck()
	card.writeToFile("_decktest")

	cards := readFromFile("_decktest")

	if len(cards) != 52 {
		t.Errorf("Expected 52 cards but got %v ", len(cards))
	}

	os.Remove("_decktest")
}
