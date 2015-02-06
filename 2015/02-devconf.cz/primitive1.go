// +build ignore

package main

import "fmt"

func main() {
	// START1 OMIT
	names := []string{
		"Michael",
		"Jan",
		"Sean",
		"Silvia",
	}
	for i, name := range names {
		fmt.Printf("%q is the %d name\n", name, i)
	}
	for i := range names {
		fmt.Printf("%q is the %d name\n", names[i], i)
	}
	for _, name := range names {
		fmt.Printf("%q is the ... name\n", name)
	}
	// STOP1 OMIT
}
