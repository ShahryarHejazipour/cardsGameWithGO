package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

//Create a new type of 'deck'
//which is a slice of string
type deck []string

func newDeck() deck {

	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "clubs"}
	cardValues := []string{"Ace", "Two", "Three", "four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {

			cards = append(cards, suit+" of "+value)
		}

	}

	return cards
}

//Comment For Git Test
func (d deck) print() {
	for index, card := range d {

		fmt.Println(index, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(fileName string) error {

	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)

}

func newDeckFromFile(fileName string) deck {

	bd, err := ioutil.ReadFile(fileName)

	if err != nil {

		//Option #1 - log the error and return a call to newDeck()
		//Option #2 - log the error and entirely quite the program

		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bd), ",")

	return deck(s)

}

func (d deck) shuffle() {

	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]

	}

}
