package main

func main() {

	//cards := deck{"Ace of Diamonds", newCard()}

	//	fmt.Println(cards)

	//	cards = append(cards, "Ace of Hearts")

	//	cards.print()
	cards := newDeck()
	cards.print()
	//fmt.Println(cards)

}

func newCard() string {
	return "Five Of Diamonds"
}
