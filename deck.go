package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}
	cards := deck{}

	for _, suit := range cardSuits {
		for _, val := range cardValue {
			cards = append(cards, suit+" Of "+val)
		}
	}

	return cards
}

func deal(d deck, count int) (deck, deck) {
	if count > len(d) {
		return d, deck{}
	}

	return d[:count], d[count:]
}

func (d deck) Print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func deckFrom(str string) deck {
	return deck(strings.Split(str, ","))
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func NewDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), ",")
}

func (d deck) Suffle() {
	num := time.Now().UnixNano()
	r := rand.New(rand.NewSource(num))

	for i := range d {
		random := r.Intn(len(d) - 1)
		d[i], d[random] = d[random], d[i]
	}
}
