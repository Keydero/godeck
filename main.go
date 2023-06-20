package main

import "fmt"

func main() {
	cards := newDeck()

	//cards.print()

	//fmt.Println(cards[:3])

	hand, remainingCards := deal(cards, 5)
	remainingCards.print()

	fmt.Println(hand.toString())
	cards.saveToFile("my_cards")

}
