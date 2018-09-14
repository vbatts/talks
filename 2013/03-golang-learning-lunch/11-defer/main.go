package main

import "fmt"

func main() {
	// START OMIT
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	fmt.Println("HOOTY HOO!")
	// STOP OMIT
}
