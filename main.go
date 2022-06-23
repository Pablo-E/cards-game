package main

import "fmt"

func main() {
	cards := newDeck()

	//hand, _ := deal(cards, 5)

	fmt.Println(cards.toString())
	//hand.print()
}
