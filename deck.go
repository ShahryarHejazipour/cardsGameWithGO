package main

import "fmt"

//Create a new type of 'deck'
//which is a slice of string
type deck []string

func newDeck() deck {

	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "clubs"}
	cardValues := []string{"Ace", "Two", "Three", "four"}

	for _, siut := range cardSuits {
		for _, value := range cardValues {

			cards = append(cards, siut+" of "+value)
		}

	}

	return cards
}

func (d deck) print() {
	for index, card := range d {

		fmt.Println(index, card)
	}
}
