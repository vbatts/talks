// +build ignore

package main

import "fmt"

func main() {
	// START1 OMIT
	fmt.Println("I've got:")
	for card := range ReadEmAndWeep() {
		fmt.Printf(" %s of %s\n", card.Value, card.Suite)
	}
	// STOP1 OMIT
}

type Card struct {
	Suite, Value string
}

// START2 OMIT
func ReadEmAndWeep() <-chan Card {
	c := make(chan Card)
	go func() {
		for _, card := range []Card{
			Card{"Ace", "Jack"},
			Card{"Ace", "Queen"},
			Card{"Ace", "King"},
			Card{"Ace", "Ace"},
		} {
			c <- card
		}
		close(c)
	}()
	return c
}

// STOP2 OMIT
