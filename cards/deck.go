package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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

func (d Deck) deal(numHand int) (Deck, Deck) {
	hand := d[:numHand]
	remainingCards := d[numHand:]

	return hand, remainingCards
}

func (d Deck) toString() string {
	deckAsString := []string(d) // type convert(cast)

	res := strings.Join(deckAsString, ",") // https://golang.org/pkg/io/ioutil/#WriteFile

	return res
}

func (d Deck) saveToDisk(filename string) error {
	content := []byte(d.toString()) // type cast

	return ioutil.WriteFile(filename, content, 0666) // https://golang.org/pkg/io/ioutil/#WriteFile
}
