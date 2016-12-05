package main

import "fmt"

func main() {
  var words string

  words = "Hello World!"

  for i := 0; i < 10; i++ {
    fmt.Printf("%s\n", words)
  }
}
