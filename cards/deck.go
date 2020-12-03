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
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
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
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666) //anyone can read and write the file "0666"

}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// return strings.Split(string(bs), ",")
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	/* SOURCE TAKES A INT64 NUMBER WHICH IS PROVIDED BY UNIXNOW() OF TIME OBJECT.
	SO EVERYTIME IT WILL TAKE THE CURENT TIME AND CREATE A SOURCE.
	R TAKES SOURCE AS ARGUMENT TO SEED A RANDOM NUMBER.
	SO WE CAN USE R AS A RAND OBJECT TO CALL INTN WHICH WILL CREATE RANDOM NUMBERS OF DIFFERENT TYPES
	AND DO NOT CREATE THE SAME RANDOMIZATION.
	*/
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
