package main

import (
	"fmt"
)

func main() {
	var card = newCard("Ace of Spades")
 
	// slice
	cards := Deck {newCard("Seven of Hearts"), newCard("Four of Clubs")}

	// append new element (creates and returns a new array, does not modify existing)
	cards = append(cards, newCard("5 of Diamonds"))
	

	fmt.Println(card, cards)

	cards.print()

}

func newCard(suit string) string {
	// return "Five of Diamonds"
	return suit
}
