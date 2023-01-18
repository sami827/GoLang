package main

// import "fmt"

func main() {

	cards := newDeck()

	hand, remainningCards := deal(cards, 2)

	hand.print()
	remainningCards.print()

}
