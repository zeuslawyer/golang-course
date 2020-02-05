package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	var i *int

	length := len(d)
	i = &length

	println(" POINTERS:   here is i ", i, *i)

	if length != 54 {
		t.Errorf("New Deck has produced an incorrect number of cards.  Should produce 54, but instead produced: %d", length)
	}

	if *i != length {
		t.Errorf(" the pointer i has the wrong value")
	}

}

func TestFirstAndLastCards(t *testing.T) {
	d := newDeck()

	first := d[0]
	last := d[len(d)-1]

	if first != "Two of Spades" {
		t.Errorf("Expected: 'Two of Spades'. Actual: %s", first)
	}
	if last != "Joker 2" {
		t.Errorf("Expected: 'Ace of Hearts'. Actual: %s", last)
	}

}

func TestSaveAndLoadIO(t *testing.T) {
	filename := "_test-save-to-disk.txt"

	// set up - remove any leftover files
	os.Remove(filename)

	d := newDeck()
	err := d.saveToDisk(filename)

	if err != nil {
		t.Errorf("Error Saving File:  %s", err)
	}

	loaded := loadDeckFromFile(filename)

	if len(d) != len(loaded) {
		t.Errorf("Error loading deck.  Its length is NOT the same as d at %d cards", len(loaded))
	}

	// cleanup file
	os.Remove(filename)
}
