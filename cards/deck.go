package main

import (
	"fmt"
)

// create a new type called "Deck"
// Deck is a slice of string types

type Deck []string

func newDeck() Deck {
	deck := Deck{}

	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	cardSuits := []string{"Spades", "Diamond", "Clubs", "Hearts"} 

	for _, val := range cardValues {
		for _, suit := range cardSuits {
			deck = append(deck, val+" of "+suit)
		}
	}

	deck = append(deck, "Joker 1", "Joker 2")

	return deck

}

// wiring up functions to Deck type using receivers
// receivers are not named...

func (d Deck) print() {
	for ind, card := range d {
		fmt.Println(ind, card)
	}

	fmt.Println("\n\nOk, you're done.")

}
