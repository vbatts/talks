// +build ignore

package main

// START1 OMIT
// exported
type Foo struct {
	Bar string // exported
	baz bool   // private
}

// private
type bif struct {
	Harf, Buz int64 // doesn't matter
}

// STOP1 OMIT

func main() {
}
