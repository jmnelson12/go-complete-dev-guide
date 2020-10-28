package main

func main() {
	// cards := newDeck()
	// myHand, deck := deal(cards, 4)
	// deck.print()
	// fmt.Println("----")
	// myHand.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
