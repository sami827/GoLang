package main

// import "fmt"

func main() {

	cards := newDeck()

	hand, remainningCards := deal(cards, 5)

	hand.print()
	remainningCards.print()

}
