package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 54 {
		t.Errorf("New Deck has produced an incorrect number of cards.  Should produce 54, but instead produced: %d", len(d))
	}

}
