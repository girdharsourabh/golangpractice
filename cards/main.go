package main

import "fmt"

func main() {

	cards := newDeck()
	//cards.printDeck()

	//hand, remainingDeck := deal(cards, 5)

	//hand.printDeck()
	fmt.Println("########")
	//remainingDeck.printDeck()
	fmt.Println("########")
	//cards.printDeck()
	fmt.Println(cards.toString())

	//cards.saveToFile("output.txt")

	//newCards := readFromFile("output.txt")
	cards.shuffle()
	cards.printDeck()
}
