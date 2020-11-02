package main

func main() {
	cards := newDeckFromFile("myCards")
	cards.print()
	//	cards.saveToFile("myCards")

	/*hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()*/
}
