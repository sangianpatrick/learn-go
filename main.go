package main

func main() {
	// var card string = "Ace of Spades"
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 7)
	// hand.print()
	// remainingCards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// cards.saveToFile("my_cards.txt")

	cards := newDeckFromFile("my_cards.txt")
	cards.shuffle()
	cards.print()
}
