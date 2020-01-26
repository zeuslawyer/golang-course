package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

// create a new type called "Deck"
// Deck is a 'slice' of string types - i.e.  array of string types

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

/**
returns multiple values
*/
func (d Deck) deal(numHand int) (Deck, Deck) {
	hand := d[:numHand]           // slice the slice and return
	remainingCards := d[numHand:] // slice  the slice and return

	return hand, remainingCards
}

func (d Deck) toString() string {
	deckAsString := []string(d) // type convert(cast) deck into a string slice (array)

	deck := strings.Join(deckAsString, ",") // https://golang.org/pkg/io/ioutil/#WriteFile

	return deck
}

func (d Deck) saveToDisk(filename string) error {
	content := []byte(d.toString()) // type cast deck as a string, into []byte

	return ioutil.WriteFile(filename, content, 0666) // https://golang.org/pkg/io/ioutil/#WriteFile
}

func loadDeckFromFile(filename string) Deck {

	byteSlice, err := ioutil.ReadFile(filename) // https://golang.org/pkg/io/ioutil/#ReadFile

	if err != nil {
		log.Fatal("Error loading file:", err)
		os.Exit(1) // https://golang.org/pkg/os/
	} else {
		log.Printf(" <<< File succesfully loaded >>>: ")
	}

	// fmt.Printf("File contents in a byte array/slice: %s", byteSlice)
	// convert []bytes to string
	sliceAsString := string(byteSlice)

	// convert string to string[]
	stringArr := strings.Split(sliceAsString, ",")
	deck := Deck(stringArr)
	return deck

}

func (d Deck) shuffle() {
	for i := range d {
		random := rand.Intn(len(d) - 1)
		// swap
		d[i], d[random] = d[random], d[i]
	}
}
