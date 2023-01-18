package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := deck{"Spades", "Diamond", "Hearts", "Clubs"}

	cardValues := deck{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {

			cards = append(cards, value+" of "+suit, "\n")
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Printf(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
