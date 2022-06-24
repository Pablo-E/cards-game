package main

func main() {
	cards := newDeck()

	cards.shuffle()

	hand, _ := deal(cards, 5)

	hand.print()
}
