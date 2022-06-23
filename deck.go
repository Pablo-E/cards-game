package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	for _, suit := range cardSuits {
		for i := 1; i <= 12; i++ {
			switch i {
			case 1:
				cards = append(cards, fmt.Sprintf("A of %v", suit))
			case 10:
				cards = append(cards, fmt.Sprintf("J of %v", suit))
			case 11:
				cards = append(cards, fmt.Sprintf("Q of %v", suit))
			case 12:
				cards = append(cards, fmt.Sprintf("K of %v", suit))
			default:
				cards = append(cards, fmt.Sprintf("%v of %v", i, suit))
			}
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
