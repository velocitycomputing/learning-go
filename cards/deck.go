package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"♥", "♦", "♠", "♣"}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	// "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// type conversion
	return strings.Join([]string(d), ",")

}

func saveToFile(d deck, filename string) {
	os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() deck {
	// rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
	// fmt.Println(d)
	return d
}
