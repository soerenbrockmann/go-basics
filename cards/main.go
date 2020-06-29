package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards := newDeckFromFile("my_cards")
	// cards.print()
	// Both are the same
	// var card string = "Ace of Spades"
	// Dynamic typing, Infer
	// card := newCard()
	// fmt.Println(newCard())
	// printState()
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// fmt.Println(cards.toString())
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// cards.print()
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// fmt.Println(cards)
	// greeting := "Hi there"
	// fmt.Println([]byte(greeting))
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
