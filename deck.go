package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Tree", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" Of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckfromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(byteSlice), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func rand(r int) int {
	//return 3 % 4
	tk := time.Now()
	return int(tk.Second() % r)
}
