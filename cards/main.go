package main

import (
	"fmt"
)

func main() {
	var cards Deck

	cards = newDeck()

	// cards.print()

	// hand, cards := cards.deal(5)
	// hand.print()
	// cards.print()

	// fmt.Println(len(cards))
	res := cards.saveToDisk("testFileSave")
	fmt.Println("response from saving file", res)

	cardsFromFile := loadDeckFromFile("testFileSave")
	fmt.Println("response from loading file:\n", cardsFromFile)


	cardsFromFile.shuffle()
	fmt.Println("Shuffled cards from file...\n", cardsFromFile)

}
