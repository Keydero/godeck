package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}

	cards = append(cards, "Six Of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(cards)
}

func newCard() string {
	return "Five Of Diamonds"
}
