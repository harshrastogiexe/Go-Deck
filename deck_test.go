package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected Deck Of Length 52, But Got %v", len(d))
	}
	firstCard, lastCard := "Clubs Of Ace", "Spades Of King"

	if d[0] != firstCard {
		t.Errorf(`Expected First Card To Be "%v", But Got "%v"`, firstCard, d[0])
	}

	if d[51] != lastCard {
		t.Errorf(`Expected First Card To Be "%v", But Got "%v"`, lastCard, d[51])
	}
}


func TestToSaveDeckAndReadDeck(t *testing.T) {
	filename := "_decktest"
	os.Remove(filename)

	d := newDeck()

	d.saveToFile(filename)
	cards := NewDeckFromFile(filename)

	if len(cards) != 52 {
		t.Errorf("Expected 52 Cards Got %v", len(cards))
	}

	os.Remove(filename)
}
