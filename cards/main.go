package main

// import (
// 	"fmt"
// )

func main() {
	var cards Deck

	cards = newDeck()

	// cards.print()

	hand, cards := cards.deal(5)
	hand.print()
	cards.print()

	// fmt.Println(len(cards))

}
