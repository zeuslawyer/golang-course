package main

import (
	"fmt"
)

func main() {
	var card = newCard("Ace of Spades")

	// slice
	cards := []string{newCard("Seven of Hearts"), newCard("Four of Clubs")}

	// append new element (creates and returns a new array, does not modify existing)
	cards = append(cards, newCard("5 of Diamonds"))
	
	// var card string = newcard()

	fmt.Println(card, cards)

	for ind, card := range cards {
		fmt.Println(ind, card)
	}
}

func newCard(suit string) string {
	// return "Five of Diamonds"
	return suit
}
