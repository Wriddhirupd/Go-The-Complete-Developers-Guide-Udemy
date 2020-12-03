package main

import "fmt"

func main() {
	//var card string = "Ace of spades"
	// card := "Ace of spades" // := for new initialization
	// card = "five of diamonds"

	// card := newCard()
	// fmt.Println(card)

	// card := deck{"Ace of Spades", newCard(), "Six of diamonds"}  //Type declaration
	// card = append(card, "Seven of cards")
	// fmt.Println(card)

	// for i, car := range card {
	// 	fmt.Println(i, " Mr "+car)
	// }
	cards := newDeck()
	cards.print()
	fmt.Println()

	// //MULTIPLE RETURN VALUES
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// fmt.Println()
	// remainingDeck.print()

	// CONCAT TO STRING
	// fmt.Println(cards.toString())

	// WRITE TO FILE
	// cards.saveToFile("my_Cards")

	// READ FROM FILE
	// d := newDeckFromFile("my_Cards")
	// d.print()

	//SHUFFLE CARDS DECK
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "five of diamonds"
}
