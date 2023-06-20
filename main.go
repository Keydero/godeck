package main

func main() {
	cards := newDeck()

	//cards.print()

	//fmt.Println(cards[:3])

	hand, remainingCards := deal(cards, 5)
	remainingCards.print()
	hand.print()

}
