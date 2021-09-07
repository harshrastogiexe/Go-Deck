package main

func main() {
	cards := NewDeckFromFile("mycards.txt")
	cards.Suffle()
	cards[1:10].Print()
}
