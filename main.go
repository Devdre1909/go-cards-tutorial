package main

func main() {
	cards := newDeckFromFile("mydecktxt")
	cards.print()
}
