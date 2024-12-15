package main

var deckSize int

func main() {

	// cards := newDeck()

	//	Iterating the slice of array
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// fmt.Println("=================")
	// remainingDeck.print()

	// fmt.Println(cards)

	// greeting := "Hello there"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// fmt.Println(cards.toString())

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Ace of spades"
// }
