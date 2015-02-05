// +build ignore

package main

import "fmt"

func main() {
	// START1 OMIT
	infos := map[string]Info{
		"Michael": Info{
			City:   "Happyville",
			Status: Open,
		},
		"Jan": Info{
			City:   "Confusville",
			Status: Closed,
		},
		"Sean": Info{
			City:   "Septober",
			Status: Complicated,
		},
		"Silvia": Info{
			City:   "Curiousville",
			Status: Curios,
		},
	}
	for name, info := range infos {
		fmt.Printf("%q is %s in %q\n", name, info.Status, info.City)
	}
	// STOP1 OMIT
}

type Info struct {
	City   string
	Status Status
}
type Status int

var (
	Open        = Status(0)
	Closed      = Status(1)
	Complicated = Status(2)
	Curios      = Status(3)
)

func (s Status) String() string {
	switch s {
	case Open:
		return "open"
	case Closed:
		return "closed"
	case Complicated:
		return "complicated"
	}
	return "hurr"
}
