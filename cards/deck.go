package main

import (
	"fmt"
)

// create a new type called "Deck"
// Deck is a slice of string types

type Deck []string

// wiring up functions to Deck type using receivers
// receivers are not named...

func (d Deck) print()  {
	 for ind, card := range d	{
		 fmt.Println(ind, card)
	 }

	 fmt.Println("\n\nOk, you're done.")

}

