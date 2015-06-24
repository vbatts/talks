// +build OMIT
package main

import (
	"io"
	"os"
)

// START1 OMIT
type Message string

func (m Message) WriteTo(w io.Writer) (int, error) {
	return w.Write([]byte(m))
}

func (m Message) Reset() {
	m = ""
}

func main() {
	m := Message("Hello World\n")
	m.WriteTo(os.Stdout)
	m.Reset()
	m.WriteTo(os.Stdout)
}

// STOP1 OMIT
