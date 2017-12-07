package main

func main() {
	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	cards := newDeckfromFile("my_cards")
	cards.shuffle()
	/*hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()*/
}
